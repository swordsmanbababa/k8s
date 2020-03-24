package disTensorflow

import (
	"fmt"
	appsbetav1 "k8s.io/api/apps/v1beta1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1beta "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
	"encoding/json"
)

var podName string="dist-tensor-worker"

func Create_gpu_worker1_deploy (deploymentclient v1beta.DeploymentInterface)  {
	var podName string="dist-gpu-tensor-worker"
	fmt.Println("create_deploy")
	var r apiv1.ResourceRequirements
	//j := `{"limits": {"cpu":"4", "memory": "16Gi"}}}`
	j := `{"limits": {"cpu":"4", "memory": "16Gi","alpha.kubernetes.io/nvidia-gpu"="1"}, "requests": {"cpu":"4", "memory": "16Gi","alpha.kubernetes.io/nvidia-gpu"="1"}}`
	json.Unmarshal([]byte(j), &r)
	
	va:=`{"mountPath":"/data/ts","readOnly":"false","name":"datats"}`
	var volumeAmount apiv1.VolumeMount
	json.Unmarshal([]byte(va), &volumeAmount)
	

	deploy := &appsbetav1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: podName+"1",
		},
		Spec: appsbetav1.DeploymentSpec{
			Replicas: int32Ptr2(1),
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"name": podName,
						"role":"worker",
					},
				},
				Spec: apiv1.PodSpec{
					Containers:[]apiv1.Container{
							{
								Name:  "ts-worker",
								Image: "registry.cn-hangzhou.aliyuncs.com/swordsman/tensorflow:1.9.0-gpu-py3",
							//	Image: "registry.cn-hangzhou.aliyuncs.com/swordsman/tensorflow:1.9.0-devel-py3",
							//	Image: "registry.cn-hangzhou.aliyuncs.com/swordsman/ts:19-cpu",
							//    Image: "tensorflow/tensorflow:1.9.0-py3",
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

func Create_gpu_worker2_deploy (deploymentclient v1beta.DeploymentInterface)  {
	var podName string="dist-gpu-tensor-worker"
	fmt.Println("create_deploy")
	var r apiv1.ResourceRequirements
//	j := `{"limits": {"cpu":"4", "memory": "16Gi"}}}`
	j := `{"limits": {"cpu":"4", "memory": "16Gi","alpha.kubernetes.io/nvidia-gpu"="1"}, "requests": {"cpu":"4", "memory": "16Gi","alpha.kubernetes.io/nvidia-gpu"="1"}}`
	json.Unmarshal([]byte(j), &r)
	
	va:=`{"mountPath":"/data/ts","readOnly":"false","name":"datats"}`
	var volumeAmount apiv1.VolumeMount
	json.Unmarshal([]byte(va), &volumeAmount)
	

	deploy := &appsbetav1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: podName+"2",
		},
		Spec: appsbetav1.DeploymentSpec{
			Replicas: int32Ptr2(1),
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"name": podName,
						"role":"worker",
					},
				},
				Spec: apiv1.PodSpec{
					Containers:[]apiv1.Container{
							{
								Name:  "ts-worker",
							 //   Image: "tensorflow/tensorflow:1.9.0-py3",
							    Image: "registry.cn-hangzhou.aliyuncs.com/swordsman/tensorflow:1.9.0-gpu-py3",
							//	Image: "registry.cn-hangzhou.aliyuncs.com/swordsman/tensorflow:1.9.0-devel-py3",
							//	Image: "registry.cn-hangzhou.aliyuncs.com/swordsman/ts:19-cpu",
							//    Image: "tensorflow/tensorflow:1.9.0-py3",
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

func Create_GPU_deploy (deploymentclient v1beta.DeploymentInterface)  {
	var podName string="dist-gpu-tensor-worker"
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
						"role":"worker",
					},
				},
				Spec: apiv1.PodSpec{
					Containers:[]apiv1.Container{
							{
								Name:  "ts-worker",
								Image: "registry.cn-hangzhou.aliyuncs.com/swordsman/tensorflow:1.9.0-gpu-py3",
							//	Image: "registry.cn-hangzhou.aliyuncs.com/swordsman/ts:19-cpu",
							//    Image: "registry.cn-hangzhou.aliyuncs.com/swordsman/ts:1.9-gpu",
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
	fmt.Println("Success create gpu_worker deployment",result.GetObjectMeta().GetName())
}

func int32Ptr2(i int32) *int32 { return &i }

func list_deploy(deploymentclient v1beta.DeploymentInterface)  {
	
	deploy, _ := deploymentclient.List(metav1.ListOptions{})
	for _, i := range deploy.Items {
		fmt.Printf("%s have %d replices", i.Name, *i.Spec.Replicas)
	}
}

func Delete_deploy(deploymentclient v1beta.DeploymentInterface){
	deletepolicy := metav1.DeletePropagationForeground
	err := deploymentclient.Delete(podName,&metav1.DeleteOptions{PropagationPolicy:&deletepolicy})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("delete workers successful")
	}
}
func Delete_GPU_deploy(deploymentclient v1beta.DeploymentInterface){
	var podName string="dist-gpu-tensor-worker"
	deletepolicy := metav1.DeletePropagationForeground
	err := deploymentclient.Delete(podName+"1",&metav1.DeleteOptions{PropagationPolicy:&deletepolicy})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("delete gpu workers1 successful")
	}
	err2 := deploymentclient.Delete(podName+"2",&metav1.DeleteOptions{PropagationPolicy:&deletepolicy})
	if err2 != nil {
		panic(err2)
	} else {
		fmt.Println("delete gpu worker2 successful")
	}
}

func watch_deploy(deploymentclient v1beta.DeploymentInterface)  {
	w, _ := deploymentclient.Watch(metav1.ListOptions{})
	for {
		select {
			case e := <- w.ResultChan():
				fmt.Println(e.Type,e.Object)
		}
	}
}

func update_deploy(deploymentclient v1beta.DeploymentInterface,theUpdatedName string)  {
	result, err := deploymentclient.Get(theUpdatedName,metav1.GetOptions{})
	if err != nil {
		panic(err)
	}

	result.Spec.Replicas = int32Ptr2(2)
	deploymentclient.Update(result)
}