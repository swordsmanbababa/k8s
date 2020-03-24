// +build ignore
package main

import (
	"fmt"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8sctl"
)

type srv struct {
	service corev1.ServiceInterface
}


func main_is() {
	clientset:=k8sctl.GetClientSet()
	var s srv
	s.service = clientset.CoreV1().Services(apiv1.NamespaceDefault)
	s.create_service()
	//s.delete_service()
	//s.list_service()
//	s.update_service()
//	s.watch_service()
}

func (s *srv) create_service ()  {
	service_yaml := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "nginx",
		},
		Spec: apiv1.ServiceSpec{
			Selector: map[string]string{
				"app": "nginx",
			},
			Ports: []apiv1.ServicePort{
				{
					Name: "nginx",
					Port: 88,
					TargetPort: intstr.IntOrString{
						Type: intstr.Int,
						IntVal: 80,
					},
					Protocol: apiv1.ProtocolTCP,
				},
			},
		},
	}
	service, err := s.service.Create(service_yaml)
	if err != nil {
		panic(err)
	}else {
		fmt.Printf("%s is created successful", service.Name)
	}
}

func (s *srv) delete_service()  {
	err := s.service.Delete("nginx",&metav1.DeleteOptions{})
	if err !=nil {
		panic(err)
	} else {
		fmt.Printf("delete successful")
	}
}

func (s *srv) list_service ()  {
	servicelist, err := s.service.List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, i :=range servicelist.Items {
		fmt.Printf("%s \n", i.Name)
	}
}

func (s *srv) update_service()  {
	service_yaml, err := s.service.Get("nginx",metav1.GetOptions{})
	if err != nil {
		panic(err)
	}
	service_yaml.Spec.Ports = []apiv1.ServicePort{
		{
			Port: 98,
			TargetPort: intstr.IntOrString{
				Type: intstr.Int,
				IntVal: 80,
			},
			Protocol: apiv1.ProtocolTCP,
		},
	}

	service, err := s.service.Update(service_yaml)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("the service %s  update successful", service.Name)
	}
}

func (s *srv) watch_service()  {
	watch_service, err := s.service.Watch(metav1.ListOptions{})
	if err !=nil {
		panic(err)
	}
	select {
		case e := <-watch_service.ResultChan():
			fmt.Println(e.Type)
	}
}