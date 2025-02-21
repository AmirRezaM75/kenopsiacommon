package middlewares

import (
	"github.com/amirrezam75/kenopsiacommon/services"
	"github.com/amirrezam75/kenopsiauser"
	"net/http"
	"strings"
)

type UserRepository interface {
	FindById(id string) (kenopsiauser.User, error)
}

type Authenticate struct {
	userRepository UserRepository
}

func NewAuthenticateMiddleware(userRepository UserRepository) Authenticate {
	return Authenticate{userRepository: userRepository}
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

		user, err := a.userRepository.FindById(claims.Subject)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := services.ContextService{}.WithUser(r.Context(), &user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
