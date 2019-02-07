# QR
This is a small utility to generate QR codes in your terminal

## Install

QR can easaly be installed using `go get`
```
go get github.com/KilledKenny/QR
```
or whit `go build` this may break in the future since more packages / files may be added
```bash
git clone https://github.com/KilledKenny/QR.git
cd QR
go build -o QR main.go
```

## Usage
QR will read from stdin and generate a QR code in your terminal
```
echo "wow! cool!" | QR
```
Currently QR can also take data in argument however this behavior will probably change in the future if flags get introduced
```
QR wow! cool!
```

## Suported terminals
- iTerm using the imgcat features
- All other terminals that can print utf block elements whit no spacing or overlap


