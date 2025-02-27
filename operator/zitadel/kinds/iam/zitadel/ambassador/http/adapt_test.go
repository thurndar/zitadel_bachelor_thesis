package http

import (
	"testing"

	"github.com/caos/orbos/mntr"
	kubernetesmock "github.com/caos/orbos/pkg/kubernetes/mock"
	"github.com/caos/orbos/pkg/labels"
	"github.com/caos/orbos/pkg/labels/mocklabels"
	"github.com/caos/zitadel/operator/zitadel/kinds/iam/zitadel/configuration"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	apixv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func SetReturnResourceVersion(
	k8sClient *kubernetesmock.MockClientInt,
	group,
	version,
	kind,
	namespace,
	name string,
	resourceVersion string,
) {
	ret := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"metadata": map[string]interface{}{
				"resourceVersion": resourceVersion,
			},
		},
	}
	k8sClient.EXPECT().GetNamespacedCRDResource(group, version, kind, namespace, name).MinTimes(1).MaxTimes(1).Return(ret, nil)
}

func TestHttp_Adapt(t *testing.T) {
	monitor := mntr.Monitor{}
	namespace := "test"
	url := "url"
	dns := &configuration.DNS{
		Domain:    "",
		TlsSecret: "",
		Subdomains: &configuration.Subdomains{
			Accounts: "",
			API:      "",
			Console:  "",
			Issuer:   "",
		},
	}
	k8sClient := kubernetesmock.NewMockClientInt(gomock.NewController(t))

	k8sClient.EXPECT().CheckCRD("mappings.getambassador.io").MinTimes(1).MaxTimes(1).Return(&apixv1beta1.CustomResourceDefinition{}, nil)

	group := "getambassador.io"
	version := "v2"
	kind := "Mapping"

	cors := map[string]interface{}{
		"origins":         "*",
		"methods":         "POST, GET, OPTIONS, DELETE, PUT",
		"headers":         "*",
		"credentials":     true,
		"exposed_headers": "*",
		"max_age":         "86400",
	}

	componentLabels := mocklabels.Component
	endSessionName := labels.MustForName(componentLabels, EndsessionName)
	endsession := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": group + "/" + version,
			"kind":       kind,
			"metadata": map[string]interface{}{
				"labels":    labels.MustK8sMap(endSessionName),
				"name":      endSessionName.Name(),
				"namespace": namespace,
			},
			"spec": map[string]interface{}{
				"connect_timeout_ms": 30000,
				"host":               ".",
				"prefix":             "/oauth/v2/endsession",
				"rewrite":            "",
				"service":            url,
				"timeout_ms":         30000,
				"cors":               cors,
			},
		},
	}
	SetReturnResourceVersion(k8sClient, group, version, kind, namespace, EndsessionName, "")
	k8sClient.EXPECT().ApplyNamespacedCRDResource(group, version, kind, namespace, EndsessionName, endsession).MinTimes(1).MaxTimes(1)

	uploadName := labels.MustForName(componentLabels, Upload)
	upload := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": group + "/" + version,
			"kind":       kind,
			"metadata": map[string]interface{}{
				"labels":    labels.MustK8sMap(uploadName),
				"name":      uploadName.Name(),
				"namespace": namespace,
			},
			"spec": map[string]interface{}{
				"connect_timeout_ms": 30000,
				"host":               ".",
				"prefix":             "/assets/v1",
				"rewrite":            "",
				"service":            url,
				"timeout_ms":         30000,
				"cors":               cors,
			},
		},
	}
	SetReturnResourceVersion(k8sClient, group, version, kind, namespace, Upload, "")
	k8sClient.EXPECT().ApplyNamespacedCRDResource(group, version, kind, namespace, Upload, upload).MinTimes(1).MaxTimes(1)

	issuerName := labels.MustForName(componentLabels, IssuerName)
	issuer := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": group + "/" + version,
			"kind":       kind,
			"metadata": map[string]interface{}{
				"labels":    labels.MustK8sMap(issuerName),
				"name":      issuerName.Name(),
				"namespace": namespace,
			},
			"spec": map[string]interface{}{
				"connect_timeout_ms": 30000,
				"host":               ".",
				"prefix":             "/.well-known/openid-configuration",
				"rewrite":            "/oauth/v2/.well-known/openid-configuration",
				"service":            url,
				"timeout_ms":         30000,
				"cors":               cors,
			},
		},
	}
	SetReturnResourceVersion(k8sClient, group, version, kind, namespace, IssuerName, "")
	k8sClient.EXPECT().ApplyNamespacedCRDResource(group, version, kind, namespace, IssuerName, issuer).MinTimes(1).MaxTimes(1)

	authorizeName := labels.MustForName(componentLabels, AuthorizeName)
	authorize := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": group + "/" + version,
			"kind":       kind,
			"metadata": map[string]interface{}{
				"labels":    labels.MustK8sMap(authorizeName),
				"name":      authorizeName.Name(),
				"namespace": namespace,
			},
			"spec": map[string]interface{}{
				"connect_timeout_ms": 30000,
				"host":               ".",
				"prefix":             "/oauth/v2/authorize",
				"rewrite":            "",
				"service":            url,
				"timeout_ms":         30000,
				"cors":               cors,
			},
		},
	}
	SetReturnResourceVersion(k8sClient, group, version, kind, namespace, AuthorizeName, "")
	k8sClient.EXPECT().ApplyNamespacedCRDResource(group, version, kind, namespace, AuthorizeName, authorize).MinTimes(1).MaxTimes(1)

	oauthName := labels.MustForName(componentLabels, OauthName)
	oauth := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": group + "/" + version,
			"kind":       kind,
			"metadata": map[string]interface{}{
				"labels":    labels.MustK8sMap(oauthName),
				"name":      oauthName.Name(),
				"namespace": namespace,
			},
			"spec": map[string]interface{}{
				"connect_timeout_ms": 30000,
				"host":               ".",
				"prefix":             "/oauth/v2/",
				"rewrite":            "",
				"service":            url,
				"timeout_ms":         30000,
				"cors":               cors,
			},
		},
	}
	SetReturnResourceVersion(k8sClient, group, version, kind, namespace, OauthName, "")
	k8sClient.EXPECT().ApplyNamespacedCRDResource(group, version, kind, namespace, OauthName, oauth).MinTimes(1).MaxTimes(1)

	mgmtName := labels.MustForName(componentLabels, MgmtName)
	mgmt := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": group + "/" + version,
			"kind":       kind,
			"metadata": map[string]interface{}{
				"labels":    labels.MustK8sMap(mgmtName),
				"name":      mgmtName.Name(),
				"namespace": namespace,
			},
			"spec": map[string]interface{}{
				"connect_timeout_ms": 30000,
				"host":               ".",
				"prefix":             "/management/v1/",
				"rewrite":            "",
				"service":            url,
				"timeout_ms":         30000,
				"cors":               cors,
			},
		},
	}
	SetReturnResourceVersion(k8sClient, group, version, kind, namespace, MgmtName, "")
	k8sClient.EXPECT().ApplyNamespacedCRDResource(group, version, kind, namespace, MgmtName, mgmt).MinTimes(1).MaxTimes(1)

	adminRName := labels.MustForName(componentLabels, AdminRName)
	adminR := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": group + "/" + version,
			"kind":       kind,
			"metadata": map[string]interface{}{
				"labels":    labels.MustK8sMap(adminRName),
				"name":      adminRName.Name(),
				"namespace": namespace,
			},
			"spec": map[string]interface{}{
				"connect_timeout_ms": 30000,
				"host":               ".",
				"prefix":             "/admin/v1",
				"rewrite":            "",
				"service":            url,
				"timeout_ms":         30000,
				"cors":               cors,
			},
		},
	}
	SetReturnResourceVersion(k8sClient, group, version, kind, namespace, AdminRName, "")
	k8sClient.EXPECT().ApplyNamespacedCRDResource(group, version, kind, namespace, AdminRName, adminR).MinTimes(1).MaxTimes(1)

	authRName := labels.MustForName(componentLabels, AuthRName)
	authR := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": group + "/" + version,
			"kind":       kind,
			"metadata": map[string]interface{}{
				"labels":    labels.MustK8sMap(authRName),
				"name":      authRName.Name(),
				"namespace": namespace,
			},
			"spec": map[string]interface{}{
				"connect_timeout_ms": 30000,
				"host":               ".",
				"prefix":             "/auth/v1/",
				"rewrite":            "",
				"service":            url,
				"timeout_ms":         30000,
				"cors":               cors,
			},
		},
	}
	SetReturnResourceVersion(k8sClient, group, version, kind, namespace, AuthRName, "")
	k8sClient.EXPECT().ApplyNamespacedCRDResource(group, version, kind, namespace, AuthRName, authR).MinTimes(1).MaxTimes(1)

	openAPIName := labels.MustForName(componentLabels, OpenAPIName)
	openAPI := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": group + "/" + version,
			"kind":       kind,
			"metadata": map[string]interface{}{
				"labels":    labels.MustK8sMap(openAPIName),
				"name":      openAPIName.Name(),
				"namespace": namespace,
			},
			"spec": map[string]interface{}{
				"connect_timeout_ms": 30000,
				"host":               ".",
				"prefix":             "/openapi/v2/swagger",
				"rewrite":            "",
				"service":            url,
				"timeout_ms":         30000,
			},
		},
	}
	SetReturnResourceVersion(k8sClient, group, version, kind, namespace, OpenAPIName, "")
	k8sClient.EXPECT().ApplyNamespacedCRDResource(group, version, kind, namespace, OpenAPIName, openAPI).MinTimes(1).MaxTimes(1)

	query, _, err := AdaptFunc(monitor, componentLabels, namespace, url, dns)
	assert.NoError(t, err)
	queried := map[string]interface{}{}
	ensure, err := query(k8sClient, queried)
	assert.NoError(t, err)
	assert.NoError(t, ensure(k8sClient))
}

