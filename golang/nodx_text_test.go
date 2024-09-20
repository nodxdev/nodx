package nodx

import (
	"bytes"
	"testing"
)

func TestNodeText(t *testing.T) {
	t.Run("Text render", func(t *testing.T) {
		text := "Hello, World!"

		n := newNodeText(text)
		buf := new(bytes.Buffer)
		err := n.Render(buf)
		if err != nil {
			t.Error(err)
		}

		if buf.String() != text {
			t.Errorf("Expected %q, got %q", text, buf.String())
		}
	})

	t.Run("Text render string", func(t *testing.T) {
		text := "Hello, World!"

		n := newNodeText(text)
		s, err := n.RenderString()
		if err != nil {
			t.Error(err)
		}

		if s != text {
			t.Errorf("Expected %q, got %q", text, s)
		}
	})

	t.Run("Text render bytes", func(t *testing.T) {
		text := "Hello, World!"

		n := newNodeText(text)
		b, err := n.RenderBytes()
		if err != nil {
			t.Error(err)
		}

		if string(b) != text {
			t.Errorf("Expected %q, got %q", text, string(b))
		}
	})

	t.Run("Empty text render", func(t *testing.T) {
		n := newNodeText("")
		buf := new(bytes.Buffer)
		err := n.Render(buf)
		if err != nil {
			t.Error(err)
		}

		if buf.String() != "" {
			t.Errorf("Expected empty string, got %q", buf.String())
		}
	})
}
