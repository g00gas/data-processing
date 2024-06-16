FROM golang:1.22

WORKDIR /app

COPY . .
RUN ls
RUN go mod download && go mod verify

RUN CGO_ENABLED=0 GOOS=linux go build -o /backend

EXPOSE 8080

CMD ["/backend"]
