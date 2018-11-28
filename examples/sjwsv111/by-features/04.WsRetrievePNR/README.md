# WsRetrievePNR: Retrieve PNR Method

Retrieve PNR (WsRetrievePNR) is a service method to retrieve PNR information from Sriwijaya Air Web Service (SOAP) v.111 [[1](https://wsp.sriwijayaair.co.id:11443/wsdl.eticketv111/index.php)].

> In this example you will learn about using  Retrieve PNR Method (WsRetrievePNR) with Go (using sjwsdk111 package). 

The following are the sections available in this guide.

- [What you'll build](#what-youll-build)
- [Prerequisites](#prerequisites)
- [Implementation](#implementation)
- [Build and Running](#build-and-running)

## What you’ll build
Let’s make a real world simple application for retreiving PNR information by using Sriwijaya Air Web Services Endpoint. Following diagram demonstrates the retrieve PNR use case.

![Retrieve PNR Diagram](images/04.WsRetrievePNR-2.png "Retrieve PNR Diagram")


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
    │   ├── 01.WsRetrievePNR
    │   │   ├── README.md
    │   │   ├── build_and_run.sh
    │   │   └── main.go
    └── wsdl
        └── wsp-wsdl.eticketv111.wsdl
```

- Create the above directories in your local machine and also create empty `main.go` and `build_and_run.sh` files.

- Download Sriwijaya Air Web Service WSDL and saved to `wsp-wsdl.eticketv111.wsdl`.


### Developing the application

Let's make a simple application for retrieving PNR information using `sjwsdk111` package. 

##### Main code for WsRetrievePNR (main.go)
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
	callWsRetrievePNR(sjClient)
}

// callWsRetrievePNR is a function to call WsRetrievePNR method
func callWsRetrievePNR(s *sjwsdk111.SoapSJClient) {
	params := []byte(
		`
			<Username xsi:type="xsd:string">SRIWIJAWA_AGENT_USERNAME</Username>
			<Password xsi:type="xsd:string">SRIWIJAWA_AGENT_PASSWORD</Password>
			<BookingCode xsi:type="xsd:string">ACTIVE_PNR_CODE_OR_BOOKING_CODE_FROM_WS_GENERATE_PNR_RESPONSE</BookingCode>
			`)
	wsResp, errC := s.CallWsRetrievePNR(params, false)

	if errC != nil {
		fmt.Println(errC)
		return
	}

	// Access response variable
	// fmt.Println()
	// fmt.Println("ReturnData-WsRetrievePNR:")
	// fmt.Printf("%#v\n", wsResp.Return)

	// Marshal response variable to XML
	myXML, _ := xml.MarshalIndent(wsResp, " ", "  ")
	fmt.Println(string(myXML))
}

```

##### Bash code for building and running the example application (build_and_run.sh)
```bash
echo "Clean..."
rm ./WsRetrievePNR
echo "Build..."
go build -o WsRetrievePNR main.go 
echo "Build Done."
echo "Run..."
./WsRetrievePNR > WsRetrievePNR-Result.xml
echo "Done."

```


## Build and Running

You can build and running by execute the "build_and_run.sh" bash files. 

```bash
   $ sh build_and_run.sh 
```

After the application is running, you will get the xml response in `WsRetrievePNR-Result.xml` files.

## Sample Response

```xml
 <WsRetrievePNRResponse>
   <return>
     <Username>SRIWIJAWA_AGENT_USERNAME</Username>
     <BookingCode>ABMNYZ</BookingCode>
     <YourItineraryDetails>
       <ReservationDetails>
         <BookingCode>ABMNYZ</BookingCode>
         <BookingDate>18 Sep 2018 08:01 (GMT+7)</BookingDate>
         <BalanceDue>0</BalanceDue>
         <BalanceDueRemarks>*Extra Cover Insurance (STI) not include in balance due.</BalanceDueRemarks>
         <CurrencyCode>IDR</CurrencyCode>
         <Time>18 Sep 2018 18:18 (GMT+7)</Time>
         <TimeDescription>DateOfIssue</TimeDescription>
         <Status>Confirm</Status>
       </ReservationDetails>
       <PassengerDetails>
         <item>
           <No>1</No>
           <Suffix>Mr</Suffix>
           <FirstName>ANGKASA</FirstName>
           <LastName>SRIWIJAYA</LastName>
           <SeatQty>1</SeatQty>
           <TicketNumber>977012345678901C1, C2</TicketNumber>
           <SpecialRequest>N/A</SpecialRequest>
         </item>
         <item>
           <No>2</No>
           <Suffix>Mstr</Suffix>
           <FirstName>MAS</FirstName>
           <LastName>SRIWIJAYA</LastName>
           <SeatQty>1</SeatQty>
           <TicketNumber>977012345678902C1, C2</TicketNumber>
           <SpecialRequest>N/A</SpecialRequest>
         </item>
         <item>
           <No>3</No>
           <Suffix>Inf</Suffix>
           <FirstName>ANANDA</FirstName>
           <LastName>SRIWIJAYA</LastName>
           <SeatQty>0</SeatQty>
           <TicketNumber>977012345678903C1, C2</TicketNumber>
           <SpecialRequest>N/A</SpecialRequest>
         </item>
       </PassengerDetails>
       <ItineraryDetails>
         <Journey>
           <item>
             <Segment>
               <item>
                 <FlownDate>18-SEP-18</FlownDate>
                 <FlightNo>SJ097</FlightNo>
                 <CityFrom>TKG</CityFrom>
                 <CityTo>CGK</CityTo>
                 <CityFromName>Lampung</CityFromName>
                 <CityToName>Jakarta</CityToName>
                 <StdLT>12:40 LT</StdLT>
                 <StaLT>13:25 LT</StaLT>
                 <ReservationStatus>RR</ReservationStatus>
                 <Class>L</Class>
                 <CheckInStatus>NO</CheckInStatus>
               </item>
             </Segment>
           </item>
           <item>
             <Segment>
               <item>
                 <FlownDate>28-SEP-18</FlownDate>
                 <FlightNo>SJ096</FlightNo>
                 <CityFrom>CGK</CityFrom>
                 <CityTo>TKG</CityTo>
                 <CityFromName>Jakarta</CityFromName>
                 <CityToName>Lampung</CityToName>
                 <StdLT>11:25 LT</StdLT>
                 <StaLT>12:05 LT</StaLT>
                 <ReservationStatus>RR</ReservationStatus>
                 <Class>Q</Class>
                 <CheckInStatus>NO</CheckInStatus>
               </item>
             </Segment>
           </item>
         </Journey>
       </ItineraryDetails>
       <PaymentDetails>
         <BasicFare>1500000</BasicFare>
         <Others>525000</Others>
         <Sti></Sti>
         <Total>2025000</Total>
         <Nta>1881900</Nta>
         <CurrencyCode>IDR</CurrencyCode>
       </PaymentDetails>
       <ContactList>
         <item>
           <Type>Phone</Type>
           <Description>Main</Description>
           <Value>081234987650</Value>
         </item>
         <item>
           <Type>Phone</Type>
           <Description>Main</Description>
           <Value>02740123456</Value>
         </item>
         <item>
           <Type>Email</Type>
           <Description>Work</Description>
           <Value>angkasa.sriwijaya@gmail.com</Value>
         </item>
       </ContactList>
       <AgentDetails>
         <BookedBy>SRIWIJAWA_AGENT_USERNAME</BookedBy>
         <IssuedBy>SRIWIJAWA_AGENT_USERNAME</IssuedBy>
       </AgentDetails>
       <BookingRemarks>
        <item>
           <CommentText>rc pleasee core gender pax1 (MRS ANGKASA) to (MR ANGKASA)-tq</CommentText>
           <CreatedBy>SRIWIJAWA_CS_USERNAME</CreatedBy>
           <CreatedDate>18 Sep 2018 18:28 (GMT+7)</CreatedDate>
           <IpAddress>10.10.10.10</IpAddress>
         </item>
         <item>
           <CommentText>rqst erman cc as rmk</CommentText>
           <CreatedBy>SRIWIJAWA_CS_USERNAME</CreatedBy>
           <CreatedDate>18 Sep 2018 18:48 (GMT+7)</CreatedDate>
           <IpAddress>10.10.10.11</IpAddress>
         </item>
       </BookingRemarks>
       <AdditionalInformation></AdditionalInformation>
     </YourItineraryDetails>
     <ErrorCode>RETRIEVE0000</ErrorCode>
     <ErrorMessage>Success.</ErrorMessage>
   </return>
 </WsRetrievePNRResponse>
```