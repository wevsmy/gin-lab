/*"""
@Version: V1.0
@Author: wevsmy
@License: Apache Licence
@Contact: wevsmy@gmail.com
@Site: blog.weii.ink
@Software: GoLand
@File: config.go
@Time: 2019/10/9 下午4:15
*/

package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var Config *config

type config struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

// 解析配置文件
func init() {
	yamlFile, err := ioutil.ReadFile("./conf/config.yaml")
	if err != nil {
		log.Fatalf("conf/config.yaml read err: %v", err)
	}
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		log.Fatalf("yaml.Unmarshal err: %v", err)
	}
	// 默认配置
	if Config.Host == "" {
		Config.Host = "localhost"
	} else if Config.Port == "" {
		Config.Port = "80"
	}
}
