# syntax=docker/dockerfile:1

ARG PROD

FROM golang:1.19.5-alpine3.17

RUN if [ "$PROD" = "True" ]; then \
    apk add --no-cache \
        python3 \
        py3-pip \
    && pip3 install --upgrade pip \
    && pip3 install --no-cache-dir \
        awscli \
    && rm -rf /var/cache/apk/*; \
    fi

WORKDIR /app

COPY . .
RUN go mod download

RUN if [ "$PROD" = "True" ]; then \
    aws s3 cp s3://nonya-aws-bucket/discord-party-bot/config.yml ./internal/config/; \
    fi

RUN go build -o /discord-party-bot ./cmd

CMD [ "/discord-party-bot" ]