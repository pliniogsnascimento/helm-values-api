FROM golang:1.18.4-alpine as base

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .

FROM bitnami/kubectl:1.20.9 as kubectl

COPY --from=base /app/main /app/main

ENTRYPOINT ["/app/main"]
