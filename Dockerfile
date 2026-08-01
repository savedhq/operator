# syntax=docker/dockerfile:1
# Packaging only: CI compiles the binary first and drops it in
# dist/manager-<arch>. Rationale: pm/docs/man/go-cicd.md.
FROM gcr.io/distroless/static-debian12:nonroot

# BuildKit sets TARGETARCH from --platform, so this stays arch-agnostic even
# though CI builds arm64 only. The build context preserves the executable bit
# (it was Actions *artifacts* that dropped it), so no `chmod +x` is needed.
ARG TARGETARCH
COPY dist/manager-${TARGETARCH} /manager

USER 65532:65532
ENTRYPOINT ["/manager"]
