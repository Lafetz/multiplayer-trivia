run:
	go run ./cmd/main.go
tailwind:
	tailwindcss -i ./internal/web/static/css/input.css -o ./internal/web/static/css/styles.css --watch