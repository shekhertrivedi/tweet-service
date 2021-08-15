# base image
FROM alpine
COPY tweet-service .
EXPOSE 8412
ENTRYPOINT [ "./tweet-service" ]
