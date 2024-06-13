# ref: https://github.com/GoogleContainerTools/distroless/blob/main/examples/go/Dockerfile

FROM golang:1.22.3 as build

WORKDIR /go/src/app
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o /go/bin/app ./cmd/api-server

FROM gcr.io/distroless/static-debian12

COPY --from=build /go/bin/app /

EXPOSE 8888

CMD ["/app"]