// +build ignore
package disTensorflow

import (
	"fmt"
	apiv1 "k8s.io/api/core/v1"
	    "k8s.io/client-go/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type srv_interface struct {
	service corev1.ServiceInterface
}


func Get_disService_interface(clientset *kubernetes.Clientset) srv_interface{
	var s srv_interface
	s.service = clientset.CoreV1().Services(apiv1.NamespaceDefault)
//	s.create_service()
	//s.delete_service()
	//s.list_service()
//	s.update_service()
//	s.watch_service()
   return s
}
var srvName string="dist-tensorflow-wk-service"
func (s *srv_interface) Create_GPU_service ()  {
	var srvName string="dist-tensorflow-wk-gpu-service"
	var podName string="dist-gpu-tensor-worker"
	service_yaml := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: srvName,
			Labels: map[string]string{
						"name": podName,
			},
		},
		Spec: apiv1.ServiceSpec{
			Selector: map[string]string{
				"name": podName,
			},
			Ports: []apiv1.ServicePort{
				{
					Name: podName,
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
		fmt.Printf("%s is created gpu successful", service.Name)
	}
}
func (s *srv_interface) Create_service ()  {
	service_yaml := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "dist-tensorflow-wk-service",
			Labels: map[string]string{
						"name": podName,
			},
		},
		Spec: apiv1.ServiceSpec{
			Selector: map[string]string{
				"name": podName,
			},
			Ports: []apiv1.ServicePort{
				{
					Name: podName,
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

func (s *srv_interface) Delete_GPU_service()  {
		var srvName string="dist-tensorflow-wk-gpu-service"
	err := s.service.Delete(srvName,&metav1.DeleteOptions{})
	if err !=nil {
		panic(err)
	} else {
		fmt.Printf("delete successful")
	}
}

func (s *srv_interface) Delete_service()  {
	err := s.service.Delete("dist-tensorflow-wk-service",&metav1.DeleteOptions{})
	if err !=nil {
		panic(err)
	} else {
		fmt.Printf("delete successful")
	}
}

func (s *srv_interface) list_service ()  {
	servicelist, err := s.service.List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, i :=range servicelist.Items {
		fmt.Printf("%s \n", i.Name)
	}
}

func (s *srv_interface) update_service()  {
	service_yaml, err := s.service.Get("dist-tensorflow-wk-service",metav1.GetOptions{})
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

func (s *srv_interface) watch_service()  {
	watch_service, err := s.service.Watch(metav1.ListOptions{})
	if err !=nil {
		panic(err)
	}
	select {
		case e := <-watch_service.ResultChan():
			fmt.Println(e.Type)
	}
}