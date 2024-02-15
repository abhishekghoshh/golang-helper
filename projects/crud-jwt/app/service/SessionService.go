package service

import (
	"crud/app/config"
	"crud/app/db"
	"crud/app/domain"
	"crud/app/dto"
	errr "crud/app/error"
	"crud/app/logger"
	"crud/app/util"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type SessionService struct {
	userRepo  *db.UserRepo
	tokenRepo *db.TokenRepo
	encoder   *util.PasswordEncoder
}

func NewSessionService(userRepo *db.UserRepo, tokenRepo *db.TokenRepo, encoder *util.PasswordEncoder) *SessionService {
	return &SessionService{
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
		encoder:   encoder,
	}
}
func (sessionService *SessionService) Login(loginRequest *dto.LoginRequest, loginResponse *dto.LoginResponse) *errr.AppError {
	if err := util.DecodeLoginRequest(loginRequest); err != nil {
		return err
	}
	var user domain.User
	if err := sessionService.userRepo.GetUserByUserName(&user, loginRequest.UserName); err != nil {
		return err
	}
	if matched, err := sessionService.encoder.Match(loginRequest.Password, user.Password); err != nil || !matched {
		if err != nil {
			return err
		} else {
			return errr.NewAuthenticationError("Password did not match")
		}
	}

	loggedInUserInAnotherSession := domain.LoggedInUser{}
	if err := sessionService.tokenRepo.GetUserEntryByUserName(loginRequest.UserName, &loggedInUserInAnotherSession); err == nil {
		tokenValidityRequest, tokenValidityResponse := dto.TokenValidityRequest{Token: loggedInUserInAnotherSession.AccessToken}, dto.TokenValidityResponse{}
		if sessionService.VerifyToken(&tokenValidityRequest, &tokenValidityResponse); tokenValidityResponse.Message != "" {
			return errr.NewAuthenticationError("Already logged in a different session")
		}
		sessionService.tokenRepo.RemoveUserEntry(loginRequest.UserName)
	}

	accessTokenSecret := config.Get("user.jwt.accessTokenSecret")
	accessTokenDuration, _ := time.ParseDuration(config.Get("user.jwt.accessTokenDuration"))
	signedAccessToken, _ := sessionService.signedString(user.UserName, user.Role(), accessTokenDuration, "access_token", accessTokenSecret)

	refreshTokenSecret := config.Get("user.jwt.refreshTokenSecret")
	refreshTokenDuration, _ := time.ParseDuration(config.Get("user.jwt.refreshTokenduration"))
	signedRefreshToken, _ := sessionService.signedString(user.UserName, user.Role(), refreshTokenDuration, "refresh_token", refreshTokenSecret)

	loggedInUser := domain.NewLoggedInUser(user.UserName, signedAccessToken, signedRefreshToken)
	if err := sessionService.tokenRepo.SaveTokenAfterLogin(loggedInUser); err != nil {
		return err
	}
	loginResponse.Create(signedAccessToken, signedRefreshToken)
	return nil
}

func (sessionService *SessionService) Logout(multiValuedHeaders map[string][]string, logoutResponse *dto.LogoutResponse) *errr.AppError {
	accessToken, err := sessionService.getAccessTokenFromHeader(multiValuedHeaders)
	if err != nil {
		return err
	}
	var claims *domain.TokenClaim
	if claims, err = sessionService.verifyAccessToken(accessToken); err != nil {
		if strings.Contains(err.Error(), "token is expired") {
			go sessionService.deleteTokenInAsync(accessToken, sessionService.tokenRepo)
		}
		return err
	}
	if err := sessionService.tokenRepo.RemoveUserEntry(claims.Username); err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return errr.NewAuthenticationError("You have already logged out")
		}
		return err
	}
	logoutResponse.Create("Logged out successfully")
	return nil
}

func (sessionService *SessionService) VerifyToken(tokenValidityRequest *dto.TokenValidityRequest, tokenValidityResponse *dto.TokenValidityResponse) *errr.AppError {
	claims, err := sessionService.verifyAccessToken(tokenValidityRequest.Token)
	if err != nil {
		return err
	}
	tokenValidityResponse.Create("token is valid " + time.Unix(claims.ExpiresAt, 0).Format("02-01-2006, 15:04:05"))
	return nil
}
func (sessionService *SessionService) RefreshToken(tokenRefreshRequest *dto.TokenRefreshRequest, tokenRefreshResponse *dto.TokenRefreshResponse) *errr.AppError {
	return nil
}
func (sessionService *SessionService) GetCurrentAccount(multiValuedHeaders map[string][]string, accountDto *dto.AccountDTO) *errr.AppError {
	accessToken, err := sessionService.getAccessTokenFromHeader(multiValuedHeaders)
	if err != nil {
		return err
	}
	claims, err := sessionService.verifyAccessToken(accessToken)
	if err != nil {
		return err
	}
	user := domain.User{}
	if err := sessionService.userRepo.GetUserWithAccountByUserName(&user, claims.Username); err != nil {
		return err
	}
	accountDto.Create(*user.Account, user)
	return nil
}
func (sessionService *SessionService) verifyAccessToken(accessToken string) (*domain.TokenClaim, *errr.AppError) {
	accessTokenSecret := config.Get("user.jwt.accessTokenSecret")
	var jwtToken *jwt.Token
	var err *errr.AppError
	if jwtToken, err = sessionService.fetchToken(accessToken, accessTokenSecret); err != nil {
		return nil, err
	}
	claims := jwtToken.Claims.(*domain.TokenClaim)
	loggedInUser := domain.LoggedInUser{}
	if err := sessionService.tokenRepo.GetUserEntryByUserName(claims.Username, &loggedInUser); err != nil {
		return nil, errr.NewBadRequestError("token is not valid")
	}
	if !strings.EqualFold(loggedInUser.AccessToken, accessToken) {
		return nil, errr.NewBadRequestError("token is not valid")
	}
	return claims, nil
}
func (*SessionService) getAccessTokenFromHeader(multiValuedHeaders map[string][]string) (string, *errr.AppError) {
	if len(multiValuedHeaders["Authorization"]) == 0 || !strings.HasPrefix(multiValuedHeaders["Authorization"][0], "Bearer ") {
		return "", errr.NewBadRequestError("No Bearer Authorization header is present")
	}
	return strings.TrimPrefix(multiValuedHeaders["Authorization"][0], "Bearer "), nil
}
func (*SessionService) deleteTokenInAsync(accessToken string, tokenRepo *db.TokenRepo) {
	if e := tokenRepo.RemoveUserEntryByAccessToken(accessToken); e != nil {
		logger.Error("Error while deleting the access token " + accessToken + " is " + e.Error())
	}
}
func (sessionService *SessionService) fetchToken(token, tokenSecret string) (*jwt.Token, *errr.AppError) {
	jwtToken, err := sessionService.jwtTokenFromString(token, tokenSecret)
	if err != nil {
		return nil, errr.NewBadRequestError(err.Error())
	} else if nil != jwtToken && !jwtToken.Valid {
		return nil, errr.NewBadRequestError("Access token is not valid ")
	}
	return jwtToken, nil
}
func (sessionService *SessionService) jwtTokenFromString(tokenString, tokenSecret string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &domain.TokenClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})
	if err != nil {
		logger.Error("Error while parsing token: " + err.Error())
		return nil, err
	}
	return token, nil
}
func (*SessionService) signedString(username string, role string, tokenDuration time.Duration, tokenType string, secret string) (string, error) {
	tokenClaim := domain.CreateTokenClaims(tokenType, username, role, tokenDuration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaim)
	return token.SignedString([]byte(secret))
}
