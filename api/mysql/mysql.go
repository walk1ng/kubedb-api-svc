package mysql

import (
	"fmt"
	"kubedb-api-svc/cli"
	"net/http"

	"github.com/gin-gonic/gin"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubedb.dev/apimachinery/apis/kubedb/v1alpha1"
)

// ListMysql func list all mysql pod in specified namespace
func ListMysql(c *gin.Context) {
	ns := c.Param("ns")

	myList, err := cli.KubedbClientSet.MySQLs(ns).List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("have %d instance in ns %s", len(myList.Items), ns),
	})
}

// GetMysql func get a mysql pod with name in the specified namespace
func GetMysql(c *gin.Context) {
	ns := c.Param("ns")
	name := c.Param("name")

	mysql, err := cli.KubedbClientSet.MySQLs(ns).Get(name, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("mysql instance %s running on cluster %s", mysql.GetName(), mysql.GetClusterName()),
	})
}

// CreateMysql func create mysql pod in specified namespace
func CreateMysql(c *gin.Context) {
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

	createdMy, err := cli.KubedbClientSet.MySQLs(ns).Create(my)
	if err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("create mysql instance %s successfully", createdMy.GetName()),
	})
}

// DeleteMysql Func delete mysql pod with name in the specified namespace
func DeleteMysql(c *gin.Context) {
	ns := c.Param("ns")
	name := c.Param("name")
	err := cli.KubedbClientSet.MySQLs(ns).Delete(name, &metav1.DeleteOptions{})
	if err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("delete instance %s successfully", name),
	})
}
