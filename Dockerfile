FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o inventory .

RUN chmod +x inventory 

CMD ["./inventory"]
