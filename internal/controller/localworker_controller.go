package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	savedv1alpha1 "github.com/savedhq/operator/api/v1alpha1"
)

type LocalWorkerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=worker.saved.sh,resources=localworkers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=worker.saved.sh,resources=localworkers/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=worker.saved.sh,resources=localworkers/finalizers,verbs=update

func (r *LocalWorkerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = logf.FromContext(ctx)

	return ctrl.Result{}, nil
}

func (r *LocalWorkerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&savedv1alpha1.LocalWorker{}).
		Named("localworker").
		Complete(r)
}
