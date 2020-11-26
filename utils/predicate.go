// B"H

package utils

// credit https://blog.risingstack.com/golang-tutorial-for-nodejs-developers-getting-started/#nethttp
// https://sdk.operatorframework.io/docs/building-operators/golang/references/event-filtering/

import (
	// logf "sigs.k8s.io/controller-runtime/pkg/log"

	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// UsePredicate is
func UsePredicate() predicate.Predicate {
	conf := ReadConfig()

	isAnnotatedIngress := predicate.Funcs{
		UpdateFunc: func(e event.UpdateEvent) bool {
			log.V(1).Info("Update Function ")
			_, ok := e.ObjectNew.(*extensionsv1beta1.Ingress)

			if !ok {
				return false
			}

			expose, _ := e.MetaNew.GetAnnotations()[conf.ExposeAnnotation]
			doExpose := expose == "true"
			log.V(1).Info("Update Function", "expose", expose)

			return doExpose
			// oldIngress, ok := e.ObjectOld.(*extensionsv1beta1.Ingress)
			// log.V(1).Info("In progress", "is old ingress exist: ", ok)
			// if !ok {
			// 	return false
			// }
			// newIngress, ok := e.ObjectNew.(*extensionsv1beta1.Ingress)
			// log.V(1).Info("In progress", "is new ingress exist: ", ok)
			// if !ok {
			// 	return false
			// }
			// // if newIngress.Type != util.TLSSecret {
			// //     return false
			// // }
			// // log.V(1).Info("In progress", "oldIngress: ", oldIngress)
			// // log.V(1).Info("In progress", "newIngress: ", newIngress)
			// oldValue, _ := e.MetaOld.GetAnnotations()["expose.dns"]
			// newValue, _ := e.MetaNew.GetAnnotations()["expose.dns"]
			// // old := oldValue == "true"
			// new := newValue == "true"

			// log.V(1).Info("In progress", "oldValue: ", oldValue)
			// log.V(1).Info("In progress", "newValue: ", newValue)

			// log.V(1).Info("In progress", "e.MetaOld.GetAnnotations(): ", e.MetaOld.GetAnnotations())
			// log.V(1).Info("In progress", "e.MetaNew.GetAnnotations(): ", e.MetaNew.GetAnnotations())
			// if the content has changed we trigger if the annotation is there
			// if !reflect.DeepEqual(newIngress, oldIngress) {
			// 	log.V(1).Info("In progress", "!reflect.DeepEqual(newIngress, oldIngress): ", !reflect.DeepEqual(newIngress, oldIngress))
			// 	return true
			// }
			// otherwise we trigger if the annotation has changed
			// return new
		},
		CreateFunc: func(e event.CreateEvent) bool {
			log.V(1).Info("Create Function ")
			_, ok := e.Object.(*extensionsv1beta1.Ingress)

			if !ok {
				return false
			}

			expose, _ := e.Meta.GetAnnotations()[conf.ExposeAnnotation]
			doExpose := expose == "true"
			log.V(1).Info("Create Function", "expose", expose)

			return doExpose
		},
		DeleteFunc: func(e event.DeleteEvent) bool {
			log.V(1).Info("Delete Function ")
			_, ok := e.Object.(*extensionsv1beta1.Ingress)

			if !ok {
				return false
			}

			expose, _ := e.Meta.GetAnnotations()[conf.ExposeAnnotation]
			doExpose := expose == "true"
			log.V(1).Info("Delete Function", "expose", expose)

			return doExpose
		},
		GenericFunc: func(e event.GenericEvent) bool {
			log.V(1).Info("Generic Function ")
			_, ok := e.Object.(*extensionsv1beta1.Ingress)

			if !ok {
				return false
			}

			expose, _ := e.Meta.GetAnnotations()[conf.ExposeAnnotation]
			doExpose := expose == "true"
			log.V(1).Info("Generic Function", "expose", expose)

			return doExpose
		},
	}
	return isAnnotatedIngress
}
