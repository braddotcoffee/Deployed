FROM golang:1.14

RUN mkdir /root/.ssh
ADD ./secrets/id_rsa /root/.ssh/id_rsa

RUN mkdir /app
ADD . /app

WORKDIR /app

RUN go build -o main .

ENV PORT=3000
ENV PRODUCTION=true
EXPOSE 3000

CMD ["/app/main"]