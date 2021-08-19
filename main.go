package main

import (
	"github.com/tmax-cloud/cloud-credential-api-server/billing"

	"k8s.io/klog"

	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/billing", serveBilling)

	klog.Info("Starting Cloud Credential server...")
	klog.Flush()

	if err := http.ListenAndServe(":80", mux); err != nil {
		klog.Errorf("Failed to listen and serve AWS COST server: %s", err)
	}
	klog.Info("Terminate Cloud Credential server")
}

func serveBilling(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		billing.Get(res, req)
	case http.MethodPut:
	case http.MethodOptions:
	default:
		//error
	}
}
