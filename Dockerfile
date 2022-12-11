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
<<<<<<< HEAD
COPY bin/ bin/
COPY .env .env
COPY configs/ configs/
=======
COPY --from=icalialabs/wkhtmltopdf:alpine /bin/wkhtmltopdf /bin/wkhtmltopdf
COPY --from=builder /backend/main /backend/main
COPY --from=builder /backend/configs/config.yml /backend/configs/config.yml
COPY --from=builder /backend/data/image /backend/data/image
COPY --from=builder /backend/static /backend/static
COPY --from=builder /backend/.env /backend/.env
>>>>>>> 8dc5c9e (added resume to pdf convert)
RUN apk add --update musl-dev libwebp-dev gcc
CMD ["bin/main"]
