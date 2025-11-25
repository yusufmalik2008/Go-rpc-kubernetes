FROM golang:1.21-alpine as builder

WORKDIR /app
COPY . .
RUN go build -o rpc-server main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/rpc-server .
EXPOSE 1234
CMD ["./rpc-server"]
