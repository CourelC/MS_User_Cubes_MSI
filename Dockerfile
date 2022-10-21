# syntax=docker/dockerfile:1

FROM golang:1.18

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /dist

EXPOSE 3000

#Exec le prog
 #RUN go run main.go

# Run the executable
CMD [ "/dist" ]

