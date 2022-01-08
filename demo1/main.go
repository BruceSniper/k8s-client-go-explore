package main

import (
	"context"
	"fmt"
	"github.com/owenliang/k8s-client-go/common"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func main() {
	var (
		clientSet *kubernetes.Clientset
		podsList  *coreV1.PodList
		err       error
	)

	// 初始化k8s客户端
	if clientSet, err = common.InitClient(); err != nil {
		goto FAIL
	}

	// 获取default命名空间下的所有POD
	if podsList, err = clientSet.CoreV1().Pods("default").List(context.TODO(), metaV1.ListOptions{}); err != nil {
		goto FAIL
	}
	fmt.Println(*podsList)

	return

FAIL:
	fmt.Println(err)
	return
}
