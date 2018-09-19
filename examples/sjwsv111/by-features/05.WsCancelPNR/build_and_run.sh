echo "Clean..."
rm ./WsCancelPNR
echo "Build..."
go build -o WsCancelPNR main.go 
echo "Build Done..."
echo "Run..."
./WsCancelPNR > WsCancelPNR-Result.xml
echo "Done."