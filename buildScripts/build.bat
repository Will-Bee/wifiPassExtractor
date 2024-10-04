@echo off
echo Building...
go build -o passExtract.exe ../main.go ../passExtract.go ../GUI.go
echo build done
pause