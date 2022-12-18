FROM alpine
WORKDIR /backend
COPY bin/main bin/main
COPY scripts/runMain.sh scripts/runMain.sh
COPY .env .env
COPY configs/ configs/
RUN apk add --update musl-dev libwebp-dev gcc
RUN sed -i 's/\r//g' scripts/runMain.sh
CMD ["sh", "scripts/runMain.sh"]
