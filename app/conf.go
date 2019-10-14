/*"""
@Version: V1.0
@Author: wevsmy
@License: Apache Licence
@Contact: wevsmy@gmail.com
@Site: blog.weii.ink
@Software: GoLand
@File: conf.go
@Time: 2019/10/12 上午11:18
*/

package app

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

// app 应用程序配置文件入口
type config struct {
	Host  string `yaml:"host"`
	Port  string `yaml:"port"`
	MySQL mysql  `yaml:"mysql"`
}

// config 接口配置
type configInterface interface {
	Init() // config init 接口
}

// mysql配置参数
type mysql struct {
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	DBName     string `yaml:"db_name"`
	Parameters string `yaml:"parameters"`
}

// config init 默认接口业务逻辑实现
func (c *config) Init() {
	// 默认配置信息
	c = new(config)
	c.Host = "localhost"
	c.Port = "8080"

	// 配置文件读写
	u, _ := user.Current()
	filePath := filepath.Join(u.HomeDir, ".GinLabConfig", "config.yaml")
	//filePath := "./config/config.yaml"
	if err := c.writeConfig(filePath); err != nil {
		log.Fatalf("%s write err: %v", filePath, err)
	}
	if err := c.readConfig(filePath); err != nil {
		log.Fatalf("%s read err: %v", filePath, err)
	}
	// 配置项目运行在herokuApp上的Port
	herokuAppPort := os.Getenv("PORT")
	if herokuAppPort != "" {
		log.Printf("herokuAppPort %s", herokuAppPort)
		c.Port = herokuAppPort
	}
	// 赋值全局Config变量
	App.Config = c
}

// 读取配置文件
func (c *config) readConfig(filePath string) (err error) {
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return err
	}
	return nil
}

// 写入配置文件
// 不存在写入默认，存在从新写入
func (c *config) writeConfig(filePath string) (err error) {
	d, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(filePath, os.O_RDWR, os.ModePerm)
	defer func() { _ = f.Close() }()
	if err != nil && os.IsNotExist(err) {
		fileDir, _ := filepath.Split(filePath)
		err = os.MkdirAll(fileDir, os.ModePerm)
		if err != nil {
			return err
		}
		f, err = os.Create(filePath)
		defer func() { _ = f.Close() }()
		if err != nil {
			return err
		}
	}
	if _, err = f.Write(d); err != nil {
		return err
	}
	err = f.Close()
	return err
}
