FROM ubuntu:22.04
ARG GO_VERSION=1.22.0
WORKDIR /build
RUN apt update && apt install -y wget make build-essential m4

RUN wget https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz \
	&& rm -rf go${GO_VERSION}.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin
ADD montgomery.go .
ADD montgomery_test.go .
ADD go.mod .
ADD setup.sh .
ADD test.sh .
RUN chmod +x setup.sh && chmod +x test.sh && ./setup.sh

# ENTRYPOINT [ "tail", "-f", "/dev/null" ]