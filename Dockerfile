# syntax=docker/dockerfile:1
# Multi-stage: buildx cross-compiles the binary itself, so CI needs no `compile`
# job and no artifact round-trip. That coupling is what broke every Go image on
# 2026-07-30: the Actions artifact storage quota filled, the upload step failed,
# `compile` failed with it, so `build` never ran and nothing reached GHCR.
#
# --platform=$BUILDPLATFORM pins this stage to the runner's NATIVE arch and lets
# GOARCH do the cross-compiling, so the toolchain never runs under QEMU. That is
# exactly what this repo's earlier "~40min for a single build" note was avoiding,
# a build stage *without* the pin. The pin is load-bearing, do not remove it.
#
# The builder tag must track go.mod's `go` directive: CI sets GOTOOLCHAIN=local,
# so a builder older than the directive fails rather than fetching a toolchain.
FROM --platform=$BUILDPLATFORM golang:1.26.5-alpine AS build
WORKDIR /src

# Dependencies first: this layer is reused whenever go.mod/go.sum are unchanged.
COPY go.mod go.sum ./
RUN go mod download

COPY . .
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} \
    go build -trimpath -ldflags="-s -w" -o /out/manager cmd/main.go

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /out/manager /manager
USER 65532:65532
ENTRYPOINT ["/manager"]
