// Copyright 2018 The ClientSDK Team Authors. All rights reserved.
// Use of this source code is governed by a Apache 2.0-style
// license that can be found in the LICENSE file.

// Author: ClientSDK Team (muharihar)

// Package sjwsdk111 is an UnOfficial SOAP client SDK for Sriwijaya Air Web Services (SOAP)
// WSDL Version:  https://wsp.sriwijayaair.co.id:11443/wsdl.eticketv111/index.php?wsdl
package sjwsdk111

import "encoding/xml"

// WsSearchFlightResponse represent WsSearchFlight soap response
type WsSearchFlightResponse struct {
	XMLName xml.Name `xml:"WsSearchFlightResponse"`
	Return  struct {
		Username     string               `xml:"Username"`
		Adult        string               `xml:"Adult"`
		Child        string               `xml:"Child"`
		Infant       string               `xml:"Infant"`
		TripDetail   *respTripDetailArray `xml:"TripDetail"`
		SearchKey    string               `xml:"SearchKey"`
		ErrorCode    string               `xml:"ErrorCode"`
		ErrorMessage string               `xml:"ErrorMessage"`
	} `xml:"return"`
}

type respTripDetailArray struct {
	Item []*struct {
		CityFrom    string                `xml:"CityFrom"`
		CityTo      string                `xml:"CityTo"`
		Category    string                `xml:"Category"`
		FlightRoute *respFlightRouteArray `xml:"FlightRoute"`
	} `xml:"item"`
}

type respFlightRouteArray struct {
	Item []*struct {
		CityFrom         string                   `xml:"CityFrom"`
		CityTo           string                   `xml:"CityTo"`
		Std              string                   `xml:"Std"`
		Sta              string                   `xml:"Sta"`
		FlightTime       string                   `xml:"FlightTime"`
		Segments         *respSegmentsArray       `xml:"Segments"`
		ClassesAvailable *respClassAvailableArray `xml:"ClassesAvailable"`
	} `xml:"item"`
}

type respSegmentsArray struct {
	Item []*struct {
		CarrierCode      string         `xml:"CarrierCode"`
		NoFlight         string         `xml:"NoFlight"`
		DepartureStation string         `xml:"DepartureStation"`
		ArrivalStation   string         `xml:"ArrivalStation"`
		Std              string         `xml:"Std"`
		Sta              string         `xml:"Sta"`
		Legs             *respLegsArray `xml:"Legs"`
	} `xml:"item"`
}

type respLegsArray struct {
	Item []*struct {
		DepartureStation string `xml:"DepartureStation"`
		ArrivalStation   string `xml:"ArrivalStation"`
		Std              string `xml:"Std"`
		Sta              string `xml:"Sta"`
	} `xml:"item"`
}

type respClassAvailableArray struct {
	Item []*struct {
		Item []*struct {
			Key          string             `xml:"Key"`
			Availability string             `xml:"Availability"`
			Class        string             `xml:"Class"`
			SeatAvail    string             `xml:"SeatAvail"`
			Price        string             `xml:"Price"`
			PriceDetail  *priceDetailArrays `xml:"PriceDetail"`
			Currency     string             `xml:"Currency"`
			StatusAvail  string             `xml:"StatusAvail"`
		} `xml:"item"`
	} `xml:"item"`
}

type priceDetailArrays struct {
	Item []*struct {
		PaxCategory   string               `xml:"PaxCategory"`
		Total1        string               `xml:"Total_1"`
		Nta1          string               `xml:"Nta_1"`
		FareComponent *fareComponentArrays `xml:"FareComponent"`
	} `xml:"item"`
}

type fareComponentArrays struct {
	Item []*struct {
		FareChargeTypeCode string `xml:"FareChargeTypeCode"`
		FareChargeTypeDesc string `xml:"FareChargeTypeDesc"`
		Amount             string `xml:"Amount"`
		CurrencyCode       string `xml:"CurrencyCode"`
	} `xml:"item"`
}

