/*
 * -----------------------------------------------------------------------------
 * Proje: tcknvkn-go
 * Dosya: normalize_test.go
 * Açıklama: Ortak normalize yardımcıları için birim testleri içerir.
 * Oluşturma Tarihi: 2026-04-24
 * Lisans: MIT
 * Site: https://www.tcknvkn.com
 * -----------------------------------------------------------------------------
 */
package tcknvkn

import "testing"

// TestOnlyDigits yalnızca rakam alma davranışını test eder.
func TestOnlyDigits(t *testing.T) {
	t.Parallel()

	got := onlyDigits(" 1a0-0 0.3_6*1#09 ")
	if got != "100036109" {
		t.Fatalf("expected normalized digits, got %q", got)
	}
}

// TestToDigits rakam dizisine dönüştürme davranışını test eder.
func TestToDigits(t *testing.T) {
	t.Parallel()

	got := toDigits("12345")
	want := []int{1, 2, 3, 4, 5}

	if len(got) != len(want) {
		t.Fatalf("expected %d digits, got %d", len(want), len(got))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("expected digit %d at index %d, got %d", want[i], i, got[i])
		}
	}
}

// TestAllSameDigits örüntü kontrolünü test eder.
func TestAllSameDigits(t *testing.T) {
	t.Parallel()

	if !allSameDigits("11111") {
		t.Fatal("expected identical pattern to be true")
	}
	if allSameDigits("12111") {
		t.Fatal("expected mixed pattern to be false")
	}
	if allSameDigits("") {
		t.Fatal("expected empty value to be false")
	}
}

// containsError hata listesinde hedef metni arar.
func containsError(errors []string, target string) bool {
	for _, err := range errors {
		if err == target {
			return true
		}
	}
	return false
}
