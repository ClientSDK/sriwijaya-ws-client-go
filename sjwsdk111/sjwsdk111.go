// Copyright 2018 The ClientSDK Team Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0-style
// license that can be found in the LICENSE file.

// Author: ClientSDK Team (muharihar)

package sjwsdk111

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"strings"
	"sync"

	"golang.org/x/net/html/charset"
)

const (
	userAgent = "Apache-HttpClient/4.1.1 (java 1.5)"
)

// NewSoapSJClient return new *SoapSJClient to handle the requests with the WSDL
func NewSoapSJClient(httpClient *http.Client, wsdlLocation string, locationType string) (*SoapSJClient, error) {

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	wsdlDef, err := getWsdlDefinitions(wsdlLocation, locationType)
	if err != nil {
		return nil, err
	}

	// set to https
	svcAddrLocation := wsdlDef.Services.Ports[0].SoapAddress.AttrLocation
	svcAddrLocation = strings.Replace(svcAddrLocation, "http://", "https://", -1)

	sjClient := &SoapSJClient{
		clientHTTP:             httpClient,
		WSDL:                   wsdlLocation,
		URL:                    "urn:sj_service",
		ServiceAddressLocation: svcAddrLocation,
		Definitions:            wsdlDef,
		userAgent:              userAgent,
	}

	return sjClient, nil
}

// HeaderParams holds params specific to the header
type HeaderParams map[string]string

// SoapSJClient struct hold all the informations about WSDL,
// request and response of the server
type SoapSJClient struct {
	clientMu   sync.Mutex   // clientMu protects the client during calls that modify the CheckRedirect func.
	clientHTTP *http.Client // HTTP client used to communicate with the API.

	WSDL                   string
	URL                    string
	ServiceAddressLocation string
	Method                 string
	SoapAction             string
	Params                 interface{}
	HeaderName             string
	HeaderParams           HeaderParams
	Definitions            *wsdlDefinitions
	SoapResponse           []byte
	Header                 []byte
	Body                   []byte

	userAgent string
	payload   []byte
}

// GetLastRequest returns the last request
func (c *SoapSJClient) GetLastRequest() []byte {
	return c.payload
}

// SetUserAgent is a function to set (HTTPClient) UserAgent
func (c *SoapSJClient) SetUserAgent(value string) {
	if value != "" {
		c.userAgent = value
	}
}

// Call call's the method m with Params p
func (c *SoapSJClient) Call(method string, params []byte, debug bool) (err error) {
	c.Method = fmt.Sprintf("urn:%s", method)
	c.Params = params
	c.SoapAction = GetSoapActionMethod(method, c.Definitions)

	c.payload, err = xml.MarshalIndent(c, "", "    ")
	if err != nil {
		return err
	}
	c.payload = []byte(strings.Replace(string(c.payload), "#PARAMETERS#", string(params), -1))

	soapResp, err := c.doRequest(c.ServiceAddressLocation, debug)
	if err != nil {
		return err
	}
	c.SoapResponse = soapResp

	var soap soapEnvelopeResponse

	//err = xml.Unmarshal(soapResp, &soap)
	reader := bytes.NewReader(soapResp)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&soap)

	c.Body = soap.Body.Contents
	c.Header = soap.Header.Contents

	return err
}

