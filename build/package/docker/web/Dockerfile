# pulling a lightweight version of golang
FROM golang:1.11-alpine
RUN apk --update add --no-cache git

ENV GOPATH /go
# Copy the local package files to the container's workspace.
ADD . /go/src/wesite-health-analyzer
WORKDIR /go/src/wesite-health-analyzer

RUN git config --global url."git://".insteadOf https://

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go get  github.com/gorilla/mux && \
  go get github.com/gorilla/handlers && \
  go get gopkg.in/mgo.v2 && \
  go get gopkg.in/mgo.v2/bson

RUN go build ./src/main.go

RUN chmod +x ./main

# Run the command by default when the container starts.
ENTRYPOINT ["./main"]

# Document that the service listens on port 9000.
EXPOSE 9000