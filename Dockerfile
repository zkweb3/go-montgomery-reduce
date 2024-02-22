FROM ubuntu:22.04
ARG GO_VERSION=1.22.0
WORKDIR /bin
RUN apt update && apt install -y wget make build-essential m4

RUN wget https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz \
	&& rm -rf go${GO_VERSION}.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin

ADD setup.sh .
RUN chmod +x setup.sh
ENTRYPOINT [ "./setup.sh" ]