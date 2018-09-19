// Copyright 2018 The ClientSDK Team Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0-style
// license that can be found in the LICENSE file.

// Author: ClientSDK Team (muharihar)

// Package sjwsdk111 is an UnOfficial SOAP client SDK for Sriwijaya Air Web Services (SOAP)
package sjwsdk111

import (
	"fmt"
	"log"
)

// DebugHTTP is function to debug HTTP Request/Response
func DebugHTTP(data []byte, err error) {
	if err == nil {
		fmt.Printf("%s\n\n", data)
	} else {
		log.Fatalf("%s\n\n", err)
	}
}

// DebugHTTPRequest is a function to debug HTTP Request
func DebugHTTPRequest(data []byte, err error) {
	fmt.Println("[DEBUG]:: DebugHTTPRequest")
	fmt.Println("=============================")
	fmt.Println("Request: ")
	fmt.Println("-----------------------------")
	DebugHTTP(data, err)
	fmt.Println("-----------------------------")
	fmt.Println("End-Request: ")
}

// DebugHTTPResponse is a function to debug HTTP Response
func DebugHTTPResponse(data []byte, err error) {
	fmt.Println("[DEBUG]:: DebugHTTPResponse")
	fmt.Println("=============================")
	fmt.Println("Response: ")
	fmt.Println("-----------------------------")
	DebugHTTP(data, err)
	fmt.Println("-----------------------------")
	fmt.Println("End-Response: ")
}
