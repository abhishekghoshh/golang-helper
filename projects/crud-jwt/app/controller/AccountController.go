package controller

import (
	"crud/app/dto"
	"crud/app/service"
	"crud/app/util"

	"net/http"

	"github.com/gorilla/mux"
)

type AccountController struct {
	router         *mux.Router
	accountService *service.AccountService
	sessionService *service.SessionService
}

func NewAccountController(router *mux.Router, accountService *service.AccountService, sessionService *service.SessionService) *AccountController {
	return &AccountController{
		router:         router,
		accountService: accountService,
		sessionService: sessionService,
	}
}

func (accountController *AccountController) AddRoutes() *AccountController {
	util.AddRoute(accountController.router, "/accounts", accountController.getAccounts, http.MethodGet, "GetAccounts")
	util.AddRoute(accountController.router, "/account/{accountId:[0-9]+}", accountController.getAccount, http.MethodGet, "GetAccountById")
	util.AddRoute(accountController.router, "/account", accountController.addAccount, http.MethodPost, "AddAccount")
	util.AddRoute(accountController.router, "/account", accountController.updateAccount, http.MethodPut, "UpdateAccount")
	util.AddRoute(accountController.router, "/account/{accountId:[0-9]+}", accountController.deleteAccount, http.MethodDelete, "MethodDelete")
	util.AddRoute(accountController.router, "/auth/login", accountController.login, http.MethodPost, "Login")
	util.AddRoute(accountController.router, "/auth/logout", accountController.logout, http.MethodDelete, "Logout")
	util.AddRoute(accountController.router, "/auth/verify", accountController.verify, http.MethodPost, "Verify")
	util.AddRoute(accountController.router, "/token/refresh", accountController.refresh, http.MethodPut, "RefreshToken")
	util.AddRoute(accountController.router, "/account/me", accountController.currentAccount, http.MethodGet, "MyAccount")
	return accountController
}

func (accountController *AccountController) getAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := accountController.accountService.GetAccounts()
	util.Write(w, r, accounts, http.StatusOK, err)
}
func (accountController *AccountController) getAccount(w http.ResponseWriter, r *http.Request) {
	accountDto := dto.AccountDTO{}
	err := accountController.accountService.GetAccount(mux.Vars(r), &accountDto)
	util.Write(w, r, accountDto, http.StatusOK, err)
}
func (accountController *AccountController) addAccount(w http.ResponseWriter, r *http.Request) {
	accountDto, accountOperation := dto.AccountDTO{}, dto.AccountOperation{}
	if !util.JsonDecode(w, r, &accountDto) {
		return
	}
	err := accountController.accountService.AddAccount(&accountDto, &accountOperation)
	util.Write(w, r, accountOperation, http.StatusCreated, err)
}
func (accountController *AccountController) deleteAccount(w http.ResponseWriter, r *http.Request) {
	accountOperation := dto.AccountOperation{}
	err := accountController.accountService.DeleteAccount(mux.Vars(r), &accountOperation)
	util.Write(w, r, accountOperation, http.StatusOK, err)
}
func (accountController *AccountController) updateAccount(w http.ResponseWriter, r *http.Request) {
	accountDto, accountOperation := dto.AccountDTO{}, dto.AccountOperation{}
	if !util.JsonDecode(w, r, &accountDto) {
		return
	}
	err := accountController.accountService.UpdateAccount(&accountDto, &accountOperation)
	util.Write(w, r, accountOperation, http.StatusOK, err)
}
func (accountController *AccountController) login(w http.ResponseWriter, r *http.Request) {
	loginRequest, loginResponse := dto.LoginRequest{}, dto.LoginResponse{}
	if !util.JsonDecode(w, r, &loginRequest) {
		return
	}
	err := accountController.sessionService.Login(&loginRequest, &loginResponse)
	util.Write(w, r, loginResponse, http.StatusOK, err)
}
func (accountController *AccountController) logout(w http.ResponseWriter, r *http.Request) {
	logoutResponse := dto.LogoutResponse{}
	err := accountController.sessionService.Logout(r.Header, &logoutResponse)
	util.Write(w, r, logoutResponse, http.StatusOK, err)
}
func (accountController *AccountController) refresh(w http.ResponseWriter, r *http.Request) {
	tokenRefreshRequest, tokenRefreshResponse := dto.TokenRefreshRequest{}, dto.TokenRefreshResponse{}
	if !util.JsonDecode(w, r, &tokenRefreshRequest) {
		return
	}
	err := accountController.sessionService.RefreshToken(&tokenRefreshRequest, &tokenRefreshResponse)
	util.Write(w, r, tokenRefreshResponse, http.StatusOK, err)
}
func (accountController *AccountController) verify(w http.ResponseWriter, r *http.Request) {
	tokenValidityRequest, tokenValidityResponse := dto.TokenValidityRequest{}, dto.TokenValidityResponse{}
	if !util.JsonDecode(w, r, &tokenValidityRequest) {
		return
	}
	err := accountController.sessionService.VerifyToken(&tokenValidityRequest, &tokenValidityResponse)
	util.Write(w, r, tokenValidityResponse, http.StatusOK, err)
}
func (accountController *AccountController) currentAccount(w http.ResponseWriter, r *http.Request) {
	accountDto := dto.AccountDTO{}
	err := accountController.sessionService.GetCurrentAccount(r.Header, &accountDto)
	util.Write(w, r, accountDto, http.StatusOK, err)
}
