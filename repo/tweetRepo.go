package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/shekhertrivedi/tweet-service/model"
)

// TweetRepo repository
type TweetRepo interface {
	Create(context.Context, *model.Tweet) (*model.Tweet, error)
	Update(context.Context, *model.Tweet) error
	Delete(context.Context, string) error
	Get(context.Context, string) (*model.Tweet, error)
	GetAll(ctx context.Context) ([]*model.Tweet, error)
}

// InitTweetRepo initialize tweet repo
func InitTweetRepo() TweetRepo {
	repo := make(map[string]*model.Tweet)
	return &tweetRepoImpl{
		Repo: repo,
	}
}

type tweetRepoImpl struct {
	Repo map[string]*model.Tweet
}

func (tr *tweetRepoImpl) Create(ctx context.Context, tweet *model.Tweet) (*model.Tweet, error) {
	if len(tweet.ID) == 0 {
		tweet.ID = uuid.New().String()
	}
	tr.Repo[tweet.ID] = tweet
	return tweet, nil
}

func (tr *tweetRepoImpl) Update(ctx context.Context, tweet *model.Tweet) error {
	tr.Repo[tweet.ID] = tweet
	return nil
}

func (tr *tweetRepoImpl) Delete(ctx context.Context, tweetID string) error {
	delete(tr.Repo, tweetID)
	return nil
}

func (tr *tweetRepoImpl) Get(ctx context.Context, tweetID string) (*model.Tweet, error) {
	return tr.Repo[tweetID], nil
}

func (tr *tweetRepoImpl) GetAll(ctx context.Context) ([]*model.Tweet, error) {
	resp := make([]*model.Tweet, 0)
	for _, v := range tr.Repo {
		resp = append(resp, v)
	}
	return resp, nil
}
