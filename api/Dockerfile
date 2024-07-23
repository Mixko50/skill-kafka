FROM golang:alpine AS builder

WORKDIR /go/src
COPY . .
RUN go build -o skill_api

FROM alpine:latest AS runner
COPY --from=builder /go/src/app .
ENTRYPOINT ["./skill_api"]