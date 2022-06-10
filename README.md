# Mock-Server
This mock server can be used for testing HTTP requests. The server can return the headers of an HTTP request, echo a JSON body back, or have text sent over an HTTP request from a file.

100% written in Go

**Go is not required to run this mock server. Download the binary from the link below.**

[Download](https://github.com/akhil-datla/Mock-Server/releases/)

**NOTE: Download the test.txt or create one and save it into the same directory as the mock server program for the default /file API**

To run the mock server from Terminal, type the following command:
```
./server
```
If you want to change the port and/or the path to read a file for the "/file" API, add the port flag and/or the path flag:
```
./server -port=<PORT> -path=<PATH TO NEW FILE>
```

HTTP Routes:

To get the headers of an HTTP request, call this API:
```
http://<DOMAIN>:<PORT>/header
```
To echo the JSON body of an HTTP request, call this API:\
**Note: Only works with JSON sent in the body of an HTTP request**
```
http://<DOMAIN>:<PORT>/body
```
To have text from a file transmitted over an HTTP Request, call this API:\
**Note: This program must be able to read this file in order for this API to properly work and the default file to read from is the test.txt which should be in the same directory.**
```
http://<DOMAIN>:<PORT>/body
```
