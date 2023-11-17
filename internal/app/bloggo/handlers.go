package bloggo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/bkohler93/bootdev-blog-aggregator/internal/database"
	"github.com/bkohler93/bootdev-blog-aggregator/internal/helpers"
	"github.com/go-chi/chi/v5"
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
		helpers.RespondWithError(w, http.StatusInternalServerError, "error creating feed")
		return
	}

	feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    u.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "error creating feed follow from new feed")
	}

	f := databaseFeedToFeed(feed)
	ff := databaseFeedFollowToFeedFollow(feedFollow)

	resBody := struct {
		Feed       Feed       `json:"feed"`
		FeedFollow FeedFollow `json:"feed_follow"`
	}{
		Feed:       f,
		FeedFollow: ff,
	}

	helpers.RespondWithJSON(w, http.StatusCreated, resBody)
}

func (cfg *apiConfig) handlerGetAllFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetAllFeeds(r.Context())
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "could not retrieve feeds")
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, feeds)
}

func (cfg *apiConfig) handlerPostFeedFollow(w http.ResponseWriter, r *http.Request, u User) {
	parameters := struct {
		FeedID uuid.UUID `json:"feed_id"`
	}{}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := decoder.Decode(&parameters)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "invalid body, must provide `feed_id`")
		return
	}

	feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    u.ID,
		FeedID:    parameters.FeedID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "error creating feed follow")
		return
	}

	f := databaseFeedFollowToFeedFollow(feedFollow)

	helpers.RespondWithJSON(w, http.StatusCreated, f)
}

func (cfg *apiConfig) handleDeleteFeedFollow(w http.ResponseWriter, r *http.Request, u User) {
	feedFollowID := chi.URLParam(r, "feedFollowID")
	feedFollowUID := uuid.MustParse(feedFollowID)

	feedFollow, err := cfg.DB.GetFeedFollow(r.Context(), feedFollowUID)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "couldn't retrieve feed follow with follow_id")
		return
	}

	if feedFollow.UserID != u.ID {
		helpers.RespondWithError(w, http.StatusUnauthorized, "not authorized to delete that feed_follow")
		return
	}

	err = cfg.DB.DeleteFeedFollow(r.Context(), feedFollowUID)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "could not delete feed follow")
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, nil)
}

func (cfg *apiConfig) handleGetUserFeedFollows(w http.ResponseWriter, r *http.Request, u User) {
	feedFollows, err := cfg.DB.GetUserFeedFollows(r.Context(), u.ID)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "could not retrieve user's feed follows")
	}

	helpers.RespondWithJSON(w, http.StatusOK, feedFollows)
}

func (cfg *apiConfig) handleGetUserPosts(w http.ResponseWriter, r *http.Request, u User) {
	limit := r.URL.Query().Get("limit")
	l, err := strconv.Atoi(limit)
	if err != nil {
		l = 10
	}

	posts, err := cfg.DB.GetUserPosts(r.Context(), database.GetUserPostsParams{
		UserID: u.ID,
		Limit:  int32(l),
	})
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("could not retrieve posts", err))
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, posts)
}
