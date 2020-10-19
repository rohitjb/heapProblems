# Container image that runs your code
FROM alpine:3.10

# Add tools
RUN apk add --no-cache bash curl gzip tar

# Add adr tools
RUN \
    curl -L https://github.com/npryce/adr-tools/archive/3.0.0.tar.gz -o adr-tools.tar.gz && \
    tar zxvf adr-tools.tar.gz && \
    mv adr-tools-3.0.0/src/ /adr-tools

# Copies your code file from your action repository to the filesystem path `/` of the container
COPY entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh

# Code file to execute when the docker container starts up (`entrypoint.sh`)
ENTRYPOINT ["/entrypoint.sh"]
