//@Version: V1.0
//@Author: wevsmy
//@License: Apache Licence
//@Contact: wevsmy@gmail.com
//@Site: blog.weii.ink
//@Software: GoLand
//@File: pay.go
//@Time: 2019/10/15 上午9:47

package utils

import (
	"github.com/smartwalle/alipay/v3"
	"os/user"
	"path/filepath"
)

type PayClient struct {
	AliPay *alipay.Client
	alipay.TradeWapPay
}

func (c *PayClient) Init() {
	c = new(PayClient)
	privateKey := "MIIEowIBAAKCAQEAsLq8J3lNDNYfYeNEkW/t/YYPfEK550yXdHE5xKXZUE3+A0SDG6AsAvaB9s0HFTCJ6sp+ROoZ2x5EeW9/9vN7tccIXCEBVmOWCfnLcUJiCBtQF75OStvx5BEuFCVS6UiaSz3plSZc4eqDMUrM6NBYB17mAjPF2nppl/F9LP4h3kvDneaOTueYaU+kV6RNKg2rNtfbNNwHZxu1baSSQr0g2JVMqbtuvKrUGVuaOtsVxuAZ3ATQzitKMNH8gdQDK9KO7J9u6TNEIdHwc8KVW11nVMSzGS+YitMMQrD2n4VMDm4ryGaHf4oqtZqmwnzgz2mdEEl3CCzYnhVlBJcYBywfZQIDAQABAoIBADRl+Tle7qhaqA1W29KfNBnR8K9v/TyF6fXdSDp0zdzQcvq3CoRbVhE+00PAgFQZAxs5FH0MR8Q+0iCLHY6znD9GFsVyB7p8ZlTo2hfnjbdHmdSgWQg69Bohud7BPjbqjsy2O5Y8PJfC90jbG6v/ccolqd6HSSdA2iPxtqJratFEdz9Hj71jaLYABqphLaf8hDBiWg161uV81u1KLSrs+kqrR7igyol+YiheAa/kkf88lZfZsWai3f1NUGWF2BUH86oZBJ6QRt6uYrzNVkGmcy9NK0+y3+Mbq46BwrRKmAavfzB4pwM7BU0e4ikvOJcOPEA2y+loTjVXUo8gi2mMKqECgYEA/nDLcwztbc5Ku6saYITZlNFDpWhV/rh2MTy4qcO70neQGARyW9mszD1RM4BH+TtRn+fTQGIQWip4tDeTZA6IDYf5LkYAfayTtpUhXuqcJg49Ferv4tb6T3jFonGhs9n/gdKC1UYjkflb1NnIMHgDaTlZgfZnanp7uGZMpGIkKKkCgYEAsdAD3aclDJaA8OqD482HghbQDsA+/YTR6PoPaIcXr/TVljF41JDmqcSqlH/dP2iEXECJtq15pKY/Ynh3bxtkttRgSOP0NyfuCd8M48TC7iwHz2RDYLmi3rmJ7LC5avITtOICviziXflIK52QlGF68uUc4b1AeZe3WGYgDEzUyl0CgYAkB1THtczo/40VheT2RdmJeRhbE6sZpoUV88MyRsURyFxfCkInP2t4gDY/VKrcX8nvGqSPOVOXcOwmmLgGMwiQ4fAm3UK0iPthnzxadF4oBVwg/mN5e3d2SWOy3ORI01WazHQ6PvRKd0TJnwz50ASrobNK89kw+qcKNXIk1MDKAQKBgBRcvuSWLH75iUCNipb+xWLXW/Ikf9ImcKdeY39T4RmMTx1JAw5Mna2ZUPN6hQqq3GV4Go0p5oE9bIrJQtwdZfYt8ezG9gOO9gp5WY+Hy87cifRtBe5As+8PjkTlpAYkPK99JlVC7JVYY7Ri8dicJSlFpX4QXx7Nifh8kXT3I3MdAoGBAOjSJahcQZvxK/aQhvRC4cOeNlN8ata2lw4hgeG9yX4+kaNrMCx8JFK/RpiYY2nZLGG8VvSKQPeKJaWHYUHhatB6KUv26yC+4ZzAy1hhITYZ21S2m3QW4wLCl2GWLHcc1Fpx+oIwKyQVf3KXUQ5V9AimMWTWjnGk20XZFU7qIAQy" // 必须，上一步中使用 RSA签名验签工具 生成的私钥
	client, _ := alipay.New("2016082000298417", privateKey, false)

	u, _ := user.Current()
	LoadAliPayPublicCertFromFile := filepath.Join(u.HomeDir, ".GinLabConfig", "alipayCertPublicKey_RSA2.crt")
	LoadAliPayRootCertFromFile := filepath.Join(u.HomeDir, ".GinLabConfig", "alipayRootCert.crt")
	LoadAppPublicCertFromFile := filepath.Join(u.HomeDir, ".GinLabConfig", "appCertPublicKey_2016082000298417.crt")
	client.LoadAppPublicCertFromFile(LoadAppPublicCertFromFile)       // 加载应用公钥证书
	client.LoadAliPayRootCertFromFile(LoadAliPayRootCertFromFile)     // 加载支付宝根证书
	client.LoadAliPayPublicCertFromFile(LoadAliPayPublicCertFromFile) // 加载支付宝公钥证书
	c.AliPay = client
}

var client *alipay.Client
