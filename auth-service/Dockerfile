FROM golang:alpine
WORKDIR /auth-service
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o auth-app .
EXPOSE 8080
CMD ["./auth-app"]
