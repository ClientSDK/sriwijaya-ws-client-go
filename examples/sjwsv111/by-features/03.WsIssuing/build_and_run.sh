echo "Clean..."
rm ./WsIssuing
echo "Build..."
go build -o WsIssuing main.go 
echo "Build Done..."
echo "Run..."
./WsIssuing > WsIssuing-Result.xml
echo "Done."