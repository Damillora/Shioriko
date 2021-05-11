FROM golang:1.16-alpine AS build

WORKDIR /go/src/shioriko
COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o /shioriko
RUN mkdir -p /web && cp -r web/static web/template /web

FROM node:14-alpine AS node_build
WORKDIR /src
COPY . .
WORKDIR /src/web/app
RUN yarn install && yarn build

FROM scratch AS runtime

WORKDIR /app
COPY --from=build /shioriko /app/
COPY --from=node_build /src/web/static/ /app/web/static/
COPY --from=node_build /src/web/template/ /app/web/template/

ENTRYPOINT ["/app/shioriko"]
