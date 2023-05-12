FROM golang:1.19
COPY . /app
WORKDIR /app
RUN go build -o app
CMD ["/app/app"]