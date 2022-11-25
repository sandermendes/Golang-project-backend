# Start from golang base image
FROM golang:alpine as builder

# Add info
LABEL maintainer="Sander Mendes <sandermendes@gmail.com>"

# Install git for dependencies
# RUN apk update && apk add --no-cache git

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/main/main.go

# Start the App
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary
COPY --from=builder /app/main .
COPY --from=builder /app/config.json .

#Expose port 8080
EXPOSE 8080

# Run app
CMD ["./main"]