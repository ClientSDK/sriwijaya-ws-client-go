echo "Clean..."
rm ./WsAccountStatement
echo "Build..."
go build -o WsAccountStatement main.go 
echo "Build Done..."
echo "Run..."
./WsAccountStatement > WsAccountStatement-Result.xml
echo "Done."