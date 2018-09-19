echo "Clean..."
rm ./WsRouteOperate
echo "Build..."
go build -o WsRouteOperate main.go 
echo "Build Done..."
echo "Run..."
./WsRouteOperate > WsRouteOperate-Result.xml
echo "Done."