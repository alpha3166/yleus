FROM golang:alpine

RUN apk update && \
    apk upgrade

ARG USERNAME=me
RUN addgroup -g 1000 $USERNAME && \
    adduser -u 1000 -G $USERNAME -D $USERNAME
RUN apk add sudo && \
    echo "$USERNAME ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/$USERNAME

RUN apk add git less

USER $USERNAME

RUN go install golang.org/x/tools/gopls@latest
