FROM golang:alpine
ARG USERNAME=me
RUN addgroup -g 1000 $USERNAME && \
    adduser -u 1000 -G $USERNAME -D $USERNAME
RUN apk --update-cache add sudo && \
    echo "$USERNAME ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/$USERNAME
RUN apk add git less
USER $USERNAME
RUN go install golang.org/x/tools/gopls@latest
