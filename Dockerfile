FROM gcr.io/distroless/static-debian12:nonroot
ARG TARGETARCH
COPY dist/manager-${TARGETARCH} /manager
USER 65532:65532
ENTRYPOINT ["/manager"]
