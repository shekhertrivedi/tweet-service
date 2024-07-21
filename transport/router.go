package transport

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	//"github.com/joho/godotenv"
	"github.com/shekhertrivedi/tweet-service/model"
	"github.com/shekhertrivedi/tweet-service/service"
)

var (
	tweetService service.TweetService
)

// InitApp initialize service and router
func InitApp() {

	os.Setenv("PORT", "8412")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("Error occured while reading port..!!")
	}

	tweetService = service.InitTweetService()

	router := mux.NewRouter()
	router.HandleFunc("/v0/tweets", CreateTweet).Methods("POST")
	router.HandleFunc("/v0/tweets/{id}", UpdateTweet).Methods("PUT")
	router.HandleFunc("/v0/tweets/{id}", DeleteTweet).Methods("DELETE")
	router.HandleFunc("/v0/tweets/{id}", GetTweet).Methods("GET")
	router.HandleFunc("/v0/tweets", GetAllTweets).Methods("GET")

	// Apply CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodPatch},
		AllowedHeaders: []string{"*"},
	})
	routerCORS := c.Handler(router)

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Println(fmt.Sprintf("Error occured while loading env file %v", err))
	// }

	log.Printf("Server is listening on %v ", port)

	if err := http.ListenAndServe(":"+port, routerCORS); err != nil {
		log.Panicf("Failed to start the application %v", err)
	}

}

// CreateTweet create tweet
func CreateTweet(w http.ResponseWriter, r *http.Request) {
	log.Printf("create tweet..!!")
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
	res, _ := json.MarshalIndent(tweet, "", "")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

// UpdateTweet update tweet
func UpdateTweet(w http.ResponseWriter, r *http.Request) {
	log.Printf("update tweet..!!")
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
	log.Printf("delete tweet..!!")
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
	log.Printf("get tweet..!!")
	m := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	tweet, err := tweetService.GetTweet(context.Background(), m["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(getErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}
	b, _ := json.MarshalIndent(tweet, "", "")
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
	b, _ := json.MarshalIndent(tweets, "", "")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func getErrorResponse(status int, message string) []byte {
	res := model.ErrorResponse{
		Code:    status,
		Message: message,
	}
	b, _ := json.MarshalIndent(res, "", "")
	return b
}
