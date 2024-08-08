package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func Test_http(t *testing.T) {
	http.HandleFunc("/", readBodyOnce)
	fmt.Println("http://127.0.0.1:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		return
	}
}

func readBodyOnce(w http.ResponseWriter, r *http.Request) {
	if data, err := json.Marshal(r.Host); err != nil {
		return
	} else {
		w.Write(data)
	}
}

func Test_http_client(t *testing.T) {
	url := "http://192.168.112.242:8000/api/ocr?engine=rp&type=1"
	method := "POST"
	filename := "F:\\goProject\\renewstudy\\web\\1.jpg"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open(filepath.Base(filename))
	defer file.Close()
	part1, errFile1 := writer.CreateFormFile("file", filepath.Base(filename))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		fmt.Println(errFile1)
		return
	}
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer tjxnbo2l562v6qlbvx8wbsn")
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "192.168.112.242:8000")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Type", "multipart/form-data")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
