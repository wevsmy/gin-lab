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
	"os"
	"path"
)

var Config *config

type config struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

// 解析配置文件
func init() {
	filePath := "./configs/config.yaml"
	err := createFileWithDir(filePath)
	if err != nil {
		log.Fatalf("filePath: %s create err: %v", filePath, err)
	}

	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("%s read err: %v", filePath, err)
	}
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		log.Fatalf("yaml.Unmarshal err: %v", err)
	}

	// 配置项目运行在herokuApp上的Port
	herokuAppPort := os.Getenv("PORT")
	if herokuAppPort != "" {
		log.Printf("herokuAppPort %s", herokuAppPort)
		Config.Port = herokuAppPort
	}
}

// 创建默认配置文件
func createDefaultConfig(filePath string) (err error) {

	// 默认配置信息
	c := config{
		Host: "localhost",
		Port: "8080",
	}

	d, err := yaml.Marshal(&c)
	if err != nil {
		return err
	}
	f, err := os.Create(filePath)
	defer func() { _ = f.Close() }()
	if err != nil {
		return err
	}
	if _, err = f.Write(d); err != nil {
		return err
	}
	err = f.Close()
	return err
}

// 创建带路径的文件,不存在创建,存在不做任何操作
func createFileWithDir(filePath string) (err error) {
	f, err := os.Open(filePath)
	defer func() { _ = f.Close() }()
	if err != nil && os.IsNotExist(err) {
		fileDir, _ := path.Split(filePath)
		err = os.MkdirAll(fileDir, os.ModePerm)
		if err != nil {
			return err
		}
		err = createDefaultConfig(filePath)
		if err != nil {
			return err
		}
	}
	err = f.Close()
	return err
}
