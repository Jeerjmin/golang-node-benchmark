FROM golang:latest AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o gin-app ./cmd/main.go


FROM alpine:latest
ENV GOMAXPROCS=1
RUN apk --no-cache add libc6-compat
WORKDIR /root/
COPY --from=builder /app/gin-app .
RUN ls -la ./gin-app
EXPOSE 8080

CMD ["./gin-app"]