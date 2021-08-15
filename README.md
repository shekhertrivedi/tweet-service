# tweet-service

Upon changes into the application, build the binary using build.sh

Build the docker image

    docker build ccshekhar/tweet-service:latest .

Push the image to the docker hub

    docker push ccshekhar/tweet-service:latest

User docker run to successively execute the service (exposed on port 8412)


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
