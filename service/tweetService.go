package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/shekhertrivedi/tweet-service/model"
	"github.com/shekhertrivedi/tweet-service/repo"
)

// TweetService interface
type TweetService interface {
	CreateTweet(context.Context, *model.Tweet) (*model.Tweet, error)
	UpdateTweet(context.Context, *model.Tweet) error
	DeleteTweet(context.Context, string) error
	GetTweet(context.Context, string) (*model.Tweet, error)
	GetAllTweets(context.Context) ([]*model.Tweet, error)
}

// InitTweetService initialize tweet service
func InitTweetService() TweetService {
	repo := repo.InitTweetRepo()
	return &tweetServiceImpl{
		tr: repo,
	}
}

type tweetServiceImpl struct {
	tr repo.TweetRepo
}

func (ts *tweetServiceImpl) CreateTweet(ctx context.Context, tweet *model.Tweet) (*model.Tweet, error) {
	if len(tweet.ID) == 0 {
		tweet.ID = uuid.New().String()
	}
	return ts.tr.Create(ctx, tweet)
}

func (ts *tweetServiceImpl) UpdateTweet(ctx context.Context, tweet *model.Tweet) error {
	if len(tweet.ID) == 0 {
		return fmt.Errorf("invalid ID %v", tweet)
	}
	return ts.tr.Update(ctx, tweet)
}

func (ts *tweetServiceImpl) DeleteTweet(ctx context.Context, tweetID string) error {
	if len(tweetID) == 0 {
		return fmt.Errorf("invalid ID %v", tweetID)
	}
	return ts.tr.Delete(ctx, tweetID)
}

func (ts *tweetServiceImpl) GetTweet(ctx context.Context, tweetID string) (*model.Tweet, error) {
	if len(tweetID) == 0 {
		return nil, fmt.Errorf("invalid ID %v", tweetID)
	}
	return ts.tr.Get(ctx, tweetID)
}

func (ts *tweetServiceImpl) GetAllTweets(ctx context.Context) ([]*model.Tweet, error) {
	return ts.tr.GetAll(ctx)
}
