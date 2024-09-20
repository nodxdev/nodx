package nodx

import "testing"

func TestEscapeHTML(t *testing.T) {
	t.Run("Escape HTML", func(t *testing.T) {
		input := "<script>alert('XSS')</script>"
		expected := "&#60;script&#62;alert(&#39;XSS&#39;)&#60;/script&#62;"
		got := escapeHTML(input)
		if got != expected {
			t.Errorf("Expected %q, got %q", expected, got)
		}
	})

	t.Run("All characters", func(t *testing.T) {
		input := "test-&<>'\"-test"
		expected := "test-&#38;&#60;&#62;&#39;&#34;-test"
		got := escapeHTML(input)
		if got != expected {
			t.Errorf("Expected %q, got %q", expected, got)
		}
	})

	t.Run("No escape", func(t *testing.T) {
		input := "test"
		expected := "test"
		got := escapeHTML(input)
		if got != expected {
			t.Errorf("Expected %q, got %q", expected, got)
		}
	})

	t.Run("Empty string", func(t *testing.T) {
		input := ""
		expected := ""
		got := escapeHTML(input)
		if got != expected {
			t.Errorf("Expected %q, got %q", expected, got)
		}
	})

	t.Run("Escape with special characters", func(t *testing.T) {
		input := "test👌áéíóú-&<>'\"-👌test"
		expected := "test👌áéíóú-&#38;&#60;&#62;&#39;&#34;-👌test"
		got := escapeHTML(input)
		if got != expected {
			t.Errorf("Expected %q, got %q", expected, got)
		}
	})

	t.Run("No escape with special characters", func(t *testing.T) {
		input := "test👌áéíóú"
		expected := "test👌áéíóú"
		got := escapeHTML(input)
		if got != expected {
			t.Errorf("Expected %q, got %q", expected, got)
		}
	})

	t.Run("Single special character", func(t *testing.T) {
		input := "👌"
		expected := "👌"
		got := escapeHTML(input)
		if got != expected {
			t.Errorf("Expected %q, got %q", expected, got)
		}
	})
}
