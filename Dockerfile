# =============================================================================
# Stage: base
#
# Build the base image with the necessary tools so that it can be used in the 
# build stage without having to install them again. 
# =============================================================================
FROM golang:alpine AS base

WORKDIR /app

RUN apk update && apk upgrade && apk add --no-cache ca-certificates \
    && update-ca-certificates

# =============================================================================
# Stage: build
#
# Download project dependencies and build the project.
# =============================================================================
FROM base AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy the source code into the container
COPY . .

# Build the application 
RUN CGO_ENABLED=0 go build -v -o driftwood ./cmd/driftwood.go


# =============================================================================
# Stage: release
#
# Copy the built binary and the necessary certificates to a scratch image to
# reduce the image size.
# =============================================================================
FROM scratch AS release


# Copy the certificates and user/group files from the builder stage
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/driftwood /driftwood

# Set the default Lua directory to `/lua`, allowing volume mounting
VOLUME /lua

# Run the bot
CMD ["/driftwood"]