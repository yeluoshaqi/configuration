package config

import (
	"admin-conf/config-srv/dao"
)


func CreateApp(appName, description string) (interface{}, error) {
	app, err := dao.GetDao().CreateApp(appName, description)
	if err != nil {
		//log.Error("[CreateApp]", err)
		return nil, err
	}

	return app, nil
}

func QueryApp(appName string) (interface{}, error) {
	app, err := dao.GetDao().QueryApp(appName)
	if err != nil {
		//log.Error("[QueryApp]", err)
		return nil, err
	}

	return app, nil
}

func DeleteApp(appName string) error {
	err := dao.GetDao().DeleteApp(appName)
	if err != nil {
		//log.Error("[DeleteApp] delete app:%s error: %s", req.GetAppName(), err.Error())
	}
	return nil
}

func ListApps() (interface{}, error) {
	apps, err := dao.GetDao().ListApps()
	if err != nil {
		//log.Error("[ListApps]", err)
		return nil, err
	}

	return apps, nil
}

func CreateCluster(appName, clusterName, description string) (interface{}, error) {
	cluster, err := dao.GetDao().CreateCluster(appName, clusterName, description)
	if err != nil {
		//log.Error("[CreateCluster]", err)
		return nil, err
	}
	return cluster, nil
}

func QueryCluster(appName, clusterName string) (interface{}, error) {
	cluster, err := dao.GetDao().QueryCluster(appName, clusterName)
	if err != nil {
		//log.Error("[QueryCluster]", err)
		return nil, err
	}

	return cluster, nil
}

func DeleteCluster(appName, clusterName string) error {
	err := dao.GetDao().DeleteCluster(appName, clusterName)
	if err != nil {
		//log.Error("[DeleteCluster] delete cluster:%s-%s error: %s", req.GetAppName(), req.GetClusterName(), err.Error())
	}
	return nil
}

func ListClusters(appName string) (interface{}, error) {
	clusters, err := dao.GetDao().ListClusters(appName)
	if err != nil {
		//log.Error("[ListClusters]", err)
		return nil, err
	}
	return clusters, nil
}

func CreateNamespace(appName, clusterName, namespaceName, format, description string) (interface{}, error) {
	namespace, err := dao.GetDao().CreateNamespace(
		appName,
		clusterName,
		namespaceName,
		format,
		description)
	if err != nil {
		//log.Error("[CreateNamespace]", err)
		return nil, err
	}
	return namespace, nil
}

func QueryNamespace(appName, clusterName, namespaceName string) (interface{}, error) {
	namespace, err := dao.GetDao().QueryNamespace(appName, clusterName, namespaceName)
	if err != nil {
		//log.Error("[QueryNamespace]", err)
		return nil, err
	}
	return namespace, nil
}

func DeleteNamespace(appName, clusterName, namespaceName string) error {
	err := dao.GetDao().DeleteNamespace(appName, clusterName, namespaceName)
	if err != nil {
		//log.Error("[DeleteNamespace] delete namespace:%s-%s-%s error: %s", req.GetAppName(), req.GetClusterName(), req.GetNamespaceName(), err.Error())
	}
	return nil
}

func ListNamespaces(appName, clusterName string) (interface{}, error) {
	namespaces, err := dao.GetDao().ListNamespaces(appName, clusterName)
	if err != nil {
		//log.Error("[ListNamespaces]", err)
		return nil, err
	}
	return namespaces, nil
}

func UpdateConfig(appName, clusterName, namespaceName, value string) error {
	return dao.GetDao().UpdateConfig(appName, clusterName, namespaceName, value)
}

func ReleaseConfig(appName, clusterName, namespaceName, tag, comment string) error {
	return nil
}

func ListReleaseHistory(appName, clusterName, namespaceName string) (interface{}, error) {
	return  nil, nil
}

func Rollback(appName, clusterName, namespaceName, tag string) (err error) {

	return nil
}
