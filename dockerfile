FROM golang:1.23.5-alpine  

ENV WORKDIR=/app \
    GOCACHE=/tmp/.cache

WORKDIR ${WORKDIR}

COPY . .

RUN apk add --no-cache build-base

RUN go mod tidy

RUN go install github.com/cosmtrek/air@v1.49.0

CMD [ "air" ]