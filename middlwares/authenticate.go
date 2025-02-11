package middlewares

import (
	"context"
	"github.com/amirrezam75/kenopsiacommon/models"
	"github.com/amirrezam75/kenopsiacommon/services"
	"net/http"
	"strings"
)

type UserService interface {
	FindById(id string) (*models.User, error)
}

type Authenticate struct {
	userService UserService
}

func NewAuthenticateMiddleware(userService UserService) Authenticate {
	return Authenticate{userService: userService}
}

func (a Authenticate) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")

		if authorizationHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token := strings.Replace(authorizationHeader, "Bearer ", "", 1)

		claims, err := services.JsonWebTokenService{}.Parse(token)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if claims.Subject == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		user, err := a.userService.FindById(claims.Subject)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "authUser", user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
