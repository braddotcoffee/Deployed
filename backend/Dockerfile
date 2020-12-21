FROM golang:1.14

RUN mkdir /app
ADD . /app

WORKDIR /app

RUN go build -o main .

ENV PORT=3000
EXPOSE 3000

CMD ["/app/main"]