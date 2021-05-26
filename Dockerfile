FROM golang:1.16 as dev

WORKDIR /go/gocoinapi

COPY . .

RUN chmod -R 755 ./scripts/wait.sh

RUN go get ./...

EXPOSE 8080

HEALTHCHECK --interval=1m --timeout=30s --start-period=1m --retries=3 \
  CMD curl -f http://localhost:8080/general/healthcheck || exit 1

CMD ["go", "run", "/go/gocoinapi/main.go"]




FROM golang:1.16 as build

WORKDIR /go/gocoinapi

COPY . .

COPY --from=dev /go/gocoinapi .

RUN chmod -R 755 ./scripts/wait.sh

RUN CGO_ENABLED=0 GOOS=linux go build -a .




FROM golang:1.16-alpine as prod

COPY --from=build /go/gocoinapi /gocoinapi

EXPOSE 8080

HEALTHCHECK --interval=1m --timeout=30s --start-period=1m --retries=3 \
  CMD curl -f http://localhost:8080/general/healthcheck || exit 1

ENTRYPOINT ["/gocoinapi"]