# pulling a lightweight version of golang
FROM golang:1.11-alpine
RUN apk --update add --no-cache git

ENV GOPATH /go
# Copy the local package files to the container's workspace.

ADD . /go
WORKDIR /go

# copy crontabs for root user
RUN echo "$WORKDIR"
COPY $WORKDIR/jobs/cronjobs /etc/crontabs/root

RUN git config --global url."git://".insteadOf https://

RUN go build ./jobs/probs.go

RUN chmod +x ./probs

# start crond with log level 8 in foreground, output to stderr
CMD ["crond", "-f", "-d", "8"]