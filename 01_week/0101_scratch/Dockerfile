# Start by building the application.
FROM golang:1.19-alpine as build

WORKDIR /go/src/app
COPY . .
RUN addgroup -S -g 20001 appgroup && \
    adduser -S -u 10001 -g 20001 appuser
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/app.bin cmd/main.go

FROM scratch

COPY --from=build /etc/group /etc/group
COPY --from=build /etc/passwd /etc/passwd
ADD --chown=appuser:appgroup /upload /upload
COPY --from=build --chown=appuser:appgroup --chmod=755 /go/bin/app.bin /go/bin/app.bin
USER appuser:appgroup
VOLUME /upload
EXPOSE 9999

ENTRYPOINT ["/go/bin/app.bin"]