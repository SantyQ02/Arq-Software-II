FROM golang:alpine

WORKDIR /api
COPY .env ./.env
COPY . .
# COPY utils/initializers/test.env.docker ./utils/initializers/test.env

RUN go mod tidy
RUN go build -o /docker-go

CMD ["/docker-go"]