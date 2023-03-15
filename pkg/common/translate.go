package common

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/thedevsaddam/gojsonq"
)

const YOUDAO_TRANSLATE_URL = "http://fanyi.youdao.com"

func TranslateDescription(sourceContent string) string {
	// 定义url，例如：http://fanyi.youdao.com/translate?&doctype=json&type=EN2ZH_CN&i=hello
	url := YOUDAO_TRANSLATE_URL + "translate?&doctype=json&type=EN2ZH_CN&i=" + sourceContent

	// 获取http请求
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("请求错误", err.Error())
		// 错误返回源内容
		return sourceContent
	}

	// 获取内容
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Body 读取错误", err.Error())
		// 错误返回源内容
		return sourceContent
	}
	// 解析json内容，提取翻译内容，例如：
	// {
	// 	"type": "EN2ZH_CN",
	// 	"errorCode": 0,
	// 	"elapsedTime": 25,
	// 	"translateResult": [
	// 	  [
	// 		{
	// 		  "src": "version is the level this availability applies to",
	// 		  "tgt": "版本是这个可用性适用于水平"
	// 		}
	// 	  ]
	// 	]
	// }
	return gojsonq.New().JSONString(string(data)).Find(".translateResult|.[]|.[].tgt").(string)
}
