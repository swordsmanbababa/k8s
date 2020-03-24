package singleNodeTensorflow

import (
	"fmt"
	appsbetav1 "k8s.io/api/apps/v1beta1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1beta "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
	"encoding/json"
)

var cpuPodName string="single-ts-cpu"

var gpuPodName string="single-ts-gpu"

func Create_singlets_deploy (deploymentclient v1beta.DeploymentInterface)  {
	fmt.Println("create_deploy")
	var r apiv1.ResourceRequirements
	j := `{"limits": {"cpu":"4", "memory": "16Gi"}}}`
	json.Unmarshal([]byte(j), &r)
	
	va:=`{"mountPath":"/data/ts","readOnly":"false","name":"datats"}`
	var volumeAmount apiv1.VolumeMount
	json.Unmarshal([]byte(va), &volumeAmount)
	

	deploy := &appsbetav1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: cpuPodName,
		},
		Spec: appsbetav1.DeploymentSpec{
			Replicas: int32Ptr2(1),
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"name": cpuPodName,
					},
				},
				Spec: apiv1.PodSpec{
					Containers:[]apiv1.Container{
							{
								Name:  cpuPodName,
							//	Image: "registry.cn-hangzhou.aliyuncs.com/swordsman/tensorflow:1.9.0-devel-py3",
							  Image: "registry.cn-hangzhou.aliyuncs.com/swordsman/ts:19-cpu",
							//	Image: "tensorflow/tensorflow:1.9.0-py3",
								Ports: []apiv1.ContainerPort{
									{
										Name:          "port",
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
				},
			},
		},
	}

	fmt.Println("that is creating deployment....")
	result, err := deploymentclient.Create(deploy)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success create worker deployment",result.GetObjectMeta().GetName())
}

func Create_Single_ts_GPU_deploy (deploymentclient v1beta.DeploymentInterface)  {
	var podName string="single-ts-gpu"
	fmt.Println("create_gpu_deploy")
	var r apiv1.ResourceRequirements
	j := `{"limits": {"cpu":"4", "memory": "16Gi","alpha.kubernetes.io/nvidia-gpu"="1"}, "requests": {"cpu":"4", "memory": "16Gi","alpha.kubernetes.io/nvidia-gpu"="1"}}`
	json.Unmarshal([]byte(j), &r)
	
	va:=`{"mountPath":"/data/ts","readOnly":"false","name":"datats"}`
	var volumeAmount apiv1.VolumeMount
	json.Unmarshal([]byte(va), &volumeAmount)
	

	deploy := &appsbetav1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: podName,
		},
		Spec: appsbetav1.DeploymentSpec{
			Replicas: int32Ptr2(1),
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"name": podName,
					},
				},
				Spec: apiv1.PodSpec{
					Containers:[]apiv1.Container{
							{
								Name:  podName,
							//	Image: "registry.cn-hangzhou.aliyuncs.com/swordsman/tensorflow:1.9.0-devel-py3",
							//	Image: "registry.cn-hangzhou.aliyuncs.com/swordsman/ts:19-cpu",
							    Image: "registry.cn-hangzhou.aliyuncs.com/swordsman/ts:1.9-gpu",
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
						"scheduler":"gpu",
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
	fmt.Println("Success create worker deployment",result.GetObjectMeta().GetName())
}

func int32Ptr2(i int32) *int32 { return &i }

func Delete_deploy(deploymentclient v1beta.DeploymentInterface){
	deletepolicy := metav1.DeletePropagationForeground
	err := deploymentclient.Delete(cpuPodName,&metav1.DeleteOptions{PropagationPolicy:&deletepolicy})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("delete workers successful")
	}
}
func Delete_Single_GPU_deploy(deploymentclient v1beta.DeploymentInterface){
	deletepolicy := metav1.DeletePropagationForeground
	err := deploymentclient.Delete(gpuPodName,&metav1.DeleteOptions{PropagationPolicy:&deletepolicy})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("delete workers successful")
	}
}



