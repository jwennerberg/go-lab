FROM registry.access.redhat.com/ubi8/ubi-minimal:latest
MAINTAINER Johan Wennerberg <jwennerb@redhat.com>
COPY hello /hello
EXPOSE 8080
ENTRYPOINT ["/hello"]
