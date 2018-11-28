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
	callWsSearchFlight(sjClient)
}

// callWsSearchFlight is a function to call WsSearchFlight method
func callWsSearchFlight(s *sjwsdk111.SoapSJClient) {
	params := []byte(
		`
			<Username xsi:type="xsd:string">SRIWIJAWA_AGENT_USERNAME</Username>
			<Password xsi:type="xsd:string">SRIWIJAWA_AGENT_PASSWORD</Password>
			<ReturnStatus xsi:type="xsd:string">YES</ReturnStatus>
			<CityFrom xsi:type="xsd:string">CGK</CityFrom>
			<CityTo xsi:type="xsd:string">DPS</CityTo>
			<DepartDate xsi:type="xsd:string">01-Feb-19</DepartDate>
			<ReturnDate xsi:type="xsd:string">13-Feb-19</ReturnDate>
			<PromoCode xsi:type="xsd:string"></PromoCode>
			<Adult xsi:type="xsd:string">1</Adult>
			<Child xsi:type="xsd:string">1</Child>
			<Infant xsi:type="xsd:string">1</Infant>
			`)
	wsResp, errC := s.CallWsSearchFlight(params, false)

	if errC != nil {
		fmt.Println(errC)
		return
	}

	// Access response variable
	// fmt.Println()
	// fmt.Println("ReturnData-WsSearchFlight:")
	// fmt.Printf("%#v\n", wsResp.Return)

	// Marshal response variable to XML
	myXML, _ := xml.MarshalIndent(wsResp, " ", "  ")
	fmt.Println(string(myXML))
}
