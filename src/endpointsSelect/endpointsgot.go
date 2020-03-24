package endpointsSelect

import(
 "k8s.io/client-go/kubernetes"
)


func GetEndpoints(clientset *kubernetes.Clientset){
	endPoints:err:= clientset.CoreV1().Endpoints("").List(metav1.ListOptions{})
	endpointSubset,err:=clientset.CoreV1().EndpointSubset("").List(metav1.ListOptions{})
//	_, err = clientset.CoreV1().Endpoints("default").Get("example-xxxxx", metav1.GetOptions{})
}