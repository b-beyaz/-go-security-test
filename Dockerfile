FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN mkdir -p /out \
    && CGO_ENABLED=0 GOOS=linux go build -o /out/app main.go

FROM alpine:3.24.1

WORKDIR /app

RUN addgroup -S app && adduser -S app -G app

COPY --from=builder /out/app ./app

RUN chown -R app:app /app

USER app:app

EXPOSE 8080

CMD ["./app"]
