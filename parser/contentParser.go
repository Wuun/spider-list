package parser

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"wutao-ali/util"
)

func ContentParse() {
	form := util.Parse()
	url := form[1][1]
	fmt.Println(url)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("User-Agent", "PostmanRuntime/7.17.1")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Postman-Token", "ca604a0f-c0b6-439b-bfee-850b4dd7ff5a,ad71e099-daea-4c1b-aa77-871e5f9aed80")
	req.Header.Add("Host", "detail.1688.com")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	body, _ := ioutil.ReadAll(res.Body)
	reader := transform.NewReader(bytes.NewReader(body),simplifiedchinese.GBK.NewDecoder())
	result,err := ioutil.ReadAll(reader)
	if err != nil {
		return
	}
	fmt.Println(res)
	fmt.Println(string(result))
}
//func determineEncodings(r io.Reader) []byte {
//	OldReader := bufio.NewReader(r)
//	bytes, err := OldReader.Peek(1024)
//	if err != nil {
//		panic(err)
//	}
//	e, _, _ := charset.DetermineEncoding(bytes, "")
//	reader := transform.NewReader(OldReader, e.NewDecoder())
//	all, err := ioutil.ReadAll(reader)
//	if err != nil {
//		panic(err)
//	}
//	return all
//}


//
//func ConvertToString(src string, srcCode string, tagCode string) string {
//	srcCoder := mahonia.NewDecoder(srcCode)
//	srcResult := srcCoder.ConvertString(src)
//	tagCoder := mahonia.NewDecoder(tagCode)
//	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
//	result := string(cdata)
//	return result
//}
