package main

import (
	"fmt"
	"github.com/owenliang/k8s-client-go/common"
	nginxV1 "github.com/owenliang/k8s-client-go/demo9/pkg/apis/nginx_controller/v1"
	"github.com/owenliang/k8s-client-go/demo9/pkg/client/clientset/versioned"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

func main() {
	var (
		restConf     *rest.Config
		crdClientSet *versioned.Clientset
		nginx        *nginxV1.Nginx
		err          error
	)

	// 读取admin.conf, 生成客户端基本配置
	if restConf, err = common.GetRestConf(); err != nil {
		goto FAIL
	}

	// 创建CRD的client
	if crdClientSet, err = versioned.NewForConfig(restConf); err != nil {
		goto FAIL
	}

	// 获取CRD的nginx对象
	if nginx, err = crdClientSet.MycompanyV1().Nginxes("default").Get("my-nginx", v1.GetOptions{}); err != nil {
		goto FAIL
	}

	fmt.Println(nginx)

	return

FAIL:
	fmt.Println(err)
	return
}
