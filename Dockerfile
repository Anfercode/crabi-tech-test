FROM golang:1.21.3 AS builder

WORKDIR /code

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY src ./src
COPY app ./app

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o crabi_tech_test ./app/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /code/crabi_tech_test .

EXPOSE 8080

CMD ["./crabi_tech_test"]