# first stage
FROM golang:alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

# second stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/internal/pkg/config/* ./internal/pkg/config/
COPY --from=builder /app/internal/pkg/script/migration/* ./internal/pkg/script/migration/

CMD ["./main"]

EXPOSE 8080