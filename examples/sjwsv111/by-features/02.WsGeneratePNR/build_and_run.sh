echo "Clean..."
rm ./WsGeneratePNR
echo "Build..."
go build -o WsGeneratePNR main.go 
echo "Build Done..."
echo "Run..."
./WsGeneratePNR > WsGeneratePNR-Result.xml
echo "Done."