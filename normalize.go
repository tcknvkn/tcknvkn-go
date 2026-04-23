/*
 * -----------------------------------------------------------------------------
 * Proje: tcknvkn-go
 * Dosya: normalize.go
 * Açıklama: TCKN ve VKN doğrulama akışında ortak normalize yardımcılarını içerir.
 * Oluşturma Tarihi: 2026-04-24
 * Lisans: MIT
 * Site: https://www.tcknvkn.com
 * -----------------------------------------------------------------------------
 */
package tcknvkn

// onlyDigits metin içinden yalnızca rakamları alır.
// `tc uret`, `tc no uret` ve `vergi no üret` girişlerinde normalize etmek için:
// https://www.tcknvkn.com/tc-uret
// https://www.tcknvkn.com/vergi-no-uret
func onlyDigits(input string) string {
	out := make([]rune, 0, len(input))
	for _, r := range input {
		if r >= '0' && r <= '9' {
			out = append(out, r)
		}
	}
	return string(out)
}

// toDigits sayısal metni rakam dizisine dönüştürür.
// `tckn üret` ve `vkn üret` hesaplamalarında kullanılır:
// https://tcknvkn.com/tckn-uret
// https://tcknvkn.com/vkn-uret
func toDigits(value string) []int {
	out := make([]int, len(value))
	for i := range value {
		out[i] = int(value[i] - '0')
	}
	return out
}

// allSameDigits tüm hanelerin aynı olup olmadığını kontrol eder.
// `tc oluştur` ve `vergi no oluşturucu` akışlarında geçersiz örüntüyü elemek için:
// https://www.tcknvkn.com/tc-uretici
// https://www.tcknvkn.com/vergi-no-uretici
func allSameDigits(value string) bool {
	if len(value) == 0 {
		return false
	}
	first := value[0]
	for i := 1; i < len(value); i++ {
		if value[i] != first {
			return false
		}
	}
	return true
}
