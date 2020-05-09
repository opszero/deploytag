FROM golang:latest

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update -y && \
    apt-get install -y apt-transport-https curl gnupg gnupg2 lsb-release software-properties-common git ca-certificates build-essential postgresql-client python3-pip jq netcat gettext

# Install kubectl
RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl && \
    chmod +x ./kubectl && \
    mv ./kubectl /usr/local/bin/kubectl

# Install Helm
ENV HELM_VERSION=v3.2.1
RUN curl https://raw.githubusercontent.com/kubernetes/helm/master/scripts/get > get_helm.sh
RUN chmod 700 get_helm.sh
RUN ./get_helm.sh --version $HELM_VERSION

# Install Docker
RUN curl -fsSL https://download.docker.com/linux/debian/gpg | apt-key add -
RUN apt-key fingerprint 0EBFCD88
RUN add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/debian $(lsb_release -cs) stable"
RUN apt-get update -y
RUN apt-get install docker-ce -y

# Install awscli
RUN pip3 install --upgrade awscli

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

RUN cp /usr/local/bin/kubectl /bin/kubectl
COPY . .
RUN go build -o /bin/deploytag

# TODO Remove the src dependencies

ENTRYPOINT /bin/deploytag