// CallWsSearchFlight is a function to call WsSearchFlight Method
func (c *SoapSJClient) CallWsSearchFlight(params []byte, debug bool) (*WsSearchFlightResponse, error) {
	err := c.Call("WsSearchFlight", params, debug)
	if err != nil {
		return nil, err
	}

	var resp *WsSearchFlightResponse

	err = xml.Unmarshal(c.Body, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Return.ErrorCode != "SEARCH0000" {
		err = fmt.Errorf("[ERROR]:: [ErrorCode]: %s; [ErrorMessage]: %s", resp.Return.ErrorCode, resp.Return.ErrorMessage)

		return nil, err
	}

	return resp, nil
}

// CallWsGeneratePNR is a function to call WsGeneratePNR Method
func (c *SoapSJClient) CallWsGeneratePNR(params []byte, debug bool) (*WsGeneratePNRResponse, error) {
	err := c.Call("WsGeneratePNR", params, debug)
	if err != nil {
		return nil, err
	}

	var resp *WsGeneratePNRResponse

	err = xml.Unmarshal(c.Body, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Return.ErrorCode != "GENERATE0000" {
		err = fmt.Errorf("[ERROR]:: [ErrorCode]: %s; [ErrorMessage]: %s", resp.Return.ErrorCode, resp.Return.ErrorMessage)

		return nil, err
	}

	return resp, nil
}

// CallWsIssuing is a function to call WsIssuing Method
func (c *SoapSJClient) CallWsIssuing(params []byte, debug bool) (*WsIssuingResponse, error) {
	err := c.Call("WsIssuing", params, debug)
	if err != nil {
		return nil, err
	}

	var resp *WsIssuingResponse

	err = xml.Unmarshal(c.Body, &resp)
	if err != nil {
		return nil, err
	}

	//if resp.Return.ErrorCode != "GENERATE0000" {
	if resp.Return.ErrorMessage != "Success." {
		err = fmt.Errorf("[ERROR]:: [ErrorCode]: %s; [ErrorMessage]: %s", resp.Return.ErrorCode, resp.Return.ErrorMessage)

		return nil, err
	}

	return resp, nil
}

// CallWsRetrievePNR is a function to call WsRetrievePNR Method
func (c *SoapSJClient) CallWsRetrievePNR(params []byte, debug bool) (*WsRetrievePNRResponse, error) {
	err := c.Call("WsRetrievePNR", params, debug)
	if err != nil {
		return nil, err
	}

	var resp *WsRetrievePNRResponse

	err = xml.Unmarshal(c.Body, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Return.ErrorCode != "RETRIEVE0000" {
		err = fmt.Errorf("[ERROR]:: [ErrorCode]: %s; [ErrorMessage]: %s", resp.Return.ErrorCode, resp.Return.ErrorMessage)

		return nil, err
	}

	return resp, nil
}

// CallWsCancelPNR is a function to call WsCancelPNR Method
func (c *SoapSJClient) CallWsCancelPNR(params []byte, debug bool) (*WsCancelPNRResponse, error) {
	err := c.Call("WsCancelPNR", params, debug)
	if err != nil {
		return nil, err
	}

	var resp *WsCancelPNRResponse

	err = xml.Unmarshal(c.Body, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Return.ErrorCode != "CANCEL0000" {
		err = fmt.Errorf("[ERROR]:: [ErrorCode]: %s; [ErrorMessage]: %s", resp.Return.ErrorCode, resp.Return.ErrorMessage)

		return nil, err
	}

	return resp, nil
}

// CallWsCreditBalance is a function to call WsCreditBalance Method
func (c *SoapSJClient) CallWsCreditBalance(params []byte, debug bool) (*WsCreditBalanceResponse, error) {
	err := c.Call("WsCreditBalance", params, debug)
	if err != nil {
		return nil, err
	}

	var resp *WsCreditBalanceResponse

	err = xml.Unmarshal(c.Body, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Return.ErrorCode != "CR_BALANCE0000" {
		err = fmt.Errorf("[ERROR]:: [ErrorCode]: %s; [ErrorMessage]: %s", resp.Return.ErrorCode, resp.Return.ErrorMessage)

		return nil, err
	}

	return resp, nil
}

// CallWsAccountStatement is a function to call WsCreditBalance Method
func (c *SoapSJClient) CallWsAccountStatement(params []byte, debug bool) (*WsAccountStatementResponse, error) {
	err := c.Call("WsAccountStatement", params, debug)
	if err != nil {
		return nil, err
	}

	var resp *WsAccountStatementResponse

	err = xml.Unmarshal(c.Body, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Return.ErrorCode != "ACC_STMT0000" {
		err = fmt.Errorf("[ERROR]:: [ErrorCode]: %s; [ErrorMessage]: %s", resp.Return.ErrorCode, resp.Return.ErrorMessage)

		return nil, err
	}

	return resp, nil
}

// CallWsRouteOperate is a function to call WsRouteOperate Method
func (c *SoapSJClient) CallWsRouteOperate(params []byte, debug bool) (*WsRouteOperateResponse, error) {
	err := c.Call("WsRouteOperate", params, debug)
	if err != nil {
		return nil, err
	}

	var resp *WsRouteOperateResponse

	err = xml.Unmarshal(c.Body, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Return.ErrorCode != "ROPERATE0000" {
		err = fmt.Errorf("[ERROR]:: [ErrorCode]: %s; [ErrorMessage]: %s", resp.Return.ErrorCode, resp.Return.ErrorMessage)

		return nil, err
	}

	return resp, nil
}

// doRequest makes new request to the server using the c.Method, c.URL and the body.
// body is enveloped in Call method
func (c *SoapSJClient) doRequest(urlWS string, debug bool) ([]byte, error) {

	req, err := http.NewRequest("POST", urlWS, bytes.NewBuffer(c.payload))
	if err != nil {
		return nil, err
	}

	client := c.clientHTTP

	req.ContentLength = int64(len(c.payload))

	req.Header.Add("Content-Type", "text/xml;charset=UTF-8")
	req.Header.Add("Accept", "text/xml")
	req.Header.Add("SOAPAction", fmt.Sprintf("%s", c.SoapAction))
	req.Header.Add("User-Agent", c.userAgent)

	if debug {
		DebugHTTPRequest(httputil.DumpRequestOut(req, true))
		//DebugHTTPRequest(httputil.DumpRequest(req, true))
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err == nil && debug == true {
		DebugHTTPResponse(httputil.DumpResponse(resp, true))
	}

	return ioutil.ReadAll(resp.Body)
}

// soapEnvelopeResponse struct
type soapEnvelopeResponse struct {
	XMLName struct{} `xml:"Envelope"`
	Header  soapHeaderResponse
	Body    soapBodyResponse
}

// soapHeaderResponse struct
type soapHeaderResponse struct {
	XMLName  struct{} `xml:"Header"`
	Contents []byte   `xml:",innerxml"`
}

// soapBodyResponse struct
type soapBodyResponse struct {
	XMLName  struct{} `xml:"Body"`
	Contents []byte   `xml:",innerxml"`
}
