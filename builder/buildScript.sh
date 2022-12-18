go build -o bin/main cmd/main.go
go build -o bin/auth auth_microservice/main.go
go build -o bin/mail mail_microservice/main.go
echo "successfully building"
