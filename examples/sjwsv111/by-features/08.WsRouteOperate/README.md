# WsRouteOperate: Route Operate Method

Route Operate (WsRouteOperate) is a service method to retrieve routes information from Sriwijaya Air Web Service (SOAP) v.111 [[1](https://wsp.sriwijayaair.co.id:11443/wsdl.eticketv111/index.php)].

> In this example you will learn about using  Route Information Method (WsRouteOperate) with Go (using sjwsdk111 package). 

The following are the sections available in this guide.

- [What you'll build](#what-youll-build)
- [Prerequisites](#prerequisites)
- [Implementation](#implementation)
- [Build and Running](#build-and-running)

## What you’ll build
Let’s make a real world simple application for retrieving route information using Sriwijaya Air Web Services Endpoint. Following diagram demonstrates the route operate use case.

![Route Operate Diagram](images/08.WsRouteOperate-2.png "Route Operate Diagram")


## Prerequisites

- [Sriwijaya Air Web Service (SOAP) Client for Go (sjwsdk111 GoLang package) ](https://github.com/ClientSDK/sriwijaya-ws-client-go)

```Go
go get github.com/ClientSDK/sriwijaya-ws-client-go/sjwsdk111
```

- A Text Editor or an IDE

### Sriwijaya Air Agent requirements
- Sriwijaya Air Agent Credential Account ([Agent Application](https://agent.sriwijayaair.co.id/SJ-Eticket/login.php?action=in))
- Sriwijaya Air Web Service Access (IP Whitelist) ( [Production](https://wsp.sriwijayaair.co.id:11443/wsdl.eticketv111/index.php), [Development](https://wsx.sriwijayaair.co.id:11443/wsdl.eticketv111/index.php) )
- WSDL File ( [Production](https://wsp.sriwijayaair.co.id:11443/wsdl.eticketv111/index.php?wsdl), [Development](https://wsx.sriwijayaair.co.id:11443/wsdl.eticketv111/index.php?wsdl) )

## Implementation

> If you want to skip the basics, you can download the git repo and directly move to the "Build and Running" section by skipping  "Implementation" section.

### Example structure

Go is a complete programming language that supports custom project structures. Let's use the following package structure for this example.

```
sjwsv111
    ├── by-features
    │   ├── 01.WsRouteOperate
    │   │   ├── README.md
    │   │   ├── build_and_run.sh
    │   │   └── main.go
    └── wsdl
        └── wsp-wsdl.eticketv111.wsdl
```

- Create the above directories in your local machine and also create empty `main.go` and `build_and_run.sh` files.

- Download Sriwijaya Air Web Service WSDL and saved to `wsp-wsdl.eticketv111.wsdl`.


### Developing the application

Let's make a simple application for retrieving route information using `sjwsdk111` package. 

##### Main code for WsRouteOperate (main.go)
```go
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

```

##### Bash code for building and running the example application (build_and_run.sh)
```bash
echo "Clean..."
rm ./WsRouteOperate
echo "Build..."
go build -o WsRouteOperate main.go 
echo "Build Done."
echo "Run..."
./WsRouteOperate > WsRouteOperate-Result.xml
echo "Done."

```


## Build and Running

You can build and running by execute the "build_and_run.sh" bash files. 

```bash
   $ sh build_and_run.sh 
```

After the application is running, you will get the xml response in `WsRouteOperate-Result.xml` files.

## Sample Response

```xml
 <WsRouteOperateResponse>
   <return>
     <Username>SRIWIJAWA_AGENT_USERNAME</Username>
     <RouteOperates>
       <item>
         <CityFrom>JOG</CityFrom>
         <CityFromName>Yogyakarta</CityFromName>
         <CityFromCountry>Indonesia</CityFromCountry>
         <ApoNameFrom>Adi Sucipto</ApoNameFrom>
         <TimeZoneFrom>7</TimeZoneFrom>
         <CityTo>CGK</CityTo>
         <CityToName>Jakarta</CityToName>
         <CityToCountry>Indonesia</CityToCountry>
         <ApoNameTo>Soekarno Hatta International Airport</ApoNameTo>
         <TimeZoneTo>7</TimeZoneTo>
         <StatusRoute>D</StatusRoute>
       </item>
       <item>
         <CityFrom>PEN</CityFrom>
         <CityFromName>Penang</CityFromName>
         <CityFromCountry>Malaysia</CityFromCountry>
         <ApoNameFrom>Penang International Airport</ApoNameFrom>
         <TimeZoneFrom>8</TimeZoneFrom>
         <CityTo>KNO</CityTo>
         <CityToName>Medan</CityToName>
         <CityToCountry>Indonesia</CityToCountry>
         <ApoNameTo>Kualanamu </ApoNameTo>
         <TimeZoneTo>7</TimeZoneTo>
         <StatusRoute>I</StatusRoute>
       </item>
     </RouteOperates>
     <ErrorCode>ROPERATE0000</ErrorCode>
     <ErrorMessage>Success.</ErrorMessage>
   </return>
 </WsRouteOperateResponse>
```