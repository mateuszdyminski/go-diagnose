FROM golang:1.13.4

# Delve Configuration
RUN go get -u -v github.com/go-delve/delve/cmd/dlv
EXPOSE 40000

# App Configuration
WORKDIR /go/src/github.com/mateuszdyminski/go-diagnose
ADD main.go .
RUN go build -gcflags="all=-N -l" -o app .
EXPOSE 8081

CMD [ "dlv", "exec", "/go/src/github.com/mateuszdyminski/go-diagnose/app", "--listen=:40000", "--headless=true", "--api-version=2", "--log" ]