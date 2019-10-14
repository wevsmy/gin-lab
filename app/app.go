/*"""
@Version: V1.0
@Author: wevsmy
@License: Apache Licence
@Contact: wevsmy@gmail.com
@Site: blog.weii.ink
@Software: GoLand
@File: app.go.go
@Time: 2019/10/11 下午7:43
*/

package app

import (
	"fmt"
	_ "gin-lab/app/models"
	"gin-lab/app/router"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay/v3"
	_ "github.com/smartwalle/alipay/v3"
	"os/user"
	"path/filepath"
)

// 全局结构体变量
var App struct {
	Config *config
}

// app 应用
func Run() {
	App.Config.Init()
	r := gin.New()
	router.Router(r)
	_ = r.Run(":" + App.Config.Port)
}

func TestAliPay() {
	privateKey := "MIIEowIBAAKCAQEAsLq8J3lNDNYfYeNEkW/t/YYPfEK550yXdHE5xKXZUE3+A0SDG6AsAvaB9s0HFTCJ6sp+ROoZ2x5EeW9/9vN7tccIXCEBVmOWCfnLcUJiCBtQF75OStvx5BEuFCVS6UiaSz3plSZc4eqDMUrM6NBYB17mAjPF2nppl/F9LP4h3kvDneaOTueYaU+kV6RNKg2rNtfbNNwHZxu1baSSQr0g2JVMqbtuvKrUGVuaOtsVxuAZ3ATQzitKMNH8gdQDK9KO7J9u6TNEIdHwc8KVW11nVMSzGS+YitMMQrD2n4VMDm4ryGaHf4oqtZqmwnzgz2mdEEl3CCzYnhVlBJcYBywfZQIDAQABAoIBADRl+Tle7qhaqA1W29KfNBnR8K9v/TyF6fXdSDp0zdzQcvq3CoRbVhE+00PAgFQZAxs5FH0MR8Q+0iCLHY6znD9GFsVyB7p8ZlTo2hfnjbdHmdSgWQg69Bohud7BPjbqjsy2O5Y8PJfC90jbG6v/ccolqd6HSSdA2iPxtqJratFEdz9Hj71jaLYABqphLaf8hDBiWg161uV81u1KLSrs+kqrR7igyol+YiheAa/kkf88lZfZsWai3f1NUGWF2BUH86oZBJ6QRt6uYrzNVkGmcy9NK0+y3+Mbq46BwrRKmAavfzB4pwM7BU0e4ikvOJcOPEA2y+loTjVXUo8gi2mMKqECgYEA/nDLcwztbc5Ku6saYITZlNFDpWhV/rh2MTy4qcO70neQGARyW9mszD1RM4BH+TtRn+fTQGIQWip4tDeTZA6IDYf5LkYAfayTtpUhXuqcJg49Ferv4tb6T3jFonGhs9n/gdKC1UYjkflb1NnIMHgDaTlZgfZnanp7uGZMpGIkKKkCgYEAsdAD3aclDJaA8OqD482HghbQDsA+/YTR6PoPaIcXr/TVljF41JDmqcSqlH/dP2iEXECJtq15pKY/Ynh3bxtkttRgSOP0NyfuCd8M48TC7iwHz2RDYLmi3rmJ7LC5avITtOICviziXflIK52QlGF68uUc4b1AeZe3WGYgDEzUyl0CgYAkB1THtczo/40VheT2RdmJeRhbE6sZpoUV88MyRsURyFxfCkInP2t4gDY/VKrcX8nvGqSPOVOXcOwmmLgGMwiQ4fAm3UK0iPthnzxadF4oBVwg/mN5e3d2SWOy3ORI01WazHQ6PvRKd0TJnwz50ASrobNK89kw+qcKNXIk1MDKAQKBgBRcvuSWLH75iUCNipb+xWLXW/Ikf9ImcKdeY39T4RmMTx1JAw5Mna2ZUPN6hQqq3GV4Go0p5oE9bIrJQtwdZfYt8ezG9gOO9gp5WY+Hy87cifRtBe5As+8PjkTlpAYkPK99JlVC7JVYY7Ri8dicJSlFpX4QXx7Nifh8kXT3I3MdAoGBAOjSJahcQZvxK/aQhvRC4cOeNlN8ata2lw4hgeG9yX4+kaNrMCx8JFK/RpiYY2nZLGG8VvSKQPeKJaWHYUHhatB6KUv26yC+4ZzAy1hhITYZ21S2m3QW4wLCl2GWLHcc1Fpx+oIwKyQVf3KXUQ5V9AimMWTWjnGk20XZFU7qIAQy" // 必须，上一步中使用 RSA签名验签工具 生成的私钥
	client, err := alipay.New("2016082000298417", privateKey, false)

	u, _ := user.Current()
	LoadAliPayPublicCertFromFile := filepath.Join(u.HomeDir, ".GinLabConfig", "alipayCertPublicKey_RSA2.crt")
	LoadAliPayRootCertFromFile := filepath.Join(u.HomeDir, ".GinLabConfig", "alipayRootCert.crt")
	LoadAppPublicCertFromFile := filepath.Join(u.HomeDir, ".GinLabConfig", "appCertPublicKey_2016082000298417.crt")
	client.LoadAppPublicCertFromFile(LoadAppPublicCertFromFile)       // 加载应用公钥证书
	client.LoadAliPayRootCertFromFile(LoadAliPayRootCertFromFile)     // 加载支付宝根证书
	client.LoadAliPayPublicCertFromFile(LoadAliPayPublicCertFromFile) // 加载支付宝公钥证书

	// 将 key 的验证调整到初始化阶段
	if err != nil {
		fmt.Println(err)
		return
	}

	w := alipay.TradeWapPay{}
	w.NotifyURL = "http://xxx"
	w.ReturnURL = "http://xxx"
	w.Subject = "标题"
	w.OutTradeNo = "传递一个唯一单号"
	w.TotalAmount = "10.00"
	w.ProductCode = "QUICK_WAP_WAY"

	url, err := client.TradeWapPay(w)

	if err != nil {
		fmt.Println(err)
	}

	payURL := url.String()
	fmt.Println(payURL)
	// 这个 payURL 即是用于支付的 URL，可将输出的内容复制，到浏览器中访问该 URL 即可打开支付页面

	p := alipay.TradePagePay{}
	p.NotifyURL = "http://xxx"
	p.ReturnURL = "http://xxx"
	p.Subject = "标题"
	p.OutTradeNo = "传递一个唯一单号"
	p.TotalAmount = "10.00"
	p.ProductCode = "QUICK_WAP_WAY"

	p.GoodsDetail = append(p.GoodsDetail, &alipay.GoodsDetail{
		GoodsId: "123",
		AliPayGoodsId:"2",
		GoodsName:"3",
	})
	//=

	url, err = client.TradePagePay(p)

	if err != nil {
		fmt.Println(err)
	}

	payURL = url.String()
	fmt.Println(payURL)
}
