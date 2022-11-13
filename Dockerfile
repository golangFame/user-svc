FROM golang:1.19-alpine
COPY . /opt/user-svc
WORKDIR /opt/user-svc


RUN #go get
RUN go build -o /go/bin/user-svc ./cmd/...
#COPY server.crt /opt/user-svc
RUN rm -rf /opt/user-svc

ENTRYPOINT ["/go/bin/user-svc"]
