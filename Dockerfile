FROM debian:stable-slim as base
# Define User and Group
ARG USER=zdeploy
ARG GROUP=zdeploy
ARG UID=2000
ARG GID=2000
ARG HOME=/home/zdeploy

RUN apt-get install tzdata
RUN apt-get update \
  && apt-get upgrade -y \
  && apt-get install -y ca-certificates \
  && apt-get install -y telnet \
  && apt-get install -y curl \
  && apt-get install -y redis-tools

ENV TZ=Asia/Ho_Chi_Minh
# Add User group and install deps
RUN groupadd -g $GID $GROUP && \
useradd -d $HOME -u $UID -s /bin/false --gid $GID $USER && \
mkdir -p $HOME && \
chown -R $GROUP:$USER $HOME && \
    echo $TZ > /etc/timezone
WORKDIR $HOME
USER $USER

FROM golang:1.17 as builder
ENV GO111MODULE=on
ARG CACHE_DIR=/tmp
# Working directory
WORKDIR /app
# Copy files
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

# Build app
RUN --mount=type=cache,id=cache-go,target=$CACHE_DIR/.cache CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o spotify ./cmd/main.go
CMD chmod +x /app/spotify

# final stage
FROM base
# Copy binary from builder
WORKDIR /app

COPY --from=builder /app/spotify /app/spotify
COPY --from=builder /app/config /app/config

# List expose port(s)
EXPOSE 8080
# Run server command
ENTRYPOINT ["./spotify"]
