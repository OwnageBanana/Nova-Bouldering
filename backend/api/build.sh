go mod tidy
go build ./cmd/main.go
sudo mv main /var/www/nova-bouldering-api
sudo systemctl restart nova-bouldering.service
