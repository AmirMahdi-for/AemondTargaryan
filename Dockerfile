FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .

# مرحله تولید
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]