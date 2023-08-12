FROM golang:1.20-alpine3.17 as builder

WORKDIR /src

COPY . .

RUN go mod download

RUN go build -o app

FROM alpine:3.17

COPY --from=builder /src/app .

CMD ["/app"]
