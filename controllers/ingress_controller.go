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

	// "k8s.io/apimachinery/pkg/types"
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
	// log := r.Log.WithValues("ingress", req.NamespacedName)

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
			r.Log.Info("Ingress resource not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		r.Log.Error(err, "Failed to get ingress")
		return ctrl.Result{}, err
	}

	// Check if the deployment already exists, if not create a new one
	// found := &appsv1.Deployment{}
	// err = r.Get(ctx, types.NamespacedName{Name: ingress.Name, Namespace: ingress.Namespace}, ingress)
	// if err != nil && errors.IsNotFound(err) {
	// 	// Define a new deployment
	// 	// dep := r.deploymentForMemcached(ingress)
	// 	// log.Info("Log ingress round", "ingress.Namespace", ingress.Namespace, "ingress.Name", ingress.Name)
	// 	// err = r.Create(ctx, dep)
	// 	// if err != nil {
	// 	// 	log.Error(err, "Failed to create new Deployment", "Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)
	// 	// 	return ctrl.Result{}, err
	// 	// }
	// 	// Deployment created successfully - return and requeue
	// 	return ctrl.Result{Requeue: false}, nil
	// } else if err != nil {
	// 	log.Error(err, "Failed to get ingress")
	// 	return ctrl.Result{}, err
	// }

	// Ensure the deployment size is the same as the spec
	currentIngress := extensionsv1beta1.Ingress{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Ingress",
			APIVersion: "extensions/v1beta1",
		},
		ObjectMeta: ingress.ObjectMeta,
	}

	// ###########################################################################################################################
	// ###########################################################################################################################
	// ###########################################################################################################################

	// if ingress.Spec.Backend != nil {
	// 	currentIngress.Spec = netv1beta.IngressSpec{
	// 		Backend: &netv1beta.IngressBackend{
	// 			ServiceName: ingress.Spec.Backend.ServiceName,
	// 			ServicePort: ingress.Spec.Backend.ServicePort,
	// 		},
	// 	}
	// }

	// for _, tls := range ingress.Spec.TLS {
	// 	currentIngress.Spec.TLS = append(currentIngress.Spec.TLS, netv1beta.IngressTLS{
	// 		Hosts:      tls.Hosts,
	// 		SecretName: tls.SecretName,
	// 	})
	// }
	var hosts []string

	for _, rule := range ingress.Spec.Rules {
		// httpIngressPaths := make([]netv1beta.HTTPIngressPath, len(rule.HTTP.Paths))
		// for i, path := range rule.HTTP.Paths {
		// 	httpIngressPaths[i].Backend.ServicePort = path.Backend.ServicePort
		// 	httpIngressPaths[i].Backend.ServiceName = path.Backend.ServiceName
		// 	httpIngressPaths[i].Path = path.Path

		// }
		// hosts := make([]string, len(ingress.Spec.Rules))
		// var hosts []string
		hosts = append(hosts, rule.Host)
		// currentIngress.Spec.Rules = append(currentIngress.Spec.Rules, netv1beta.IngressRule{
		// hosts[i]
		// 	Host: rule.Host,
		// 	IngressRuleValue: netv1beta.IngressRuleValue{
		// 		HTTP: &netv1beta.HTTPIngressRuleValue{
		// 			Paths: httpIngressPaths,
		// 		},
		// 	},
	}

	// ###########################################################################################################################
	// ###########################################################################################################################
	// ###########################################################################################################################

	lables := currentIngress.ObjectMeta.Labels
	// metaData := ingress
	// log = r.Log.WithValues("Lables", oldIngress.ObjectMeta.Labels)
	// log.Info("Ingress metadata ", ":", oldIngress.ObjectMeta.Labels)
	// fmt.Printf("%+v", oldIngress.ObjectMeta)

	// if lables.exposed.dns != nil && errors.IsNotFound(err) {

	// r.Log.Info("\n\n\nLet's try just print")
	// r.Log.Info("\n\n\n", "Lables", labels)

	// for _, f := range AddToManagerFuncs {
	// 	if err := f(m); err != nil {
	// 		return err
	// 	}
	// }

	for key, value := range lables {
		if key == "expose.dns" && value == "true" {
			// expose.dns
			r.Log.Info("\n\n\nThis ingress should ----   ---- be updated.")
			r.Log.Info("\n", key, value)

			// r.Log.Info("\n", ingress.Spec)

			for _, host := range hosts {
				r.Log.Info("\n", "host: ", host)
			}

			return ctrl.Result{}, nil
		}
	}

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
