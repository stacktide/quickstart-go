# Any copyright is dedicated to the Public Domain.
# https://creativecommons.org/publicdomain/zero/1.0/

FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/greetd

FROM scratch

WORKDIR /

COPY --from=builder /app/greetd /greetd

USER 1001:1001

EXPOSE 8080

ENTRYPOINT ["/greetd"]
