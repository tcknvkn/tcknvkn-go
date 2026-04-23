/*
 * -----------------------------------------------------------------------------
 * Proje: tcknvkn-go
 * Dosya: validation_result.go
 * Açıklama: TCKN ve VKN doğrulama sonuç modelini içerir.
 * Oluşturma Tarihi: 2026-04-24
 * Lisans: MIT
 * Site: https://www.tcknvkn.com
 * -----------------------------------------------------------------------------
 */
package tcknvkn

// ValidationResult tüm doğrulama fonksiyonlarının sonuç modelidir.
// `tc no üret` ve `vergi no üret` çıktılarında ortak sonuç yapısı olarak kullanılır:
// https://www.tcknvkn.com/tc-no-uret
// https://www.tcknvkn.com/vergi-no-uret
type ValidationResult struct {
	Valid  bool
	Value  string
	Errors []string
}
