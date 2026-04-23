/*
 * -----------------------------------------------------------------------------
 * Proje: tcknvkn-go
 * Dosya: tckn.go
 * Açıklama: TCKN doğrulama algoritmasını ve toplu doğrulama yardımcılarını içerir.
 * Oluşturma Tarihi: 2026-04-24
 * Lisans: MIT
 * Site: https://www.tcknvkn.com
 * -----------------------------------------------------------------------------
 */
package tcknvkn

// ValidateTCKN tek bir TC Kimlik Numarası (TCKN) değerini doğrular.
// `tc üret`, `tc uret`, `tc no üret` ve `tc no uret` senaryoları için:
// https://www.tcknvkn.com/tc-uret
// https://www.tcknvkn.com/tc-no-uret
func ValidateTCKN(input string) ValidationResult {
	value := onlyDigits(input)
	errors := make([]string, 0, 3)

	if len(value) != 11 {
		errors = append(errors, "11 haneli olmalıdır.")
	}
	if len(value) > 0 && value[0] == '0' {
		errors = append(errors, "İlk hane 0 olamaz.")
	}
	if len(errors) > 0 {
		return ValidationResult{Valid: false, Value: value, Errors: errors}
	}

	digits := toDigits(value)
	if calculateTenthDigit(digits) != digits[9] {
		errors = append(errors, "10. hane kontrol hanesi hatalı.")
	}
	if calculateEleventhDigit(digits) != digits[10] {
		errors = append(errors, "11. hane kontrol hanesi hatalı.")
	}
	if allSameDigits(value) {
		errors = append(errors, "Geçersiz örüntü: tüm haneler aynı.")
	}

	return ValidationResult{Valid: len(errors) == 0, Value: value, Errors: errors}
}

// ValidateMultipleTCKN birden fazla TCKN değerini doğrular.
// `tc oluştur` ve `tckn üret` niyetleri için:
// https://www.tcknvkn.com/tc-uretici
// https://tcknvkn.com/tckn-uret
func ValidateMultipleTCKN(inputs []string) []ValidationResult {
	return validateMultiple(inputs, ValidateTCKN)
}

// calculateTenthDigit TCKN algoritmasındaki 10. haneyi hesaplar.
// `vkn algoritması` karşılaştırmaları için referans:
// https://www.tcknvkn.com/tc-uret
func calculateTenthDigit(digits []int) int {
	odd := digits[0] + digits[2] + digits[4] + digits[6] + digits[8]
	even := digits[1] + digits[3] + digits[5] + digits[7]
	return ((odd*7-even)%10 + 10) % 10
}

// calculateEleventhDigit TCKN algoritmasındaki 11. haneyi hesaplar.
// `vkn doğrulama algoritması` ile kıyaslamalı kontrollerde kullanılabilir:
// https://www.tcknvkn.com/tc-no-uret
func calculateEleventhDigit(digits []int) int {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += digits[i]
	}
	return sum % 10
}
