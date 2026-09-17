# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image

FROM golang:1.27 AS builder

# Add Maintainer Info
LABEL maintainer="tizips <tizips@163.com>"

# Enable Go modules proxy for mainland China
ENV GOPROXY=https://goproxy.cn,direct

# Set the Current Working Directory inside the container
WORKDIR /build

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

WORKDIR /build/server/admin

# Build the Go app
RUN CGO_ENABLED=0 go build -o application .

# Start a new stage from scratch
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# Copy the Pre-build binary file from the previous stage
COPY --from=builder /build/server/admin/application .
COPY --from=builder /build/server/admin/conf ./conf
COPY --from=builder /build/server/admin/migration ./migration

# Expose port 8080 to the outside world
EXPOSE 9600

# Command to run the executable
CMD ["/app/application", "server"]
