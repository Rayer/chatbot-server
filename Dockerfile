FROM golang:alpine as build-env
WORKDIR /ChatbotAPIs
ADD . /ChatbotAPIs
RUN cd /ChatbotAPIs/server && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server

FROM scratch
WORKDIR /app
COPY --from=build-env /ChatbotAPIs/server/server /app/
CMD /bin/sh
EXPOSE 8080
ENTRYPOINT /app/server
