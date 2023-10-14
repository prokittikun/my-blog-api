# ใช้รูปภาพของ Go 1.20 หรือเวอร์ชันล่าสุดที่มีอยู่
FROM golang:1.20

# สร้างไดเร็กทอรีของแอพพลิเคชันใน Docker และตั้งค่าที่อยู่ที่ใช้ในโปรเจค
WORKDIR /app

COPY . .

# Expose port 8080 to the outside world
EXPOSE 8080
# Build the Go app
RUN go build -o main .
# รันแอพพลิเคชัน
CMD ["./main"]