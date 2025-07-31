FROM golang:latest AS builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
# Constrói o executável estaticamente, otimizado para Linux
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch
WORKDIR /
COPY --from=builder /app/main .
EXPOSE 8080
ENTRYPOINT ["/main"]