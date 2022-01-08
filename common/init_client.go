package common

import (
	"io/ioutil"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// InitClient 初始化k8s客户端
func InitClient() (clientSet *kubernetes.Clientset, err error) {
	var (
		restConf *rest.Config
	)

	if restConf, err = GetRestConf(); err != nil {
		return
	}

	// 生成clientSet配置
	if clientSet, err = kubernetes.NewForConfig(restConf); err != nil {
		goto END
	}
END:
	return
}

// GetRestConf 获取k8s restful client配置
func GetRestConf() (restConf *rest.Config, err error) {
	var (
		kubeConfig []byte
	)

	// 读kubeConfig文件
	if kubeConfig, err = ioutil.ReadFile("./admin.conf"); err != nil {
		goto END
	}
	// 生成rest client配置
	if restConf, err = clientcmd.RESTConfigFromKubeConfig(kubeConfig); err != nil {
		goto END
	}
END:
	return
}
