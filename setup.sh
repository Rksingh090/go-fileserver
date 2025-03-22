CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/kmfileserver .
sudo systemctl stop kmfileserver.service 
sudo cp bin/kmfileserver /usr/local/bin/kmfileserver
sudo chmod +x /usr/local/bin/kmfileserver
sudo systemctl restart kmfileserver.service 
