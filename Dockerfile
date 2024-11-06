FROM golang:1.21 as builder
WORKDIR /app
COPY . .
RUN go build -o app ./cmd/app

FROM alpine:3.18
COPY --from=builder /app/app /app
EXPOSE 8080
CMD ["/app"]