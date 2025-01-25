
#FROM golang AS build
FROM golang

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

#RUN go build -o api ./cmd/main/main.go

#FROM alpine:latest

#WORKDIR /root/

#COPY --from=build /app/api .

EXPOSE 8080

CMD ["go", "run", "./cmd/main/main.go"]
#CMD ["./api"]
