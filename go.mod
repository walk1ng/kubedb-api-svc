module kubedb-api-svc

go 1.14

require github.com/gin-gonic/gin v1.6.3

require (
	github.com/go-openapi/spec v0.19.4 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/google/go-cmp v0.4.0 // indirect
	github.com/googleapis/gnostic v0.3.1 // indirect
	github.com/onsi/ginkgo v1.11.0 // indirect
	github.com/onsi/gomega v1.8.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/crypto v0.0.0-20200429183012-4b2356b1ed79 // indirect
	golang.org/x/net v0.0.0-20200506145744-7e3656a0809f // indirect
	golang.org/x/sys v0.0.0-20200501145240-bc7a7d42d5c3 // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	gomodules.xyz/jsonpatch/v2 v2.0.1 // indirect
	google.golang.org/appengine v1.6.5 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	k8s.io/api v0.17.4
	k8s.io/apiextensions-apiserver v0.17.4 // indirect
	k8s.io/apimachinery v0.17.4
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/utils v0.0.0-20191114200735-6ca3b61696b6 // indirect
	kubedb.dev/apimachinery v0.13.0-rc.0

)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.3.2+incompatible // Required by OLM
	k8s.io/apimachinery => github.com/sangyuruo/apimachinery-1 v0.17.5-beta.0.0.20200508080956-56dd00734ae5
	k8s.io/client-go => k8s.io/client-go v0.17.4 // Required by prometheus-operator
	// 引入 kubedb
	kubedb.dev/apimachinery => github.com/kubedb/apimachinery v0.13.0-rc.0
)
