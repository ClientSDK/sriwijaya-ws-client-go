// Copyright 2018 The ClientSDK Team Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0-style
// license that can be found in the LICENSE file.

// Author: ClientSDK Team (muharihar)
package main

import (
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"

	"github.com/ClientSDK/sriwijaya-ws-client-go/sjwsdk111"
)

func makeHTTPClient() *http.Client {
	// Access via proxy if needed
	proxyURL, _ := url.Parse("http://proxy-ip-address:proxy-port")
	//proxyURL, _ := url.Parse("http://proxy-user:proxy-password@proxy-ip-address:proxy-port")

	// Initite transport with proxy and skip TLS (if needed)
	tr := &http.Transport{
		Proxy:           http.ProxyURL(proxyURL),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// Initiate transport without proxy and skip TLS (if needed)
	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }

	httpClient := &http.Client{Transport: tr}

	return httpClient
}

func main() {

	// Initiate http client
	httpClient := makeHTTPClient()

	// Initiate NewSoapSJClient version 111
	sjClient, err := sjwsdk111.NewSoapSJClient(httpClient, "../../wsdl/wsp-wsdl.eticketv111.wsdl", "file")
	if err != nil {
		fmt.Println(err)
	}

	// call Sriwijaya web service operation
	callWsRouteOperate(sjClient)
}

// callWsRouteOperate is a function to call WsRouteOperate method
func callWsRouteOperate(s *sjwsdk111.SoapSJClient) {
	params := []byte(
		`
			<Username xsi:type="xsd:string">SRIWIJAWA_AGENT_USERNAME</Username>
			<Password xsi:type="xsd:string">SRIWIJAWA_AGENT_PASSWORD</Password>
			`)
	wsResp, errC := s.CallWsRouteOperate(params, false)

	if errC != nil {
		fmt.Println(errC)
		return
	}

	// Access response variable
	// fmt.Println()
	// fmt.Println("ReturnData-WsRouteOperate:")
	// fmt.Printf("%#v\n", wsResp.Return)

	// Marshal response variable to XML
	myXML, _ := xml.MarshalIndent(wsResp, " ", "  ")
	fmt.Println(string(myXML))
}
