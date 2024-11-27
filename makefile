build:
	docker build -t task-manager-server .
run:
	docker run -d -p 8083:8083 task-manager-server
test:
	 go test ./repository ./handler -v
lint:
	 golangci-lint run

