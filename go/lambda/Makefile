build:
	rm .deploy.zip || true
	GOOS=linux GOARCH=amd64 go build -o bootstrap
	zip -9 .deploy.zip bootstrap
