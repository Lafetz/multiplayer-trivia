# Trivia Multiplayer

A web-based application developed using Go and HTMX. Players can engage in real-time trivia competitions with other players.

## ☑ FE

- [x] User Authentication
- [x] Modify timer and number of questions
- [x] Choose categories
- [ ] CI/CD
- [ ] Monitoring
- [ ] Leaderboards
- [ ] Power-ups

## Demo

Check out the live demo [**Here**](https://showdown-trivia-game-1.onrender.com/home). note that since this demo is hosted on a free version, it may take some time to start up initially.

## Built with

- Go
- net/http
- gorilla/websocket
- Htmx
- templ
- Tailwind
- MongoDB

## Getting Started

### With Makefile

To run the application using the provided Makefile, you can follow these steps:

1. Ensure all prerequisites are met, including Go, Tailwind, Templ, and Air.

2. Make sure you have a `.env` file containing the necessary environment variables:

```sh
  DB_URL= Database URL.
  PORT= Port number.
  LOG_LEVEL=debug||info||warn||error
  ENV=dev||prod
```

### With Docker Compose

To run the application using Docker, you can follow these steps:

1. Navigate to the main directory.

2. Use the following command to build and start the containers:

   ```sh
   docker compose up --build
   ```
