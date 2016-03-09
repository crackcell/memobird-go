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
//"fmt"
)

//===================================================================
// Public APIs
//===================================================================

type RestApi struct {
	appKey string
	devId  string
}

func NewRestApi(appKey string) *RestApi {
	return &RestApi{appKey: appKey, devId: ""}
}

//===================================================================
// Private
//===================================================================

func (this *RestApi) SetUserBind(userId, devId string) Result {}

func (this *RestApi) PrintPaper(content, devId string) Result {}

func (this *RestApi) GetPrintStatus(contentId string) Result {}
