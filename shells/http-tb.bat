@echo off
echo "HTTP Test"
echo "Build Start."

echo "Building Libraries..."
go build lib/net/http/http.go
go install lib/net/http/http.go
echo "Finished Library Building."

echo "Start Build Test Program..."
go build test/testhttp.go
echo "Finished."

