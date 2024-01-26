FROM golang:1.21

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./

# Install app dependencies
RUN go mod download

COPY . .

# Check if /app/logs folder exists before creating
RUN if [ ! -d "/app/logs" ]; then mkdir /app/logs; fi

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /main

EXPOSE 8000

# Run
CMD ["/main"]
