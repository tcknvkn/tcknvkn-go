/*
 * -----------------------------------------------------------------------------
 * Proje: tcknvkn-go
 * Dosya: tckn_test.go
 * Açıklama: TCKN doğrulama fonksiyonları için varyasyonlu birim testlerini içerir.
 * Oluşturma Tarihi: 2026-04-24
 * Lisans: MIT
 * Site: https://www.tcknvkn.com
 * -----------------------------------------------------------------------------
 */
package tcknvkn

import "testing"

// TestValidateTCKN TCKN doğrulama senaryolarını test eder.
func TestValidateTCKN(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name      string
		input     string
		wantValid bool
		wantValue string
		wantErrs  []string
	}{
		{
			name:      "valid raw input",
			input:     "10000000146",
			wantValid: true,
			wantValue: "10000000146",
		},
		{
			name:      "valid normalized input",
			input:     "100-000 00146",
			wantValid: true,
			wantValue: "10000000146",
		},
		{
			name:      "reject invalid length",
			input:     "12345",
			wantValid: false,
			wantValue: "12345",
			wantErrs:  []string{"11 haneli olmalıdır."},
		},
		{
			name:      "reject leading zero",
			input:     "01234567890",
			wantValid: false,
			wantValue: "01234567890",
			wantErrs:  []string{"İlk hane 0 olamaz."},
		},
		{
			name:      "reject checksum mismatch on 11th digit",
			input:     "10000000145",
			wantValid: false,
			wantValue: "10000000145",
			wantErrs:  []string{"11. hane kontrol hanesi hatalı."},
		},
		{
			name:      "reject checksum mismatch on 10th digit",
			input:     "10000000156",
			wantValid: false,
			wantValue: "10000000156",
			wantErrs:  []string{"10. hane kontrol hanesi hatalı."},
		},
		{
			name:      "reject identical pattern",
			input:     "11111111111",
			wantValid: false,
			wantValue: "11111111111",
			wantErrs:  []string{"Geçersiz örüntü: tüm haneler aynı."},
		},
		{
			name:      "collects multiple errors for short leading zero",
			input:     "0",
			wantValid: false,
			wantValue: "0",
			wantErrs:  []string{"11 haneli olmalıdır.", "İlk hane 0 olamaz."},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := ValidateTCKN(tc.input)
			if got.Valid != tc.wantValid {
				t.Fatalf("expected valid=%v, got valid=%v with errors=%v", tc.wantValid, got.Valid, got.Errors)
			}
			if got.Value != tc.wantValue {
				t.Fatalf("expected value=%q, got value=%q", tc.wantValue, got.Value)
			}
			for _, wantErr := range tc.wantErrs {
				if !containsError(got.Errors, wantErr) {
					t.Fatalf("expected error %q in %v", wantErr, got.Errors)
				}
			}
		})
	}
}

// TestValidateMultipleTCKN toplu TCKN doğrulama akışını test eder.
func TestValidateMultipleTCKN(t *testing.T) {
	t.Parallel()

	results := ValidateMultipleTCKN([]string{
		"10000000146",
		"100-000 00145",
		"11111111111",
	})

	if len(results) != 3 {
		t.Fatalf("expected 3 results, got %d", len(results))
	}
	if !results[0].Valid {
		t.Fatalf("expected first result to be valid, got %v", results[0].Errors)
	}
	if results[1].Valid {
		t.Fatalf("expected second result to be invalid")
	}
	if results[1].Value != "10000000145" {
		t.Fatalf("expected normalized value, got %q", results[1].Value)
	}
	if results[2].Valid {
		t.Fatalf("expected third result to be invalid")
	}
}

// TestValidateMultipleTCKNEmpty toplu doğrulamada boş girdi davranışını test eder.
func TestValidateMultipleTCKNEmpty(t *testing.T) {
	t.Parallel()

	results := ValidateMultipleTCKN([]string{})
	if len(results) != 0 {
		t.Fatalf("expected 0 results, got %d", len(results))
	}
}
