run: 
	@go mod vendor
	@echo "building binary..."
	@go build -o internal/bin/http internal/cmd/http/app.go
	@echo "running service" 
	@./internal/bin/http
