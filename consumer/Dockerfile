FROM golang:alpine AS builder

WORKDIR /go/src
COPY . .
RUN go build -o app

FROM alpine:latest AS runner
COPY --from=builder /go/src/app .
ENTRYPOINT ["./app"]