FROM --platform=linux/amd64 ubuntu:24.04
LABEL org.opencontainers.image.authors="support@halon.io"

ARG HALON_REPO_USER
ENV HALON_REPO_USER=$HALON_REPO_USER
ARG HALON_REPO_PASS
ENV HALON_REPO_PASS=$HALON_REPO_PASS

RUN apt-get update && apt-get install -y wget gnupg apt-transport-https

RUN wget -qO - https://raw.githubusercontent.com/halon/pgp-keys/master/pgp-keys/7F0A73B5.gpg | tee /usr/share/keyrings/halon-archive-keyring.gpg > /dev/null
RUN echo "deb [arch=`dpkg --print-architecture` signed-by=/usr/share/keyrings/halon-archive-keyring.gpg] https://repo.halon.io/ noble stable" >> /etc/apt/sources.list.d/halon.list
RUN echo "machine repo.halon.io login ${HALON_REPO_USER} password ${HALON_REPO_PASS}" >> /etc/apt/auth.conf
RUN apt-get update && apt-get install -y halon=6.7.1

RUN /usr/bin/install -d /var/run/halon
ENV LD_LIBRARY_PATH=$LD_LIBRARY_PATH

RUN apt-get update && apt-get install -y build-essential cmake  git curl

RUN mkdir /tmp/halon
RUN cd /tmp/halon \
    && curl -O -L https://go.dev/dl/go1.25.0.linux-amd64.tar.gz \
    && echo "2852af0cb20a13139b3448992e69b868e50ed0f8a1e5940ee1de9e19a123b613 go1.25.0.linux-amd64.tar.gz" | sha256sum -c \
    && tar -C /usr/local -xzf go1.25.0.linux-amd64.tar.gz

ENV PATH="/usr/local/go/bin/:/root/go/bin/:${PATH}"

RUN rm -rf /tmp/halon

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

RUN apt-get update && apt-get install -y protobuf-compiler

RUN git clone -b v1.74.2 --depth 1 https://github.com/grpc/grpc-go \
    && cd grpc-go/examples/helloworld \
    && protoc --go_out=. --go_opt=paths=source_relative \
            --go-grpc_out=. --go-grpc_opt=paths=source_relative \
            helloworld/helloworld.proto \
    && go build -o greeter_server greeter_server/main.go \
    && go build -o greeter_client greeter_client/main.go

COPY .devcontainer/entrypoint.sh /entrypoint.sh
RUN chmod 755 /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

RUN apt-get update && apt-get install -y supervisor
COPY .devcontainer/supervisord.conf /etc/supervisor/conf.d/supervisord.conf
CMD ["/usr/bin/supervisord"]