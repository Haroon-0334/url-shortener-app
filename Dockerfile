# build stage
FROM golang:1.21-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o /bin/url-shortener ./app/cmd/server

# runtime stage
FROM gcr.io/distroless/static
COPY --from=build /bin/url-shortener /bin/url-shortener
EXPOSE 8080
USER 65532:65532
ENTRYPOINT ["/bin/url-shortener"]