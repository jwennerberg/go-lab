FROM registry.access.redhat.com/ubi8/go-toolset:1.16.12 AS builder

RUN mkdir -p $HOME/go/src/hello
WORKDIR $HOME/go/src/hello
COPY *.go .

RUN go mod init && \
    go mod tidy && \
    go build -o hello

# -- App stage --

FROM registry.access.redhat.com/ubi8/ubi-minimal:8.5
MAINTAINER Johan Wennerberg <jwennerb@redhat.com>

WORKDIR /home
COPY --from=builder /opt/app-root/src/go/src/hello/hello .

USER 1001
EXPOSE 8080

CMD ["./hello"]
