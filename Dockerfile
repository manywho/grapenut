FROM golang:1.11-alpine AS build

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64

RUN apk add --no-cache git ca-certificates

ADD . /app
WORKDIR /app

RUN go test -timeout 30s -coverprofile=/tmp/go-code-coverage
RUN go build -ldflags '-w -extldflags "-static"' -o /grafana-alert-check .

FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /grafana-alert-check /grafana-alert-check

CMD [ "/grafana-alert-check" ]