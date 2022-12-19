FROM golang:alpine as builder
WORKDIR /backend
COPY . .
RUN apk add --update musl-dev libwebp-dev gcc
RUN go mod tidy
RUN go build -o main cmd/main.go
FROM alpine:3.8
# Needed for wkhtmltopdf
RUN apk add --no-cache \
 libstdc++ \
 libx11 \
 libxrender \
 libxext \
 libssl1.0 \
 libressl-dev \
 ca-certificates \
 fontconfig \
 freetype \
 ttf-dejavu \
 ttf-droid \
 ttf-freefont \
 ttf-liberation
WORKDIR /backend
COPY bin/ bin/
COPY .env .env
COPY configs/ configs/
RUN apk add --update musl-dev libwebp-dev gcc
CMD ["bin/main"]
