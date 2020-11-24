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

	"github.com/go-logr/logr"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"

	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	// networkingv1beta1 "k8s.io/api/networking/v1beta1"
	// "fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IngressReconciler reconciles a Ingress object
type IngressReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=extensions,resources=ingresses,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=extensions,resources=ingresses/status,verbs=get;update;patch

//Reconcile is
func (r *IngressReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("ingress", req.NamespacedName)

	// your logic here
	// ###########################################################################################################################
	// Lookup the Memcached instance for this reconcile request
	ingress := &extensionsv1beta1.Ingress{}
	err := r.Get(ctx, req.NamespacedName, ingress)

	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			log.Info("Ingress resource not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		log.Error(err, "Failed to get ingress")
		return ctrl.Result{}, err
	}

	// Check if the deployment already exists, if not create a new one
	// found := &appsv1.Deployment{}
	err = r.Get(ctx, types.NamespacedName{Name: ingress.Name, Namespace: ingress.Namespace}, ingress)
	if err != nil && errors.IsNotFound(err) {
		// Define a new deployment
		// dep := r.deploymentForMemcached(ingress)
		// log.Info("Log ingress round", "ingress.Namespace", ingress.Namespace, "ingress.Name", ingress.Name)
		// err = r.Create(ctx, dep)
		// if err != nil {
		// 	log.Error(err, "Failed to create new Deployment", "Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)
		// 	return ctrl.Result{}, err
		// }
		// Deployment created successfully - return and requeue
		return ctrl.Result{Requeue: false}, nil
	} else if err != nil {
		log.Error(err, "Failed to get ingress")
		return ctrl.Result{}, err
	}

	// Ensure the deployment size is the same as the spec
	oldIngress := extensionsv1beta1.Ingress{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Ingress",
			APIVersion: "extensions/v1beta1",
		},
		ObjectMeta: ingress.ObjectMeta,
	}
	// metaData := ingress
	// log = r.Log.WithValues("Lables", oldIngress.ObjectMeta.Labels)
	// log.Info("Ingress metadata ", ":", oldIngress.ObjectMeta.Labels)
	// fmt.Printf("%+v", oldIngress.ObjectMeta)

	r.Log.Info("\n\n\nLet's try just print")
	r.Log.Info("\n\n\n", "Lables", oldIngress.ObjectMeta.Labels)

	// Ensure the deployment size is the same as the spec
	// size := memcached.Spec.Size
	// if *found.Spec.Replicas != size {
	// 	found.Spec.Replicas = &size
	// 	err = r.Update(ctx, found)
	// 	if err != nil {
	// 		log.Error(err, "Failed to update Deployment", "Deployment.Namespace", found.Namespace, "Deployment.Name", found.Name)
	// 		return ctrl.Result{}, err
	// 	}
	// 	// Spec updated - return and requeue
	// 	return ctrl.Result{Requeue: true}, nil
	// }

	// Update the Memcached status with the pod names
	// List the pods for this memcached's deployment
	// podList := &corev1.PodList{}
	// listOpts := []client.ListOption{
	// 	client.InNamespace(memcached.Namespace),
	// 	client.MatchingLabels(labelsForMemcached(memcached.Name)),
	// }
	// if err = r.List(ctx, podList, listOpts...); err != nil {
	// 	log.Error(err, "Failed to list pods", "Memcached.Namespace", memcached.Namespace, "Memcached.Name", memcached.Name)
	// 	return ctrl.Result{}, err
	// }
	// podNames := getPodNames(podList.Items)

	// // Update status.Nodes if needed
	// if !reflect.DeepEqual(podNames, memcached.Status.Nodes) {
	// 	memcached.Status.Nodes = podNames
	// 	err := r.Status().Update(ctx, memcached)
	// 	if err != nil {
	// 		log.Error(err, "Failed to update Memcached status")
	// 		return ctrl.Result{}, err
	// 	}
	// }

	// #######################################################################################################################################################

	return ctrl.Result{}, nil
}

// https://sdk.operatorframework.io/docs/building-operators/golang/references/event-filtering/
// func ignoreDeletionPredicate() predicate.Predicate {
// 	return predicate.Funcs{
// 		UpdateFunc: func(e event.UpdateEvent) bool {
// 			// Ignore updates to CR status in which case metadata.Generation does not change
// 			return e.MetaOld.GetGeneration() != e.MetaNew.GetGeneration()
// 		},
// 		DeleteFunc: func(e event.DeleteEvent) bool {
// 			// Evaluates to false if the object has been confirmed deleted.
// 			return !e.DeleteStateUnknown
// 		},
// 	}
// }

//SetupWithManager is
func (r *IngressReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&extensionsv1beta1.Ingress{}).
		Complete(r)
}
