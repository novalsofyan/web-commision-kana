# Build stage
FROM golang:1.25-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api/*.go

# Final stage
FROM alpine:latest

# Install ca-certificates untuk HTTPS & tzdata untuk timezone
RUN apk --no-cache add ca-certificates tzdata

# Set timezone ke Jakarta
ENV TZ=Asia/Jakarta

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
