FROM golang:alpine as builder
WORKDIR /backend
COPY . .
RUN apk add --update musl-dev libwebp-dev gcc
RUN go mod tidy
RUN go build -o main cmd/main.go
FROM alpine
WORKDIR /backend
COPY --from=builder /backend/main /backend/main
COPY --from=builder /backend/configs/config.yml /backend/configs/config.yml
RUN apk add --update musl-dev libwebp-dev gcc
CMD ["./main"]
