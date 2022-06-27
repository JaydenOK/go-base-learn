package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

// net/http post demo

func Test_post(t *testing.T) {
	url := "http://127.0.0.1:9090/post"
	// 表单数据
	//contentType := "application/x-www-form-urlencoded"
	//data := "name=枯藤&age=18"
	// json
	contentType := "application/json"
	data := `{"name":"枯藤","age":18}`
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))
}

//server 端接收
//func postHandler(w http.ResponseWriter, r *http.Request) {
//	defer r.Body.Close()
//	// 1. 请求类型是application/x-www-form-urlencoded时解析form数据
//	r.ParseForm()
//	fmt.Println(r.PostForm) // 打印form数据
//	fmt.Println(r.PostForm.Get("name"), r.PostForm.Get("age"))
//	// 2. 请求类型是application/json时从r.Body读取数据
//	b, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		fmt.Println("read request.Body failed, err:%v\n", err)
//		return
//	}
//	fmt.Println(string(b))
//	answer := `{"status": "ok"}`
//	w.Write([]byte(answer))
//}
