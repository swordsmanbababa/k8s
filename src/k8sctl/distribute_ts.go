package k8sctl

import (
	"fmt"
	appsbetav1 "k8s.io/api/apps/v1beta1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1beta "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
	//"encoding/json"
)

func Create_disdeploy (deploymentclient v1beta.DeploymentInterface)  {
	fmt.Println("create_disdeploy")
//	var r apiv1.ResourceRequirements
//	j := `{"limits": {"cpu":"2", "memory": "8Gi"}}}`
//	json.Unmarshal([]byte(j), &r)

	deploy := &appsbetav1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "tensorflow-cpu",
		},
		Spec: appsbetav1.DeploymentSpec{
			Replicas: int32Ptr2(1),
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "tensorflow-cpu-tmp",
					},
				},
				Spec: apiv1.PodSpec{
					Containers:[]apiv1.Container{
						{	Name: "tensorflow-cpu-po",
							Image: "registry.cn-hangzhou.aliyuncs.com/swordsman/ts:19-cpu",
	//						Resources: r,
						},
					},
				},
			},
		},
	}

	fmt.Println("that is creating diats deployment....")
	result, err := deploymentclient.Create(deploy)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success create deployment",result.GetObjectMeta().GetName())
}

func int32Ptr2(i int32) *int32 { return &i }

func list_disdeploy(deploymentclient v1beta.DeploymentInterface)  {
	deploy, _ := deploymentclient.List(metav1.ListOptions{})
	for _, i := range deploy.Items {
		fmt.Printf("%s have %d replices", i.Name, *i.Spec.Replicas)
	}
}

func Delete_disdeploy(deploymentclient v1beta.DeploymentInterface){
	deletepolicy := metav1.DeletePropagationForeground
	err := deploymentclient.Delete("tensorflow-cpu",&metav1.DeleteOptions{PropagationPolicy:&deletepolicy})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("delete successful")
	}
}

func watch_disdeploy(deploymentclient v1beta.DeploymentInterface)  {
	w, _ := deploymentclient.Watch(metav1.ListOptions{})
	for {
		select {
			case e := <- w.ResultChan():
				fmt.Println(e.Type,e.Object)
		}
	}
}

func update_disdeploy(deploymentclient v1beta.DeploymentInterface)  {
	result, err := deploymentclient.Get("nginx",metav1.GetOptions{})
	if err != nil {
		panic(err)
	}

	result.Spec.Replicas = int32Ptr2(2)
	deploymentclient.Update(result)
}