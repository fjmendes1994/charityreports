package cmd

import (
	"github.com/go-chi/chi"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"net/http"
)

func StartHttServer() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		// creates the in-cluster config
		config, err := rest.InClusterConfig()
		if err != nil {
			w.Write([]byte("error: " + err.Error()))
		}
		// creates the clientset
		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			w.Write([]byte("error: " + err.Error()))
		}
		pods, err := clientset.CoreV1().Pods("default").List(metav1.ListOptions{})
		if err != nil {
			w.Write([]byte("error: " + err.Error()))
		}

		//// Examples for error handling:
		//// - Use helper functions like e.g. errors.IsNotFound()
		//// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
		//_, err = clientset.CoreV1().Pods("default").Get("example-xxxxx", metav1.GetOptions{})
		//if errors.IsNotFound(err) {
		//	fmt.Printf("Pod not found\n")
		//} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		//	fmt.Printf("Error getting pod %v\n", statusError.ErrStatus.Message)
		//} else if err != nil {
		//	panic(err.Error())
		//} else {
		//	fmt.Printf("Found pod\n")
		//}

		w.Write([]byte(pods.String()))

	})
	http.ListenAndServe(":8080", r)

}
