# Support setting various labels on the final image
ARG COMMIT=""
ARG VERSION=""
ARG BUILDNUM=""

# Build Biz in a stock Go builder container
FROM golang:1.16-buster as builder

ADD . /go-ethereum
RUN cd /go-ethereum && go run build/ci.go install ./cmd/biz

# Pull Biz into a second stage deploy alpine container
FROM ubuntu:18.04

COPY --from=builder /go-ethereum/build/bin/biz /usr/local/bin/

EXPOSE 8545 8546 30303 30303/udp
ENTRYPOINT ["biz"]

# Add some metadata labels to help programatic image consumption
ARG COMMIT=""
ARG VERSION=""
ARG BUILDNUM=""

LABEL commit="$COMMIT" version="$VERSION" buildnum="$BUILDNUM"
