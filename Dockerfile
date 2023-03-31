FROM golang:alpine
WORKDIR /flight-service
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV PORT 8000
ENV DATABASE_NAME XWS
ENV JWT_SECRET secret
ENV DATABASE_CONNECTION_STRING mongodb+srv://xws:lozinka123@cluster0.yqfgbck.mongodb.net/?retryWrites=true&w=majority
RUN go build -o flight-app .
EXPOSE 8000
CMD ["./flight-app"]
