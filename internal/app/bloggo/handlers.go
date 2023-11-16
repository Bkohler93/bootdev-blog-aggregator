package bloggo

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/bkohler93/bootdev-blog-aggregator/internal/database"
	"github.com/bkohler93/bootdev-blog-aggregator/internal/helpers"
	"github.com/google/uuid"
)

func (cfg apiConfig) postUsers(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)

	parameters := struct {
		Name string `json:"name"`
	}{}
	decoder.Decode(&parameters)

	name := parameters.Name
	id := uuid.New()
	createdAt := time.Now()
	updatedAt := time.Now()

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        id,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Name:      name,
	})
	if err != nil {
		log.Println("failed to create user in database", err)
	}

	resUser := databaseUserToUser(user)

	helpers.RespondWithJSON(w, http.StatusCreated, resUser)
}

func getReadiness(w http.ResponseWriter, r *http.Request) {
	res := map[string]interface{}{
		"status": "ok",
	}
	helpers.RespondWithJSON(w, http.StatusOK, res)
}

func (cfg apiConfig) getUser(w http.ResponseWriter, r *http.Request) {
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

	resUser := databaseUserToUser(user)

	helpers.RespondWithJSON(w, http.StatusOK, resUser)
}
