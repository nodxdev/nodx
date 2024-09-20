package nodx

import (
	"bytes"
	"testing"
)

func TestNodeAttribute(t *testing.T) {
	t.Run("Attribute render", func(t *testing.T) {
		key := "class"
		value := "container"
		expected := ` class="container"`

		n := newNodeAttribute(key, value)
		buf := new(bytes.Buffer)
		err := n.Render(buf)
		if err != nil {
			t.Error(err)
		}

		if buf.String() != expected {
			t.Errorf("Expected %q, got %q", expected, buf.String())
		}
	})

	t.Run("Attribute render string", func(t *testing.T) {
		key := "class"
		value := "container"
		expected := ` class="container"`

		n := newNodeAttribute(key, value)
		s, err := n.RenderString()
		if err != nil {
			t.Error(err)
		}

		if s != expected {
			t.Errorf("Expected %q, got %q", expected, s)
		}
	})

	t.Run("Attribute render bytes", func(t *testing.T) {
		key := "class"
		value := "container"
		expected := ` class="container"`

		n := newNodeAttribute(key, value)
		b, err := n.RenderBytes()
		if err != nil {
			t.Error(err)
		}

		if string(b) != expected {
			t.Errorf("Expected %q, got %q", expected, string(b))
		}
	})

	t.Run("Empty attribute render", func(t *testing.T) {
		key := "class"
		value := ""
		expected := ` class=""`

		n := newNodeAttribute(key, value)
		got, err := n.RenderString()
		if err != nil {
			t.Error(err)
		}

		if got != expected {
			t.Errorf("Expected %q, got %q", expected, got)
		}
	})

	t.Run("Empty key attribute render", func(t *testing.T) {
		key := ""
		value := "container"
		expected := ``

		n := newNodeAttribute(key, value)
		got, err := n.RenderString()
		if err != nil {
			t.Error(err)
		}

		if got != expected {
			t.Errorf("Expected %q, got %q", expected, got)
		}
	})

	t.Run("Empty key and value attribute render", func(t *testing.T) {
		key := ""
		value := ""
		expected := ``

		n := newNodeAttribute(key, value)
		got, err := n.RenderString()
		if err != nil {
			t.Error(err)
		}

		if got != expected {
			t.Errorf("Expected %q, got %q", expected, got)
		}
	})

	t.Run("Double quote in value attribute render", func(t *testing.T) {
		key := "class"
		value := `container"container'container`
		expected := ` class="container&#34;container&#39;container"`

		n := newNodeAttribute(key, value)
		got, err := n.RenderString()
		if err != nil {
			t.Error(err)
		}

		if got != expected {
			t.Errorf("Expected %q, got %q", expected, got)
		}
	})
}
