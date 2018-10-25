// Copyright 2018 The ClientSDK Team Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0-style
// license that can be found in the LICENSE file.

// Author: ClientSDK Team (muharihar)

package sjwsdk111

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"strings"
)

var tokens []xml.Token

// MarshalXML envelope the body and encode to xml
func (c SoapSJClient) MarshalXML(e *xml.Encoder, _ xml.StartElement) error {

	// assign SoapAction
	c.SoapAction = GetSoapActionMethod(c.Method[4:], c.Definitions)

	tokens = []xml.Token{}

	//start envelope
	if c.Definitions == nil {
		return fmt.Errorf("definitions is nil")
	}

	startEnvelope()
	if len(c.HeaderParams) > 0 {
		startHeader(c.HeaderName, c.Definitions.Types.XsdSchema.AttrTargetNamespace)
		for k, v := range c.HeaderParams {
			t := xml.StartElement{
				Name: xml.Name{
					Space: "",
					Local: k,
				},
			}

			tokens = append(tokens, t, xml.CharData(v), xml.EndElement{Name: t.Name})
		}

		endHeader(c.HeaderName)
	} else {
		startHeader("", "")
		endHeader("")
	}

	//err := startBody(c.Method, c.Definitions.Types.XsdSchema.AttrTargetNamespace)
	err := startBody(c.Method, "http://schemas.xmlsoap.org/soap/encoding/")

	if err != nil {
		return err
	}

	//start params
	startParam(c.Method, c.Definitions)

	// we use external xml parameters
	// recursiveEncode(c.Params)

	//end params
	endParam(c.Method)

	//end envelope
	endBody(c.Method)
	endEnvelope()

	for _, t := range tokens {
		err := e.EncodeToken(t)
		if err != nil {
			return err
		}
	}

	return e.Flush()
}

func recursiveEncode(hm interface{}) {
	v := reflect.ValueOf(hm)

	switch v.Kind() {
	case reflect.Map:
		for _, key := range v.MapKeys() {
			t := xml.StartElement{
				Name: xml.Name{
					Space: "",
					Local: key.String(),
				},
			}

			tokens = append(tokens, t)
			recursiveEncode(v.MapIndex(key).Interface())
			tokens = append(tokens, xml.EndElement{Name: t.Name})
		}
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			recursiveEncode(v.Index(i).Interface())
		}
	case reflect.String:
		content := xml.CharData(v.String())
		tokens = append(tokens, content)
	}

}

func startEnvelope() {
	e := xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: "soapenv:Envelope",
		},
		Attr: []xml.Attr{
			{Name: xml.Name{Space: "", Local: "xmlns:xsi"}, Value: "http://www.w3.org/2001/XMLSchema-instance"},
			{Name: xml.Name{Space: "", Local: "xmlns:xsd"}, Value: "http://www.w3.org/2001/XMLSchema"},
			{Name: xml.Name{Space: "", Local: "xmlns:soapenv"}, Value: "http://schemas.xmlsoap.org/soap/envelope/"},
			{Name: xml.Name{Space: "", Local: "xmlns:urn"}, Value: "urn:sj_service"},
		},
	}

	tokens = append(tokens, e)
}

func endEnvelope() {
	e := xml.EndElement{
		Name: xml.Name{
			Space: "",
			Local: "soapenv:Envelope",
		},
	}

	tokens = append(tokens, e)
}

func startHeader(m, n string) {
	h := xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: "soapenv:Header",
		},
	}

	if m == "" || n == "" {
		tokens = append(tokens, h)
		return
	}

	r := xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: m,
		},
		Attr: []xml.Attr{
			{Name: xml.Name{Space: "", Local: "xmlns"}, Value: n},
		},
	}

	tokens = append(tokens, h, r)

	return
}

func endHeader(m string) {
	h := xml.EndElement{
		Name: xml.Name{
			Space: "",
			Local: "soapenv:Header",
		},
	}

	if m == "" {
		tokens = append(tokens, h)
		return
	}

	r := xml.EndElement{
		Name: xml.Name{
			Space: "",
			Local: m,
		},
	}

	tokens = append(tokens, r, h)
}

// startToken initiate body of the envelope
func startBody(m, n string) error {
	b := xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: "soapenv:Body",
		},
	}

	if m == "" || n == "" {
		return fmt.Errorf("method or namespace is empty")
	}

	r := xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: m,
		},
		Attr: []xml.Attr{
			{Name: xml.Name{Space: "", Local: "soapenv:encodingStyle"}, Value: n},
		},
	}

	tokens = append(tokens, b, r)

	return nil
}

// endToken close body of the envelope
func endBody(m string) {
	b := xml.EndElement{
		Name: xml.Name{
			Space: "",
			Local: "soapenv:Body",
		},
	}

	r := xml.EndElement{
		Name: xml.Name{
			Space: "",
			Local: m,
		},
	}

	tokens = append(tokens, r, b)
}

func startParam(m string, d *wsdlDefinitions) {

	xsiType := ""
	for _, dPO := range d.PortTypes.Operations {
		if dPO.AttrName == m[4:] {
			xsiType = dPO.Inputs.AttrMessage

			for _, dMsg := range d.Messages {
				if dMsg.AtrrName == xsiType[4:] {
					xsiType = dMsg.Parts[0].AttrType
					xsiType = "urn:" + xsiType[4:]
				}
			}
		}
	}

	r := xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: "param",
		},
		Attr: []xml.Attr{
			{Name: xml.Name{Space: "", Local: "xsi:type"}, Value: xsiType},
			{Name: xml.Name{Space: "", Local: "xmlns:urn"}, Value: "urn:webservice"},
		},
	}

	tokens = append(tokens, r)

	paramContent := xml.CharData("#PARAMETERS#")
	tokens = append(tokens, paramContent)

	/*
		if xsi_type != "" {
			for _, dSchema := range d.Types.XsdSchema.ComplexTypes {
				if dSchema.AttrName == xsi_type[4:] {
					for _, dSeqEl := range dSchema.Sequence.Elements {
						es := xml.StartElement{
							Name: xml.Name{Space: "", Local: dSeqEl.AttrName},
							Attr: []xml.Attr{
								{Name: xml.Name{Space: "", Local: "xsi.type"}, Value: dSeqEl.AttrType},
							},
						}
						ee := xml.EndElement{
							Name: xml.Name{Space: "", Local: dSeqEl.AttrName},
						}
						tokens = append(tokens, es, ee)
					}
				}
			}
		}
	*/
}

func endParam(m string) {
	r := xml.EndElement{
		Name: xml.Name{
			Space: "",
			Local: "param",
		},
	}

	tokens = append(tokens, r)
}

// GetSoapActionMethod retrieve Soap Action Method
func GetSoapActionMethod(m string, d *wsdlDefinitions) string {
	result := ""

	for _, dBOp := range d.Bindings[0].Operations {
		if strings.ToLower(dBOp.AttrName) == strings.ToLower(m) {
			result = dBOp.SoapOperation.AttrSoapAction
		}
	}

	return result
}
