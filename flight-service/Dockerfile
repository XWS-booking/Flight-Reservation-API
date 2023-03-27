FROM golang:alpine
WORKDIR /flight-service
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app .
EXPOSE 8080
CMD ["./app"]
