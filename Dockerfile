FROM golang:alpine as builder
WORKDIR /backend
COPY . .
RUN go mod tidy
RUN go build -o main cmd/main.go
FROM alpine
WORKDIR /backend
COPY --from=builder /backend/main /backend/main
COPY --from=builder /backend/configs/config.yml /backend/configs/config.yml
CMD ["./main"]