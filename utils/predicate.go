// B"H

package utils

// credit https://blog.risingstack.com/golang-tutorial-for-nodejs-developers-getting-started/#nethttp
// https://sdk.operatorframework.io/docs/building-operators/golang/references/event-filtering/

import (
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// UsePredicate is operator function that return predicate.Funcs object contains Create, Update, Delete and Generic functions
// That executed accordingly to the current event.
// The code is basicly repeintng it self since the objects types and functions are with diffrents names.
// Log function type update, delete etc.., check if object exist, check if object have expose annotation.
func UsePredicate() predicate.Predicate {
	conf := ReadConfig()

	isAnnotatedIngress := predicate.Funcs{
		UpdateFunc: func(e event.UpdateEvent) bool {
			log.V(1).Info("predicate", "function", "Update")
			_, ok := e.ObjectNew.(*extensionsv1beta1.Ingress)

			if !ok {
				return false
			}

			expose, _ := e.MetaNew.GetAnnotations()[conf.ExposeAnnotation]
			doExpose := expose == "true"
			log.V(1).Info("predicate", "function", "Update", "expose", expose)

			return doExpose
		},
		CreateFunc: func(e event.CreateEvent) bool {
			log.V(1).Info("predicate", "function", "Create")
			_, ok := e.Object.(*extensionsv1beta1.Ingress)

			if !ok {
				return false
			}

			expose, _ := e.Meta.GetAnnotations()[conf.ExposeAnnotation]
			doExpose := expose == "true"
			log.V(1).Info("predicate", "function", "Create", "expose", expose)

			return doExpose
		},
		DeleteFunc: func(e event.DeleteEvent) bool {
			log.V(1).Info("predicate", "function", "Delete  ")
			_, ok := e.Object.(*extensionsv1beta1.Ingress)

			if !ok {
				return false
			}

			expose, _ := e.Meta.GetAnnotations()[conf.ExposeAnnotation]
			doExpose := expose == "true"
			log.V(1).Info("predicate", "function", "Delete", "expose", expose)

			return doExpose
		},
		GenericFunc: func(e event.GenericEvent) bool {
			log.V(1).Info("predicate", "function", "Generic")
			_, ok := e.Object.(*extensionsv1beta1.Ingress)

			if !ok {
				return false
			}

			expose, _ := e.Meta.GetAnnotations()[conf.ExposeAnnotation]
			doExpose := expose == "true"
			log.V(1).Info("predicate", "function", "Generic", "expose", expose)

			return doExpose
		},
	}
	return isAnnotatedIngress
}
