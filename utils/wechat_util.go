package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"hmshop/global"
	"io/ioutil"
	"net/http"
)

func GetOpenID(code string) (string, error) {

	fmt.Println(code)
	//调用微信接口服务，获取返回数据
	// 构造请求 URL，微信登录验证接口
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		global.AppConfig.Wechat.AppId, global.AppConfig.Wechat.Secret, code)
	// 使用 Go 自带的 net/http 发起 HTTP 请求
	resp, err := http.Get(url)
	if err != nil {
		global.Log.Error(err)
		return "", errors.New("failed to connect to WeChat API")
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		global.Log.Error(err)
		return "", errors.New("failed to read WeChat API response")
	}
	// 检查微信返回的 JSON 数据
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		global.Log.Error(err)
		return "", errors.New("failed to parse WeChat response")
	}
	fmt.Println(result, "-----------")
	// 微信返回的数据中包含 openid 和 session_key
	openid, ok := result["openid"].(string)
	if !ok {
		global.Log.Error(errors.New("openid not found"))
		return "", errors.New("openid not found")
	}
	return openid, nil
}
