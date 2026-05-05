FROM golang:alpine

WORKDIR /app

# Copy semua file ke dalam folder /app di kontainer
COPY . .

# Download semua library (termasuk Gin dan MySQL driver)
RUN go mod download

# Build aplikasinya dan PAKSA namanya jadi "main"
RUN go build -o main .

# Buka pintu port 8080
EXPOSE 8080

# Jalankan file "main" yang udah dibikin tadi
CMD ["./main"]