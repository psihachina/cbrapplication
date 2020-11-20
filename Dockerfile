FROM golang:latest

RUN mkdir /app
COPY ./ /app
WORKDIR /app
RUN ls
RUN go build -o main ./cmd/soapserver/
EXPOSE 8080
CMD ["./main"]