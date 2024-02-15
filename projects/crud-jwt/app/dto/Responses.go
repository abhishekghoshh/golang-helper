package dto

type AccountOperation struct {
	UserId string `json:"userId,omitempty"`
	Messge string `json:"message,omitempty"`
}

func (accountOperation *AccountOperation) CreateAccountOperation(userId string, message string) {
	accountOperation.UserId = userId
	accountOperation.Messge = message
}

type LoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (loginResponse *LoginResponse) Create(accessToken, refreshToken string) {
	loginResponse.AccessToken = accessToken
	loginResponse.RefreshToken = refreshToken
}

type LogoutResponse struct {
	Message string `json:"message"`
}

func (logoutResponse *LogoutResponse) Create(message string) {
	logoutResponse.Message = message
}

type TokenValidityRequest struct {
	Token string `json:"token"`
}

type TokenValidityResponse struct {
	Message string `json:"message"`
}

func (tokenValidityResponse *TokenValidityResponse) Create(message string) {
	tokenValidityResponse.Message = message
}

type TokenRefreshRequest struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type TokenRefreshResponse struct {
	AccessToken string `json:"accessToken"`
}

func (tokenRefreshRequest *TokenRefreshResponse) Create(token string) {
	tokenRefreshRequest.AccessToken = token
}
