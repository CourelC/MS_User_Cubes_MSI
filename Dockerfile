# syntax=docker/dockerfile:1

FROM golang:1.18

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /dist

# This container exposes port 8080 to the outside world
#EXPOSE $ANALYTICS_MS_PORT
EXPOSE 3000

# Run the executable
CMD [ "/dist" ]
