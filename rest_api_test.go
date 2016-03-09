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
 * @file rest_api_test.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Wed Mar  9 19:42:38 2016
 *
 **/

package memobird

import (
	"fmt"
	"testing"
)

//===================================================================
// Public APIs
//===================================================================

func TestRestApiPrintPaper(t *testing.T) {
	rest := NewRestApi("appkey")
	fmt.Println(rest.PrintPaper("devid", Text("我的名字是Menglong TAN")))
}

//===================================================================
// Private
//===================================================================
