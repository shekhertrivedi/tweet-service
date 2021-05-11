# base image
FROM alpine
COPY ./build/. .
EXPOSE 8412
ENTRYPOINT [ "./tweet-service" ]
