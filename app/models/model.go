/*"""
@Version: V1.0
@Author: wevsmy
@License: Apache Licence
@Contact: wevsmy@gmail.com
@Site: blog.weii.ink
@Software: GoLand
@File: model.go
@Time: 2019/10/11 下午6:20
*/

package models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"time"
)

// 关系型数据库模型,也就是结构体
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               // 忽略本字段
}

func init() {
	// Migrate the schema
	//dbs.db.AutoMigrate(&Product{})
	//
	//// 创建
	//db.Create(&Product{Code: "L1212", Price: 1000})
	//
	//// 读取
	//var product Product
	//db.First(&product, 1)                   // 查询id为1的product
	//db.First(&product, "code = ?", "L1212") // 查询code为l1212的product
	//
	//// 更新 - 更新product的price为2000
	//db.Model(&product).Update("Price", 2000)
	//
	//// 删除 - 删除product
	//db.Delete(&product)
}
