/***************************************************************
 *
 * Copyright (c) 2016, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the BSD licence
 *
 **************************************************************/

/**
 *
 *
 * @file rest_api.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Wed Mar  9 19:08:48 2016
 *
 **/

package memobird

import (
	"encoding/base64"
	"fmt"
	"github.com/doumadou/mahonia"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

//===================================================================
// Public APIs
//===================================================================

const API_URL string = "http://open.memobird.cn/"

type RestApi struct {
	appKey string
	devId  string
}

func NewRestApi(appKey string) *RestApi {
	return &RestApi{appKey: appKey, devId: ""}
}

func Text(text string) string {
	enc := mahonia.NewEncoder("gbk")
	if ret, ok := enc.ConvertStringOK(text); ok {
		return "T:" + base64.StdEncoding.EncodeToString([]byte(ret))
	} else {
		panic(ok)
	}
}

/*
func (this *RestApi) SetUserBind(userId, devId string) Result {
	prefix := API_URL + "home/setuserbind"

}
*/

func (this *RestApi) PrintPaper(devId, content string) Result {
	url := this.getPrefix("/home/printpaper") +
		"&memobirdID=" + devId +
		"&printcontent=" + content
	fmt.Println(url)
	if body, err := this.get(url); err != nil {
		panic(err)
	} else {
		fmt.Println(body)
	}
	return Result{}
}

/*
func (this *RestApi) GetPrintStatus(contentId string) Result {
	prefix := API_URL + "home/getprintstatus"
}
*/

//===================================================================
// Private
//===================================================================

func (this *RestApi) getPrefix(subUrl string) string {
	subUrl = strings.Trim(subUrl, "/")
	t := time.Now()
	timestamp := t.Format("2006-01-02") + "%20" + t.Format("15:04:05")
	return fmt.Sprintf(
		API_URL+subUrl+"?ak=%s&timestamp=%s",
		this.appKey, timestamp)
}

func (this *RestApi) get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, nil
}
