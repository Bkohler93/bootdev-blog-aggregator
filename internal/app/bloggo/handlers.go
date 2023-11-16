package bloggo

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/bkohler93/bootdev-blog-aggregator/internal/database"
	"github.com/bkohler93/bootdev-blog-aggregator/internal/helpers"
	"github.com/google/uuid"
)

func (cfg apiConfig) handlerPostUsers(w http.ResponseWriter, r *http.Request) {
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

func handlerGetErr(w http.ResponseWriter, r *http.Request) {
	helpers.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}

func handlerGetReadiness(w http.ResponseWriter, r *http.Request) {
	res := map[string]interface{}{
		"status": "ok",
	}
	helpers.RespondWithJSON(w, http.StatusOK, res)
}

func (cfg apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, u User) {
	helpers.RespondWithJSON(w, http.StatusOK, u)
}

func (cfg *apiConfig) handlerPostFeed(w http.ResponseWriter, r *http.Request, u User) {

	parameters := struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}{}

	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&parameters)

	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		Url:       parameters.Url,
		Name:      parameters.Name,
		UserID:    u.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		log.Println("error creating feed", err)
		helpers.RespondWithError(w, http.StatusInternalServerError, "error creating feed")
		return
	}

	f := databaseFeedToFeed(feed)

	helpers.RespondWithJSON(w, http.StatusCreated, f)
}

func (cfg *apiConfig) handlerGetAllFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetAllFeeds(r.Context())
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "could not retrieve feeds")
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, feeds)
}
