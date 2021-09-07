# tweet-service

Upon changes into the application, build the binary using build.sh

Build the docker image

    docker build -t ccshekhar/tweet-service:latest .

Push the image to the docker hub

    docker push ccshekhar/tweet-service:latest

User docker run to successively execute the service (exposed on port 8412)

    docker run -p 8412:8412 ccshekhar/tweet-service:latest


# Using gitlab container registry

login to gitlab container registry

    docker login registry.gitlab.com

pull the image from dockerhub

    docker pull ccshekhar/tweet-service:latest

tag the image with gitlab reference

    docker tag ccshekhar/tweet-service:latest registry.gitlab.com/trial-cc1/codpeipes-ci

push the image to gitlab container registry

    docker push registry.gitlab.com/trial-cc1/codpeipes-ci


service provides crud operations for tweets

APIS:

Create Tweet
POST {{host}}/v0/tweets

Update a tweet
PUT {{host}}/v0/tweets/{{tweetId}}

Get a tweet
GET {{host}}/v0/tweets/{{tweetId}}

Get All tweets
GET {{host}}/v0/tweets

Delete a tweet
DELETE {{host}}/v0/tweets/{{tweetId}}
