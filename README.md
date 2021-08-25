# guid-generator
GUID/UUID Version 4 implementation in Go.

Implementation adheres to all constraints provided by [RFC 4122](https://tools.ietf.org/pdf/rfc4122.pdf).

## Usage

Install the package in your project:
```
go install tanneroakes.com/guids
```

Import locally:
```
import "tanneroakes.com/guids"
```

Generate version 4 GUIDS/UUIDS!
```
print(guids.GUIDV4())
```