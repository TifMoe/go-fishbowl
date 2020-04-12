# Build Stage
FROM latest:1.13 AS build-stage

LABEL app="build-go-fishbowl"
LABEL REPO="https://github.com/tifmoe/go-fishbowl"

ENV PROJPATH=/go/src/github.com/tifmoe/go-fishbowl

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/tifmoe/go-fishbowl
WORKDIR /go/src/github.com/tifmoe/go-fishbowl

RUN make build-alpine

# Final Stage
FROM golang

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/tifmoe/go-fishbowl"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/go-fishbowl/bin

WORKDIR /opt/go-fishbowl/bin

COPY --from=build-stage /go/src/github.com/tifmoe/go-fishbowl/bin/go-fishbowl /opt/go-fishbowl/bin/
RUN chmod +x /opt/go-fishbowl/bin/go-fishbowl

# Create appuser
RUN adduser -D -g '' go-fishbowl
USER go-fishbowl

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/go-fishbowl/bin/go-fishbowl"]
