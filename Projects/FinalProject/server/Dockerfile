FROM golang:1.16-alpine
WORKDIR /app
COPY movies.go /app
COPY movieDatabase.json /app
COPY logs.txt /app
COPY go.mod /app
RUN go build -o server
EXPOSE 8080
CMD ["./server"]
