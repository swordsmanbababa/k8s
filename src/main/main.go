package main



import (
	"fmt"
	apiv1 "k8s.io/api/core/v1"
    "disTensorflow"
    "k8sctl"
    "singleNodeTensorflow"
    "cluster"
    "singletsWithNFS"
)

func main() {
	
//	fmt.Println("1 got....")
//	execCmd()
//	fmt.Println("finish got....")
	//  creat_single_ts()
	//delete_single_ts()
	// creat_single_gpu_ts()
	// delete_single_gpu_ts()
	
	//  creat_single_ts_nfs()
	 //delete_single_ts_nfs()
	//  creat_single_gpu_ts_nfs()
	//delete_single_gpu_ts_nfs()
	
	
	
// create_Cluster()
	//delete_Cluster()
      create_GPU_Cluster()
     //delete_GPU_Cluster()
	
	 //create_GPU_service()
	//delete_GPU_service()
	
	//  create_service()
	//delete_service()
 //   clientset:=k8sctl.GetClientSet()
//	fmt.Println("clientset got....")
//	deploymentclient := clientset.AppsV1beta1().Deployments(apiv1.NamespaceDefault)

	//fmt.Println("clientset 1....")
	//disTensorflow.Create_deploy(deploymentclient)
	//fmt.Println("clientset 2...")
	//disTensorflow.Delete_deploy(deploymentclient)
	//go list_deploy(deploymentclient)
	//go update_deploy(deploymentclient)
	//watch_deploy(deploymentclient)
	
//	disTensorflow.Create_PS_deploy(deploymentclient)
	
	//s:=disTensorflow.Get_disService_interface(clientset)
//	s.Create_service()
	//s.delete_service()
	//s.list_service()
//	s.update_service()
//	s.watch_service()
//
//   s:=disTensorflow.Get_disPsService_interface(clientset)
//	s.Create_Ps_service()
 //  test()
}

func test(){
	 k8sctl.Test()
}

func execCmd(){
	fmt.Println("3 got....")
	var podname string="single-ts-cpu-56d4f4d976-5mh9s"
	var containername string="single-ts-cpu"
	var namesp string="default"
	cluster.ExecCommand(podname,containername, namesp,`"mkdir","/a"`)
//	p:=exec.InitExecOpints()
//	p.Run()
}

func create_GPU_service(){
	clientset:=k8sctl.GetClientSet()
	ws:=disTensorflow.Get_disService_interface(clientset)
	ws.Create_GPU_service()
}
func delete_GPU_service(){
	clientset:=k8sctl.GetClientSet()
	ws:=disTensorflow.Get_disService_interface(clientset)
	ws.Delete_GPU_service()
	 
}
func create_Cluster(){
     clientset:=k8sctl.GetClientSet()
     deploymentclient := clientset.AppsV1beta1().Deployments(apiv1.NamespaceDefault)
     //   disTensorflow.Create_deploy(deploymentclient)
     disTensorflow.Create_PS_deploy(deploymentclient)
     
	
}
func create_GPU_Cluster(){
     clientset:=k8sctl.GetClientSet()
     deploymentclient := clientset.AppsV1beta1().Deployments(apiv1.NamespaceDefault)
	 disTensorflow.Create_gpu_worker1_deploy(deploymentclient)
	 disTensorflow.Create_gpu_worker2_deploy(deploymentclient)
     disTensorflow.Create_PS_deploy(deploymentclient)
	 
}


func create_service(){
	clientset:=k8sctl.GetClientSet()
	ws:=disTensorflow.Get_disService_interface(clientset)
	ws.Create_service()
	ps:=disTensorflow.Get_disPsService_interface(clientset)
	ps.Create_Ps_service()	
}

func delete_service(){
	clientset:=k8sctl.GetClientSet()
	ws:=disTensorflow.Get_disService_interface(clientset)
	ws.Delete_service()
	ps:=disTensorflow.Get_disPsService_interface(clientset)
	ps.Delete_Ps_service()	
}
func delete_Cluster(){
	 clientset:=k8sctl.GetClientSet()
	 deploymentclient := clientset.AppsV1beta1().Deployments(apiv1.NamespaceDefault)
	 disTensorflow.Delete_deploy(deploymentclient)
	 disTensorflow.Delete_PS_deploy(deploymentclient)
    
}

func delete_GPU_Cluster(){
	 clientset:=k8sctl.GetClientSet()
	 deploymentclient := clientset.AppsV1beta1().Deployments(apiv1.NamespaceDefault)
//	 disTensorflow.Delete_deploy(deploymentclient)
	 disTensorflow.Delete_PS_deploy(deploymentclient)
     disTensorflow.Delete_GPU_deploy(deploymentclient)
}
func creat_single_ts(){
	clientset:=k8sctl.GetClientSet()
	 deploymentclient := clientset.AppsV1beta1().Deployments(apiv1.NamespaceDefault)
	 singleNodeTensorflow.Create_singlets_deploy(deploymentclient)
	
}
func delete_single_ts(){
		fmt.Println("3 got....")
	 clientset:=k8sctl.GetClientSet()
	 deploymentclient := clientset.AppsV1beta1().Deployments(apiv1.NamespaceDefault)
	 singleNodeTensorflow.Delete_deploy(deploymentclient)
	 
}

func creat_single_gpu_ts(){
	clientset:=k8sctl.GetClientSet()
	 deploymentclient := clientset.AppsV1beta1().Deployments(apiv1.NamespaceDefault)
	 singleNodeTensorflow.Create_Single_ts_GPU_deploy(deploymentclient)
	
}
func delete_single_gpu_ts(){
		fmt.Println("3 got....")
	 clientset:=k8sctl.GetClientSet()
	 deploymentclient := clientset.AppsV1beta1().Deployments(apiv1.NamespaceDefault)
	 singleNodeTensorflow.Delete_Single_GPU_deploy(deploymentclient)
	 
}

func creat_single_ts_nfs(){
	clientset:=k8sctl.GetClientSet()
	 deploymentclient := clientset.AppsV1beta1().Deployments(apiv1.NamespaceDefault)
	 singletsWithNFS.Create_singlets_deploy(deploymentclient)
	
}
func delete_single_ts_nfs(){
		fmt.Println("3 got....")
	 clientset:=k8sctl.GetClientSet()
	 deploymentclient := clientset.AppsV1beta1().Deployments(apiv1.NamespaceDefault)
	 singletsWithNFS.Delete_deploy(deploymentclient)
	 
}

func creat_single_gpu_ts_nfs(){
	clientset:=k8sctl.GetClientSet()
	 deploymentclient := clientset.AppsV1beta1().Deployments(apiv1.NamespaceDefault)
	 singletsWithNFS.Create_Single_ts_GPU_deploy(deploymentclient)
	
}
func delete_single_gpu_ts_nfs(){
		fmt.Println("3 got....")
	 clientset:=k8sctl.GetClientSet()
	 deploymentclient := clientset.AppsV1beta1().Deployments(apiv1.NamespaceDefault)
	 singletsWithNFS.Delete_Single_GPU_deploy(deploymentclient)
	 
}