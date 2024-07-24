run:
	@echo "Loading environment variables from .env file"
	@set -o allexport; source ./load_env.sh; set +o allexport; \
	echo "Running Go application"; \
	go run ./cmd/web/main.go
tailwind:
	tailwindcss -i ./internal/web/static/css/input.css -o ./internal/web/static/css/styles.css --watch
templ:
	templ generate --watch
air:
	@echo "Loading environment variables from .env file"
	@set -o allexport; source ./load_env.sh; set +o allexport; \
	echo "Running air"; \
	air -c .air.toml
test:
	go test  ./... 

coverage:
	go test  -coverprofile=coverage.out ./... ;
	go tool cover -func=coverage.out
build:
	go build -o ./bin/web ./cmd/web

# go test $(go list ./... | grep -v /views/)  -coverprofile=coverage.out ./... ;
