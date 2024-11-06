FROM golang:1.21 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/app

FROM alpine:3.18
COPY --from=builder /app/app /app/app
COPY --from=builder /app/configs/config.yaml /app/configs/config.yaml
EXPOSE 8080
CMD ["/app/app"]