package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/guotie/gogb2312"

	iconv "github.com/djimenez/iconv-go"
)

func fetch(url string) []byte {
	fmt.Println("Fetch Url", url)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Http get err:", err)
		return nil
	}
	if resp.StatusCode != 200 {
		fmt.Println("Http status code:", resp.StatusCode)
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read error", err)
		return nil
	}
	return body
}

func main() {
	body := fetch("http://live.500.com")
	fmt.Println(string(body))
	output, err, _, _ := gogb2312.ConvertGB2312(body)
	if err != nil {
		log.Panic("failed to Convert ", err)
	}
	fmt.Println(string(output))
	return
	out := make([]byte, 2*len(body))
	out = out[:]
	iconv.Convert(body, out, "gb2312", "utf-8")
	fmt.Println(string(out))
}
