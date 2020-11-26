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

https://github.com/operator-framework/operator-sdk/blob/master/testdata/go/memcached-operator/controllers/memcached_controller.go
https://sdk.operatorframework.io/docs/building-operators/golang/tutorial/#generating-crd-manifests
https://github.com/jaegertracing/jaeger-operator/blob/master/pkg/ingress/ingress.go

*/

package controllers

import (
	"context"
	"github.com/go-logr/logr"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"bytes"
	"encoding/json"
	"github.com/smuel1414/ingresses-changes/utils"

	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

// IngressReconciler reconciles a Ingress object
type IngressReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// IngressData is
type IngressData struct {
	Name   string `json:"name"`
	Host   string `json:"host"`
	Expose bool   `json:"expose"`
}

// +kubebuilder:rbac:groups=extensions,resources=ingresses,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=extensions,resources=ingresses/status,verbs=get;update;patch

var log = logf.Log.WithName("controllers.Ingress")

//Reconcile is
func (r *IngressReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {

	conf := utils.ReadConfig()
	ctx := context.Background()

	// your logic here
	// ###########################################################################################################################
	// Lookup the instance for this reconcile request
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

	// ###########################################################################################################################
	// ###########################################################################################################################
	// ###########################################################################################################################

	var hosts []string

	for _, rule := range ingress.Spec.Rules {
		host = rule.Host)

		currentHostData := IngressData{
			Name:   ingress.Name,
			Host:   host,
			Expose: true,
		}

		payloadBuf := new(bytes.Buffer)
		json.NewEncoder(payloadBuf).Encode(currentHostData)

		utils.MakePostRequest(conf.IngressesHandlerAddress, payloadBuf)

	}


	// ###########################################################################################################################
	// ###########################################################################################################################
	// ###########################################################################################################################
	return ctrl.Result{}, nil
}

// https://sdk.operatorframework.io/docs/building-operators/golang/references/event-filtering/

//SetupWithManager is
func (r *IngressReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&extensionsv1beta1.Ingress{}).
		WithEventFilter(utils.UsePredicate()).
		Complete(r)
}
