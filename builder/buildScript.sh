echo "start building..."
go build -o /backend/bin/main /backend/cmd/main.go
go build -o /backend/bin/auth /backend/auth_microservice/main.go
go build -o /backend/bin/mail /backend/mail_microservice/main.go
echo "successfully building"
