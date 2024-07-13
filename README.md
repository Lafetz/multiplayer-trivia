# Trivia Multiplayer

A web-based application developed using Go and HTMX. Players can engage in real-time trivia competitions with other players.players have the ability to choose categories, modify game settings including timer duration and the number of questions.

## Demo

Check out the live demo [**Here**](https://showdown-trivia-game-1.onrender.com/home). note that since this demo is hosted on a free version, it may take some time to start up initially.

## Built with

- Go
- Templ
- Htmx
- Tailwind
- MongoDB

## Getting Started

### With Makefile

To run the application using the provided Makefile, you can follow these steps:

1. Ensure all prerequisites are met, including Go, Tailwind, Templ, and Air.

2. Make sure you have a `.env` file containing the necessary environment variables:

```sh
  WS_HOST=WebSocket host address. e.g connect:wss://exmple.com or connect:ws://localhost:8080
  DB_URL= Database URL.
  PORT= Port number.
  LOG_LEVEL=info||warn
  ENV=dev||prod
```

### With Docker

To run the application using Docker, you can follow these steps:

1. Build the Docker image:

   ```sh
   docker build -t trivia-multiplayer .
   ```

2. Run the Docker container, either supplying the .env file or using environment variables:

   Supplying .env file:

   ```sh
      docker run --env-file .env -p 8080:8080 trivia-multiplayer
   ```

   Using environment variables:

   ```sh
      docker run -e WS_HOST="wss://example.com" -e DB_URL="your_database_url" -e PORT="8080" -p 8080:8080 trivia-multiplayer
   ```