// WsGeneratePNRResponse represent WsGeneratePNR soap response
type WsGeneratePNRResponse struct {
	XMLName xml.Name `xml:"WsGeneratePNRResponse"`
	Return  struct {
		Username             string                     `xml:"Username"`
		BookingCode          string                     `xml:"BookingCode"`
		YourItineraryDetails *yourItineraryDetailsArray `xml:"YourItineraryDetails"`
		ErrorCode            string                     `xml:"ErrorCode"`
		ErrorMessage         string                     `xml:"ErrorMessage"`
	} `xml:"return"`
}

// WsIssuingResponse represent WsIssuing soap response
type WsIssuingResponse struct {
	XMLName xml.Name `xml:"WsIssuingResponse"`
	Return  struct {
		Username             string                     `xml:"Username"`
		BookingCode          string                     `xml:"BookingCode"`
		YourItineraryDetails *yourItineraryDetailsArray `xml:"YourItineraryDetails"`
		ErrorCode            string                     `xml:"ErrorCode"`
		ErrorMessage         string                     `xml:"ErrorMessage"`
	} `xml:"return"`
}

// WsRetrievePNRResponse represent WsRetrievePNR soap response
type WsRetrievePNRResponse struct {
	XMLName xml.Name `xml:"WsRetrievePNRResponse"`
	Return  struct {
		Username             string                     `xml:"Username"`
		BookingCode          string                     `xml:"BookingCode"`
		YourItineraryDetails *yourItineraryDetailsArray `xml:"YourItineraryDetails"`
		ErrorCode            string                     `xml:"ErrorCode"`
		ErrorMessage         string                     `xml:"ErrorMessage"`
	} `xml:"return"`
}

type yourItineraryDetailsArray struct {
	ReservationDetails    *reservationDetailsArray    `xml:"ReservationDetails"`
	PassengerDetails      *passengerDetailsArray      `xml:"PassengerDetails"`
	ItineraryDetails      *itineraryDetailsArray      `xml:"ItineraryDetails"`
	PaymentDetails        *paymentDetailsArray        `xml:"PaymentDetails"`
	ContactList           *contactListArray           `xml:"ContactList"`
	AgentDetails          *agentDetailsArray          `xml:"AgentDetails"`
	BookingRemarks        *bookingRemarksArray        `xml:"BookingRemarks"`
	AdditionalInformation *additionalInformationArray `xml:"AdditionalInformation"`
}

type reservationDetailsArray struct {
	BookingCode       string `xml:"BookingCode"`
	BookingDate       string `xml:"BookingDate"`
	BalanceDue        string `xml:"BalanceDue"`
	BalanceDueRemarks string `xml:"BalanceDueRemarks"`
	CurrencyCode      string `xml:"CurrencyCode"`
	Time              string `xml:"Time"`
	TimeDescription   string `xml:"TimeDescription"`
	Status            string `xml:"Status"`
}

type passengerDetailsArray struct {
	Item []*struct {
		No             string `xml:"No"`
		Suffix         string `xml:"Suffix"`
		FirstName      string `xml:"FirstName"`
		LastName       string `xml:"LastName"`
		SeatQty        string `xml:"SeatQty"`
		TicketNumber   string `xml:"TicketNumber"`
		SpecialRequest string `xml:"SpecialRequest"`
	} `xml:"item"`
}

type itineraryDetailsArray struct {
	Journey journeyArray `xml:"Journey"`
}

type journeyArray struct {
	Item []*struct {
		Segment *segmentArray `xml:"Segment"`
	} `xml:"item"`
}

type segmentArray struct {
	Item []*struct {
		FlownDate         string `xml:"FlownDate"`
		FlightNo          string `xml:"FlightNo"`
		CityFrom          string `xml:"CityFrom"`
		CityTo            string `xml:"CityTo"`
		CityFromName      string `xml:"CityFromName"`
		CityToName        string `xml:"CityToName"`
		StdLT             string `xml:"StdLT"`
		StaLT             string `xml:"StaLT"`
		ReservationStatus string `xml:"ReservationStatus"`
		Class             string `xml:"Class"`
		CheckInStatus     string `xml:"CheckInStatus"`
	} `xml:"item"`
}

