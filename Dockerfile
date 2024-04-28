FROM golang:1.22.2 AS build-stage
    WORKDIR /app

    COPY go.mod go.sum ./
    RUN go mod download

    COPY ./src ./src

    RUN CGO_ENABLED=0 GOOS=linux go build -o /api ./src/

FROM build-stage AS run-test-stage
    RUN go test -v ./...

FROM scratch as run-release-stage
    WORKDIR /app

    COPY --from=build-stage /api /api

    EXPOSE 8080

    CMD ["/api"]
