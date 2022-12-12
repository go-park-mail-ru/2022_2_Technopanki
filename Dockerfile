FROM alpine
WORKDIR /backend
COPY . .
RUN apk add --update musl-dev libwebp-dev gcc
CMD ["bin/main"]
