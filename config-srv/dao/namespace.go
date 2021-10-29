package dao

import (
	"errors"

	"github.com/yeluoshaqi/config/config-srv/model"
)

func (d *Dao) NamespaceExist(appName, clusterName, namespaceName string) bool {
	return !d.client.Table("namespace").Where("app_name = ? and cluster_name = ? and namespace_name = ?",
		appName, clusterName, namespaceName).First(&model.Namespace{}).RecordNotFound()
}

func (d *Dao) CreateNamespace(appName, clusterName, namespaceName, format, description string) (*model.Namespace, error) {
	if !d.ClusterExist(appName, clusterName) {
		return nil, errors.New("cluster not found")
	}

	namespace := &model.Namespace{
		AppName:       appName,
		ClusterName:   clusterName,
		NamespaceName: namespaceName,
		Format:        format,
		Description:   description,
	}
	err := d.client.Table("namespace").Create(namespace).Error

	return namespace, err
}

func (d *Dao) QueryNamespace(appName, clusterName, namespaceName string) (namespace model.Namespace, err error) {
	err = d.client.Table("namespace").First(&namespace, "app_name = ? and cluster_name = ? and namespace_name = ?",
		appName, clusterName, namespaceName).Error
	return
}

func (d *Dao) DeleteNamespace(appName, clusterName, namespaceName string) error {
	if err := d.client.Table("namespace").Unscoped().Delete(model.Namespace{}, "app_name = ? and cluster_name = ? and namespace_name = ?",
		appName, clusterName, namespaceName).Error; err != nil {
		return err
	}

	return nil
}

func (d *Dao) ListNamespaces(appName, clusterName string) (namespaces []*model.Namespace, err error) {
	err = d.client.Table("namespace").Find(&namespaces, "app_name = ? and cluster_name = ?", appName, clusterName).Error
	return
}
