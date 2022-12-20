FROM alpine:3.8
WORKDIR /backend
COPY bin/ bin/
COPY ./static/html /usr/share/html
COPY .env .env
COPY configs/ configs/
COPY --from=icalialabs/wkhtmltopdf:alpine /bin/wkhtmltopdf /bin/wkhtmltopdf
RUN apk add --update musl-dev libwebp-dev gcc
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
CMD ["bin/main"]
