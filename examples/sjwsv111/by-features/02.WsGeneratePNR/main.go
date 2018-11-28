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
	callWsGeneratePNR(sjClient)
}

// callWsGeneratePNR is a function to call WsGeneratePNR method
func callWsGeneratePNR(s *sjwsdk111.SoapSJClient) {
	params := []byte(
		`
			<Username xsi:type="xsd:string">SRIWIJAWA_AGENT_USERNAME</Username>
			<Password xsi:type="xsd:string">SRIWIJAWA_AGENT_PASSWORD</Password>
			<Received xsi:type="xsd:string">Angkasa Sriwijaya</Received>
			<ReceivedPhone xsi:type="xsd:string">081234987650</ReceivedPhone>
			<Email xsi:type="xsd:string">angkasa.sriwijaya@gmail.com</Email>
			<SearchKey xsi:type="xsd:string">SEARCH_KEY_FROM_WS_SEARCH_FLIGHT_RESPONSE</SearchKey>
			<ExtraCoverAddOns xsi:type="xsd:string">NO</ExtraCoverAddOns>
			<AdultNames xsi:type="urn:AdultNamesArray" soapenc:arrayType="urn:InputReqNameArray[1]">
				<item xsi:type="urn:InputReqNameArray">
					<FirstName xsi:type="xsd:string">Angkasa</FirstName>
					<LastName xsi:type="xsd:string">Sriwijaya</LastName>
					<Suffix xsi:type="xsd:string">MRS</Suffix>
				</item>
			</AdultNames>
			<ChildNames xsi:type="urn:ChildNamesArray" soapenc:arrayType="urn:InputReqNameArray[1]">
				<item xsi:type="urn:InputReqNameArray">
					<FirstName xsi:type="xsd:string">Mas</FirstName>
					<LastName xsi:type="xsd:string">Sriwijaya</LastName>
					<Suffix xsi:type="xsd:string">MSTR</Suffix>
					<Dob xsi:type="xsd:string">2010-10-10</Dob>
				</item>
			</ChildNames>
			<InfantNames xsi:type="urn:InfantNamesArray" soapenc:arrayType="urn:InputReqArrayInf[1]">
				<item xsi:type="urn:InputReqArrayInf">
					<FirstName xsi:type="xsd:string">Ananda</FirstName>
					<LastName xsi:type="xsd:string">Sriwijaya</LastName>
					<Suffix xsi:type="xsd:string">INF</Suffix>
					<Dob xsi:type="xsd:string">2017-07-17</Dob>
					<AdultRefference xsi:type="xsd:string">1</AdultRefference>
				</item>
			</InfantNames>
			<Keys xsi:type="urn:InputReqArrayKey" soapenc:arrayType="urn:InputReqArrayKeys[5]">
				<item xsi:type="urn:InputReqArrayKeys">
					<Key xsi:type="xsd:string">SELECTED_DEPARTURE_SEGMENT_CLASS_KEY_01:T:S</Key>
					<Category xsi:type="xsd:string">Departure</Category>
				</item>
				<item xsi:type="urn:InputReqArrayKeys">
					<Key xsi:type="xsd:string">SELECTED_DEPARTURE_SEGMENT_CLASS_KEY_02_IF_CONNECTING:T:S</Key>
					<Category xsi:type="xsd:string">Departure</Category>
				</item>
				<item xsi:type="urn:InputReqArrayKeys">
					<Key xsi:type="xsd:string">SELECTED_DEPARTURE_SEGMENT_CLASS_KEY_03_IF_CONNECTING:T:S</Key>
					<Category xsi:type="xsd:string">Departure</Category>
				</item>
				<item xsi:type="tns:InputReqArrayKeys">
					<Key xsi:type="xsd:string">SELECTED_RETURN_SEGMENT_CLASS_KEY_01_IF_ROUNTRIP:Q:S</Key>
					<Category xsi:type="xsd:string">Return</Category>
				</item>
				<item xsi:type="tns:InputReqArrayKeys">
					<Key xsi:type="xsd:string">SELECTED_RETURN_SEGMENT_CLASS_KEY_02_IF_ROUNTRIP_CONNECTING:Q:S</Key>
					<Category xsi:type="xsd:string">Return</Category>
				</item>
			</Keys>
			`)
	wsResp, errC := s.CallWsGeneratePNR(params, false)

	if errC != nil {
		fmt.Println(errC)
		return
	}

	// Access response variable
	// fmt.Println()
	// fmt.Println("ReturnData-WsGeneratePNR:")
	// fmt.Printf("%#v\n", wsResp.Return)

	// Marshal response variable to XML
	myXML, _ := xml.MarshalIndent(wsResp, " ", "  ")
	fmt.Println(string(myXML))
}
