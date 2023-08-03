ARG GO_VERSION=1.20.7
ARG BASE_DEBIAN_DISTRO="bullseye"
ARG GOLANG_IMAGE="golang:${GO_VERSION}-${BASE_DEBIAN_DISTRO}"

FROM --platform=$BUILDPLATFORM ${GOLANG_IMAGE} AS base
FROM base AS build

RUN mkdir -p /build
WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

ENV CGO_ENABLED=0
RUN go build -trimpath -ldflags '-s' -o /build/bin/ ./cmd/servantd

FROM scratch

LABEL servant.type="controller"

COPY --from=build /build/bin/* /usr/local/bin/

EXPOSE 8080
HEALTHCHECK --interval=60s --timeout=10s --start-period=5s --start-interval=5s --retries=3 \
  CMD ["servantd", "healthz"]
ENTRYPOINT ["servantd"]
