/*
Copyright 2022.

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

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	k8sErrors "k8s.io/apimachinery/pkg/api/errors"

	messagesv1beta1 "github.com/kubbee/kafka-messaging-topology-operator/api/v1beta1"
)

// KafkaProducerReconciler reconciles a KafkaProducer object
type KafkaProducerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=messages.kubbee.tech,resources=kafkaproducers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=messages.kubbee.tech,resources=kafkaproducers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=messages.kubbee.tech,resources=kafkaproducers/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the KafkaProducer object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
func (r *KafkaProducerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := ctrl.LoggerFrom(ctx)
	logger.Info("Reconcile")

	kafkaProducer := &messagesv1beta1.KafkaProducer{}

	if err := r.Get(ctx, req.NamespacedName, kafkaProducer); err != nil {
		if k8sErrors.IsNotFound(err) {
			if !kafkaProducer.ObjectMeta.DeletionTimestamp.IsZero() {
				logger.Info("Was marked for deletion.")
				return reconcile.Result{}, nil
			}
		}
		return reconcile.Result{}, nil
	}

	return r.declareTopic(ctx, req, kafkaProducer)
}

// declateTopic This function is reponsible to orchestrate the topic creation
func (r *KafkaProducerReconciler) declareTopic(ctx context.Context, req ctrl.Request, kafkaProducer *messagesv1beta1.KafkaProducer) (ctrl.Result, error) {
	logger := ctrl.LoggerFrom(ctx)
	logger.Info("declareKafkaTopic")

	// TODO implements the creation topic

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *KafkaProducerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&messagesv1beta1.KafkaProducer{}).
		Complete(r)
}
