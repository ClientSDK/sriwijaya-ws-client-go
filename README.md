# sriwijaya-ws-client-go
Sriwijaya Air Web Service (SOAP) Client (Go)

### Repository Structure
```
.
└── sriwijaya-ws-client-go
    ├── LICENSE
    ├── README.md
    ├── examples
    │   └── sjwsv111
    │       ├── by-features
    │       │   ├── 01.WsSearchFlight
    │       │   ├── 02.WsGeneratePNR
    │       │   ├── 03.WsIssuing
    │       │   ├── 04.WsRetrievePNR
    │       │   ├── 05.WsCancelPNR
    │       │   ├── 06.WsAccountStatement
    │       │   ├── 07.WsCreditBalance
    │       │   ├── 08.WsRouteOperate
    │       │   └── README.md
    │       ├── by-scenario
    │       │   ├── Domestic
    │       │   │   ├── 1.1.OneWayDirect-1.0.0
    │       │   │   ├── 1.2.OneWayDirect-1.1.0
    │       │   │   ├── 1.3.OneWayDirect-1.1.1
    │       │   │   ├── 2.1.OneWayConnecting-1.0.0
    │       │   │   ├── 2.1.OneWayConnecting-1.1.0
    │       │   │   ├── 2.1.OneWayConnecting-1.1.1
    │       │   │   ├── 3.1.RoundTripDirect-1.0.0
    │       │   │   ├── 3.2.RoundTripDirect-1.1.0
    │       │   │   ├── 3.3.RoundTripDirect-1.1.1
    │       │   │   ├── 4.1.RoundTripConnecting-1.0.0
    │       │   │   ├── 4.2.RoundTripConnecting-1.1.0
    │       │   │   └── 4.3.RoundTripConnecting-1.1.1
    │       │   └── International
    │       └── wsdl
    │           ├── wsp-wsdl.eticketv111.wsdl
    │           └── wsx-wsdl.eticketv111.wsdl
    └── sjwsdk111
        ├── encode.go
        ├── helpers.go
        ├── response_types.go
        ├── sjwsdk111.go
        └── wsdl.go
```
