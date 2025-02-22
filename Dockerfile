FROM golang:1.23-alpine AS build

WORKDIR /go/src/shioriko
COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o /shioriko -ldflags '-extldflags "-static"' -tags timetzdata
RUN mkdir -p /web && cp -r web/static /web

FROM node:20-alpine AS node_build
WORKDIR /src
COPY . .
WORKDIR /src/web/app
RUN npm ci && npm run build

FROM scratch AS runtime

WORKDIR /app
COPY --from=build /shioriko /app/
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=node_build /src/web/static/ /app/web/static/

ENTRYPOINT ["/app/shioriko"]
