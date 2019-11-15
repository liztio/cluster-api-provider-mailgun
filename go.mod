module github.com/liztio/cluster-api-provider-mailgun

go 1.12

require (
	github.com/go-logr/logr v0.1.0
	github.com/gobuffalo/envy v1.7.1 // indirect
	github.com/mailgun/mailgun-go v2.0.0+incompatible
	github.com/onsi/ginkgo v1.10.1
	github.com/onsi/gomega v1.7.0
	k8s.io/apimachinery v0.0.0-20190817020851-f2f3a405f61d
	k8s.io/client-go v0.0.0-20190918200256-06eb1244587a
	sigs.k8s.io/cluster-api v0.2.6-0.20191114201936-ce4ea5c522d5
	sigs.k8s.io/controller-runtime v0.3.0
)
