FROM golang:1.21 as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./app/main.go

FROM scratch

COPY --from=builder /app/main /main

EXPOSE 9000

CMD ["/main"]
