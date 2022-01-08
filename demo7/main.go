package main

import (
	"context"
	"fmt"
	"github.com/owenliang/k8s-client-go/common"
	coreV1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	var (
		clientSet *kubernetes.Clientset
		tailLines int64
		req       *rest.Request
		res       rest.Result
		logs      []byte
		err       error
	)

	// 初始化k8s客户端
	if clientSet, err = common.InitClient(); err != nil {
		goto FAIL
	}

	// 生成获取POD日志请求
	req = clientSet.CoreV1().Pods("default").GetLogs("nginx-deployment-5c689d88bb-sm9nl", &coreV1.PodLogOptions{Container: "nginx", TailLines: &tailLines})

	// req.Stream()也可以实现Do的效果

	// 发送请求
	if res = req.Do(context.TODO()); res.Error() != nil {
		err = res.Error()
		goto FAIL
	}

	// 获取结果
	if logs, err = res.Raw(); err != nil {
		goto FAIL
	}

	fmt.Println("容器输出:", string(logs))
	return

FAIL:
	fmt.Println(err)
	return
}
