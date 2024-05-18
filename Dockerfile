FROM golang:1.22.3-alpine3.18 as builder

ENV APP_HOME /go/src/web

WORKDIR "${APP_HOME}"


COPY . .

RUN go mod download
RUN go mod verify
RUN go build -o ./bin/web ./cmd/web

FROM golang:1.22.3-alpine3.18 
ENV APP_HOME /go/src/web
RUN mkdir -p "${APP_HOME}"
WORKDIR "${APP_HOME}"
ARG PORT
ARG DB_URL
ARG HOST_URL

ENV PORT ${Port}
ENV DB_URL=${DB_URL}
ENV HOST_URL=${HOST_URL}
COPY --from=builder "$APP_HOME"/bin/web $APP_HOME

EXPOSE 8080

CMD ["./web"]