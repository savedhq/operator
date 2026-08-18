# operator

Kubernetes operator for saved.sh, CRDs and controllers. Scaffolded with
[kubebuilder](https://book.kubebuilder.io) v4.13.1 (`go.kubebuilder.io/v4`).

**Status: scaffold only.** The project layout, CI and one placeholder CRD exist.
No reconcile logic has been written yet, `LocalWorkerReconciler.Reconcile` is the
generated stub.

## Scope

Today the operator is a **manager layer for `local-worker`**: it runs customer-side
backup workers in-cluster and hands each one its `config.yaml`. Because that file
carries a WorkOS org-scoped API key, it lives in a **Secret**, never a ConfigMap,
see [`auth#machine-authentication-workos-org-scoped`](../pm/arch/auth.md).

Later this grows to backing things up from Kubernetes directly (PVCs, directories,
files). Those get their **own API group**, `backup.saved.sh`, rather than being bolted
onto `worker.saved.sh`, the same split Kubernetes itself uses between `apps` and
`storage`.

**The design lives in [`kubernetes.md`](../pm/arch/kubernetes.md), read it first.** In
particular: Kubernetes deploys and declares, it never schedules and never executes a
backup, and a worker Deployment is pinned to one replica because a worker id is a task
queue.

Before adding anything, check `../pm/planning/m1.md` for the current milestone.
Scope decisions belong in [`kubernetes.md`](../pm/arch/kubernetes.md) first.
`../pm/arch/infra.md` is **our own** cluster, which is a different subject.

## API

| Group / version            | Kind          | Purpose                                      |
|----------------------------|---------------|----------------------------------------------|
| `worker.saved.sh/v1alpha1` | `LocalWorker` | One customer-side backup worker + its config |

`LocalWorkerSpec` is a placeholder shape (`image`, `configSecretRef`), the minimum
needed to express intent, not the final API. Sample:
`config/samples/worker_v1alpha1_localworker.yaml`.

## Development

```sh
make build          # manifests + generate + fmt + vet, then compile cmd/main.go
make test           # the above, plus the envtest suite
make lint           # golangci-lint (custom build, plugins from .custom-gcl.yml)
make run            # run the controller against your current kubeconfig
make docker-build   # cross-compile into dist/, then package the image
```

The image (`ghcr.io/savedhq/operator`) runs the controller itself, it is not the
`local-worker` image, which is built in its own repo and referenced by
`LocalWorker.spec.image`.

**The Dockerfile packages a prebuilt binary; it does not compile.** `dist/manager-<arch>`
must exist before `docker build`. `make docker-build` and the CI `compile` job both
handle that. Compiling in-image instead makes buildx run the Go toolchain under QEMU
for the non-native arch, which took ~40min per build; cross-compiling natively takes
~30s for both arches.

`make manifests` and `make generate` write into `config/` and `api/`, that output is
**committed**, and CI fails if it drifts from what the types produce.

Deploying to a cluster (`make install` / `make deploy` / `make undeploy`) follows the
stock kubebuilder flow; see the
[quick start](https://book.kubebuilder.io/quick-start.html).

## CI

- `.github/workflows/ci.yml`: `lint` and `test` on every push and PR, then
  `compile` (one native job per arch) and `build`, which pushes a multi-arch image to
  `ghcr.io/savedhq/operator` on pushes to `main` and on `v*` tags (`:<short-sha>` +
  `:latest` on main, `:<tag>` on a tag). Every job is capped at
  `timeout-minutes: 10`.
- `.github/workflows/scan.yml`, gitleaks on push, PR and weekly.

E2E tests (`make test-e2e`) need a local Kind cluster and are **not** wired into CI
yet, there is no controller behavior to exercise.
