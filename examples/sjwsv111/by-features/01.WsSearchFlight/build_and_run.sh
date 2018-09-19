echo "Clean..."
rm ./WsSearchFlight
echo "Build..."
go build -o WsSearchFlight main.go 
echo "Build Done..."
echo "Run..."
./WsSearchFlight > WsSearchFlight-Result.xml
echo "Done."