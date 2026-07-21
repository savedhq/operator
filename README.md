# operator

Kubernetes operator for saved.sh — CRDs and controllers. Scaffolded with
[kubebuilder](https://book.kubebuilder.io) v4.13.1 (`go.kubebuilder.io/v4`).

**Status: scaffold only.** The project layout, CI and one placeholder CRD exist.
No reconcile logic has been written yet — `LocalWorkerReconciler.Reconcile` is the
generated stub.

## Scope

Today the operator is a **manager layer for `local-worker`**: it runs customer-side
backup workers in-cluster and hands each one its `config.yaml`. Because that file
carries a WorkOS org-scoped API key, it lives in a **Secret**, never a ConfigMap —
see `../pm/docs/arch/auth.md` and `.claude/rules/security.md`.

Later this grows to backing things up from Kubernetes directly (PVCs, directories,
files). Those get their **own API group** (e.g. `backup.saved.sh`) rather than being
bolted onto `worker.saved.sh` — the same split Kubernetes itself uses between `apps`
and `storage`.

Before adding anything, check `../pm/planning/planning.md` for the current milestone
and `../pm/docs/arch/infra.md` for how this fits the rest of the infrastructure.
Scope decisions belong in an `arch/` doc first.

## API

| Group / version            | Kind          | Purpose                                      |
|----------------------------|---------------|----------------------------------------------|
| `worker.saved.sh/v1alpha1` | `LocalWorker` | One customer-side backup worker + its config |

`LocalWorkerSpec` is a placeholder shape (`image`, `configSecretRef`) — the minimum
needed to express intent, not the final API. Sample:
`config/samples/worker_v1alpha1_localworker.yaml`.

## Development

```sh
make build      # manifests + generate + fmt + vet, then compile cmd/main.go
make test       # the above, plus the envtest suite
make lint       # golangci-lint (custom build, plugins from .custom-gcl.yml)
make run        # run the controller against your current kubeconfig
```

`make manifests` and `make generate` write into `config/` and `api/` — that output is
**committed**, and CI fails if it drifts from what the types produce.

Deploying to a cluster (`make install` / `make deploy` / `make undeploy`) follows the
stock kubebuilder flow; see the
[quick start](https://book.kubebuilder.io/quick-start.html).

## CI

- `.github/workflows/ci.yml` — `lint` and `test` on every push and PR, then `build`
  pushes a multi-arch image to `ghcr.io/savedhq/operator` on pushes to `main` and on
  `v*` tags (`:<short-sha>` + `:latest` on main, `:<tag>` on a tag).
- `.github/workflows/scan.yml` — gitleaks on push, PR and weekly.

E2E tests (`make test-e2e`) need a local Kind cluster and are **not** wired into CI
yet — there is no controller behavior to exercise.
