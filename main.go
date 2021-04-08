/*
 * @Descripttion:
 * @Author: Zhang YaSong
 * @version:
 * @Date: 2021-04-02 16:42:58
 * @LastEditors: Zhang YaSong
 * @LastEditTime: 2021-04-08 16:24:49
 */
package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", UserFileDownloadCommonService)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func UserFileDownloadCommonService(c *gin.Context) {
	url := "https://api.sinagin.com/bgfile/docs/202104/1/file-文件/6c324a53-3dc2-4578-a583-1788929d175e.xlsx"
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bts, _ := ioutil.ReadAll(resp.Body)
	read := bytes.NewReader(bts)

	fileName := "1.xlsx"
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Cache-Control", "no-cache")

	read.WriteTo(c.Writer)
}
