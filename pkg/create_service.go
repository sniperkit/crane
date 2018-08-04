package crane

/*
import (
	"github.com/docker/libcompose/config"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/unversioned"
)

func createService(shortName string, service *config.ServiceConfig, rc *api.ReplicationController) *api.Service {
	ports := make([]api.ServicePort, len(rc.Spec.Template.Spec.Containers[0].Ports))
	for i, port := range rc.Spec.Template.Spec.Containers[0].Ports {
		ports[i].Port = port.ContainerPort
	}

	srv := &api.Service{
		TypeMeta: unversioned.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: api.ObjectMeta{
			Name:      shortName,
			Namespace: "default",
		},
		Spec: api.ServiceSpec{
			Selector: map[string]string{"service": shortName},
			Ports:    ports,
		},
	}

	return srv
}
*/
