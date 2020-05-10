package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rest "k8s.io/client-go/rest"
	"kubedb.dev/apimachinery/apis/kubedb/v1alpha1"
	kubedb "kubedb.dev/apimachinery/client/clientset/versioned/typed/kubedb/v1alpha1"
)

var config *rest.Config
var kubedbClientSet *kubedb.KubedbV1alpha1Client
var err error

func init() {
	config, err = rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	kubedbClientSet, err = kubedb.NewForConfig(config)
}

func main() {
	route := gin.Default()

	route.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong!")
	})

	my := route.Group("/my")
	{
		// 获取指定命名空间的 mysql 实例列表
		my.GET("/:ns", listMysql)

		// 获取指定命名空间，指定名称的 mysql 实例
		my.GET("/:ns/:name", getMysql)

		// 创建 mysql 实例在指定命名空间
		my.POST("/:ns", createMysql)

		// 删除指定命名空间和命名的 mysql 实例
		my.DELETE("/:ns/:name", deleteMysql)
	}

	route.Run()
}

func listMysql(c *gin.Context) {
	ns := c.Param("ns")

	myList, err := kubedbClientSet.MySQLs(ns).List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("have %d instance in ns %s", len(myList.Items), ns),
	})
}

func getMysql(c *gin.Context) {

	ns := c.Param("ns")
	name := c.Param("name")

	mysql, err := kubedbClientSet.MySQLs(ns).Get(name, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("mysql instance %s running on cluster %s", mysql.GetName(), mysql.GetClusterName()),
	})
}

func createMysql(c *gin.Context) {
	ns := c.Param("ns")
	name := "mysql-demo"
	storageClass := "hostpath"
	requests := make(v1.ResourceList)
	requests["storage"] = resource.MustParse("1Gi")

	my := &v1alpha1.MySQL{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: ns,
			Name:      name,
		},
		Spec: v1alpha1.MySQLSpec{
			Version:     "8.0-v2",
			StorageType: v1alpha1.StorageTypeDurable,
			Storage: &v1.PersistentVolumeClaimSpec{
				StorageClassName: &storageClass,
				AccessModes: []v1.PersistentVolumeAccessMode{
					v1.ReadWriteOnce,
				},
				Resources: v1.ResourceRequirements{
					Requests: requests,
				},
			},
			TerminationPolicy: v1alpha1.TerminationPolicyPause,
		},
	}

	createdMy, err := kubedbClientSet.MySQLs(ns).Create(my)
	if err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("create mysql instance %s successfully", createdMy.GetName()),
	})
}

func deleteMysql(c *gin.Context) {
	ns := c.Param("ns")
	name := c.Param("name")
	err := kubedbClientSet.MySQLs(ns).Delete(name, &metav1.DeleteOptions{})
	if err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("delete instance %s successfully", name),
	})
}
