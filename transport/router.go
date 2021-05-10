package transport

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shekhertrivedi/tweet-service/model"
	"github.com/shekhertrivedi/tweet-service/service"
)

var (
	tweetService service.TweetService
)

// InitApp initialize service and router
func InitApp() {

	tweetService = service.InitTweetService()

	router := mux.NewRouter()
	router.HandleFunc("/v0/tweets", CreateTweet).Methods("POST")
	router.HandleFunc("/v0/tweets/{id}", UpdateTweet).Methods("PUT")
	router.HandleFunc("/v0/tweets/{id}", DeleteTweet).Methods("DELETE")
	router.HandleFunc("/v0/tweets/{id}", GetTweet).Methods("GET")
	router.HandleFunc("/v0/tweets", GetAllTweets).Methods("GET")

	log.Println("Server is listening on port 8412...!!!")

	http.ListenAndServe(":8412", router)

}

// CreateTweet create tweet
func CreateTweet(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(getErrorResponse(http.StatusBadRequest, "invalid create request"))
		return
	}
	tweet := &model.Tweet{}
	json.Unmarshal(b, tweet)
	tweet, err = tweetService.CreateTweet(context.Background(), tweet)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(getErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}
	res, _ := json.Marshal(tweet)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

// UpdateTweet update tweet
func UpdateTweet(w http.ResponseWriter, r *http.Request) {
	m := mux.Vars(r)
	b, err := ioutil.ReadAll(r.Body)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(getErrorResponse(http.StatusBadRequest, "invalid update request"))
		return
	}
	tweet := &model.Tweet{}
	json.Unmarshal(b, tweet)
	tweet.ID = m["id"]
	err = tweetService.UpdateTweet(context.Background(), tweet)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(getErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// DeleteTweet delete tweet
func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	m := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	err := tweetService.DeleteTweet(context.Background(), m["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(getErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

// GetTweet get tweet
func GetTweet(w http.ResponseWriter, r *http.Request) {
	m := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	tweet, err := tweetService.GetTweet(context.Background(), m["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(getErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}
	b, _ := json.Marshal(tweet)
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

// GetAllTweets get all tweets
func GetAllTweets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tweets, err := tweetService.GetAllTweets(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(getErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}
	b, _ := json.Marshal(tweets)
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func getErrorResponse(status int, message string) []byte {
	res := model.ErrorResponse{
		Code:    status,
		Message: message,
	}
	b, _ := json.Marshal(res)
	return b
}
