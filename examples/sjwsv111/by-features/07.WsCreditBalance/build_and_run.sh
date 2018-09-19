echo "Clean..."
rm ./WsCreditBalance
echo "Build..."
go build -o WsCreditBalance main.go 
echo "Build Done..."
echo "Run..."
./WsCreditBalance > WsCreditBalance-Result.xml
echo "Done."