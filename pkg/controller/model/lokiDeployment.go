package model

import (
	"github.com/ronak-agarwal/loki-operator/pkg/apis/loki/v1alpha1"
	v1 "k8s.io/api/apps/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// LokiDeployment ...
func LokiDeployment(cr *v1alpha1.Loki, configHash, dsHash string) *v1.Deployment {
	return &v1.Deployment{
		ObjectMeta: v12.ObjectMeta{
			Name:      LokiDeploymentName,
			Namespace: cr.Namespace,
		},
		Spec: getDeploymentSpec(cr, nil, configHash, "", dsHash),
	}
}

func getDeploymentSpec(cr *v1alpha1.Loki, annotations map[string]string, configHash, plugins, dsHash string) v1.DeploymentSpec {
	return v1.DeploymentSpec{}
}
