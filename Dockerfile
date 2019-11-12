FROM golang:1.13 as first
WORKDIR /app
COPY . .
#RUN go get -d -v ./...
#RUN go install -v ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v ./cmd/application

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=first /app .
CMD ["./application"]