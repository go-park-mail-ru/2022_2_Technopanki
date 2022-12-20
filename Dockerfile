FROM alpine
WORKDIR /backend
COPY bin/ bin/
COPY .env .env
COPY configs/ configs/
RUN apk add --update musl-dev libwebp-dev gcc
CMD ["bin/main"]
