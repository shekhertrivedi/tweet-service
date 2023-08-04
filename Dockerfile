ARG GO_BASE_IMAGE=golang:1.19-alpine
ARG RUNTIME_BASE_IMAGE=alpine

FROM ${GO_BASE_IMAGE} as build

WORKDIR /usr/src/app

# Copy the current directory contents into the container at /app
COPY go.mod go.sum ./

# Download any necessary dependencies
RUN go mod download

COPY . .

ARG VERSION
# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w -X 'github.com/shekhertrivedi/tweet-service/build.Date=$(date)' -X 'github.com/shekhertrivedi/tweet-service/build.Version=${VERSION}' " -o /go/bin/app

# Final stage
FROM ${RUNTIME_BASE_IMAGE} as api-runner

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /app

# Copy the compiled Go app from the build stage into the new image
COPY --from=build /go/bin/app .

# Set the default command to run when the container starts
ENTRYPOINT ["./app"]

# run unit test
FROM golang:1.20 AS unit-test
WORKDIR /usr/src/app
COPY . .
ENTRYPOINT [ "make", "test" ]