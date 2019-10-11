/*"""
@Version: V1.0
@Author: wevsmy
@License: Apache Licence
@Contact: wevsmy@gmail.com
@Site: blog.weii.ink
@Software: GoLand
@File: db.go
@Time: 2019/10/11 下午6:50
*/

package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func init() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer func() { _ = db.Close() }()
}
