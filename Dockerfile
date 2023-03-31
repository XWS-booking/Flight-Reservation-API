FROM golang:alpine
WORKDIR /flight-service
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o flight-app .
EXPOSE 8000
CMD ["./flight-app"]
