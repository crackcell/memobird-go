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
 * @file memobird.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Wed Mar  9 19:19:09 2016
 *
 **/

package memobird

import (
//"fmt"
)

//===================================================================
// Public APIs
//===================================================================

type Result struct {
	ReturnCode    int
	ReturnMessage string
	Data          []string
}

//===================================================================
// Private
//===================================================================
