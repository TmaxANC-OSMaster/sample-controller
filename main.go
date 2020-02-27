package main

import (
	"flag"
	"time"

	kubeinformers "k8s.io/client-go/informers"

	clientset "generated/clientset/versioned"
	informers "generated/informers/externalversions"
)

var (
	masterURL  string
	kubeconfig string
)

func main() {
	stopCh := signals.SetupSignalHandler()

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
	}

	exampleClient, err := clientset.NewForConfig(cfg)
	if err != nil {
	}

	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(kubeClient, time.Second*30)
	exampleInformerFactory := informers.NewSharedInformerFactory(exampleClient, time.Second*30)

	controller := NewController(kubeClient, exampleClient, kubeInformerFactory.Apps.V1().Deployments(), exampleInformerFactory.Samplecontroller().V1apha1().Foos())

	// initialize requested informers
	// runs all registered informers in a dedicated goroutine
	kubeInformerFactory.Start(stopCh)
	exampleInformerFactory.Start(stopCh)

	if err = controller.Run(2, stopCh); err != nil {

	}
}
