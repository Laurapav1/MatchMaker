# Start med en Go-baseret base image
FROM golang:1.20 AS builder

# Sæt arbejdsbiblioteket
WORKDIR /app

# Kopier go mod og sum filer
COPY go.mod go.sum ./

# Download afhængigheder
RUN go mod download

# Kopier resten af koden
COPY . .

# Byg applikationen
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Kør en mindre runtime image
FROM gcr.io/distroless/base-debian11

WORKDIR /root/

# Kopier den kompilerede binær fra build-stadiet
COPY --from=builder /app/main .

# Åbn porten din applikation kører på
EXPOSE 8080

# Definer startkommando
CMD ["./main"]