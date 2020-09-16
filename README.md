# perlcrc-go
Ported Perl's [String::CRC32](https://metacpan.org/pod/String::CRC32) module to go for strings.

Example:
```go
import (
    "github.com/limpasha/perlcrc-go/v1"
)

func main(){
    var email = "user@gmail.com"
    var crc32 = perlcrc.CRC32String(email)

    fmt.Println(crc32)
}
```
prints `3326238302`. That's it :)