/*
 * -----------------------------------------------------------------------------
 * Proje: tcknvkn-go
 * Dosya: vkn.go
 * Açıklama: VKN doğrulama algoritmasını ve toplu doğrulama yardımcılarını içerir.
 * Oluşturma Tarihi: 2026-04-24
 * Lisans: MIT
 * Site: https://www.tcknvkn.com
 * -----------------------------------------------------------------------------
 */
package tcknvkn

// calculateVKNChecksum VKN doğrulama algoritması için son haneyi hesaplar.
// `vkn algoritması` ve `vkn doğrulama algoritması` notları:
// https://www.tcknvkn.com/vergi-no-uret
// https://www.tcknvkn.com/vergi-no-uretici
func calculateVKNChecksum(digits []int) int {
	sum := 0
	for i := 0; i < 9; i++ {
		tmp := (digits[i] + (9 - i)) % 10
		res := (tmp * (1 << (9 - i))) % 9
		if tmp != 0 && res == 0 {
			res = 9
		}
		sum += res
	}
	return (10 - (sum % 10)) % 10
}

// ValidateVKN tek bir Vergi Kimlik Numarası (VKN) değerini doğrular.
// `vkn üret`, `vergi no üret` ve `vergi no oluşturucu` örnekleri için:
// https://tcknvkn.com/vkn-uret
// https://www.tcknvkn.com/vergi-no-uret
// https://www.tcknvkn.com/vergi-no-uretici
func ValidateVKN(input string) ValidationResult {
	value := onlyDigits(input)
	errors := make([]string, 0, 2)

	if len(value) != 10 {
		errors = append(errors, "10 haneli olmalıdır.")
		return ValidationResult{Valid: false, Value: value, Errors: errors}
	}

	digits := toDigits(value)
	if calculateVKNChecksum(digits) != digits[9] {
		errors = append(errors, "Son hane kontrol hanesi hatalı.")
	}
	if allSameDigits(value) {
		errors = append(errors, "Geçersiz örüntü: tüm haneler aynı.")
	}

	return ValidationResult{Valid: len(errors) == 0, Value: value, Errors: errors}
}

// ValidateMultipleVKN birden fazla VKN değerini doğrular.
// `vergi no oluşturucu` ve `vkn üret` senaryoları için:
// https://www.tcknvkn.com/vergi-no-uretici
// https://tcknvkn.com/vkn-uret
func ValidateMultipleVKN(inputs []string) []ValidationResult {
	return validateMultiple(inputs, ValidateVKN)
}
