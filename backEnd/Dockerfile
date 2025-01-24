
FROM golang

RUN mkdir -p /home/backEnd

WORKDIR /home/backEnd

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 8080

#RUN go build -o app .
#CMD ["./app"]

CMD ["go", "run", "./cmd/main/main.go"]
