FROM golang:1.23
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o producer-service main.go
CMD [ "./producer-service" ]