func TestHttp_Adapt2(t *testing.T) {
	monitor := mntr.Monitor{}
	namespace := "test"
	url := "url"
	dns := &configuration.DNS{
		Domain:    "domain",
		TlsSecret: "tls",
		Subdomains: &configuration.Subdomains{
			Accounts: "accounts",
			API:      "api",
			Console:  "console",
			Issuer:   "issuer",
		},
	}
	k8sClient := kubernetesmock.NewMockClientInt(gomock.NewController(t))

	k8sClient.EXPECT().CheckCRD("mappings.getambassador.io").MinTimes(1).MaxTimes(1).Return(&apixv1beta1.CustomResourceDefinition{}, nil)

	group := "getambassador.io"
	version := "v2"
	kind := "Mapping"

	cors := map[string]interface{}{
		"origins":         "*",
		"methods":         "POST, GET, OPTIONS, DELETE, PUT",
		"headers":         "*",
		"credentials":     true,
		"exposed_headers": "*",
		"max_age":         "86400",
	}

	componentLabels := mocklabels.Component

	endsessionName := labels.MustForName(componentLabels, EndsessionName)
	endsession := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": group + "/" + version,
			"kind":       kind,
			"metadata": map[string]interface{}{
				"labels":    labels.MustK8sMap(endsessionName),
				"name":      endsessionName.Name(),
				"namespace": namespace,
			},
			"spec": map[string]interface{}{
				"connect_timeout_ms": 30000,
				"host":               "accounts.domain",
				"prefix":             "/oauth/v2/endsession",
				"rewrite":            "",
				"service":            url,
				"timeout_ms":         30000,
				"cors":               cors,
			},
		},
	}
	SetReturnResourceVersion(k8sClient, group, version, kind, namespace, EndsessionName, "")
	k8sClient.EXPECT().ApplyNamespacedCRDResource(group, version, kind, namespace, EndsessionName, endsession).MinTimes(1).MaxTimes(1)

	uploadName := labels.MustForName(componentLabels, Upload)
	upload := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": group + "/" + version,
			"kind":       kind,
			"metadata": map[string]interface{}{
				"labels":    labels.MustK8sMap(uploadName),
				"name":      uploadName.Name(),
				"namespace": namespace,
			},
			"spec": map[string]interface{}{
				"connect_timeout_ms": 30000,
				"host":               "api.domain",
				"prefix":             "/assets/v1",
				"rewrite":            "",
				"service":            url,
				"timeout_ms":         30000,
				"cors":               cors,
			},
		},
	}
	SetReturnResourceVersion(k8sClient, group, version, kind, namespace, Upload, "")
	k8sClient.EXPECT().ApplyNamespacedCRDResource(group, version, kind, namespace, Upload, upload).MinTimes(1).MaxTimes(1)

	issuerName := labels.MustForName(componentLabels, IssuerName)
	issuer := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": group + "/" + version,
			"kind":       kind,
			"metadata": map[string]interface{}{
				"labels":    labels.MustK8sMap(issuerName),
				"name":      issuerName.Name(),
				"namespace": namespace,
			},
			"spec": map[string]interface{}{
				"connect_timeout_ms": 30000,
				"host":               "issuer.domain",
				"prefix":             "/.well-known/openid-configuration",
				"rewrite":            "/oauth/v2/.well-known/openid-configuration",
				"service":            url,
				"timeout_ms":         30000,
				"cors":               cors,
			},
		},
	}
	SetReturnResourceVersion(k8sClient, group, version, kind, namespace, IssuerName, "")
	k8sClient.EXPECT().ApplyNamespacedCRDResource(group, version, kind, namespace, IssuerName, issuer).MinTimes(1).MaxTimes(1)

	authorizeName := labels.MustForName(componentLabels, AuthorizeName)
	authorize := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": group + "/" + version,
			"kind":       kind,
			"metadata": map[string]interface{}{
				"labels":    labels.MustK8sMap(authorizeName),
				"name":      authorizeName.Name(),
				"namespace": namespace,
			},
			"spec": map[string]interface{}{
				"connect_timeout_ms": 30000,
				"host":               "accounts.domain",
				"prefix":             "/oauth/v2/authorize",
				"rewrite":            "",
				"service":            url,
				"timeout_ms":         30000,
				"cors":               cors,
			},
		},
	}
	SetReturnResourceVersion(k8sClient, group, version, kind, namespace, AuthorizeName, "")
	k8sClient.EXPECT().ApplyNamespacedCRDResource(group, version, kind, namespace, AuthorizeName, authorize).MinTimes(1).MaxTimes(1)

	oauthName := labels.MustForName(componentLabels, OauthName)
	oauth := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": group + "/" + version,
			"kind":       kind,
			"metadata": map[string]interface{}{
				"labels":    labels.MustK8sMap(oauthName),
				"name":      oauthName.Name(),
				"namespace": namespace,
			},
			"spec": map[string]interface{}{
				"connect_timeout_ms": 30000,
				"host":               "api.domain",
				"prefix":             "/oauth/v2/",
				"rewrite":            "",
				"service":            url,
				"timeout_ms":         30000,
				"cors":               cors,
			},
		},
	}
	SetReturnResourceVersion(k8sClient, group, version, kind, namespace, OauthName, "")
	k8sClient.EXPECT().ApplyNamespacedCRDResource(group, version, kind, namespace, OauthName, oauth).MinTimes(1).MaxTimes(1)

	mgmtName := labels.MustForName(componentLabels, MgmtName)
	mgmt := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": group + "/" + version,
			"kind":       kind,
			"metadata": map[string]interface{}{
				"labels":    labels.MustK8sMap(mgmtName),
				"name":      mgmtName.Name(),
				"namespace": namespace,
			},
			"spec": map[string]interface{}{
				"connect_timeout_ms": 30000,
				"host":               "api.domain",
				"prefix":             "/management/v1/",
				"rewrite":            "",
				"service":            url,
				"timeout_ms":         30000,
				"cors":               cors,
			},
		},
	}
	SetReturnResourceVersion(k8sClient, group, version, kind, namespace, MgmtName, "")
	k8sClient.EXPECT().ApplyNamespacedCRDResource(group, version, kind, namespace, MgmtName, mgmt).MinTimes(1).MaxTimes(1)

	adminRName := labels.MustForName(componentLabels, AdminRName)
	adminR := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": group + "/" + version,
			"kind":       kind,
			"metadata": map[string]interface{}{
				"labels":    labels.MustK8sMap(adminRName),
				"name":      adminRName.Name(),
				"namespace": namespace,
			},
			"spec": map[string]interface{}{
				"connect_timeout_ms": 30000,
				"host":               "api.domain",
				"prefix":             "/admin/v1",
				"rewrite":            "",
				"service":            url,
				"timeout_ms":         30000,
				"cors":               cors,
			},
		},
	}
	SetReturnResourceVersion(k8sClient, group, version, kind, namespace, AdminRName, "")
	k8sClient.EXPECT().ApplyNamespacedCRDResource(group, version, kind, namespace, AdminRName, adminR).MinTimes(1).MaxTimes(1)

	authRName := labels.MustForName(componentLabels, AuthRName)
	authR := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": group + "/" + version,
			"kind":       kind,
			"metadata": map[string]interface{}{
				"labels":    labels.MustK8sMap(authRName),
				"name":      authRName.Name(),
				"namespace": namespace,
			},
			"spec": map[string]interface{}{
				"connect_timeout_ms": 30000,
				"host":               "api.domain",
				"prefix":             "/auth/v1/",
				"rewrite":            "",
				"service":            url,
				"timeout_ms":         30000,
				"cors":               cors,
			},
		},
	}
	SetReturnResourceVersion(k8sClient, group, version, kind, namespace, AuthRName, "")
	k8sClient.EXPECT().ApplyNamespacedCRDResource(group, version, kind, namespace, AuthRName, authR).MinTimes(1).MaxTimes(1)

	openAPIName := labels.MustForName(componentLabels, OpenAPIName)
	openAPI := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": group + "/" + version,
			"kind":       kind,
			"metadata": map[string]interface{}{
				"labels":    labels.MustK8sMap(openAPIName),
				"name":      openAPIName.Name(),
				"namespace": namespace,
			},
			"spec": map[string]interface{}{
				"connect_timeout_ms": 30000,
				"host":               "api.domain",
				"prefix":             "/openapi/v2/swagger",
				"rewrite":            "",
				"service":            url,
				"timeout_ms":         30000,
			},
		},
	}
	SetReturnResourceVersion(k8sClient, group, version, kind, namespace, OpenAPIName, "")
	k8sClient.EXPECT().ApplyNamespacedCRDResource(group, version, kind, namespace, OpenAPIName, openAPI).MinTimes(1).MaxTimes(1)

	query, _, err := AdaptFunc(monitor, componentLabels, namespace, url, dns)
	assert.NoError(t, err)
	queried := map[string]interface{}{}
	ensure, err := query(k8sClient, queried)
	assert.NoError(t, err)
	assert.NoError(t, ensure(k8sClient))
}
