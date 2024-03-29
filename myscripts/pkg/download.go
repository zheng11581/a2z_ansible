package download

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type LoginReq struct {
	BaseReq
	Body LoginBody
}

type BaseReq struct {
	Url    string
	Method string
}

type LoginBody struct {
	Username string
	Password string
}

type LoginResp struct {
	BaseResp
	Data LoginData
}

type BaseResp struct {
	Code    string
	Status  string
	Message string
}

type LoginData struct {
	Opst     string
	OpstData string
}

type Login struct {
	LoginReq
	LoginResp
}

func _main() {
	myJson := []byte(`{
		"code":"200",
		"message":"ok",
		"data":{
			"opst":"ready"
		},
		"status":"success"
	}`)
	var myResp LoginResp
	err := json.Unmarshal(myJson, &myResp)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("respose code: %s\n", myResp.Code)
	fmt.Printf("respose data: %s\n", myResp.Data.Opst)
	fmt.Printf("respose code: %s\n", myResp.Status)
}

func (login *Login) Login() {
	//
	bytesData, err := json.Marshal(login.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	reader := bytes.NewReader(bytesData)
	url := login.Url
	request, err := http.NewRequest(login.Method, url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	cookies := resp.Cookies()
	for _, v := range cookies {
		if v.Name == "opst" {
			login.LoginResp.Data.OpstData = v.Value
			break
		}
	}
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	json.Unmarshal(respBytes, &login.LoginResp)

}

type Download struct {
	DownloadReq
	DownloadResp
}

type DownloadReq struct {
	BaseReq
	FileName string
	SaveAs   string
	Body     DownloadBody
}

type DownloadBody struct {
}

type DownloadResp struct {
	BaseResp
	Data DownloadData
}

type DownloadData struct {
}

func (download *Download) Download() {
	// Request Body
	reqBody, err := json.Marshal(download.Body)
	if err != nil {
		log.Fatalf("读取下载参数失败：%s", err.Error())
	}
	url := download.Url + "?file_name=" + download.FileName
	method := download.Method
	// 创建一个Request
	request, err := http.NewRequest(method, url, bytes.NewReader(reqBody))
	if err != nil {
		log.Fatalf("创建下载请求失败：%s", err.Error())
	}
	// 配置Request的Headers
	request.Header.Set("Content-Type", "application/octet-stream")
	request.Header.Set("Secret", "dkhadhsEJKBG239203#^@*^&()")
	// 创建一个HTTP客户端
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("请求下载失败：%s", err.Error())
	}
	if (response.StatusCode == 403) || (response.StatusCode == 500) || (response.StatusCode == 404) {
		log.Fatalf("请求下载失败：%d", response.StatusCode)
	}

	// 读取Response Body
	respBody, err := io.ReadAll(response.Body)

	// 读取到的内容保存到文件
	file, err := os.OpenFile(download.SaveAs, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("打开下载文件失败：%s", err.Error())
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.Write(respBody)

}
