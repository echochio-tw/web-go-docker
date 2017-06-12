FROM golang:latest AS build-env
COPY web.go /go/web.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go get github.com/ajstarks/svgo
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o web

# final stage
FROM centurylink/ca-certs
COPY --from=build-env /go/web /
ENTRYPOINT ["/web"]
