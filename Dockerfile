FROM golang:1.16-alpine AS build

WORKDIR /go/src/shioriko
COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o /shioriko
RUN mkdir -p /web && cp -r web/static web/template /web

FROM scratch AS runtime
COPY --from=build /shioriko /
COPY --from=build /web /

ENTRYPOINT ["/shioriko"]
