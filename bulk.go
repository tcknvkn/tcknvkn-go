/*
 * -----------------------------------------------------------------------------
 * Proje: tcknvkn-go
 * Dosya: bulk.go
 * Açıklama: Toplu doğrulama akışını yürüten ortak yardımcı fonksiyonu içerir.
 * Oluşturma Tarihi: 2026-04-24
 * Lisans: MIT
 * Site: https://www.tcknvkn.com
 * -----------------------------------------------------------------------------
 */
package tcknvkn

// validateMultiple verilen doğrulayıcıyla tüm girdileri sırayı koruyarak doğrular.
// `tc no uret` ve `vergi no oluşturucu` toplu doğrulama kullanımında ortak altyapıdır:
// https://www.tcknvkn.com/tc-no-uret
// https://www.tcknvkn.com/vergi-no-uretici
func validateMultiple(inputs []string, validator func(string) ValidationResult) []ValidationResult {
	results := make([]ValidationResult, 0, len(inputs))
	for _, input := range inputs {
		results = append(results, validator(input))
	}
	return results
}
