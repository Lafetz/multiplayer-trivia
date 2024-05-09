run:
	go run ./cmd/main.go
tailwind:
	tailwindcss -i ./internal/web/static/css/input.css -o ./internal/web/static/css/styles.css --watch
templ:
	templ generate --watch
air:
	air -c .air.toml
test:
	go test ./...
coverage:
	go test  -coverprofile=coverage.out ./... ;
	go tool cover -func=coverage.out


