FROM golang:1.15-alpine as builder

# Add Maintainer Info
LABEL maintainer="Sam Zhou <sam@mixmedia.com>"

# Set the Current Working Directory inside the container
WORKDIR /app/es-autoclean

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go version \
 && export GO111MODULE=on \
 && go mod vendor \
 && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o es-autoclean

######## Start a new stage from scratch #######
FROM alpine:latest  

RUN wget -O /usr/local/bin/dumb-init https://github.com/Yelp/dumb-init/releases/download/v1.2.2/dumb-init_1.2.2_amd64 \
 && chmod +x /usr/local/bin/dumb-init \
 && apk add --update libintl \
 && apk add --virtual build_deps gettext \
 && apk add --no-cache tzdata \
 && cp /usr/bin/envsubst /usr/local/bin/envsubst \
 && apk del build_deps

WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/es-autoclean/es-autoclean .
COPY --from=builder /app/es-autoclean/config.json .


ENV TZ=Asia/Hong_Kong \
 SERVICE_NAME=es-autoclean \
 ES_ENDPOINT="http://192.168.33.6:9200" \
 INDEX_PREFIX="filebeat-7.3.0-" \
 KEEP_DAY=15 \
 CRON_SPEC="0 9 * * * ?"

ENTRYPOINT ["dumb-init", "--"]

CMD envsubst < /app/config.json > /app/temp.json \
 && /app/es-autoclean -c /app/temp.json
