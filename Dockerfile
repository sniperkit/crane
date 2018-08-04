####################################
## Builder - image arguments
####################################
ARG BUILDER_ALPINE_VERSION=${BUILDER_ALPINE_VERSION:-"3.7"}
ARG BUILDER_GOLANG_VERSION=${BUILDER_GOLANG_VERSION:-"1.10.3"}
ARG BUILDER_IMAGE_TAG=${BUILDER_IMAGE_TAG:-"${BUILDER_GOLANG_VERSION}-alpine${BUILDER_ALPINE_VERSION}"}
ARG BUILDER_IMAGE_NAME=${BUILDER_IMAGE_NAME:-"golang:${BUILDER_IMAGE_TAG}"}

####################################
## Builder
###################################
FROM ${BUILDER_IMAGE_NAME} AS builder
# FROM sniperkit/golang:dev-1.10.3-alpine3.7 AS builder

ARG REPO_VCS=${REPO_VCS:-"github.com"}
ARG REPO_NAMESPACE=${REPO_NAMESPACE:-"sniperkit"}
ARG REPO_PROJECT=${REPO_PROJECT:-"crane"}
ARG REPO_URI=${REPO_URI:-"${REPO_VCS}/${REPO_NAMESPACE}/${REPO_PROJECT}"}

ARG GOLANG_TOOLS_URIS=${GOLANG_TOOLS_URIS:-"github.com/mitchellh/gox \
                                            github.com/Masterminds/glide \
                                            github.com/golang/dep/cmd/dep \
                                            github.com/mattn/gom"}

WORKDIR /go/src/${REPO_URI}

## copy deps definitions
COPY Gopkg.lock Gopkg.toml ./
# COPY glide.lock glide.yaml ./

## install deps
# RUN clear && glide update --strip-vendor
# RUN clear && dep ensure -v
# COPY vendor vendor

## code
COPY pkg pkg
COPY plugin plugin

## binaries
COPY cmd/crane cmd/crane
# COPY cmd/${REPO_PROJECT} ${REPO_PROJECT}
# COPY cmd/${REPO_PROJECT}-plus ${REPO_PROJECT}-plus

## install commands
RUN for gtu in ${GOLANG_TOOLS_URIS}; do echo "[in progress]... go get -u $gtu" ; go get -u $gtu ; go install $gtu ; done \
    \
    && \
    if [ -f "glide.lock" ]; then \
        glide update --strip-vendor; \
    \
    elif [ -f "Gopkg.lock" ]; then \
        dep ensure -v; \
    fi \
    \
    && go install ./... \
    \
    && rm -fR $GOPATH/src \
    && rm -fR $GOPATH/pkg \
    \
    && ls -l $GOPATH/bin

############################################################################################################
############################################################################################################

####################################
## Builder - image arguments
####################################
ARG RUNNER_ALPINE_VERSION=${RUNNER_ALPINE_VERSION:-"3.7"}
ARG RUNNER_IMAGE_NAME=${RUNNER_IMAGE_NAME:-"alpine:${RUNNER_ALPINE_VERSION}"}

####################################
## Build
####################################
FROM alpine:3.7 AS dist

WORKDIR /usr/bin
COPY --from=builder /go/bin ./

RUN echo "\n---- DEBUG INFO -----\n" \
    ls -l /usr/bin/crane* \
    echo "\nPATH: ${PATH}\n"
