FROM golang:alpine AS builder

WORKDIR /go/src
COPY . .
RUN go build -o skill_consumer

FROM alpine:latest AS runner
COPY --from=builder /go/src/skill_consumer .
ENTRYPOINT ["./skill_consumer"]