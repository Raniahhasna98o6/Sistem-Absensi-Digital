# Build Stage
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .

# Download dependensi dan build aplikasi
RUN go mod tidy
RUN go build -o main .

# Run Stage (Menggunakan image ultra-ringan)
FROM alpine:latest

WORKDIR /app

# Copy hasil build dari stage sebelumnya
COPY --from=builder /app/main .

# Set default env untuk Azure/Docker (Port dinamis)
ENV PORT=8080

# Expose port (sebagai dokumentasi)
EXPOSE 8080

# Jalankan server
CMD ["./main"]
