# Start by building the application.
FROM golang:1.19-buster as build

WORKDIR /go/src/app

RUN groupadd --gid 20001 appgroup && \
    useradd --gid 20001 --uid 10001 appuser
COPY . .
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o /go/bin/app.bin cmd/main.go

FROM debian:buster-slim as bash

RUN apt-get update -y && \ 
    apt-get -yy -q install --no-install-recommends --no-install-suggests --fix-missing bash-static

# Now copy it into our base image.
FROM gcr.io/distroless/base-debian11

COPY --from=bash /bin/bash-static /bin/bash
COPY --from=build /etc/group /etc/group
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build --chown=appuser:appgroup /go/bin/app.bin /app.bin
USER appuser:appgroup

EXPOSE $APP_PORT

CMD ["/bin/bash", "-c", "/app.bin --port ${APP_PORT} --host ${APP_HOST} --dbUrl ${APP_DBURL}"]
