package cluster

import (
	"bytes"
	"fmt"
	"strings"
	"io"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"
   "k8s.io/client-go/kubernetes"
       restclient "k8s.io/client-go/rest"
	"k8sctl"
)
type Cluster struct {
	Config           *restclient.Config 
	KubeClient      *kubernetes.Clientset //TODO: move clients to the better place?	
}

func initCluster()*Cluster{
	var c Cluster 
	c.KubeClient,c.Config=k8sctl.GetClientSetAndConfig()
	 return &c
}  


//ExecCommand executes arbitrary command inside the pod
func  ExecCommand_q(podName string, containerName string, namespace string,command ...string) (string, error) {
	//c.setProcessName("executing command %q", strings.Join(command, " "))
    c:=initCluster()
    
  fmt.Println("4 got....")
	var (
		execOut bytes.Buffer
		execErr bytes.Buffer
	)
    pod, err := c.KubeClient.CoreV1().Pods("default").Get(podName, metav1.GetOptions{})
//	pod, err := c.KubeClient.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		return "", fmt.Errorf("could not get pod info: %v", err)
	}
	
     fmt.Println(pod)
	// iterate through all containers looking for the one running PostgreSQL.
	targetContainer := -1
	for _, cr := range pod.Spec.Containers {
		fmt.Println(cr.Name)
		fmt.Println(containerName)
		if cr.Name!=""{
			targetContainer = 1
			break
		}
	}
     fmt.Println(targetContainer)
	if targetContainer < 0 {
		return "", fmt.Errorf("could not find container to exec to")
	}
  //  pods,err := c.KubeClient.CoreV1().Pods("default").List(metav1.ListOptions{})
  
    fmt.Println("req")
	req :=  c.KubeClient.CoreV1().RESTClient().Post().
		Resource("pods").
		Name("single-ts-cpu-56d4f4d976-5mh9s").
		Namespace("default").
		SubResource("exec").
		VersionedParams(&v1.PodExecOptions{
				Container: containerName,
				Command:   []string{"bash"},
				Stdin:     true,
				Stdout:    true,
				Stderr:    true,
				TTY:       false,
			}, scheme.ParameterCodec)
 
 
			req.SetHeader("Upgrade","Websocket")
			req.SetHeader("Connection","Upgrade")
			req.SetHeader("Authorization","Bearer <ti29bq.levzeldzak5c9r9u>")
			req.SetHeader("Accept","application/json")
			req.SetHeader("Sec-WebSocket-Key","111")
			req.SetHeader("Sec-WebSocket-Version","13")
	
    fmt.Println("req.VersionedParam end")
    fmt.Println(req)
 
    
    
    
	exec, err := remotecommand.NewSPDYExecutor(c.Config, "POST", req.URL())
	if err != nil {
		return "", fmt.Errorf("failed to init executor: %v", err)
	}
   fmt.Println("exec.Stream(remotecommand.StreamOptions{")
    
    

	err = exec.Stream(remotecommand.StreamOptions{
		Stdout: &execOut,
		Stderr: &execErr,
		Tty:    false,
	})


  fmt.Println(err)
   
  fmt.Println("exec.Stream(remotecommand.StreamOptions end")
        fmt.Println( execErr.String())
       fmt.Println("stderr: %v",execOut.String())
         fmt.Println(err)
	if err != nil {
		return "", fmt.Errorf("could not execute: %v", err)
	}

	if execErr.Len() > 0 {
		return "", fmt.Errorf("stderr: %v", execErr.String())
	}
       fmt.Println("stderr: %v", execErr.String())
       fmt.Println("stderr: %v",execOut.String())
      fmt.Println("stderr:")

	return execOut.String(), nil
}




type Writer struct {
	Str []string
}

func (w *Writer) Write(p []byte) (n int, err error) {
	str := string(p)
	if len(str) > 0 {
		w.Str = append(w.Str, str)
	}
	return len(str), nil
}

func newStringReader(ss []string) io.Reader {
	formattedString := strings.Join(ss, "\n")
	reader := strings.NewReader(formattedString)
	return reader
}



func  ExecCommand(podName string, containerName string, namespace string,command string) (string, error) {
	KubeClient,config:=k8sctl.GetClientSetAndConfig()
	execRequest := KubeClient.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(podName).
		Namespace(namespace).
		SubResource("exec").
		Param("container",containerName).
		Param("command", command).
		Param("stdin", "true").
		Param("stdout", "false").
		Param("stderr", "false").
		Param("tty", "false")

	exec, err := remotecommand.NewSPDYExecutor(config, "POST", execRequest.URL())
	
    fmt.Println(err)
    
	stdIn := newStringReader([]string{"-c"})
	stdOut := new(Writer)
	stdErr := new(Writer)

	err = exec.Stream(remotecommand.StreamOptions{
		Stdin:  stdIn,
		Stdout: stdOut,
		Stderr: stdErr,
		Tty:    false,
	})

	fmt.Println(err)

	output := fmt.Sprintf("Exit Code: %v",err )
	

	return output,err
}
