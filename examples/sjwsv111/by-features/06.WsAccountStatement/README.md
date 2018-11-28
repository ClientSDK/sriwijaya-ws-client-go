# WsAccountStatement: Account Statement Method

Account Statement (WsAccountStatement) is a service method to retrieve account statement information from Sriwijaya Air Web Service (SOAP) v.111 [[1](https://wsp.sriwijayaair.co.id:11443/wsdl.eticketv111/index.php)].

> In this example you will learn about using  Account Statement Method (WsAccountStatement) with Go (using sjwsdk111 package). 

The following are the sections available in this guide.

- [What you'll build](#what-youll-build)
- [Prerequisites](#prerequisites)
- [Implementation](#implementation)
- [Build and Running](#build-and-running)

## What you’ll build
Let’s make a real world simple application for retrieving account statement information using Sriwijaya Air Web Services Endpoint. Following diagram demonstrates the account statement use case.

![Account Statement Diagram](images/06.WsAccountStatement-2.png "Account Statement Diagram")


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
    │   ├── 01.WsAccountStatement
    │   │   ├── README.md
    │   │   ├── build_and_run.sh
    │   │   └── main.go
    └── wsdl
        └── wsp-wsdl.eticketv111.wsdl
```

- Create the above directories in your local machine and also create empty `main.go` and `build_and_run.sh` files.

- Download Sriwijaya Air Web Service WSDL and saved to `wsp-wsdl.eticketv111.wsdl`.


### Developing the application

Let's make a simple application for retrieving account statement information using `sjwsdk111` package. 

##### Main code for WsAccountStatement (main.go)
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
	callWsAccountStatement(sjClient)
}

// callWsAccountStatement is a function to call WsAccountStatement method
func callWsAccountStatement(s *sjwsdk111.SoapSJClient) {
	params := []byte(
		`
			<Username xsi:type="xsd:string">SRIWIJAWA_AGENT_USERNAME</Username>
			<Password xsi:type="xsd:string">SRIWIJAWA_AGENT_PASSWORD</Password>
			<PeriodeStart xsi:type="xsd:string">18-Sep-18</PeriodeStart>
            <PeriodeEnd xsi:type="xsd:string">18-Sep-18</PeriodeEnd>
			`)
	wsResp, errC := s.CallWsAccountStatement(params, false)

	if errC != nil {
		fmt.Println(errC)
		return
	}

	// Access response variable
	// fmt.Println()
	// fmt.Println("ReturnData-WsAccountStatement:")
	// fmt.Printf("%#v\n", wsResp.Return)

	// Marshal response variable to XML
	myXML, _ := xml.MarshalIndent(wsResp, " ", "  ")
	fmt.Println(string(myXML))
}

```

##### Bash code for building and running the example application (build_and_run.sh)
```bash
echo "Clean..."
rm ./WsAccountStatement
echo "Build..."
go build -o WsAccountStatement main.go 
echo "Build Done."
echo "Run..."
./WsAccountStatement > WsAccountStatement-Result.xml
echo "Done."

```


## Build and Running

You can build and running by execute the "build_and_run.sh" bash files. 

```bash
   $ sh build_and_run.sh 
```

After the application is running, you will get the xml response in `WsAccountStatement-Result.xml` files.

## Sample Response

```xml
 <WsAccountStatementResponse>
   <return>
     <Username>SRIWIJAWA_AGENT_USERNAME</Username>
     <AccountStatementDetail>
       <item>
         <DateCreate>18 Sep 2018 08:08 (GMT+7)</DateCreate>
         <UserCreate>SRIWIJAWA_AGENT_USERNAME</UserCreate>
         <Description>Ticketing for ABMNYZ</Description>
         <Amount>1000000</Amount>
         <BalanceStatus>D</BalanceStatus>
         <LastBalance>1020188102</LastBalance>
         <Currency>IDR</Currency>
       </item>
       <item>
         <DateCreate>18 Sep 2018 08:08 (GMT+7)</DateCreate>
         <UserCreate>SRIWIJAWA_AGENT_USERNAME</UserCreate>
         <Description>Agent Commision for ABMNYZ</Description>
         <Amount>30000</Amount>
         <BalanceStatus>C</BalanceStatus>
         <LastBalance>1020218102</LastBalance>
         <Currency>IDR</Currency>
       </item>
       <item>
         <DateCreate>18 Sep 2018 08:08 (GMT+7)</DateCreate>
         <UserCreate>SRIWIJAWA_AGENT_USERNAME</UserCreate>
         <Description>Incentive for ABMNYZ</Description>
         <Amount>10000</Amount>
         <BalanceStatus>C</BalanceStatus>
         <LastBalance>1020228102</LastBalance>
         <Currency>IDR</Currency>
       </item>
     </AccountStatementDetail>
     <ErrorCode>ACC_STMT0000</ErrorCode>
     <ErrorMessage>Success.</ErrorMessage>
   </return>
 </WsAccountStatementResponse>
```