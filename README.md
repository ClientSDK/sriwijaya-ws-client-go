# Sriwijaya Air Web Service (SOAP) Client for Go [![GoDoc](https://godoc.org/github.com/ClientSDK/sriwijaya-ws-client-go?status.png)](https://godoc.org/github.com/ClientSDK/sriwijaya-ws-client-go) [![Go Report Card](https://goreportcard.com/badge/github.com/ClientSDK/sriwijaya-ws-client-go)](https://goreportcard.com/report/github.com/ClientSDK/sriwijaya-ws-client-go) 


<p align="center">
  <a href="https://github.com/ClientSDK/sriwijaya-ws-client-go">
    <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/e/ea/LOGO_SJ_VERTIKAL.png/320px-LOGO_SJ_VERTIKAL.png" alt="ClientSDK Swirijaya Air WS Client Go" width=300>
  </a>

  <h3 align="center">UnOfficial Go Package SOAP Client library for Sriwijaya Air Web Service (SOAP)</h3>

  <p align="center">
    Author:
    <br>
    <a href="https://github.com/ClientSDK"><strong>ClientSDK Team »</strong></a>
    <br>
    <br>
    <a href="https://github.com/ClientSDK/sriwijaya-ws-client-go/issues">Report issues</a>
  </p>
</p>


## Install

##### Sriwijaya Air Webservice Version 111

```bash
go get github.com/ClientSDK/sriwijaya-ws-client-go/sjwsdk111
```

## Prerequisites

- [Sriwijaya Air Web Service (SOAP) Client for Go (sjwsdk111 GoLang package) ](https://github.com/ClientSDK/sriwijaya-ws-client-go)

```bash
go get github.com/ClientSDK/sriwijaya-ws-client-go/sjwsdk111
```

- A Text Editor or an IDE

### Sriwijaya Air Agent requirements
- Sriwijaya Air Agent Credential Account ([Agent Application](https://agent.sriwijayaair.co.id/SJ-Eticket/login.php?action=in))
- Sriwijaya Air Web Service Access (IP Whitelist) ( [Production](https://wsp.sriwijayaair.co.id:11443/wsdl.eticketv111/index.php), [Development](https://wsx.sriwijayaair.co.id:11443/wsdl.eticketv111/index.php) )
- WSDL File ( [Production](https://wsp.sriwijayaair.co.id:11443/wsdl.eticketv111/index.php?wsdl), [Development](https://wsx.sriwijayaair.co.id:11443/wsdl.eticketv111/index.php?wsdl) )


## Example

##### Sriwijaya Air Webservice Version 111
- [Example By Features](examples/sjwsv111/by-features/README.md)
- [Examples By Scenarios]

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
    │       │   │   ├── ...
    │       │   │   └── 4.3.RoundTripConnecting-1.1.1
    │       │   └── International
    │       │   │   ├── 1.1.OneWayDirect-1.0.0
    │       │   │   ├── ...
    │       │   │   └── 4.3.RoundTripConnecting-1.1.1
    │       │   └── README.md
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


### Credit

- [SOAP Package for Go](https://github.com/tiaguinho/gosoap/)
- [Go library for accessing Github API](https://github.com/google/go-github)