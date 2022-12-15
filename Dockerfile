FROM golang:alpine

RUN apk update && apk add --no-cache git php python

WORKDIR /app

COPY . .

RUN go mod tidy

RUN GIN_MODE=release go build -o binary

ENTRYPOINT ["/app/binary"]