FROM golang:1.20-alpine as build

WORKDIR /app

COPY . .

# download depedencies
RUN go mod vendor

# build binary
RUN go build -o app main.go

EXPOSE 7070

CMD ["./app"]