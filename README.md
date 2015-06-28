[![Build Status](https://travis-ci.org/samwierema/go-qr-generator.svg?branch=master)](https://travis-ci.org/samwierema/go-qr-generator)

# A QR code generator written in Golang
Starts an HTTP server (listening on port 8080) that generates QR codes. Once installed and running (see below), the service accepts the following two parameters:
* ```data```: (Required) The (URL encoded) string that should be encoded in the QR code
* ```size```: (Optional) The size of the image (default: 250)

E.g. ```http://your-domain.tld:8080/?data=Hello%2C%20world&size=300```

## Installation
Download the source code and install it using the `go install` command.

Alternatively, use Docker to run the service in a container:
```
docker run -d -p 8080:8080 samwierema/go-qr-generator
```

## References
* Barcode Library: https://github.com/boombuler/barcode

## Author
* [Sam Wierema](http://wiere.ma)
