# Go-Vin
A simple library for parsing vehicle identification numbers

## Usage
```go
package main

import (
	"fmt"
	"github.com/ynori7/go-vin"
)

func main() {
	vin, _ := govin.ParseVin("1J4GZB8S1TY103658")
	fmt.Printf("%+v", vin)
}
```

Returns (note: presented here with pretty formatting):
```go
{
    Vin:      "1J4GZB8S1TY103658", 
    Region:   "NA", //NorthAmerica 
    Country:  "US", 
    ModelYear: 1996, 
    VinParts: {
        Wmi:           "1J4", 
        Vds:           "GZB8S1", 
        Vis:           "TY103658", 
        ModelYear:     "T", 
        SerialNumber:  "03658",
        Checksum:      "1",
        ChecksumValid: true,
    },
}
```

### TODO
- The model year is currently only reliable for North American VINs. It will still be parsed for others, but
it's not certain whether they follow any convention or not.