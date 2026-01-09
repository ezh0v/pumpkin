FROM golang:1.24-alpine3.21 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static" -s -w' -o /pumpkin main.go

FROM gcr.io/distroless/static-debian11
ARG APP_VERSION
ENV APP_VERSION=$APP_VERSION
COPY --from=build /pumpkin /pumpkin
USER nonroot:nonroot
WORKDIR /
CMD ["/pumpkin"]