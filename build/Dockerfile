
#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN apk add --no-cache git
RUN go get -d -v ./...
RUN go install -v ./...

#final stage
FROM alpine:latest
COPY --from=builder /go/bin /app
ENV TOKEN api-token
ENV PREFIX !w
ENTRYPOINT ./app/bot
LABEL Name=dnd-discord-bot Version=0.0.1

