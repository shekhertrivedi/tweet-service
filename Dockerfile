# base image
FROM alpine

# copy the binaries
COPY tweet-service .

# merely for documentation purpose (from image creator to runner)
EXPOSE 8412

# bootstart the service
ENTRYPOINT [ "./tweet-service" ]
