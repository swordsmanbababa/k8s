package disTensorflow

import (
	"fmt"
	appsbetav1 "k8s.io/api/apps/v1beta1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1beta "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
	"encoding/json"
)

var pspodName string="dist-tensor-ps"

func Create_PS_deploy (deploymentclient v1beta.DeploymentInterface)  {
	fmt.Println("create_deploy")
	var r apiv1.ResourceRequirements
	j := `{"limits": {"cpu":"4", "memory": "16Gi"}}}`
	json.Unmarshal([]byte(j), &r)
	
	va:=`{"mountPath":"/data/ts","readOnly":"false","name":"datats"}`
	var volumeAmount apiv1.VolumeMount
	json.Unmarshal([]byte(va), &volumeAmount)
	

	deploy := &appsbetav1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: pspodName,
		},
		Spec: appsbetav1.DeploymentSpec{
			Replicas: int32Ptr2(1),
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"name": pspodName,
						"role":"ps",
					},
				},
				Spec: apiv1.PodSpec{
					Containers:[]apiv1.Container{
							{
								Name:  "ts-ps",
							//	Image: "registry.cn-hangzhou.aliyuncs.com/swordsman/tensorflow:1.9.0-devel-py3",
							//	Image: "registry.cn-hangzhou.aliyuncs.com/swordsman/ts:19-cpu",
								Image: "tensorflow/tensorflow:1.9.0-py3",
								Ports: []apiv1.ContainerPort{
									{
										Name:          "workerport",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 2222,
								},
								},
								Resources:r,
								VolumeMounts:[]apiv1.VolumeMount{ 
									volumeAmount, 
								},
							},
					},
					Volumes:[]apiv1.Volume{
						{
							Name:"datats",
							VolumeSource:apiv1.VolumeSource{
								NFS:&apiv1.NFSVolumeSource{
									Server:"11.11.11.26",
									Path:"/data/ts",
								},
								
							},
							
						},
					},	
					NodeSelector:map[string]string{
						"kubernetes.io/hostname":"10.10.10.5",
					},		
				},
			},
		},
	}

	fmt.Println("that is creating deployment....")
	result, err := deploymentclient.Create(deploy)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success create ps deployment",result.GetObjectMeta().GetName())
}





func list_PS_deploy(deploymentclient v1beta.DeploymentInterface)  {
	deploy, _ := deploymentclient.List(metav1.ListOptions{})
	for _, i := range deploy.Items {
		fmt.Printf("%s have %d replices", i.Name, *i.Spec.Replicas)
	}
}

func Delete_PS_deploy(deploymentclient v1beta.DeploymentInterface){
	deletepolicy := metav1.DeletePropagationForeground
	err := deploymentclient.Delete(pspodName,&metav1.DeleteOptions{PropagationPolicy:&deletepolicy})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("delete  ps successful")
	}
}

func watch_PS_deploy(deploymentclient v1beta.DeploymentInterface)  {
	w, _ := deploymentclient.Watch(metav1.ListOptions{})
	for {
		select {
			case e := <- w.ResultChan():
				fmt.Println(e.Type,e.Object)
		}
	}
}

func update_PS_deploy(deploymentclient v1beta.DeploymentInterface)  {
	result, err := deploymentclient.Get("nginx",metav1.GetOptions{})
	if err != nil {
		panic(err)
	}

	result.Spec.Replicas = int32Ptr2(2)
	deploymentclient.Update(result)
}