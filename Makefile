run:
	go run main.go

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/aplefileserver .

windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/aplefileserver.exe .
