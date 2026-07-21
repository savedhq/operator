# operator

Kubernetes operator for saved.sh — CRDs and controllers.

**Status: empty.** Nothing has been built here yet and the scope is still to be
decided. This repo currently holds only a README and CI.

Before adding anything, check `pm/planning/planning.md` for the current milestone
and its scope, and `pm/docs/arch/infra.md` for how this is meant to fit with the
rest of the infrastructure. Scope decisions belong in an `arch/` doc first.

## CI

`.github/workflows/scan.yml` runs gitleaks on push, PR and weekly. There is no
build yet — add one alongside the first code.
