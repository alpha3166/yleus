FROM alpine
RUN addgroup -g 1000 me && \
    adduser -u 1000 -G me -D me && \
    addgroup me wheel
RUN apk -U add sudo && \
    echo "%wheel ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/wheel
RUN apk -U add git less
RUN apk -U add go
RUN su - me -c "go install golang.org/x/tools/gopls@latest"
