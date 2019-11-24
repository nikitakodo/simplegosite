FROM golang:1.13 as app
WORKDIR /app
COPY ./cmd ./cmd
COPY ./configs ./configs
COPY ./internal ./internal
COPY ./migrations ./migrations
COPY ./web ./web
COPY ./go.mod ./go.mod
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v ./cmd/application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v ./cmd/migration

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=app /app/application application
COPY --from=app /app/migration migration

CMD ["./application"]