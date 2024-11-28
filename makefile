build:
	docker build -t task-manager-server .
run:
	docker run -d -p 8083:8083 task-manager-server
test:
	 go test ./repository ./api/handler -v
lint:
	 golangci-lint run

