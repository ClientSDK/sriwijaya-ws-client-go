# WsSearchFlight: Search Flight Method

Search Flight Method (WsSearchFlight) is a service method to search flight schedule information from Sriwijaya Air Web Service (SOAP) v.111 [[1](https://wsp.sriwijayaair.co.id:11443/wsdl.eticketv111/index.php)].

> In this example you will learn about using  Search Flight Method (WsSearchFlight) with Go (sjwsdk111 package). 

The following are the sections available in this guide.

- [What you'll build](#what-youll-build)
- [Prerequisites](#prerequisites)
- [Implementation](#implementation)
- [Build and Running](#build-and-running)

## What you’ll build
Let’s make a real world simple application for online search flight schedule using Sriwijaya Air Web Services Endpoint. Following diagram demonstrates the Search Flight use case.

![Search Flight Diagram](images/01.WsSearchFlight-2.png "Search Flight Diagram")


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
    │   ├── 01.WsSearchFlight
    │   │   ├── README.md
    │   │   ├── build_and_run.sh
    │   │   └── main.go
    └── wsdl
        └── wsp-wsdl.eticketv111.wsdl
```

- Create the above directories in your local machine and also create empty `main.go` and `build_and_run.sh` files.

- Download Sriwijaya Air Web Service WSDL and saved to `wsp-wsdl.eticketv111.wsdl`.


### Developing the application

Let's make a simple application for searching flight schedule information using `sjwsdk111` package. 

##### Main code for WsSearchFlight (main.go)
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

```

##### Bash code for building and running the example application (build_and_run.sh)
```bash
echo "Clean..."
rm ./WsSearchFlight
echo "Build..."
go build -o WsSearchFlight main.go 
echo "Build Done..."
echo "Run..."
./WsSearchFlight > WsSearchFlight-Result.xml
echo "Done."

```


## Build and Running

You can build and running by execute the "build_and_run.sh" bash files. 

```bash
   $ sh build_and_run.sh 
```

After the application is running, you will get the xml response in `WsSearchFlight-Result.xml` files.

## Sample Response

```xml
 <WsSearchFlightResponse>
   <return>
     <Username>SRIWIJAWA_AGENT_USERNAME</Username>
     <Adult>1</Adult>
     <Child>1</Child>
     <Infant>1</Infant>
     <TripDetail>
       <item>
         <CityFrom>TKG</CityFrom>
         <CityTo>CGK</CityTo>
         <Category>Departure</Category>
         <FlightRoute>
           <item>
             <CityFrom>TKG</CityFrom>
             <CityTo>CGK</CityTo>
             <Std>18-NOV-18 06.00.00 AM</Std>
             <Sta>18-NOV-18 06.45.00 AM</Sta>
             <FlightTime>45</FlightTime>
             <Segments>
               <item>
                 <CarrierCode>SJ</CarrierCode>
                 <NoFlight>095</NoFlight>
                 <DepartureStation>TKG</DepartureStation>
                 <ArrivalStation>CGK</ArrivalStation>
                 <Std>18-NOV-18 06.00.00 AM</Std>
                 <Sta>18-NOV-18 06.45.00 AM</Sta>
                 <Legs>
                   <item>
                     <DepartureStation>TKG</DepartureStation>
                     <ArrivalStation>CGK</ArrivalStation>
                     <Std>18-NOV-18 06.00.00 AM</Std>
                     <Sta>18-NOV-18 06.45.00 AM</Sta>
                   </item>
                 </Legs>
               </item>
             </Segments>
             <ClassesAvailable>
               <item>
                 <item>
                   <Key>-</Key>
                   <Availability>0</Availability>
                   <Class>X</Class>
                   <SeatAvail>CLOSE</SeatAvail>
                   <Price>0</Price>
                   <PriceDetail></PriceDetail>
                   <Currency>IDR</Currency>
                   <StatusAvail>Normal</StatusAvail>
                 </item>
                 <item>
                    ...
                 </item>
               </item>
             </ClassesAvailable>
           </item>
           <item>
             <CityFrom>TKG</CityFrom>
             <CityTo>CGK</CityTo>
             <Std>18-NOV-18 09.00.00 AM</Std>
             <Sta>18-NOV-18 09.45.00 AM</Sta>
             <FlightTime>45</FlightTime>
             <Segments>
               <item>
                 <CarrierCode>SJ</CarrierCode>
                 <NoFlight>089</NoFlight>
                 <DepartureStation>TKG</DepartureStation>
                 <ArrivalStation>CGK</ArrivalStation>
                 <Std>18-NOV-18 09.00.00 AM</Std>
                 <Sta>18-NOV-18 09.45.00 AM</Sta>
                 <Legs>
                   <item>
                     <DepartureStation>TKG</DepartureStation>
                     <ArrivalStation>CGK</ArrivalStation>
                     <Std>18-NOV-18 09.00.00 AM</Std>
                     <Sta>18-NOV-18 09.45.00 AM</Sta>
                   </item>
                 </Legs>
               </item>
             </Segments>
             <ClassesAvailable>
               <item>
                 <item>
                   <Key>-</Key>
                   <Availability>0</Availability>
                   <Class>X</Class>
                   <SeatAvail>CLOSE</SeatAvail>
                   <Price>0</Price>
                   <PriceDetail></PriceDetail>
                   <Currency>IDR</Currency>
                   <StatusAvail>Normal</StatusAvail>
                 </item>
                 <item>
                    ...
                 </item>
                 <item>
                   <Key>SEGMENT_CLASS_KEY:T:S</Key>
                   <Availability>0</Availability>
                   <Class>T</Class>
                   <SeatAvail>OPEN</SeatAvail>
                   <Price>1193200</Price>
                   <PriceDetail>
                     <item>
                       <PaxCategory>ADT</PaxCategory>
                       <Total_1>558800</Total_1>
                       <Nta_1>523900</Nta_1>
                       <FareComponent>
                         <item>
                           <FareChargeTypeCode>BF</FareChargeTypeCode>
                           <FareChargeTypeDesc>Basic Fare</FareChargeTypeDesc>
                           <Amount>458000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>PSC</FareChargeTypeCode>
                           <FareChargeTypeDesc>PSC Fee</FareChargeTypeDesc>
                           <Amount>30000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>VAT</FareChargeTypeCode>
                           <FareChargeTypeDesc>PPN - TAX</FareChargeTypeDesc>
                           <Amount>45800</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>IWJR</FareChargeTypeCode>
                           <FareChargeTypeDesc>Government Insurance</FareChargeTypeDesc>
                           <Amount>5000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>SC</FareChargeTypeCode>
                           <FareChargeTypeDesc>Surcharge</FareChargeTypeDesc>
                           <Amount>0</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>STI</FareChargeTypeCode>
                           <FareChargeTypeDesc>Extra Cover</FareChargeTypeDesc>
                           <Amount>20000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>AC</FareChargeTypeCode>
                           <FareChargeTypeDesc>Agent Discount</FareChargeTypeDesc>
                           <Amount>-24900</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>INCV</FareChargeTypeCode>
                           <FareChargeTypeDesc>Incentive</FareChargeTypeDesc>
                           <Amount>-10000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                       </FareComponent>
                     </item>
                     <item>
                       <PaxCategory>CHD</PaxCategory>
                       <Total_1>558800</Total_1>
                       <Nta_1>523900</Nta_1>
                       <FareComponent>
                         <item>
                           <FareChargeTypeCode>BF</FareChargeTypeCode>
                           <FareChargeTypeDesc>Basic Fare</FareChargeTypeDesc>
                           <Amount>458000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>PSC</FareChargeTypeCode>
                           <FareChargeTypeDesc>PSC Fee</FareChargeTypeDesc>
                           <Amount>30000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>VAT</FareChargeTypeCode>
                           <FareChargeTypeDesc>PPN - TAX</FareChargeTypeDesc>
                           <Amount>45800</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>IWJR</FareChargeTypeCode>
                           <FareChargeTypeDesc>Government Insurance</FareChargeTypeDesc>
                           <Amount>5000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>SC</FareChargeTypeCode>
                           <FareChargeTypeDesc>Surcharge</FareChargeTypeDesc>
                           <Amount>0</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>STI</FareChargeTypeCode>
                           <FareChargeTypeDesc>Extra Cover</FareChargeTypeDesc>
                           <Amount>20000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>AC</FareChargeTypeCode>
                           <FareChargeTypeDesc>Agent Discount</FareChargeTypeDesc>
                           <Amount>-24900</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>INCV</FareChargeTypeCode>
                           <FareChargeTypeDesc>Incentive</FareChargeTypeDesc>
                           <Amount>-10000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                       </FareComponent>
                     </item>
                     <item>
                       <PaxCategory>INF</PaxCategory>
                       <Total_1>75600</Total_1>
                       <Nta_1>73600</Nta_1>
                       <FareComponent>
                         <item>
                           <FareChargeTypeCode>BF</FareChargeTypeCode>
                           <FareChargeTypeDesc>Basic Fare</FareChargeTypeDesc>
                           <Amount>46000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>VAT</FareChargeTypeCode>
                           <FareChargeTypeDesc>PPN - TAX</FareChargeTypeDesc>
                           <Amount>4600</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>IWJR</FareChargeTypeCode>
                           <FareChargeTypeDesc>Government Insurance</FareChargeTypeDesc>
                           <Amount>5000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>SC</FareChargeTypeCode>
                           <FareChargeTypeDesc>Surcharge</FareChargeTypeDesc>
                           <Amount>0</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>STI</FareChargeTypeCode>
                           <FareChargeTypeDesc>Extra Cover</FareChargeTypeDesc>
                           <Amount>20000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>AC</FareChargeTypeCode>
                           <FareChargeTypeDesc>Agent Discount</FareChargeTypeDesc>
                           <Amount>-2000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                       </FareComponent>
                     </item>
                   </PriceDetail>
                   <Currency>IDR</Currency>
                   <StatusAvail>(WL)</StatusAvail>
                 </item>
                 <item>
                    ...
                 </item>
               </item>
             </ClassesAvailable>
           </item>
           <item>
             <CityFrom>TKG</CityFrom>
             <CityTo>CGK</CityTo>
             <Std>18-NOV-18 12.25.00 PM</Std>
             <Sta>18-NOV-18 01.10.00 PM</Sta>
             <FlightTime>45</FlightTime>
             <Segments>
               <item>
                 <CarrierCode>SJ</CarrierCode>
                 <NoFlight>097</NoFlight>
                 <DepartureStation>TKG</DepartureStation>
                 <ArrivalStation>CGK</ArrivalStation>
                 <Std>18-NOV-18 12.25.00 PM</Std>
                 <Sta>18-NOV-18 01.10.00 PM</Sta>
                 <Legs>
                   <item>
                     <DepartureStation>TKG</DepartureStation>
                     <ArrivalStation>CGK</ArrivalStation>
                     <Std>18-NOV-18 12.25.00 PM</Std>
                     <Sta>18-NOV-18 01.10.00 PM</Sta>
                   </item>
                 </Legs>
               </item>
             </Segments>
             <ClassesAvailable>
               <item>
                 <item>
                   <Key>-</Key>
                   <Availability>0</Availability>
                   <Class>X</Class>
                   <SeatAvail>CLOSE</SeatAvail>
                   <Price>0</Price>
                   <PriceDetail></PriceDetail>
                   <Currency>IDR</Currency>
                   <StatusAvail>Normal</StatusAvail>
                 </item>
                 <item>
                    ...
                 </item>
               </item>
             </ClassesAvailable>
           </item>
           <item>
             <CityFrom>TKG</CityFrom>
             <CityTo>CGK</CityTo>
             <Std>18-NOV-18 02.15.00 PM</Std>
             <Sta>18-NOV-18 03.05.00 PM</Sta>
             <FlightTime>50</FlightTime>
             <Segments>
               <item>
                 <CarrierCode>SJ</CarrierCode>
                 <NoFlight>069</NoFlight>
                 <DepartureStation>TKG</DepartureStation>
                 <ArrivalStation>CGK</ArrivalStation>
                 <Std>18-NOV-18 02.15.00 PM</Std>
                 <Sta>18-NOV-18 03.05.00 PM</Sta>
                 <Legs>
                   <item>
                     <DepartureStation>TKG</DepartureStation>
                     <ArrivalStation>CGK</ArrivalStation>
                     <Std>18-NOV-18 02.15.00 PM</Std>
                     <Sta>18-NOV-18 03.05.00 PM</Sta>
                   </item>
                 </Legs>
               </item>
             </Segments>
             <ClassesAvailable>
               <item>
                 <item>
                   <Key>-</Key>
                   <Availability>0</Availability>
                   <Class>X</Class>
                   <SeatAvail>CLOSE</SeatAvail>
                   <Price>0</Price>
                   <PriceDetail></PriceDetail>
                   <Currency>IDR</Currency>
                   <StatusAvail>Normal</StatusAvail>
                 </item>
                 <item>
                 </item>
               </item>
             </ClassesAvailable>
           </item>
           <item>
             <CityFrom>TKG</CityFrom>
             <CityTo>CGK</CityTo>
             <Std>18-NOV-18 05.30.00 PM</Std>
             <Sta>18-NOV-18 06.15.00 PM</Sta>
             <FlightTime>45</FlightTime>
             <Segments>
               <item>
                 <CarrierCode>SJ</CarrierCode>
                 <NoFlight>059</NoFlight>
                 <DepartureStation>TKG</DepartureStation>
                 <ArrivalStation>CGK</ArrivalStation>
                 <Std>18-NOV-18 05.30.00 PM</Std>
                 <Sta>18-NOV-18 06.15.00 PM</Sta>
                 <Legs>
                   <item>
                     <DepartureStation>TKG</DepartureStation>
                     <ArrivalStation>CGK</ArrivalStation>
                     <Std>18-NOV-18 05.30.00 PM</Std>
                     <Sta>18-NOV-18 06.15.00 PM</Sta>
                   </item>
                 </Legs>
               </item>
             </Segments>
             <ClassesAvailable>
               <item>
                 <item>
                   <Key>-</Key>
                   <Availability>0</Availability>
                   <Class>X</Class>
                   <SeatAvail>CLOSE</SeatAvail>
                   <Price>0</Price>
                   <PriceDetail></PriceDetail>
                   <Currency>IDR</Currency>
                   <StatusAvail>Normal</StatusAvail>
                 </item>
                 <item>
                    ...
                 </item>
               </item>
             </ClassesAvailable>
           </item>
           <item>
             <CityFrom>TKG</CityFrom>
             <CityTo>CGK</CityTo>
             <Std>18-NOV-18 06.10.00 PM</Std>
             <Sta>18-NOV-18 07.00.00 PM</Sta>
             <FlightTime>50</FlightTime>
             <Segments>
               <item>
                 <CarrierCode>IN</CarrierCode>
                 <NoFlight>099</NoFlight>
                 <DepartureStation>TKG</DepartureStation>
                 <ArrivalStation>CGK</ArrivalStation>
                 <Std>18-NOV-18 06.10.00 PM</Std>
                 <Sta>18-NOV-18 07.00.00 PM</Sta>
                 <Legs>
                   <item>
                     <DepartureStation>TKG</DepartureStation>
                     <ArrivalStation>CGK</ArrivalStation>
                     <Std>18-NOV-18 06.10.00 PM</Std>
                     <Sta>18-NOV-18 07.00.00 PM</Sta>
                   </item>
                 </Legs>
               </item>
             </Segments>
             <ClassesAvailable>
               <item>
                 <item>
                   <Key>-</Key>
                   <Availability>0</Availability>
                   <Class>X</Class>
                   <SeatAvail>CLOSE</SeatAvail>
                   <Price>0</Price>
                   <PriceDetail></PriceDetail>
                   <Currency>IDR</Currency>
                   <StatusAvail>Normal</StatusAvail>
                 </item>
                 <item>
                    ...
                 </item>
               </item>
             </ClassesAvailable>
           </item>
           <item>
             <CityFrom>TKG</CityFrom>
             <CityTo>CGK</CityTo>
             <Std>18-NOV-18 08.00.00 PM</Std>
             <Sta>18-NOV-18 08.45.00 PM</Sta>
             <FlightTime>45</FlightTime>
             <Segments>
               <item>
                 <CarrierCode>SJ</CarrierCode>
                 <NoFlight>087</NoFlight>
                 <DepartureStation>TKG</DepartureStation>
                 <ArrivalStation>CGK</ArrivalStation>
                 <Std>18-NOV-18 08.00.00 PM</Std>
                 <Sta>18-NOV-18 08.45.00 PM</Sta>
                 <Legs>
                   <item>
                     <DepartureStation>TKG</DepartureStation>
                     <ArrivalStation>CGK</ArrivalStation>
                     <Std>18-NOV-18 08.00.00 PM</Std>
                     <Sta>18-NOV-18 08.45.00 PM</Sta>
                   </item>
                 </Legs>
               </item>
             </Segments>
             <ClassesAvailable>
               <item>
                 <item>
                   <Key>-</Key>
                   <Availability>0</Availability>
                   <Class>X</Class>
                   <SeatAvail>CLOSE</SeatAvail>
                   <Price>0</Price>
                   <PriceDetail></PriceDetail>
                   <Currency>IDR</Currency>
                   <StatusAvail>Normal</StatusAvail>
                 </item>
                 <item>
                    ...
                 </item>
               </item>
             </ClassesAvailable>
           </item>
         </FlightRoute>
       </item>
       <item>
         <CityFrom>CGK</CityFrom>
         <CityTo>TKG</CityTo>
         <Category>Return</Category>
         <FlightRoute>
           <item>
             <CityFrom>CGK</CityFrom>
             <CityTo>TKG</CityTo>
             <Std>28-SEP-18 07.00.00 AM</Std>
             <Sta>28-SEP-18 07.40.00 AM</Sta>
             <FlightTime>40</FlightTime>
             <Segments>
               <item>
                 <CarrierCode>SJ</CarrierCode>
                 <NoFlight>088</NoFlight>
                 <DepartureStation>CGK</DepartureStation>
                 <ArrivalStation>TKG</ArrivalStation>
                 <Std>28-SEP-18 07.00.00 AM</Std>
                 <Sta>28-SEP-18 07.40.00 AM</Sta>
                 <Legs>
                   <item>
                     <DepartureStation>CGK</DepartureStation>
                     <ArrivalStation>TKG</ArrivalStation>
                     <Std>28-SEP-18 07.00.00 AM</Std>
                     <Sta>28-SEP-18 07.40.00 AM</Sta>
                   </item>
                 </Legs>
               </item>
             </Segments>
             <ClassesAvailable>
               <item>
                 <item>
                   <Key>-</Key>
                   <Availability>0</Availability>
                   <Class>X</Class>
                   <SeatAvail>CLOSE</SeatAvail>
                   <Price>0</Price>
                   <PriceDetail></PriceDetail>
                   <Currency>IDR</Currency>
                   <StatusAvail>Normal</StatusAvail>
                 </item>
                 <item>
                    ...
                 </item>
               </item>
             </ClassesAvailable>
           </item>
           <item>
             <CityFrom>CGK</CityFrom>
             <CityTo>TKG</CityTo>
             <Std>28-SEP-18 08.00.00 AM</Std>
             <Sta>28-SEP-18 08.55.00 AM</Sta>
             <FlightTime>55</FlightTime>
             <Segments>
               <item>
                 <CarrierCode>SJ</CarrierCode>
                 <NoFlight>068</NoFlight>
                 <DepartureStation>CGK</DepartureStation>
                 <ArrivalStation>TKG</ArrivalStation>
                 <Std>28-SEP-18 08.00.00 AM</Std>
                 <Sta>28-SEP-18 08.55.00 AM</Sta>
                 <Legs>
                   <item>
                     <DepartureStation>CGK</DepartureStation>
                     <ArrivalStation>TKG</ArrivalStation>
                     <Std>28-SEP-18 08.00.00 AM</Std>
                     <Sta>28-SEP-18 08.55.00 AM</Sta>
                   </item>
                 </Legs>
               </item>
             </Segments>
             <ClassesAvailable>
               <item>
                 <item>
                   <Key>-</Key>
                   <Availability>0</Availability>
                   <Class>X</Class>
                   <SeatAvail>CLOSE</SeatAvail>
                   <Price>0</Price>
                   <PriceDetail></PriceDetail>
                   <Currency>IDR</Currency>
                   <StatusAvail>Normal</StatusAvail>
                 </item>
                 <item>
                   <Key>-</Key>
                   <Availability>0</Availability>
                   <Class>E</Class>
                   <SeatAvail>CLOSE</SeatAvail>
                   <Price>0</Price>
                   <PriceDetail></PriceDetail>
                   <Currency>IDR</Currency>
                   <StatusAvail>Normal</StatusAvail>
                 </item>
                 <item>
                   <Key>-</Key>
                   <Availability>0</Availability>
                   <Class>G</Class>
                   <SeatAvail>CLOSE</SeatAvail>
                   <Price>0</Price>
                   <PriceDetail></PriceDetail>
                   <Currency>IDR</Currency>
                   <StatusAvail>Normal</StatusAvail>
                 </item>
                 <item>
                   <Key>-</Key>
                   <Availability>0</Availability>
                   <Class>V</Class>
                   <SeatAvail>CLOSE</SeatAvail>
                   <Price>0</Price>
                   <PriceDetail></PriceDetail>
                   <Currency>IDR</Currency>
                   <StatusAvail>Normal</StatusAvail>
                 </item>
                 <item>
                   <Key>-</Key>
                   <Availability>0</Availability>
                   <Class>T</Class>
                   <SeatAvail>CLOSE</SeatAvail>
                   <Price>0</Price>
                   <PriceDetail></PriceDetail>
                   <Currency>IDR</Currency>
                   <StatusAvail>Normal</StatusAvail>
                 </item>
                 <item>
                   <Key>-</Key>
                   <Availability>0</Availability>
                   <Class>Q</Class>
                   <SeatAvail>CLOSE</SeatAvail>
                   <Price>0</Price>
                   <PriceDetail></PriceDetail>
                   <Currency>IDR</Currency>
                   <StatusAvail>Normal</StatusAvail>
                 </item>
                 <item>
                   <Key>-</Key>
                   <Availability>0</Availability>
                   <Class>N</Class>
                   <SeatAvail>CLOSE</SeatAvail>
                   <Price>0</Price>
                   <PriceDetail></PriceDetail>
                   <Currency>IDR</Currency>
                   <StatusAvail>Normal</StatusAvail>
                 </item>
                 <item>
                   <Key>-</Key>
                   <Availability>0</Availability>
                   <Class>M</Class>
                   <SeatAvail>CLOSE</SeatAvail>
                   <Price>0</Price>
                   <PriceDetail></PriceDetail>
                   <Currency>IDR</Currency>
                   <StatusAvail>Normal</StatusAvail>
                 </item>
                 <item>
                   <Key>-</Key>
                   <Availability>0</Availability>
                   <Class>L</Class>
                   <SeatAvail>CLOSE</SeatAvail>
                   <Price>0</Price>
                   <PriceDetail></PriceDetail>
                   <Currency>IDR</Currency>
                   <StatusAvail>Normal</StatusAvail>
                 </item>
                 <item>
                   <Key>-</Key>
                   <Availability>0</Availability>
                   <Class>K</Class>
                   <SeatAvail>CLOSE</SeatAvail>
                   <Price>0</Price>
                   <PriceDetail></PriceDetail>
                   <Currency>IDR</Currency>
                   <StatusAvail>Normal</StatusAvail>
                 </item>
                 <item>
                   <Key>-</Key>
                   <Availability>0</Availability>
                   <Class>H</Class>
                   <SeatAvail>CLOSE</SeatAvail>
                   <Price>0</Price>
                   <PriceDetail></PriceDetail>
                   <Currency>IDR</Currency>
                   <StatusAvail>Normal</StatusAvail>
                 </item>
                 <item>
                   <Key>-</Key>
                   <Availability>0</Availability>
                   <Class>B</Class>
                   <SeatAvail>CLOSE</SeatAvail>
                   <Price>0</Price>
                   <PriceDetail></PriceDetail>
                   <Currency>IDR</Currency>
                   <StatusAvail>Normal</StatusAvail>
                 </item>
                 <item>
                   <Key>-</Key>
                   <Availability>0</Availability>
                   <Class>W</Class>
                   <SeatAvail>CLOSE</SeatAvail>
                   <Price>0</Price>
                   <PriceDetail></PriceDetail>
                   <Currency>IDR</Currency>
                   <StatusAvail>Normal</StatusAvail>
                 </item>
                 <item>
                   <Key>-</Key>
                   <Availability>0</Availability>
                   <Class>S</Class>
                   <SeatAvail>CLOSE</SeatAvail>
                   <Price>0</Price>
                   <PriceDetail></PriceDetail>
                   <Currency>IDR</Currency>
                   <StatusAvail>Normal</StatusAvail>
                 </item>
                 <item>
                   <Key>18077711:Y:S</Key>
                   <Availability>0</Availability>
                   <Class>Y</Class>
                   <SeatAvail>OPEN</SeatAvail>
                   <Price>1303200</Price>
                   <PriceDetail>
                     <item>
                       <PaxCategory>ADT</PaxCategory>
                       <Total_1>613800</Total_1>
                       <Nta_1>578900</Nta_1>
                       <FareComponent>
                         <item>
                           <FareChargeTypeCode>BF</FareChargeTypeCode>
                           <FareChargeTypeDesc>Basic Fare</FareChargeTypeDesc>
                           <Amount>458000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>PSC</FareChargeTypeCode>
                           <FareChargeTypeDesc>PSC Fee</FareChargeTypeDesc>
                           <Amount>85000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>VAT</FareChargeTypeCode>
                           <FareChargeTypeDesc>PPN - TAX</FareChargeTypeDesc>
                           <Amount>45800</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>IWJR</FareChargeTypeCode>
                           <FareChargeTypeDesc>Government Insurance</FareChargeTypeDesc>
                           <Amount>5000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>SC</FareChargeTypeCode>
                           <FareChargeTypeDesc>Surcharge</FareChargeTypeDesc>
                           <Amount>0</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>STI</FareChargeTypeCode>
                           <FareChargeTypeDesc>Extra Cover</FareChargeTypeDesc>
                           <Amount>20000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>AC</FareChargeTypeCode>
                           <FareChargeTypeDesc>Agent Discount</FareChargeTypeDesc>
                           <Amount>-24900</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>INCV</FareChargeTypeCode>
                           <FareChargeTypeDesc>Incentive</FareChargeTypeDesc>
                           <Amount>-10000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                       </FareComponent>
                     </item>
                     <item>
                       <PaxCategory>CHD</PaxCategory>
                       <Total_1>613800</Total_1>
                       <Nta_1>578900</Nta_1>
                       <FareComponent>
                         <item>
                           <FareChargeTypeCode>BF</FareChargeTypeCode>
                           <FareChargeTypeDesc>Basic Fare</FareChargeTypeDesc>
                           <Amount>458000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>PSC</FareChargeTypeCode>
                           <FareChargeTypeDesc>PSC Fee</FareChargeTypeDesc>
                           <Amount>85000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>VAT</FareChargeTypeCode>
                           <FareChargeTypeDesc>PPN - TAX</FareChargeTypeDesc>
                           <Amount>45800</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>IWJR</FareChargeTypeCode>
                           <FareChargeTypeDesc>Government Insurance</FareChargeTypeDesc>
                           <Amount>5000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>SC</FareChargeTypeCode>
                           <FareChargeTypeDesc>Surcharge</FareChargeTypeDesc>
                           <Amount>0</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>STI</FareChargeTypeCode>
                           <FareChargeTypeDesc>Extra Cover</FareChargeTypeDesc>
                           <Amount>20000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>AC</FareChargeTypeCode>
                           <FareChargeTypeDesc>Agent Discount</FareChargeTypeDesc>
                           <Amount>-24900</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>INCV</FareChargeTypeCode>
                           <FareChargeTypeDesc>Incentive</FareChargeTypeDesc>
                           <Amount>-10000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                       </FareComponent>
                     </item>
                     <item>
                       <PaxCategory>INF</PaxCategory>
                       <Total_1>75600</Total_1>
                       <Nta_1>73600</Nta_1>
                       <FareComponent>
                         <item>
                           <FareChargeTypeCode>BF</FareChargeTypeCode>
                           <FareChargeTypeDesc>Basic Fare</FareChargeTypeDesc>
                           <Amount>46000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>VAT</FareChargeTypeCode>
                           <FareChargeTypeDesc>PPN - TAX</FareChargeTypeDesc>
                           <Amount>4600</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>IWJR</FareChargeTypeCode>
                           <FareChargeTypeDesc>Government Insurance</FareChargeTypeDesc>
                           <Amount>5000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>SC</FareChargeTypeCode>
                           <FareChargeTypeDesc>Surcharge</FareChargeTypeDesc>
                           <Amount>0</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>STI</FareChargeTypeCode>
                           <FareChargeTypeDesc>Extra Cover</FareChargeTypeDesc>
                           <Amount>20000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>AC</FareChargeTypeCode>
                           <FareChargeTypeDesc>Agent Discount</FareChargeTypeDesc>
                           <Amount>-2000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                       </FareComponent>
                     </item>
                   </PriceDetail>
                   <Currency>IDR</Currency>
                   <StatusAvail>(WL)</StatusAvail>
                 </item>
                 <item>
                   <Key>-</Key>
                   <Availability>0</Availability>
                   <Class>I</Class>
                   <SeatAvail>CLOSE</SeatAvail>
                   <Price>0</Price>
                   <PriceDetail></PriceDetail>
                   <Currency>IDR</Currency>
                   <StatusAvail>Normal</StatusAvail>
                 </item>
                 <item>
                   <Key>-</Key>
                   <Availability>0</Availability>
                   <Class>D</Class>
                   <SeatAvail>CLOSE</SeatAvail>
                   <Price>0</Price>
                   <PriceDetail></PriceDetail>
                   <Currency>IDR</Currency>
                   <StatusAvail>Normal</StatusAvail>
                 </item>
                 <item>
                   <Key>18077702:C:S</Key>
                   <Availability>0</Availability>
                   <Class>C</Class>
                   <SeatAvail>OPEN</SeatAvail>
                   <Price>2254700</Price>
                   <PriceDetail>
                     <item>
                       <PaxCategory>ADT</PaxCategory>
                       <Total_1>1067000</Total_1>
                       <Nta_1>1011500</Nta_1>
                       <FareComponent>
                         <item>
                           <FareChargeTypeCode>BF</FareChargeTypeCode>
                           <FareChargeTypeDesc>Basic Fare</FareChargeTypeDesc>
                           <Amount>870000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>PSC</FareChargeTypeCode>
                           <FareChargeTypeDesc>PSC Fee</FareChargeTypeDesc>
                           <Amount>85000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>VAT</FareChargeTypeCode>
                           <FareChargeTypeDesc>PPN - TAX</FareChargeTypeDesc>
                           <Amount>87000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>IWJR</FareChargeTypeCode>
                           <FareChargeTypeDesc>Government Insurance</FareChargeTypeDesc>
                           <Amount>5000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>SC</FareChargeTypeCode>
                           <FareChargeTypeDesc>Surcharge</FareChargeTypeDesc>
                           <Amount>0</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>STI</FareChargeTypeCode>
                           <FareChargeTypeDesc>Extra Cover</FareChargeTypeDesc>
                           <Amount>20000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>AC</FareChargeTypeCode>
                           <FareChargeTypeDesc>Agent Discount</FareChargeTypeDesc>
                           <Amount>-45500</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>INCV</FareChargeTypeCode>
                           <FareChargeTypeDesc>Incentive</FareChargeTypeDesc>
                           <Amount>-10000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                       </FareComponent>
                     </item>
                     <item>
                       <PaxCategory>CHD</PaxCategory>
                       <Total_1>1067000</Total_1>
                       <Nta_1>1011500</Nta_1>
                       <FareComponent>
                         <item>
                           <FareChargeTypeCode>BF</FareChargeTypeCode>
                           <FareChargeTypeDesc>Basic Fare</FareChargeTypeDesc>
                           <Amount>870000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>PSC</FareChargeTypeCode>
                           <FareChargeTypeDesc>PSC Fee</FareChargeTypeDesc>
                           <Amount>85000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>VAT</FareChargeTypeCode>
                           <FareChargeTypeDesc>PPN - TAX</FareChargeTypeDesc>
                           <Amount>87000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>IWJR</FareChargeTypeCode>
                           <FareChargeTypeDesc>Government Insurance</FareChargeTypeDesc>
                           <Amount>5000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>SC</FareChargeTypeCode>
                           <FareChargeTypeDesc>Surcharge</FareChargeTypeDesc>
                           <Amount>0</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>STI</FareChargeTypeCode>
                           <FareChargeTypeDesc>Extra Cover</FareChargeTypeDesc>
                           <Amount>20000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>AC</FareChargeTypeCode>
                           <FareChargeTypeDesc>Agent Discount</FareChargeTypeDesc>
                           <Amount>-45500</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>INCV</FareChargeTypeCode>
                           <FareChargeTypeDesc>Incentive</FareChargeTypeDesc>
                           <Amount>-10000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                       </FareComponent>
                     </item>
                     <item>
                       <PaxCategory>INF</PaxCategory>
                       <Total_1>120700</Total_1>
                       <Nta_1>118700</Nta_1>
                       <FareComponent>
                         <item>
                           <FareChargeTypeCode>BF</FareChargeTypeCode>
                           <FareChargeTypeDesc>Basic Fare</FareChargeTypeDesc>
                           <Amount>87000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>VAT</FareChargeTypeCode>
                           <FareChargeTypeDesc>PPN - TAX</FareChargeTypeDesc>
                           <Amount>8700</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>IWJR</FareChargeTypeCode>
                           <FareChargeTypeDesc>Government Insurance</FareChargeTypeDesc>
                           <Amount>5000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>SC</FareChargeTypeCode>
                           <FareChargeTypeDesc>Surcharge</FareChargeTypeDesc>
                           <Amount>0</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>STI</FareChargeTypeCode>
                           <FareChargeTypeDesc>Extra Cover</FareChargeTypeDesc>
                           <Amount>20000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>AC</FareChargeTypeCode>
                           <FareChargeTypeDesc>Agent Discount</FareChargeTypeDesc>
                           <Amount>-2000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                       </FareComponent>
                     </item>
                   </PriceDetail>
                   <Currency>IDR</Currency>
                   <StatusAvail>(WL)</StatusAvail>
                 </item>
               </item>
             </ClassesAvailable>
           </item>
           <item>
             <CityFrom>CGK</CityFrom>
             <CityTo>TKG</CityTo>
             <Std>28-SEP-18 11.10.00 AM</Std>
             <Sta>28-SEP-18 11.50.00 AM</Sta>
             <FlightTime>40</FlightTime>
             <Segments>
               <item>
                 <CarrierCode>SJ</CarrierCode>
                 <NoFlight>096</NoFlight>
                 <DepartureStation>CGK</DepartureStation>
                 <ArrivalStation>TKG</ArrivalStation>
                 <Std>28-SEP-18 11.10.00 AM</Std>
                 <Sta>28-SEP-18 11.50.00 AM</Sta>
                 <Legs>
                   <item>
                     <DepartureStation>CGK</DepartureStation>
                     <ArrivalStation>TKG</ArrivalStation>
                     <Std>28-SEP-18 11.10.00 AM</Std>
                     <Sta>28-SEP-18 11.50.00 AM</Sta>
                   </item>
                 </Legs>
               </item>
             </Segments>
             <ClassesAvailable>
               <item>
                 <item>
                   <Key>-</Key>
                   <Availability>0</Availability>
                   <Class>X</Class>
                   <SeatAvail>CLOSE</SeatAvail>
                   <Price>0</Price>
                   <PriceDetail></PriceDetail>
                   <Currency>IDR</Currency>
                   <StatusAvail>Normal</StatusAvail>
                 </item>
                 <item>
                   ...
                 </item>
                 <item>
                   <Key>SEGMENT_CLASS_KEY:Q:S</Key>
                   <Availability>0</Availability>
                   <Class>Q</Class>
                   <SeatAvail>OPEN</SeatAvail>
                   <Price>1303200</Price>
                   <PriceDetail>
                     <item>
                       <PaxCategory>ADT</PaxCategory>
                       <Total_1>613800</Total_1>
                       <Nta_1>578900</Nta_1>
                       <FareComponent>
                         <item>
                           <FareChargeTypeCode>BF</FareChargeTypeCode>
                           <FareChargeTypeDesc>Basic Fare</FareChargeTypeDesc>
                           <Amount>458000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>PSC</FareChargeTypeCode>
                           <FareChargeTypeDesc>PSC Fee</FareChargeTypeDesc>
                           <Amount>85000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>VAT</FareChargeTypeCode>
                           <FareChargeTypeDesc>PPN - TAX</FareChargeTypeDesc>
                           <Amount>45800</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>IWJR</FareChargeTypeCode>
                           <FareChargeTypeDesc>Government Insurance</FareChargeTypeDesc>
                           <Amount>5000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>SC</FareChargeTypeCode>
                           <FareChargeTypeDesc>Surcharge</FareChargeTypeDesc>
                           <Amount>0</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>STI</FareChargeTypeCode>
                           <FareChargeTypeDesc>Extra Cover</FareChargeTypeDesc>
                           <Amount>20000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>AC</FareChargeTypeCode>
                           <FareChargeTypeDesc>Agent Discount</FareChargeTypeDesc>
                           <Amount>-24900</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>INCV</FareChargeTypeCode>
                           <FareChargeTypeDesc>Incentive</FareChargeTypeDesc>
                           <Amount>-10000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                       </FareComponent>
                     </item>
                     <item>
                       <PaxCategory>CHD</PaxCategory>
                       <Total_1>613800</Total_1>
                       <Nta_1>578900</Nta_1>
                       <FareComponent>
                         <item>
                           <FareChargeTypeCode>BF</FareChargeTypeCode>
                           <FareChargeTypeDesc>Basic Fare</FareChargeTypeDesc>
                           <Amount>458000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>PSC</FareChargeTypeCode>
                           <FareChargeTypeDesc>PSC Fee</FareChargeTypeDesc>
                           <Amount>85000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>VAT</FareChargeTypeCode>
                           <FareChargeTypeDesc>PPN - TAX</FareChargeTypeDesc>
                           <Amount>45800</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>IWJR</FareChargeTypeCode>
                           <FareChargeTypeDesc>Government Insurance</FareChargeTypeDesc>
                           <Amount>5000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>SC</FareChargeTypeCode>
                           <FareChargeTypeDesc>Surcharge</FareChargeTypeDesc>
                           <Amount>0</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>STI</FareChargeTypeCode>
                           <FareChargeTypeDesc>Extra Cover</FareChargeTypeDesc>
                           <Amount>20000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>AC</FareChargeTypeCode>
                           <FareChargeTypeDesc>Agent Discount</FareChargeTypeDesc>
                           <Amount>-24900</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>INCV</FareChargeTypeCode>
                           <FareChargeTypeDesc>Incentive</FareChargeTypeDesc>
                           <Amount>-10000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                       </FareComponent>
                     </item>
                     <item>
                       <PaxCategory>INF</PaxCategory>
                       <Total_1>75600</Total_1>
                       <Nta_1>73600</Nta_1>
                       <FareComponent>
                         <item>
                           <FareChargeTypeCode>BF</FareChargeTypeCode>
                           <FareChargeTypeDesc>Basic Fare</FareChargeTypeDesc>
                           <Amount>46000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>VAT</FareChargeTypeCode>
                           <FareChargeTypeDesc>PPN - TAX</FareChargeTypeDesc>
                           <Amount>4600</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>IWJR</FareChargeTypeCode>
                           <FareChargeTypeDesc>Government Insurance</FareChargeTypeDesc>
                           <Amount>5000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>SC</FareChargeTypeCode>
                           <FareChargeTypeDesc>Surcharge</FareChargeTypeDesc>
                           <Amount>0</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>STI</FareChargeTypeCode>
                           <FareChargeTypeDesc>Extra Cover</FareChargeTypeDesc>
                           <Amount>20000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                         <item>
                           <FareChargeTypeCode>AC</FareChargeTypeCode>
                           <FareChargeTypeDesc>Agent Discount</FareChargeTypeDesc>
                           <Amount>-2000</Amount>
                           <CurrencyCode>IDR</CurrencyCode>
                         </item>
                       </FareComponent>
                     </item>
                   </PriceDetail>
                   <Currency>IDR</Currency>
                   <StatusAvail>(WL)</StatusAvail>
                 </item>
                 <item>
                   ...
                 </item>
               </item>
             </ClassesAvailable>
           </item>
           <item>
             <CityFrom>CGK</CityFrom>
             <CityTo>TKG</CityTo>
             <Std>28-SEP-18 02.55.00 PM</Std>
             <Sta>28-SEP-18 03.35.00 PM</Sta>
             <FlightTime>40</FlightTime>
             <Segments>
               <item>
                 <CarrierCode>SJ</CarrierCode>
                 <NoFlight>094</NoFlight>
                 <DepartureStation>CGK</DepartureStation>
                 <ArrivalStation>TKG</ArrivalStation>
                 <Std>28-SEP-18 02.55.00 PM</Std>
                 <Sta>28-SEP-18 03.35.00 PM</Sta>
                 <Legs>
                   <item>
                     <DepartureStation>CGK</DepartureStation>
                     <ArrivalStation>TKG</ArrivalStation>
                     <Std>28-SEP-18 02.55.00 PM</Std>
                     <Sta>28-SEP-18 03.35.00 PM</Sta>
                   </item>
                 </Legs>
               </item>
             </Segments>
             <ClassesAvailable>
               <item>
                 <item>
                   <Key>-</Key>
                   <Availability>0</Availability>
                   <Class>X</Class>
                   <SeatAvail>CLOSE</SeatAvail>
                   <Price>0</Price>
                   <PriceDetail></PriceDetail>
                   <Currency>IDR</Currency>
                   <StatusAvail>Normal</StatusAvail>
                 </item>
                 <item>
                    ...
                 </item>
               </item>
             </ClassesAvailable>
           </item>
           <item>
             <CityFrom>CGK</CityFrom>
             <CityTo>TKG</CityTo>
             <Std>28-SEP-18 04.15.00 PM</Std>
             <Sta>28-SEP-18 04.55.00 PM</Sta>
             <FlightTime>40</FlightTime>
             <Segments>
               <item>
                 <CarrierCode>SJ</CarrierCode>
                 <NoFlight>058</NoFlight>
                 <DepartureStation>CGK</DepartureStation>
                 <ArrivalStation>TKG</ArrivalStation>
                 <Std>28-SEP-18 04.15.00 PM</Std>
                 <Sta>28-SEP-18 04.55.00 PM</Sta>
                 <Legs>
                   <item>
                     <DepartureStation>CGK</DepartureStation>
                     <ArrivalStation>TKG</ArrivalStation>
                     <Std>28-SEP-18 04.15.00 PM</Std>
                     <Sta>28-SEP-18 04.55.00 PM</Sta>
                   </item>
                 </Legs>
               </item>
             </Segments>
             <ClassesAvailable>
               <item>
                 <item>
                   <Key>-</Key>
                   <Availability>0</Availability>
                   <Class>X</Class>
                   <SeatAvail>CLOSE</SeatAvail>
                   <Price>0</Price>
                   <PriceDetail></PriceDetail>
                   <Currency>IDR</Currency>
                   <StatusAvail>Normal</StatusAvail>
                 </item>
                 <item>
                    ...
                 </item>
               </item>
             </ClassesAvailable>
           </item>
           <item>
             <CityFrom>CGK</CityFrom>
             <CityTo>TKG</CityTo>
             <Std>28-SEP-18 05.10.00 PM</Std>
             <Sta>28-SEP-18 05.40.00 PM</Sta>
             <FlightTime>30</FlightTime>
             <Segments>
               <item>
                 <CarrierCode>IN</CarrierCode>
                 <NoFlight>098</NoFlight>
                 <DepartureStation>CGK</DepartureStation>
                 <ArrivalStation>TKG</ArrivalStation>
                 <Std>28-SEP-18 05.10.00 PM</Std>
                 <Sta>28-SEP-18 05.40.00 PM</Sta>
                 <Legs>
                   <item>
                     <DepartureStation>CGK</DepartureStation>
                     <ArrivalStation>TKG</ArrivalStation>
                     <Std>28-SEP-18 05.10.00 PM</Std>
                     <Sta>28-SEP-18 05.40.00 PM</Sta>
                   </item>
                 </Legs>
               </item>
             </Segments>
             <ClassesAvailable>
               <item>
                 <item>
                   <Key>-</Key>
                   <Availability>0</Availability>
                   <Class>X</Class>
                   <SeatAvail>CLOSE</SeatAvail>
                   <Price>0</Price>
                   <PriceDetail></PriceDetail>
                   <Currency>IDR</Currency>
                   <StatusAvail>Normal</StatusAvail>
                 </item>
                 <item>
                   ...
                 </item>
               </item>
             </ClassesAvailable>
           </item>
           <item>
             <CityFrom>CGK</CityFrom>
             <CityTo>TKG</CityTo>
             <Std>28-SEP-18 06.45.00 PM</Std>
             <Sta>28-SEP-18 07.25.00 PM</Sta>
             <FlightTime>40</FlightTime>
             <Segments>
               <item>
                 <CarrierCode>SJ</CarrierCode>
                 <NoFlight>086</NoFlight>
                 <DepartureStation>CGK</DepartureStation>
                 <ArrivalStation>TKG</ArrivalStation>
                 <Std>28-SEP-18 06.45.00 PM</Std>
                 <Sta>28-SEP-18 07.25.00 PM</Sta>
                 <Legs>
                   <item>
                     <DepartureStation>CGK</DepartureStation>
                     <ArrivalStation>TKG</ArrivalStation>
                     <Std>28-SEP-18 06.45.00 PM</Std>
                     <Sta>28-SEP-18 07.25.00 PM</Sta>
                   </item>
                 </Legs>
               </item>
             </Segments>
             <ClassesAvailable>
               <item>
                 <item>
                   <Key>-</Key>
                   <Availability>0</Availability>
                   <Class>X</Class>
                   <SeatAvail>CLOSE</SeatAvail>
                   <Price>0</Price>
                   <PriceDetail></PriceDetail>
                   <Currency>IDR</Currency>
                   <StatusAvail>Normal</StatusAvail>
                 </item>
                 <item>
                    ...
                 </item>                 
               </item>
             </ClassesAvailable>
           </item>
         </FlightRoute>
       </item>
     </TripDetail>
     <SearchKey>SEARCH_KEY_WS_SEARCH_FLIGHT_RESPONSE</SearchKey>
     <ErrorCode>SEARCH0000</ErrorCode>
     <ErrorMessage>Success.</ErrorMessage>
   </return>
 </WsSearchFlightResponse>
```