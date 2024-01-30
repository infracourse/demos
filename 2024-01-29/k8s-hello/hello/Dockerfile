FROM --platform=arm64 golang:1.21-alpine AS builder

WORKDIR /src

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o hello

#------------------------------------------------------------

FROM --platform=arm64 alpine:latest

WORKDIR /app

COPY --from=builder /src/hello .
COPY --from=builder /src/templates ./templates

CMD ["./hello"]
