package main

import (
	"context"
	"encoding/json"
	"fmt"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/external-dns/source"
)

func main() {
	fmt.Println("Using Generated client")
	cfg := ctrl.GetConfigOrDie()
	cfg.QPS = 100
	cfg.Burst = 100

	kc, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		panic(err)
	}

	// func NewIngressSource(ctx context.Context, kubeClient kubernetes.Interface, namespace, annotationFilter string, fqdnTemplate string, combineFqdnAnnotation bool, ignoreHostnameAnnotation bool, ignoreIngressTLSSpec bool, ignoreIngressRulesSpec bool, labelSelector labels.Selector) (Source, error) {

	ctx := context.Background()
	namespace := ""
	annotationFilter := ""
	fqdnTemplate := ""
	combineFqdnAnnotation := false
	ignoreHostnameAnnotation := false
	ignoreIngressTLSSpec := false
	ignoreIngressRulesSpec := false
	labelSelector := labels.SelectorFromSet(map[string]string{})

	s1, err := source.NewIngressSource(
		ctx,
		kc,
		namespace,
		annotationFilter,
		fqdnTemplate,
		combineFqdnAnnotation,
		ignoreHostnameAnnotation,
		ignoreIngressTLSSpec,
		ignoreIngressRulesSpec,
		labelSelector,
	)
	if err != nil {
		panic(err)
	}
	eps, err := s1.Endpoints(ctx)
	if err != nil {
		panic(err)
	}
	data, err := json.MarshalIndent(eps, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
