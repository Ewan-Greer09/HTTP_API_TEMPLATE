FROM golang:1.18.3-stretch AS base

ENV CGO_ENABLED=1
ENV GO111MODULE on
ENV GOOS=linux
ENV GOARCH=arm64

# System dependencies
RUN apt-get update \
    && apt-get install -y build-essential ca-certificates git \
    && update-ca-certificates

### Development with nodemon and debugger
FROM base AS dev

WORKDIR /app

COPY go.* ./

# Dependencies
RUN go mod download \
    && go mod verify

RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Install nodemon
ENV PATH /opt/communications/node_modules/.bin:$PATH

RUN curl -fsSL https://deb.nodesource.com/setup_16.x | bash -
RUN apt-get install -y nodejs npm
RUN npm install nodemon -g --loglevel notice

EXPOSE 8080
EXPOSE 2345

# as we grow services multiple nodemon files??
CMD [ "nodemon", "--config", "nodemon.json" ]