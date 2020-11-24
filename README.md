# perlcrc
Ported Perl's [String::CRC32](https://metacpan.org/pod/String::CRC32) module to Go for strings.

## Usage example
```go
package main

import (
	"fmt"

	"github.com/limpasha/perlcrc"
)

func main() {
	var email = "user@gmail.com"
	var crc32 = perlcrc.CRC32String(email)

	fmt.Println(crc32)
}
```
prints `3326238302`. That's it :)

## LICENSE
The algorithm is taken from the source code of [String::CRC32](https://metacpan.org/pod/String::CRC32) perl module. 

The author of this package disclaims all copyrights and releases it into the public domain.