type paymentDetailsArray struct {
	BasicFare    string `xml:"BasicFare"`
	Others       string `xml:"Others"`
	Sti          string `xml:"Sti"`
	Total        string `xml:"Total"`
	Nta          string `xml:"Nta"`
	CurrencyCode string `xml:"CurrencyCode"`
}

type contactListArray struct {
	Item []*struct {
		Type        string `xml:"Type"`
		Description string `xml:"Description"`
		Value       string `xml:"Value"`
	} `xml:"item"`
}

type agentDetailsArray struct {
	BookedBy string `xml:"BookedBy"`
	IssuedBy string `xml:"IssuedBy"`
}

type bookingRemarksArray struct {
	Item []*struct {
		CommentText string `xml:"CommentText"`
		CreatedBy   string `xml:"CreatedBy"`
		CreatedDate string `xml:"CreatedDate"`
		IPAddress   string `xml:"IpAddress"`
	} `xml:"item"`
}

type additionalInformationArray struct {
	Item []*struct {
		Reasons string `xml:"Reasons"`
		Value   string `xml:"Value"`
	} `xml:"item"`
}

// WsCancelPNRResponse represent WsCancelPNR soap response
type WsCancelPNRResponse struct {
	XMLName xml.Name `xml:"WsCancelPNRResponse"`
	Return  struct {
		Username     string `xml:"Username"`
		BookingCode  string `xml:"BookingCode"`
		ErrorCode    string `xml:"ErrorCode"`
		ErrorMessage string `xml:"ErrorMessage"`
	} `xml:"return"`
}

// WsAccountStatementResponse represent WsAccountStatement soap response
type WsAccountStatementResponse struct {
	XMLName xml.Name `xml:"WsAccountStatementResponse"`
	Return  struct {
		Username               string                            `xml:"Username"`
		AccountStatementDetail *respAccountStatementDetailArrays `xml:"AccountStatementDetail"`
		ErrorCode              string                            `xml:"ErrorCode"`
		ErrorMessage           string                            `xml:"ErrorMessage"`
	} `xml:"return"`
}

type respAccountStatementDetailArrays struct {
	Item []*struct {
		DateCreate    string `xml:"DateCreate"`
		UserCreate    string `xml:"UserCreate"`
		Description   string `xml:"Description"`
		Amount        string `xml:"Amount"`
		BalanceStatus string `xml:"BalanceStatus"`
		LastBalance   string `xml:"LastBalance"`
		Currency      string `xml:"Currency"`
	} `xml:"item"`
}

// WsCreditBalanceResponse represent WsCreditBalance soap response
type WsCreditBalanceResponse struct {
	XMLName xml.Name `xml:"WsCreditBalanceResponse"`
	Return  struct {
		Username      string `xml:"Username"`
		CreditBalance string `xml:"CreditBalance"`
		ErrorCode     string `xml:"ErrorCode"`
		ErrorMessage  string `xml:"ErrorMessage"`
	} `xml:"return"`
}

// WsRouteOperateResponse represent WsRouteOperate soap response
type WsRouteOperateResponse struct {
	XMLName xml.Name `xml:"WsRouteOperateResponse"`
	Return  struct {
		Username      string              `xml:"Username"`
		RouteOperates *routeOperatesArray `xml:"RouteOperates"`
		ErrorCode     string              `xml:"ErrorCode"`
		ErrorMessage  string              `xml:"ErrorMessage"`
	} `xml:"return"`
}

type routeOperatesArray struct {
	Item []*struct {
		CityFrom        string `xml:"CityFrom"`
		CityFromName    string `xml:"CityFromName"`
		CityFromCountry string `xml:"CityFromCountry"`
		ApoNameFrom     string `xml:"ApoNameFrom"`
		TimeZoneFrom    string `xml:"TimeZoneFrom"`
		CityTo          string `xml:"CityTo"`
		CityToName      string `xml:"CityToName"`
		CityToCountry   string `xml:"CityToCountry"`
		ApoNameTo       string `xml:"ApoNameTo"`
		TimeZoneTo      string `xml:"TimeZoneTo"`
		StatusRoute     string `xml:"StatusRoute"`
	} `xml:"item"`
}
