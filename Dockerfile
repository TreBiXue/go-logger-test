FROM golang:1.19

ENV MODE prod

COPY . /app
WORKDIR /app
RUN go build -o app
CMD ["/app/app"]