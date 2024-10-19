# Build the application from source
FROM golang:1.23.1-alpine AS build-stage
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o /main

# Run the tests in the container
FROM build-stage AS test-stage
RUN go test -v ./... -cover

# Deploy the application binary into a lean image
FROM golang:1.23.1-alpine AS run-stage
WORKDIR /app
COPY --from=build-stage /main /main
COPY .env .env
EXPOSE 8080
ENTRYPOINT ["/main"]