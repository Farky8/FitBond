
FROM golang

RUN mkdir -p /home/app

COPY . /home/app

#RUN go build -o app .
#CMD ["./app"]

CMD ["go", "run", "."]
