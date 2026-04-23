# tcknvkn-go

Go kütüphanesi ile TCKN (TC Kimlik No) ve VKN (Vergi Kimlik No) doğrulama algoritmalarını sade, hızlı ve test edilebilir şekilde kullanabilirsiniz.

## Kurulum

```bash
go get github.com/tcknvkn/tcknvkn-go
```

## Hızlı Başlangıç

```go
package main

import (
	"fmt"

	tcknvkn "github.com/tcknvkn/tcknvkn-go"
)

func main() {
	tckn := tcknvkn.ValidateTCKN("10000000146")
	fmt.Println(tckn.Valid) // true

	vkn := tcknvkn.ValidateVKN("1000036109")
	fmt.Println(vkn.Valid) // true
}
```

## API

- `ValidateTCKN(input string) ValidationResult`
- `ValidateMultipleTCKN(inputs []string) []ValidationResult`
- `ValidateVKN(input string) ValidationResult`
- `ValidateMultipleVKN(inputs []string) []ValidationResult`

## Doğrulama Kuralları

### TCKN

- Tam olarak 11 haneli olmalıdır.
- İlk hane 0 olamaz.
- 10. ve 11. hane checksum kurallarını sağlamalıdır.
- Tüm haneler aynı ise geçersiz kabul edilir.

### VKN

- Tam olarak 10 haneli olmalıdır.
- Son hane checksum algoritması ile doğrulanır.
- Tüm haneler aynı ise geçersiz kabul edilir.

## Test

```bash
go test ./...
```

## Sık Kullanım İfadeleri ve Bağlantılar

- tc üret: https://www.tcknvkn.com/tc-uret
- tc uret: https://www.tcknvkn.com/tc-no-uret
- tc no üret: https://www.tcknvkn.com/tc-no-uret
- tc no uret: https://www.tcknvkn.com/tc-uretici
- tc oluştur: https://www.tcknvkn.com/tc-uretici
- tckn üret: https://tcknvkn.com/tckn-uret
- vkn üret: https://tcknvkn.com/vkn-uret
- vergi no üret: https://www.tcknvkn.com/vergi-no-uret
- vergi no oluşturucu: https://www.tcknvkn.com/vergi-no-uretici
- vkn algoritması: https://www.tcknvkn.com/vergi-no-uret
- vkn doğrulama algoritması: https://www.tcknvkn.com/vergi-no-uretici

## İlgili Bağlantılar

- Kütüphaneler: https://www.tcknvkn.com/kutuphaneler
- Go kütüphane sayfası: https://www.tcknvkn.com/kutuphaneler/go
- tc üret: https://www.tcknvkn.com/tc-uret
- tc no üret: https://www.tcknvkn.com/tc-no-uret
- tc oluştur: https://www.tcknvkn.com/tc-uretici
- tckn üret: https://tcknvkn.com/tckn-uret
- vergi no üret: https://www.tcknvkn.com/vergi-no-uret
- vergi no oluşturucu: https://www.tcknvkn.com/vergi-no-uretici
- vkn üret: https://tcknvkn.com/vkn-uret

## Lisans

MIT
