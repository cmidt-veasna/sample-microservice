FROM alpine
ADD service /
ENTRYPOINT /service
EXPOSE 8080