FROM golang:1.26.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./cmd/server

FROM scratch

COPY --from=builder /server /bin/server
COPY --from=builder /app/.config.example.yml /bin/.config.yml

WORKDIR /bin

ENTRYPOINT ["/bin/server", "--config=/bin/.config.yml"]