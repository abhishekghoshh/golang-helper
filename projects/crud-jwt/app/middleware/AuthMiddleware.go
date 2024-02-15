package middleware

import (
	"net/http"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}
func (authMiddleWare *AuthMiddleware) Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		// currentRoute := mux.CurrentRoute(request)
		// pathVariables := mux.Vars(request)
		// authorization := request.Header.Get("Authorization")

		// fmt.Println("route name", currentRoute.GetName())
		// fmt.Println("pathVariables", pathVariables)
		// fmt.Println("authorization", authorization)

		next.ServeHTTP(response, request)
	})
}
