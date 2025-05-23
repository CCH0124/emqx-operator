/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package apps

import (
	"context"

	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	appsv1beta3 "github.com/emqx/emqx-operator/apis/apps/v1beta3"
)

var _ reconcile.Reconciler = &EmqxBrokerReconciler{}

// EmqxEnterpriseReconciler reconciles a EmqxEnterprise object
type EmqxEnterpriseReconciler struct {
	EmqxReconciler
}

//+kubebuilder:rbac:groups=apps.emqx.io,resources=emqxenterprises,verbs=get;list;watch;create;update;patch
//+kubebuilder:rbac:groups=apps.emqx.io,resources=emqxenterprises/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=apps.emqx.io,resources=emqxenterprises/finalizers,verbs=update

func (r *EmqxEnterpriseReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	instance := &appsv1beta3.EmqxEnterprise{}
	if err := r.Get(ctx, req.NamespacedName, instance); err != nil {
		if k8sErrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}
	return r.Do(ctx, instance)
}

// SetupWithManager sets up the controller with the Manager.
func (r *EmqxEnterpriseReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appsv1beta3.EmqxEnterprise{}).
		Complete(r)
}
