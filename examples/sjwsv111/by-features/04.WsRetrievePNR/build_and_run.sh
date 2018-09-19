echo "Clean..."
rm ./WsRetrievePNR
echo "Build..."
go build -o WsRetrievePNR main.go 
echo "Build Done..."
echo "Run..."
./WsRetrievePNR > WsRetrievePNR-Result.xml
echo "Done."