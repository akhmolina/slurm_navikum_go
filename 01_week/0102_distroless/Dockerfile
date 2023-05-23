# Start by building the application.
FROM golang:1.19-alpine as build

WORKDIR /go/src/app
RUN addgroup -S -g 20001 appgroup && \
    adduser -S -u 10001 -g 20001 appuser
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/app.bin cmd/main.go

# Now copy it into our base image.
FROM gcr.io/distroless/base-debian11

COPY --from=build /etc/group /etc/group
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build --chown=appuser:appgroup /go/bin/app.bin /app.bin
# USER appuser:appgroup

ENV APP_PORT=9000
ENV APP_HOST=1.0.0.0
ENV APP_DBURL=postgres://user:pass@localhost:5432/app
ENV GOFLAGS="-port $APP_PORT -host $APP_HOST -dbUrl $APP_DBURL"
EXPOSE $APP_PORT

CMD ["/app.bin"]
