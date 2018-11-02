// Copyright 2018 The ClientSDK Team Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0-style
// license that can be found in the LICENSE file.

// Author: ClientSDK Team (muharihar)

package sjwsdk111

import (
	"bytes"
	"encoding/xml"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html/charset"
)

// wsdlDefinitions represent Sriwijaya WSDL (wsdl.eticketv111)
type wsdlDefinitions struct {
	XMLName             xml.Name       `xml:"definitions"`
	Name                string         `xml:"name,attr"`
	AttrXmlnsSoapEnv    string         `xml:"SOAP-ENV,attr"`
	AttrXmlnsXsd        string         `xml:"xsd,attr"`
	AttrXmlnsXsi        string         `xml:"xsi,attr"`
	AttrXmlnsSoapEnc    string         `xml:"SOAP-ENC,attr"`
	AttrXmlnsTns        string         `xml:"tns,attr"`
	AttrXmlnsSoap       string         `xml:"soap,attr"`
	AttrXmlnsWsdl       string         `xml:"wsdl,attr"`
	AttrXmlns           string         `xml:"xmlns,attr"`
	AttrTargetNamespace string         `xml:"targetNamespace,attr"`
	Types               *wsdlTypes     `xml:"types"`
	Messages            []*wsdlMessage `xml:"message"`
	PortTypes           *wsdlPortTypes `xml:"portType"`
	Bindings            []*wsdlBinding `xml:"binding"`
	Services            *wsdlService   `xml:"service"`
}

type wsdlTypes struct {
	XsdSchema xsdSchema `xml:"schema"`
}

type xsdSchema struct {
	AttrTargetNamespace string            `xml:"targetNamespace,attr"`
	Imports             []*xsdImport      `xml:"import"`
	ComplexTypes        []*xsdComplexType `xml:"complexType"`
}

type xsdImport struct {
	AttrNamespace string `xml:"namespace,attr"`
}

type xsdComplexType struct {
	AttrName       string             `xml:"name,attr"`
	Sequence       *xsdSequence       `xml:"sequence,omitempty"`
	ComplexContent *xsdComplexContent `xml:"complexContent,omitempty"`
	All            *xsdComplexTypeAll `xml:"all,omitempty"`
}

type xsdSequence struct {
	Elements []*xsdElement `xml:"element"`
}

type xsdElement struct {
	AttrName      string `xml:"name,attr"`
	AttrType      string `xml:"type,attr"`
	AttrNillable  bool   `xml:"nillable,attr,omitempty"`
	AttrMinOccurs string `xml:"minOccurs,attr,omitempty"`
	AttrMaxOccurs string `xml:"maxOccurs,attr,omitempty"`
}

type xsdComplexContent struct {
	Restriction *xsdRestriction `xml:"restriction"`
}

type xsdRestriction struct {
	AttrBase  string        `xml:"base,attr"`
	Attribute *xsdAttribute `xml:"attribute"`
}

type xsdAttribute struct {
	AttrRef       string `xml:"ref,attr"`
	AttrArrayType string `xml:"arrayType,attr"`
}

type xsdComplexTypeAll struct {
	Elements []*xsdElement `xml:"element"`
}

type wsdlMessage struct {
	AtrrName string             `xml:"name,attr"`
	Parts    []*wsdlMessagePart `xml:"part"`
}

type wsdlMessagePart struct {
	AttrName string `xml:"name,attr"`
	AttrType string `xml:"type,attr"`
}

type wsdlPortTypes struct {
	AttrName   string           `xml:"name,attr"`
	Operations []*wsdlOperation `xml:"operation"`
}

type wsdlOperation struct {
	AttrName      string               `xml:"name,attr"`
	Documentation string               `xml:"documentation"`
	Inputs        *wsdlOperationInput  `xml:"input"`
	Outputs       *wsdlOperationOutput `xml:"output"`
}

type wsdlOperationInput struct {
	AttrMessage string `xml:"message,attr"`
}

type wsdlOperationOutput struct {
	AttrMessage string `xml:"message,attr"`
}

type wsdlBinding struct {
	AttrName    string                  `xml:"name,attr"`
	AttrType    string                  `xml:"type,attr"`
	SoapBinding *soapBinding            `xml:"binding"`
	Operations  []*wsdlBindingOperation `xml:"operation"`
}

type soapBinding struct {
	AttrStyle     string `xml:"style,attr"`
	AttrTransport string `xml:"transport,attr"`
}

type wsdlBindingOperation struct {
	AttrName      string                `xml:"name,attr"`
	SoapOperation *bindingSoapOperation `xml:"operation"`
	Input         *bindingSoapOprBody   `xml:"input"`
	Output        *bindingSoapOprBody   `xml:"output"`
}

type bindingSoapOperation struct {
	AttrSoapAction string `xml:"soapAction,attr"`
	AttrStyle      string `xml:"style,attr"`
}

type bindingSoapOprBody struct {
	Body struct {
		AttrUse           string `xml:"use,attr"`
		AttrNamespace     string `xml:"namespace,attr"`
		AttrEncodingStyle string `xml:"encodingStyle,attr"`
	} `xml:"body"`
}

type wsdlService struct {
	AttrName string      `xml:"name,attr"`
	Ports    []*wsdlPort `xml:"http://schemas.xmlsoap.org/wsdl/ port"`
}

type wsdlPort struct {
	AttrName    string       `xml:"name,attr"`
	AttrBinding string       `xml:"binding,attr"`
	SoapAddress *soapAddress `xml:"address"`
}

type soapAddress struct {
	AttrLocation string `xml:"location,attr"`
}

// getWsdlDefinitions sent request to the wsdl location (URL location or File location) and set definitions on struct
// URL location must accessible without proxy.
// File location is recommended for faster response.
func getWsdlDefinitions(wsdlLocation string, locationType string) (wsdl *wsdlDefinitions, err error) {

	reader, err := getWsdlContent(wsdlLocation, locationType)
	if err != nil {
		return nil, err
	}

	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&wsdl)

	return wsdl, err
}

// getWsdlContent is a function to get WSDL content
func getWsdlContent(wsdlLocation string, locationType string) (io.Reader, error) {
	switch strings.ToLower(locationType) {
	case "url":
		r, e := getWsdlFromURL(wsdlLocation)
		return r, e
	case "file":
		r, e := getWsdlFromFile(wsdlLocation)
		return r, e
	default:
		r, e := getWsdlFromFile(wsdlLocation)
		return r, e
	}
}

// getWsdlFromURL is a function to get WSDL content from URL resource
func getWsdlFromURL(wsdlURL string) (io.Reader, error) {
	r, err := http.Get(wsdlURL)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	result := r.Body

	return result, err
}

// getWsdlFromFile is a function to get WSDL content from File resource
func getWsdlFromFile(wsdlFilePath string) (io.Reader, error) {
	xmlFile, err := os.Open(wsdlFilePath)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()

	result, errF := ioutil.ReadAll(xmlFile)
	if errF != nil {
		return nil, errF
	}

	reader := bytes.NewReader([]byte(result))

	return reader, err
}
