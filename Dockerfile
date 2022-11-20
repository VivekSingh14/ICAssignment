FROM golang:alpine as builder


WORKDIR /app

COPY ./ /app
COPY go.mod go.sum ./
RUN go mod download

RUN go build -o bin/main cmd/main.go 


FROM alpine:3.14
COPY config/config-env.yaml config/

COPY --from=builder /app/bin/main .

CMD ["./main"]