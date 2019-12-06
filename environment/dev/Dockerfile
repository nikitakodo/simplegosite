FROM golang:1.13 as app
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v ./cmd/application && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v ./cmd/migration

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=app /app/application application
COPY --from=app /app/migration migration

CMD ["./application"]