# Web client
FROM node:20-alpine AS node_build
WORKDIR /src
COPY . .
WORKDIR /src/pkg/web
RUN npm ci && npm run build

# Go application
FROM golang:1.23-alpine AS build
WORKDIR /go/src/shioriko
COPY . .
COPY --from=node_build /src/pkg/web/build/ /go/src/shioriko/pkg/web/build/
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o /shioriko -ldflags '-extldflags "-static"' -tags timetzdata

FROM scratch AS runtime
WORKDIR /app
COPY --from=build /shioriko /app/
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/app/shioriko"]
