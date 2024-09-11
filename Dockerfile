FROM golang:1.23-alpine as builder
RUN apk update && apk add --no-cache git

WORKDIR /tmp/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Unit tests
RUN CGO_ENABLED=0 go test -v

# Build the Go app
RUN go build -o ./out/app .

FROM golang:1.23-alpine
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

EXPOSE 8080

CMD ["./main"]