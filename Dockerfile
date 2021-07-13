# FROM golang:1.16-alpine3.13 AS builder
FROM golang:1.16-alpine3.13

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY web/ ./web/

EXPOSE 5000

COPY shawty.go .

# RUN go build -ldflags="-w -s" -o /app/shawty /app/shawty.go

# FROM SCRATCH

# COPY --from=builder /app/shawty /app/shawty

# CMD ["/app/shawty"]

CMD ["go", "run", "shawty.go"]
