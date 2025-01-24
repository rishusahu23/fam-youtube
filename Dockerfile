# Stage 1: Build stage
FROM golang:1.23-alpine as builder

# Install build dependencies
RUN apk add --no-cache make bash

WORKDIR /go/src/github.com/rishusahu23/fam-youtube

# Copy the Go modules files
COPY go.mod go.sum ./

# Download Go modules dependencies
RUN go mod tidy
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o fam-youtube .

# Stage 2: Create the final image
FROM alpine:latest as stage-1

# Install bash, curl and other necessary packages in the runtime image
RUN apk add --no-cache bash make go curl

WORKDIR /go/src/github.com/rishusahu23/fam-youtube

# Copy the compiled binary from the builder image
COPY --from=builder /go/src/github.com/rishusahu23/fam-youtube/fam-youtube .

# Copy the Makefile
COPY --from=builder /go/src/github.com/rishusahu23/fam-youtube/Makefile .

# Copy db folder
COPY --from=builder /go/src/github.com/rishusahu23/fam-youtube /go/src/github.com/rishusahu23/fam-youtube

# Download and copy the wait-for-it script
RUN curl -sSL https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh -o /usr/local/bin/wait-for-it && \
    chmod +x /usr/local/bin/wait-for-it && \
    echo "wait-for-it script downloaded and made executable."

# Set the default command to run (wait for postgres, then start app)
CMD ["sh", "-c", "wait-for-it postgres:5432 wait-for-it elasticsearch:9200 wait-for-it elasticsearch:9300 -- make upgradeDB DB_NAME=youtube && ./fam-youtube"]
