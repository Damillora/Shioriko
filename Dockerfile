FROM golang:1.20-alpine AS build

WORKDIR /go/src/shioriko
COPY . .

RUN apk add vips-dev pkgconfig gcc musl-dev
RUN go get -d -v ./...
RUN go build -o /shioriko
RUN mkdir -p /web && cp -r web/static /web

FROM node:18-alpine AS node_build
WORKDIR /src
COPY . .
WORKDIR /src/web/app
RUN npm install -g pnpm && pnpm install && pnpm build

FROM alpine AS runtime

WORKDIR /app
RUN apk add vips
COPY --from=build /shioriko /app/
COPY --from=node_build /src/web/static/ /app/web/static/

ENTRYPOINT ["/app/shioriko"]
