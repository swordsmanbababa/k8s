// +build ignore
package disTensorflow

import (
	"fmt"
	apiv1 "k8s.io/api/core/v1"
	    "k8s.io/client-go/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)



func Get_disPsService_interface(clientset *kubernetes.Clientset) srv_interface{
	var s srv_interface
	s.service = clientset.CoreV1().Services(apiv1.NamespaceDefault)
   return s
}
var srvPsName string="dist-tensorflow-ps-service"
func (s *srv_interface) Create_Ps_service ()  {
	service_yaml := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: srvPsName,
			Labels: map[string]string{
						"name": pspodName,
			},
		},
		Spec: apiv1.ServiceSpec{
			Selector: map[string]string{
				"name": pspodName,
			},
			Ports: []apiv1.ServicePort{
				{
					Name: "dists-tensorflow-ps",
					Port: 2222,
					TargetPort: intstr.IntOrString{
						Type: intstr.Int,
						IntVal: 2222,
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

func (s *srv_interface) Delete_Ps_service()  {
	err := s.service.Delete(srvPsName,&metav1.DeleteOptions{})
	if err !=nil {
		panic(err)
	} else {
		fmt.Printf("delete successful")
	}
}

func (s *srv_interface) list_Ps_service ()  {
	servicelist, err := s.service.List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, i :=range servicelist.Items {
		fmt.Printf("%s \n", i.Name)
	}
}

func (s *srv_interface) update_Ps_service()  {
	service_yaml, err := s.service.Get(srvPsName,metav1.GetOptions{})
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

func (s *srv_interface) watch_Ps_service()  {
	watch_service, err := s.service.Watch(metav1.ListOptions{})
	if err !=nil {
		panic(err)
	}
	select {
		case e := <-watch_service.ResultChan():
			fmt.Println(e.Type)
	}
}