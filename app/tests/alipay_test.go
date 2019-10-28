//@Version: V1.0
//@Author: wevsmy
//@License: Apache Licence
//@Contact: wevsmy@gmail.com
//@Site: blog.weii.ink
//@Software: GoLand
//@File: alipay_test.go
//@Time: 2019/10/15 上午11:28

package tests

import (
	"github.com/smartwalle/alipay/v3"
	"os/user"
	"path/filepath"
	"testing"
)

var client, _ = newClient()

func newClient() (c *alipay.Client, e error) {
	privateKey := "MIIEowIBAAKCAQEAsLq8J3lNDNYfYeNEkW/t/YYPfEK550yXdHE5xKXZUE3+A0SDG6AsAvaB9s0HFTCJ6sp+ROoZ2x5EeW9/9vN7tccIXCEBVmOWCfnLcUJiCBtQF75OStvx5BEuFCVS6UiaSz3plSZc4eqDMUrM6NBYB17mAjPF2nppl/F9LP4h3kvDneaOTueYaU+kV6RNKg2rNtfbNNwHZxu1baSSQr0g2JVMqbtuvKrUGVuaOtsVxuAZ3ATQzitKMNH8gdQDK9KO7J9u6TNEIdHwc8KVW11nVMSzGS+YitMMQrD2n4VMDm4ryGaHf4oqtZqmwnzgz2mdEEl3CCzYnhVlBJcYBywfZQIDAQABAoIBADRl+Tle7qhaqA1W29KfNBnR8K9v/TyF6fXdSDp0zdzQcvq3CoRbVhE+00PAgFQZAxs5FH0MR8Q+0iCLHY6znD9GFsVyB7p8ZlTo2hfnjbdHmdSgWQg69Bohud7BPjbqjsy2O5Y8PJfC90jbG6v/ccolqd6HSSdA2iPxtqJratFEdz9Hj71jaLYABqphLaf8hDBiWg161uV81u1KLSrs+kqrR7igyol+YiheAa/kkf88lZfZsWai3f1NUGWF2BUH86oZBJ6QRt6uYrzNVkGmcy9NK0+y3+Mbq46BwrRKmAavfzB4pwM7BU0e4ikvOJcOPEA2y+loTjVXUo8gi2mMKqECgYEA/nDLcwztbc5Ku6saYITZlNFDpWhV/rh2MTy4qcO70neQGARyW9mszD1RM4BH+TtRn+fTQGIQWip4tDeTZA6IDYf5LkYAfayTtpUhXuqcJg49Ferv4tb6T3jFonGhs9n/gdKC1UYjkflb1NnIMHgDaTlZgfZnanp7uGZMpGIkKKkCgYEAsdAD3aclDJaA8OqD482HghbQDsA+/YTR6PoPaIcXr/TVljF41JDmqcSqlH/dP2iEXECJtq15pKY/Ynh3bxtkttRgSOP0NyfuCd8M48TC7iwHz2RDYLmi3rmJ7LC5avITtOICviziXflIK52QlGF68uUc4b1AeZe3WGYgDEzUyl0CgYAkB1THtczo/40VheT2RdmJeRhbE6sZpoUV88MyRsURyFxfCkInP2t4gDY/VKrcX8nvGqSPOVOXcOwmmLgGMwiQ4fAm3UK0iPthnzxadF4oBVwg/mN5e3d2SWOy3ORI01WazHQ6PvRKd0TJnwz50ASrobNK89kw+qcKNXIk1MDKAQKBgBRcvuSWLH75iUCNipb+xWLXW/Ikf9ImcKdeY39T4RmMTx1JAw5Mna2ZUPN6hQqq3GV4Go0p5oE9bIrJQtwdZfYt8ezG9gOO9gp5WY+Hy87cifRtBe5As+8PjkTlpAYkPK99JlVC7JVYY7Ri8dicJSlFpX4QXx7Nifh8kXT3I3MdAoGBAOjSJahcQZvxK/aQhvRC4cOeNlN8ata2lw4hgeG9yX4+kaNrMCx8JFK/RpiYY2nZLGG8VvSKQPeKJaWHYUHhatB6KUv26yC+4ZzAy1hhITYZ21S2m3QW4wLCl2GWLHcc1Fpx+oIwKyQVf3KXUQ5V9AimMWTWjnGk20XZFU7qIAQy" // 必须，上一步中使用 RSA签名验签工具 生成的私钥
	client, _ := alipay.New("2016082000298417", privateKey, false)

	u, _ := user.Current()
	LoadAliPayPublicCertFromFile := filepath.Join(u.HomeDir, ".GinLabConfig", "alipayCertPublicKey_RSA2.crt")
	LoadAliPayRootCertFromFile := filepath.Join(u.HomeDir, ".GinLabConfig", "alipayRootCert.crt")
	LoadAppPublicCertFromFile := filepath.Join(u.HomeDir, ".GinLabConfig", "appCertPublicKey_2016082000298417.crt")
	client.LoadAppPublicCertFromFile(LoadAppPublicCertFromFile)       // 加载应用公钥证书
	client.LoadAliPayRootCertFromFile(LoadAliPayRootCertFromFile)     // 加载支付宝根证书
	client.LoadAliPayPublicCertFromFile(LoadAliPayPublicCertFromFile) // 加载支付宝公钥证书
	return client, nil
}

