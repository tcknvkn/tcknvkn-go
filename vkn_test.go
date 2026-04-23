/*
 * -----------------------------------------------------------------------------
 * Proje: tcknvkn-go
 * Dosya: vkn_test.go
 * Açıklama: VKN doğrulama fonksiyonları için varyasyonlu birim testlerini içerir.
 * Oluşturma Tarihi: 2026-04-24
 * Lisans: MIT
 * Site: https://www.tcknvkn.com
 * -----------------------------------------------------------------------------
 */
package tcknvkn

import "testing"

// TestValidateVKN VKN doğrulama senaryolarını test eder.
func TestValidateVKN(t *testing.T) {
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
			input:     "1000036109",
			wantValid: true,
			wantValue: "1000036109",
		},
		{
			name:      "valid normalized input",
			input:     "100-003-6109",
			wantValid: true,
			wantValue: "1000036109",
		},
		{
			name:      "reject invalid length",
			input:     "1234",
			wantValid: false,
			wantValue: "1234",
			wantErrs:  []string{"10 haneli olmalıdır."},
		},
		{
			name:      "reject checksum mismatch",
			input:     "1000036108",
			wantValid: false,
			wantValue: "1000036108",
			wantErrs:  []string{"Son hane kontrol hanesi hatalı."},
		},
		{
			name:      "reject identical pattern",
			input:     "1111111111",
			wantValid: false,
			wantValue: "1111111111",
			wantErrs:  []string{"Geçersiz örüntü: tüm haneler aynı."},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := ValidateVKN(tc.input)
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

// TestValidateMultipleVKN toplu VKN doğrulama akışını test eder.
func TestValidateMultipleVKN(t *testing.T) {
	t.Parallel()

	results := ValidateMultipleVKN([]string{
		"1000036109",
		"100-003-6108",
		"1111111111",
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
	if results[1].Value != "1000036108" {
		t.Fatalf("expected normalized value, got %q", results[1].Value)
	}
	if results[2].Valid {
		t.Fatalf("expected third result to be invalid")
	}
}

// TestValidateMultipleVKNEmpty toplu doğrulamada boş girdi davranışını test eder.
func TestValidateMultipleVKNEmpty(t *testing.T) {
	t.Parallel()

	results := ValidateMultipleVKN([]string{})
	if len(results) != 0 {
		t.Fatalf("expected 0 results, got %d", len(results))
	}
}
