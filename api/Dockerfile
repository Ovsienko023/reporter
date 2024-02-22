FROM golang:1.18-bullseye AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -o reporter main.go

FROM alpine
LABEL authors="Ovsienko V."
RUN apk add git make libc6-compat
WORKDIR /app
COPY --from=builder /app/reporter /app/reporter
