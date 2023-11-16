package bloggo

import (
	"net/http"
	"strings"

	"github.com/bkohler93/bootdev-blog-aggregator/internal/helpers"
)

type authedHandler func(http.ResponseWriter, *http.Request, User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if hasKey := strings.Contains(authHeader, "ApiKey"); !hasKey {
			helpers.RespondWithError(w, http.StatusUnauthorized, "No Api Key in request headers")
			return
		}

		apiKey := strings.Trim(authHeader, "ApiKey ")

		user, err := cfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			helpers.RespondWithError(w, http.StatusNotFound, "User with that Api Key cannot be found")
			return
		}

		u := databaseUserToUser(user)

		handler(w, r, u)
	}
}
