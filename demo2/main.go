package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/owenliang/k8s-client-go/common"
	"io/ioutil"
	appsV1Beta1 "k8s.io/api/apps/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	yaml2 "k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
)

func main() {
	var (
		clientSet  *kubernetes.Clientset
		deployYaml []byte
		deployJson []byte
		deployment = appsV1Beta1.Deployment{}
		replicas   int32
		err        error
	)

	// 初始化k8s客户端
	if clientSet, err = common.InitClient(); err != nil {
		goto FAIL
	}

	// 读取YAML
	if deployYaml, err = ioutil.ReadFile("./nginx.yaml"); err != nil {
		goto FAIL
	}

	// YAML转JSON
	if deployJson, err = yaml2.ToJSON(deployYaml); err != nil {
		goto FAIL
	}

	// JSON转struct
	if err = json.Unmarshal(deployJson, &deployment); err != nil {
		goto FAIL
	}

	// 修改replicas数量为1
	replicas = 1
	deployment.Spec.Replicas = &replicas

	// 查询k8s是否有该deployment
	if _, err = clientSet.AppsV1beta1().Deployments("default").Get(context.TODO(), deployment.Name, metaV1.GetOptions{}); err != nil {
		if !errors.IsNotFound(err) {
			goto FAIL
		}
		// 不存在则创建
		if _, err = clientSet.AppsV1beta1().Deployments("default").Create(context.TODO(), &deployment, metaV1.CreateOptions{}); err != nil {
			goto FAIL
		}
	} else { // 已存在则更新
		if _, err = clientSet.AppsV1beta1().Deployments("default").Update(context.TODO(), &deployment, metaV1.UpdateOptions{}); err != nil {
			goto FAIL
		}
	}

	fmt.Println("apply成功!")
	return

FAIL:
	fmt.Println(err)
	return
}
