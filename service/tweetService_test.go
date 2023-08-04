package service

import (
	"context"
	"testing"

	"github.com/shekhertrivedi/tweet-service/model"
	"github.com/shekhertrivedi/tweet-service/repo/mocks"
	"github.com/stretchr/testify/mock"
)

func Test_tweetServiceImpl_CreateTweet(t *testing.T) {

	type args struct {
		ctx   context.Context
		tweet *model.Tweet
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				tweet: &model.Tweet{
					ID:      "test",
					Message: "test message",
				},
			},
		},
		{
			name: "success without ID",
			args: args{
				ctx: context.Background(),
				tweet: &model.Tweet{
					Message: "test message",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repoMocks := &mocks.TweetRepo{}
			repoMocks.On("Create", mock.Anything, mock.Anything).Return(nil, nil)
			ts := &tweetServiceImpl{
				tr: repoMocks,
			}
			_, err := ts.CreateTweet(tt.args.ctx, tt.args.tweet)
			if (err != nil) != tt.wantErr {
				t.Errorf("tweetServiceImpl.CreateTweet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_tweetServiceImpl_UpdateTweet(t *testing.T) {
	type args struct {
		ctx   context.Context
		tweet *model.Tweet
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				tweet: &model.Tweet{
					ID:      "test",
					Message: "test message",
				},
			},
		},
		{
			name: "ID not available",
			args: args{
				ctx: context.Background(),
				tweet: &model.Tweet{
					Message: "test message",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repoMocks := &mocks.TweetRepo{}
			repoMocks.On("Update", mock.Anything, mock.Anything).Return(nil)
			ts := &tweetServiceImpl{
				tr: repoMocks,
			}
			if err := ts.UpdateTweet(tt.args.ctx, tt.args.tweet); (err != nil) != tt.wantErr {
				t.Errorf("tweetServiceImpl.UpdateTweet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_tweetServiceImpl_DeleteTweet(t *testing.T) {

	type args struct {
		ctx     context.Context
		tweetID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "tweetID not present",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				ctx:     context.Background(),
				tweetID: "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repoMocks := &mocks.TweetRepo{}
			repoMocks.On("Delete", mock.Anything, mock.Anything).Return(nil)
			ts := &tweetServiceImpl{
				tr: repoMocks,
			}
			if err := ts.DeleteTweet(tt.args.ctx, tt.args.tweetID); (err != nil) != tt.wantErr {
				t.Errorf("tweetServiceImpl.DeleteTweet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_tweetServiceImpl_GetTweet(t *testing.T) {

	type args struct {
		ctx     context.Context
		tweetID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "tweetID not present",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				ctx:     context.Background(),
				tweetID: "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repoMocks := &mocks.TweetRepo{}
			repoMocks.On("Get", mock.Anything, mock.Anything).Return(nil, nil)
			ts := &tweetServiceImpl{
				tr: repoMocks,
			}
			_, err := ts.GetTweet(tt.args.ctx, tt.args.tweetID)
			if (err != nil) != tt.wantErr {
				t.Errorf("tweetServiceImpl.GetTweet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
