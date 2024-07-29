# Start med en Go-baseret Debian image for at bygge applikationen
FROM golang:1.20-buster AS builder

# Sæt arbejdsbiblioteket
WORKDIR /app

# Kopier go mod og sum filer
COPY go.mod go.sum ./

# Download afhængigheder
RUN go mod download

# Kopier resten af koden
COPY . .

# Byg applikationen
RUN go build -o main .

# Kør en mindre runtime image
FROM debian:bullseye-slim

WORKDIR /root/

# Kopier den kompilerede binær fra build-stadiet
COPY --from=builder /app/main .

# Åbn porten din applikation kører på
EXPOSE 8080

# Definer startkommando
CMD ["./main"]
