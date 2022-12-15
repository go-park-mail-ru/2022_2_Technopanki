FROM alpine
WORKDIR /backend
COPY /bin/main /bin/main
RUN apk add --update musl-dev libwebp-dev gcc
CMD ["bin/main"]