func TestClient_TradeWapPay(t *testing.T) {
	t.Log("========== TradeWapPay ==========")
	w := alipay.TradeWapPay{}
	w.NotifyURL = "http://xxx"
	w.ReturnURL = "http://xxx"
	w.Subject = "标题"
	w.OutTradeNo = "trade_no_2017062302112"
	w.TotalAmount = "10.00"
	w.ProductCode = "QUICK_WAP_WAY"

	url, err := client.TradeWapPay(w)
	if err != nil {
		t.FailNow()
	}
	t.Log(url)
	// 这个 payURL 即是用于支付的 URL，可将输出的内容复制，到浏览器中访问该 URL 即可打开支付页面
}

func TestClient_TradeAppPay(t *testing.T) {
	t.Log("========== TradeAppPay ==========")
	var p = alipay.TradeAppPay{}
	p.NotifyURL = "http://203.86.24.181:3000/alipay"
	p.Body = "body"
	p.Subject = "商品标题"
	p.OutTradeNo = "01010101"
	p.TotalAmount = "100.00"
	p.ProductCode = "p_1010101"
	param, err := client.TradeAppPay(p)
	if err != nil {
		t.FailNow()
	}
	t.Log(param)
}

func TestClient_TradePagePay(t *testing.T) {
	t.Log("========== TradePagePay ==========")
	var p = alipay.TradePagePay{}
	p.NotifyURL = "http://220.112.233.229:3000/alipay"
	p.ReturnURL = "http://220.112.233.229:3000"
	p.Subject = "修正了中文的 Bug"
	p.OutTradeNo = "trade_no_20170623011121d1"
	p.TotalAmount = "10.00"
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	p.GoodsDetail = []*alipay.GoodsDetail{&alipay.GoodsDetail{
		GoodsId:   "123",
		GoodsName: "xxx",
		Quantity:  1,
		Price:     13,
	}}

	url, err := client.TradePagePay(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(url)
}

func TestClient_TradePreCreate(t *testing.T) {
	t.Log("========== TradePreCreate ==========")
	var p = alipay.TradePreCreate{}
	p.OutTradeNo = "no_0001"
	p.Subject = "测试订单"
	p.TotalAmount = "10.10"

	rsp, err := client.TradePreCreate(p)
	if err != nil {
		t.Fatal(err)
	}
	if rsp.Content.Code != alipay.K_SUCCESS_CODE {
		t.Fatal(rsp.Content.Msg, rsp.Content.SubMsg)
	}
	t.Log(rsp.Content.QRCode)
}

func TestClient_TradePay(t *testing.T) {
	t.Log("========== Trade ==========")
	var p = alipay.TradePay{}
	p.OutTradeNo = "no_000111"
	p.Subject = "测试订单"
	p.TotalAmount = "10.10"
	p.Scene = "bar_code"
	p.AuthCode = "扫描用户的支付码"

	rsp, err := client.TradePay(p)
	if err != nil {
		t.Fatal(err)
	}
	if rsp.Content.Code != alipay.K_SUCCESS_CODE {
		t.Fatal(rsp.Content.Msg, rsp.Content.SubMsg)
	}
	t.Log(rsp.Content.Msg)
}

func TestClient_TradeRefund(t *testing.T) {
	t.Log("========== TradeRefund ==========")
	var p = alipay.TradeRefund{}
	p.RefundAmount = "10"
	p.OutTradeNo = "trade_no_20170623011121d1"
	rsp, err := client.TradeRefund(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", rsp.Content)
}

func TestClient_FundTransToAccountTransfer(t *testing.T) {
	t.Log("========== FundTransToAccountTransfer ==========")
	var p = alipay.FundTransToAccountTransfer{}
	p.OutBizNo = "xxxx"
	p.PayeeType = "ALIPAY_LOGONID"
	p.PayeeAccount = "xwmkjn7612@sandbox.com"
	p.Amount = "100"
	rsp, err := client.FundTransToAccountTransfer(p)
	if err != nil {
		t.Fatal(err)
	}
	if rsp.Content.Code != alipay.K_SUCCESS_CODE {
		t.Fatal(rsp.Content.Msg, rsp.Content.SubMsg)
	}
	t.Log(rsp.Content.Msg)
}

func TestClient_FundAuthOrderVoucherCreate(t *testing.T) {
	t.Log("========== FundAuthOrderVoucherCreate ==========")
	var p = alipay.FundAuthOrderVoucherCreate{}
	p.OutOrderNo = "1111"
	p.OutRequestNo = "222"
	p.OrderTitle = "eee"
	p.Amount = "1001"
	rsp, err := client.FundAuthOrderVoucherCreate(p)
	if err != nil {
		t.Fatal(err)
	}
	if rsp.Content.Code != alipay.K_SUCCESS_CODE {
		t.Fatal(rsp.Content.Msg, rsp.Content.SubMsg)
	}
	t.Log(rsp.Content.Msg)
}

func TestClient_FundAuthOrderAppFreeze(t *testing.T) {
	t.Log("========== FundAuthOrderAppFreeze ==========")
	var p = alipay.FundAuthOrderAppFreeze{}
	p.OutOrderNo = "111"
	p.OutRequestNo = "xxxxx"
	p.OrderTitle = "test"
	p.Amount = "100"
	p.ProductCode = "PRE_AUTH_ONLINE"

	rsp, err := client.FundAuthOrderAppFreeze(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rsp)
}

func TestClient_UserCertifyOpenInitialize(t *testing.T) {
	t.Log("========== UserCertifyOpenInitialize ==========")
	var p = alipay.UserCertifyOpenInitialize{}
	p.OuterOrderNo = "xxxx"
	p.BizCode = alipay.K_CERTIFY_BIZ_CODE_FACE
	p.IdentityParam.IdentityType = "CERT_INFO"
	p.IdentityParam.CertType = "IDENTITY_CARD"
	p.IdentityParam.CertName = "沙箱环境"
	p.IdentityParam.CertNo = "829297191402263571"
	p.MerchantConfig.ReturnURL = "http://127.0.0.1"
	rsp, err := client.UserCertifyOpenInitialize(p)
	if err != nil {
		t.Fatal(err)
	}
	if rsp.Content.Code != alipay.K_SUCCESS_CODE {
		t.Fatal(rsp.Content.Msg, rsp.Content.SubMsg)
	}
	t.Log(rsp.Content.CertifyId)
}

func TestClient_UserCertifyOpenCertify(t *testing.T) {
	t.Log("========== UserCertifyOpenCertify ==========")
	var p = alipay.UserCertifyOpenCertify{}
	p.CertifyId = "xxxx"
	rsp, err := client.UserCertifyOpenCertify(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rsp)
}

func TestClient_UserCertifyOpenQuery(t *testing.T) {
	t.Log("========== UserCertifyOpenQuery ==========")
	var p = alipay.UserCertifyOpenQuery{}
	p.CertifyId = "xxxx"
	rsp, err := client.UserCertifyOpenQuery(p)
	if err != nil {
		t.Fatal(err)
	}
	if rsp.Content.Code != alipay.K_SUCCESS_CODE {
		t.Fatal(rsp.Content.Msg, rsp.Content.SubMsg)
	}
	t.Log(rsp.Content.Msg)
}

func TestClient_PublicAppAuthorize(t *testing.T) {
	t.Log("========== PublicAppAuthorize ==========")
	var result, err = client.PublicAppAuthorize([]string{"auth_user"}, "http://127.0.0.1", "hhh")
	t.Log(result, err)
}

func TestClient_SystemOauthToken(t *testing.T) {
	t.Log("========== SystemOauthToken ==========")
	var p = alipay.SystemOauthToken{}
	p.GrantType = "authorization_code"
	p.Code = "647f16afe0b44c49a8eb1cb3c02aXX31"
	rsp, err := client.SystemOauthToken(p)
	if err != nil {
		t.Fatal(err)
	}
	if rsp.Content.Code != alipay.K_SUCCESS_CODE {
		t.Fatal(rsp.Content.Msg, rsp.Content.SubMsg)
	}
	t.Log(rsp.Content.UserId, rsp.Content.AccessToken)
}

func TestClient_UserInfoShare(t *testing.T) {
	t.Log("========== UserInfoShare ==========")
	var p = alipay.UserInfoShare{}
	p.AuthToken = "authusrB133e40c363934488a9c3e25e17fd9X31"
	rsp, err := client.UserInfoShare(p)
	if err != nil {
		t.Fatal(err)
	}
	if rsp.Content.Code != alipay.K_SUCCESS_CODE {
		t.Fatal(rsp.Content.Msg, rsp.Content.SubMsg)
	}
	t.Log(rsp.Content.UserId)
}

func TestClient_AppToAppAuth(t *testing.T) {
	t.Log("========== AppToAppAuth ==========")
	var result, err = client.AppToAppAuth("http://127.0.0.1")
	t.Log(result, err)
}

func TestClient_OpenAuthTokenApp(t *testing.T) {
	t.Log("========== OpenAuthTokenApp ==========")
	var p = alipay.OpenAuthTokenApp{}
	p.GrantType = "authorization_code"
	p.Code = "5a14fd7482254120a351109daedbdX31"
	rsp, err := client.OpenAuthTokenApp(p)
	if err != nil {
		t.Fatal(err)
	}
	if rsp.Content.Code != alipay.K_SUCCESS_CODE {
		t.Fatal(rsp.Content.Msg, rsp.Content.SubMsg)
	}
	t.Log(rsp.Content.AppAuthToken, rsp.Content.UserId)
}
