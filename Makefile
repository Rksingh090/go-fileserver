run:
	go run main.go

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/kmfileserver .

windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/kmfileserver.exe .

stop: linux
	sudo systemctl stop kmfileserver.service 

copy: stop
	sudo cp bin/kmfileserver /usr/local/bin/kmfileserver

perm: copy
	sudo chmod +x /usr/local/bin/kmfileserver

deploy:	perm
	sudo systemctl restart kmfileserver.service 