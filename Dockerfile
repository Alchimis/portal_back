# syntax=docker/dockerfile:1

# Stage 1: Build
FROM golang:1.20 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen
RUN go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen
COPY . .
RUN oapi-codegen -generate gorilla,types authentication/api/frontend/frontendapi.yaml > authentication/api/frontend/frontendapi.gen.go
RUN oapi-codegen -generate gorilla,types company/api/frontend/frontendapi.yaml > company/api/frontend/frontendapi.gen.go
RUN oapi-codegen -generate gorilla,types documentation/api/frontend/frontendapi.yaml > documentation/api/frontend/frontendapi.gen.go
RUN CGO_ENABLED=0 go build -o app -buildvcs=false main.go

# Stage 2: Runtime
FROM alpine
WORKDIR /app
COPY --from=builder /app/app ./
COPY --from=builder /app/authentication/db/migrations ./authentication/db/migrations
CMD ["./app"]
