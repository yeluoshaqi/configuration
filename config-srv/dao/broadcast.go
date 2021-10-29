package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/yeluoshaqi/config/config-srv/model"
)

func broadcastNewestConfig(tx *gorm.DB, appName, clusterName, namespaceName string) error {
	var releaseConfig model.Namespace
	err := tx.Table("namespace").First(&releaseConfig, "app_name = ? and cluster_name = ? and namespace_name = ?",
		appName, clusterName, namespaceName).Error
	if err != nil {
		return err
	}

	return nil
}
