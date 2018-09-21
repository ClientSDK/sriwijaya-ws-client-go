# WsCreditBalance: Credit Balance Method

Credit Balance (WsCreditBalance) is a service method to retrieve credit balance information from Sriwijaya Air Web Service (SOAP) v.111 [[1](https://wsp.sriwijayaair.co.id:11443/wsdl.eticketv111/index.php)].

> In this example you will learn about using  Credit Balance Method (WsCreditBalance) with Go (using sjwsdk111 package). 

The following are the sections available in this guide.

- [What you'll build](#what-youll-build)
- [Prerequisites](#prerequisites)
- [Implementation](#implementation)
- [Build and Running](#build-and-running)

## What you’ll build
Let’s make a real world simple application for retrieving credit balance information using Sriwijaya Air Web Services Endpoint. Following diagram demonstrates the credit balance use case.

![Credit Balance Diagram](images/07.WsCreditBalance-1.png "Credit Balance Diagram")


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
    │   ├── 01.WsCreditBalance
    │   │   ├── README.md
    │   │   ├── build_and_run.sh
    │   │   └── main.go
    └── wsdl
        └── wsp-wsdl.eticketv111.wsdl
```

- Create the above directories in your local machine and also create empty `main.go` and `build_and_run.sh` files.

- Download Sriwijaya Air Web Service WSDL and saved to `wsp-wsdl.eticketv111.wsdl`.


### Developing the application

Let's make a simple application for retrieving credit balance information using `sjwsdk111` package. 

##### Main code for WsCreditBalance (main.go)
```go
package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"

	"github.com/ClientSDK/sriwijaya-ws-client-go/sjwsdk111"
)

func main() {

	// Access via proxy if needed
	proxyURL, _ := url.Parse("http://proxy-ip-address:proxy-port")
	//proxyURL, _ := url.Parse("http://proxy-user:proxy-password@proxy-ip-address:proxy-port")

	// Initiate http client with transport
	httpClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
    
	// Initiate NewSoapSJClient version 111
	sjClient, err := sjwsdk111.NewSoapSJClient(httpClient, "../../wsdl/wsp-wsdl.eticketv111.wsdl", "file")
	if err != nil {
		fmt.Println(err)
	}

	callWsCreditBalance(sjClient)
}

// callWsCreditBalance is a function to call WsCreditBalance method
func callWsCreditBalance(s *sjwsdk111.SoapSJClient) {
	params := []byte(
		`
			<Username xsi:type="xsd:string">SRIWIJAWA_AGENT_USERNAME</Username>
			<Password xsi:type="xsd:string">SRIWIJAWA_AGENT_PASSWORD</Password>
			...
			`)
	wsResp, errC := s.CallWsCreditBalance(params, false)

	if errC != nil {
		fmt.Println(errC)
		return
	}

	// Access response variable
	// fmt.Println()
	// fmt.Println("ReturnData-WsCreditBalance:")
	// fmt.Printf("%#v\n", wsResp.Return)
    
	// Marshal response variable to XML
	myXML, _ := xml.MarshalIndent(wsResp, " ", "  ")
	fmt.Println(string(myXML))
}

```

##### Bash code for building and running the example application (build_and_run.sh)
```bash
echo "Clean..."
rm ./WsCreditBalance
echo "Build..."
go build -o WsCreditBalance main.go 
echo "Build Done."
echo "Run..."
./WsCreditBalance > WsCreditBalance-Result.xml
echo "Done."

```


## Build and Running

You can build and running by execute the "build_and_run.sh" bash files. 

```bash
   $ sh build_and_run.sh 
```

After the application is running, you will get the xml response in `WsCreditBalance-Result.xml` files.

## Sample Response

```xml
 <WsCreditBalanceResponse>
   <return>
     <Username>SRIWIJAWA_AGENT_USERNAME</Username>
     <CreditBalance>20188102</CreditBalance>
     <ErrorCode>CR_BALANCE0000</ErrorCode>
     <ErrorMessage>Success.</ErrorMessage>
   </return>
 </WsCreditBalanceResponse>
```