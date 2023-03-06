FROM golang:1.18.4-alpine as base

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN apk add build-base && GIN_MODE=release go build -o main .

FROM alpine as final

COPY --from=base /app/main /app/main
RUN apk add build-base

ENTRYPOINT ["GIN_MODE=release", "/app/main"]
