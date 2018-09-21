# WsGeneratePNR: Generate PNR (Booking) Method

Generate PNR/Booking (WsGeneratePNR) is a service method to make/generate booking (PNR) from Sriwijaya Air Web Service (SOAP) v.111 [[1](https://wsp.sriwijayaair.co.id:11443/wsdl.eticketv111/index.php)].

> In this example you will learn about using Generate PNR/Booking Method (WsGeneratePNR) with Go (using sjwsdk111 package). 

The following are the sections available in this guide.

- [What you'll build](#what-youll-build)
- [Prerequisites](#prerequisites)
- [Implementation](#implementation)
- [Build and Running](#build-and-running)

## What you’ll build
Let’s make a real world simple application for make a booking (PNR) by using Sriwijaya Air Web Services Endpoint. Following diagram demonstrates the booking use case.

![Generate PNR/Booking Diagram](images/02.WsGeneratePNR-1.png "Generate PNR/Booking Diagram")


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
    │   ├── 01.WsGeneratePNR
    │   │   ├── README.md
    │   │   ├── build_and_run.sh
    │   │   └── main.go
    └── wsdl
        └── wsp-wsdl.eticketv111.wsdl
```

- Create the above directories in your local machine and also create empty `main.go` and `build_and_run.sh` files.

- Download Sriwijaya Air Web Service WSDL and saved to `wsp-wsdl.eticketv111.wsdl`.


### Developing the application

Let's make a simple application for make a booking by using `sjwsdk111` package. 

##### Main code for WsGeneratePNR (main.go)
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
					<Key xsi:type="xsd:string">SELECTED_DEPARTURE_SEGMENT_CLASS_KEY_01</Key>
					<Category xsi:type="xsd:string">Departure</Category>
				</item>
				<item xsi:type="urn:InputReqArrayKeys">
					<Key xsi:type="xsd:string">SELECTED_DEPARTURE_SEGMENT_CLASS_KEY_02_IF_CONNECTING</Key>
					<Category xsi:type="xsd:string">Departure</Category>
				</item>
				<item xsi:type="urn:InputReqArrayKeys">
					<Key xsi:type="xsd:string">SELECTED_DEPARTURE_SEGMENT_CLASS_KEY_03_IF_CONNECTING</Key>
					<Category xsi:type="xsd:string">Departure</Category>
				</item>
				<item xsi:type="tns:InputReqArrayKeys">
					<Key xsi:type="xsd:string">SELECTED_RETURN_SEGMENT_CLASS_KEY_01_IF_ROUNDTRIP</Key>
					<Category xsi:type="xsd:string">Return</Category>
				</item>
				<item xsi:type="tns:InputReqArrayKeys">
					<Key xsi:type="xsd:string">SELECTED_RETURN_SEGMENT_CLASS_KEY_02_IF_ROUNDTRIP_CONNECTING</Key>
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

```

##### Bash code for building and running the example application (build_and_run.sh)
```bash
echo "Clean..."
rm ./WsGeneratePNR
echo "Build..."
go build -o WsGeneratePNR main.go 
echo "Build Done."
echo "Run..."
./WsGeneratePNR > WsGeneratePNR-Result.xml
echo "Done."

```


## Build and Running

You can build and running by execute the "build_and_run.sh" bash files. 

```bash
   $ sh build_and_run.sh 
```

After the application is running, you will get the xml response in `WsGeneratePNR-Result.xml` files.

## Sample Response

```xml
 <WsGeneratePNRResponse>
   <return>
     <Username>SRIWIJAWA_AGENT_USERNAME</Username>
     <BookingCode>ABMNYZ</BookingCode>
     <YourItineraryDetails>
       <ReservationDetails>
         <BookingCode>ABMNYZ</BookingCode>
         <BookingDate>18 Sep 2018 08:01 (GMT+7)</BookingDate>
         <BalanceDue>2025000</BalanceDue>
         <BalanceDueRemarks>*Extra Cover Insurance (STI) not include in balance due.</BalanceDueRemarks>
         <CurrencyCode>IDR</CurrencyCode>
         <Time>18 Sep 2018 14:14 (GMT+7)</Time>
         <TimeDescription>TimeLimit</TimeDescription>
         <Status>Hold</Status>
       </ReservationDetails>
       <PassengerDetails>
         <item>
           <No>1</No>
           <Suffix>Mrs</Suffix>
           <FirstName>ANGKASA</FirstName>
           <LastName>SRIWIJAYA</LastName>
           <SeatQty>1</SeatQty>
           <TicketNumber>N/A</TicketNumber>
           <SpecialRequest>N/A</SpecialRequest>
         </item>
         <item>
           <No>2</No>
           <Suffix>Mstr</Suffix>
           <FirstName>MAS</FirstName>
           <LastName>SRIWIJAYA</LastName>
           <SeatQty>1</SeatQty>
           <TicketNumber>N/A</TicketNumber>
           <SpecialRequest>N/A</SpecialRequest>
         </item>
         <item>
           <No>3</No>
           <Suffix>Inf</Suffix>
           <FirstName>ANANDA</FirstName>
           <LastName>SRIWIJAYA</LastName>
           <SeatQty>0</SeatQty>
           <TicketNumber>N/X</TicketNumber>
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
         <IssuedBy>-</IssuedBy>
       </AgentDetails>
       <BookingRemarks></BookingRemarks>
       <AdditionalInformation></AdditionalInformation>
     </YourItineraryDetails>
     <ErrorCode>RETRIEVE0000</ErrorCode>
     <ErrorMessage>Success.</ErrorMessage>
   </return>
 </WsGeneratePNRResponse>
```