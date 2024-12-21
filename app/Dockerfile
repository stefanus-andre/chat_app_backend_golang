FROM golang:1.23
WORKDIR /app

# Copy go.mod and go.sum files first to leverage Docker layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application files
COPY . .

# Run the build for the main package
RUN go build -o main .

# Command to run the built executable
CMD [ "/app/main" ]
