package main

import (
	"flag"
	"fmt"
	"github.com/owenliang/k8s-client-go/common"
	"github.com/owenliang/k8s-client-go/demo10/controller"
	"github.com/owenliang/k8s-client-go/demo10/pkg/client/clientSet/versioned"
	"github.com/owenliang/k8s-client-go/demo10/pkg/client/informers/externalversions"
	"github.com/owenliang/k8s-client-go/demo10/pkg/client/informers/externalversions/nginx_controller/v1"
	"k8s.io/client-go/informers"
	coreV1 "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/klog"
	"time"
)

func main() {
	var (
		restConf           *rest.Config
		crdClientSet       *versioned.Clientset
		clientSet          *kubernetes.Clientset
		informerFactory    informers.SharedInformerFactory
		crdInformerFactory externalversions.SharedInformerFactory
		podInformer        coreV1.PodInformer
		nginxInformer      v1.NginxInformer
		nginxController    *controller.NginxController
		err                error
	)

	// 日志参数
	klog.InitFlags(nil)
	flag.Set("logtostderr", "1") // 输出日志到stderr
	flag.Parse()

	// 读取admin.conf, 生成客户端基本配置
	if restConf, err = common.GetRestConf(); err != nil {
		goto FAIL
	}

	// 创建CRD的client
	if crdClientSet, err = versioned.NewForConfig(restConf); err != nil {
		goto FAIL
	}

	// 创建K8S内置的client
	if clientSet, err = kubernetes.NewForConfig(restConf); err != nil {
		goto FAIL
	}

	// 内建informer工厂
	informerFactory = informers.NewSharedInformerFactory(clientSet, time.Second*120)
	// crd Informer工厂
	crdInformerFactory = externalversions.NewSharedInformerFactory(crdClientSet, time.Second*120)

	// POD informer
	podInformer = informerFactory.Core().V1().Pods()
	// nginx informer
	nginxInformer = crdInformerFactory.Mycompany().V1().Nginxes()

	// 创建调度controller
	nginxController = &controller.NginxController{Clientset: clientSet, CrdClientset: crdClientSet, PodInformer: podInformer, NginxInformer: nginxInformer}
	nginxController.Start()

	// 等待
	for {
		time.Sleep(1 * time.Second)
	}

	return

FAIL:
	fmt.Println(err)
	return
}
