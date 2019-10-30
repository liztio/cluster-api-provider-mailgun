/*

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

package controllers

import (
	"context"
	"fmt"

	apierrors "k8s.io/apimachinery/pkg/api/errors"

	"github.com/go-logr/logr"
	"github.com/mailgun/mailgun-go"
	"sigs.k8s.io/cluster-api/util"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	infrastructurev1alpha3 "github.com/liztio/cluster-api-provider-mailgun/api/v1alpha3"
)

// MailgunClusterReconciler reconciles a MailgunCluster object
type MailgunClusterReconciler struct {
	client.Client
	Log       logr.Logger
	Mailgun   mailgun.Mailgun
	Recipient string
}

// +kubebuilder:rbac:groups=infrastructure.cluster.x-k8s.io,resources=mailgunclusters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=infrastructure.cluster.x-k8s.io,resources=mailgunclusters/status,verbs=get;update;patch

func (r *MailgunClusterReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	_ = r.Log.WithValues("mailguncluster", req.NamespacedName)

	var mgCluster infrastructurev1alpha3.MailgunCluster
	if err := r.Get(ctx, req.NamespacedName, &mgCluster); err != nil {
		// 	apierrors "k8s.io/apimachinery/pkg/api/errors"
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	cluster, err := util.GetOwnerMachine(ctx, r.Client, mgCluster.ObjectMeta)
	if err != nil {
		return ctrl.Result{}, err
	}

	if mgCluster.Status.MessageID != nil {
		// We already sent a message, so skip reconcilation
		return ctrl.Result{}, nil
	}

	subject := fmt.Sprintf("[%s] New Cluster %s requested", mgCluster.Spec.Priority, cluster.Name)
	body := fmt.Sprint("Hello! One cluster please.\n\n%s\n", mgCluster.Spec.Request)

	msg := mailgun.NewMessage(mgCluster.Spec.Requester, subject, body, r.Recipient)
	_, msgID, err := r.Mailgun.Send(msg)
	if err != nil {
		return ctrl.Result{}, err
	}

	mgCluster.Status.MessageID = &msgID
	r.Update(ctx, &mgCluster)

	return ctrl.Result{}, nil
}

func (r *MailgunClusterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&infrastructurev1alpha3.MailgunCluster{}).
		Complete(r)
}
