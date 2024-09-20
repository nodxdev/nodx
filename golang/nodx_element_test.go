package nodx

import (
	"bytes"
	"testing"
)

func TestNodeElement(t *testing.T) {
	t.Run("Basic element render", func(t *testing.T) {
		tag := "div"
		expected := "<div></div>"

		n := newNodeElement(false, tag)
		buf := new(bytes.Buffer)
		err := n.Render(buf)
		if err != nil {
			t.Error(err)
		}

		if buf.String() != expected {
			t.Errorf("Expected %q, got %q", expected, buf.String())
		}
	})

	t.Run("Basic element text render", func(t *testing.T) {
		tag := "div"
		expected := "<div></div>"

		n := newNodeElement(false, tag)
		got, err := n.RenderString()
		if err != nil {
			t.Error(err)
		}

		if got != expected {
			t.Errorf("Expected %q, got %q", expected, got)
		}
	})

	t.Run("Basic element bytes render", func(t *testing.T) {
		tag := "div"
		expected := "<div></div>"

		n := newNodeElement(false, tag)
		got, err := n.RenderBytes()
		if err != nil {
			t.Error(err)
		}

		if string(got) != expected {
			t.Errorf("Expected %q, got %q", expected, string(got))
		}
	})

	t.Run("Basic void element", func(t *testing.T) {
		tag := "img"
		expected := "<img/>"

		n := newNodeElement(true, tag)
		got, err := n.RenderString()
		if err != nil {
			t.Error(err)
		}

		if got != expected {
			t.Errorf("Expected %q, got %q", expected, got)
		}
	})

	t.Run("Element with text", func(t *testing.T) {
		tag := "p"
		children := []Node{
			newNodeText("Hello, World!"),
		}
		expected := "<p>Hello, World!</p>"

		n := newNodeElement(false, tag, children...)

		got, err := n.RenderString()
		if err != nil {
			t.Error(err)
		}

		if got != expected {
			t.Errorf("Expected %q, got %q", expected, got)
		}
	})

	t.Run("Void element with text", func(t *testing.T) {
		tag := "link"
		children := []Node{
			newNodeText("Hello, World!"),
		}
		expected := "<link/>"

		n := newNodeElement(true, tag, children...)

		got, err := n.RenderString()
		if err != nil {
			t.Error(err)
		}

		if got != expected {
			t.Errorf("Expected %q, got %q", expected, got)
		}
	})

	t.Run("Element with attributes", func(t *testing.T) {
		tag := "div"
		children := []Node{
			newNodeAttribute("class", "container"),
			newNodeAttribute("id", "main"),
		}
		expected := `<div class="container" id="main"></div>`

		n := newNodeElement(false, tag, children...)
		got, err := n.RenderString()
		if err != nil {
			t.Error(err)
		}

		if got != expected {
			t.Errorf("Expected %q, got %q", expected, got)
		}
	})

	t.Run("Void element with attributes", func(t *testing.T) {
		tag := "img"
		children := []Node{
			newNodeAttribute("src", "image.jpg"),
			newNodeAttribute("alt", "Image"),
		}
		expected := `<img src="image.jpg" alt="Image"/>`

		n := newNodeElement(true, tag, children...)
		got, err := n.RenderString()
		if err != nil {
			t.Error(err)
		}

		if got != expected {
			t.Errorf("Expected %q, got %q", expected, got)
		}
	})

	t.Run("Element with text and attributes", func(t *testing.T) {
		tag := "a"
		children := []Node{
			newNodeAttribute("href", "#"),
			newNodeAttribute("class", "btn"),
			newNodeText("Click"),
		}
		expected := `<a href="#" class="btn">Click</a>`

		n := newNodeElement(false, tag, children...)
		got, err := n.RenderString()
		if err != nil {
			t.Error(err)
		}

		if got != expected {
			t.Errorf("Expected %q, got %q", expected, got)
		}
	})

	t.Run("Void element with text and attributes", func(t *testing.T) {
		tag := "link"
		children := []Node{
			newNodeAttribute("rel", "stylesheet"),
			newNodeAttribute("href", "style.css"),
			newNodeText("This will be ignored"),
		}
		expected := `<link rel="stylesheet" href="style.css"/>`

		n := newNodeElement(true, tag, children...)
		got, err := n.RenderString()
		if err != nil {
			t.Error(err)
		}

		if got != expected {
			t.Errorf("Expected %q, got %q", expected, got)
		}
	})

	t.Run("Element with children in mixed order", func(t *testing.T) {
		tag := "div"
		children := []Node{
			newNodeAttribute("class", "container"),
			newNodeText("Hello"),
			newNodeAttribute("id", "main"),
			newNodeText(", World!"),
			newNodeAttribute("style", "color: red;"),
			newNodeText(" - nodx"),
		}
		expected := `<div class="container" id="main" style="color: red;">Hello, World! - nodx</div>`

		n := newNodeElement(false, tag, children...)
		got, err := n.RenderString()
		if err != nil {
			t.Error(err)
		}

		if got != expected {
			t.Errorf("Expected %q, got %q", expected, got)
		}
	})

	t.Run("Void element with children in mixed order", func(t *testing.T) {
		tag := "img"
		children := []Node{
			newNodeAttribute("src", "image.jpg"),
			newNodeText("This will be ignored"),
			newNodeAttribute("alt", "Image"),
			newNodeText("This will also be ignored"),
		}
		expected := `<img src="image.jpg" alt="Image"/>`

		n := newNodeElement(true, tag, children...)
		got, err := n.RenderString()
		if err != nil {
			t.Error(err)
		}

		if got != expected {
			t.Errorf("Expected %q, got %q", expected, got)
		}
	})

	t.Run("Element with nested elements", func(t *testing.T) {
		tag := "ul"
		children := []Node{
			newNodeElement(false, "li", newNodeText("Item 1")),
			newNodeElement(false, "li", newNodeText("Item 2")),
		}
		expected := "<ul><li>Item 1</li><li>Item 2</li></ul>"

		n := newNodeElement(false, tag, children...)
		got, err := n.RenderString()
		if err != nil {
			t.Error(err)
		}

		if got != expected {
			t.Errorf("Expected %q, got %q", expected, got)
		}
	})

	t.Run("Element with duplicate attributes", func(t *testing.T) {
		tag := "div"
		children := []Node{
			newNodeAttribute("class", "container"),
			newNodeAttribute("class", "main"),
		}
		expected := `<div class="container" class="main"></div>`

		n := newNodeElement(false, tag, children...)
		got, err := n.RenderString()
		if err != nil {
			t.Error(err)
		}

		if got != expected {
			t.Errorf("Expected %q, got %q", expected, got)
		}
	})

	t.Run("Element with special attributes", func(t *testing.T) {
		tag := "button"
		children := []Node{
			newNodeAttribute("data-toggle", "modal"),
			newNodeAttribute("aria-label", "Open Modal"),
		}
		expected := `<button data-toggle="modal" aria-label="Open Modal"></button>`

		n := newNodeElement(false, tag, children...)
		got, err := n.RenderString()
		if err != nil {
			t.Error(err)
		}

		if got != expected {
			t.Errorf("Expected %q, got %q", expected, got)
		}
	})

	t.Run("Element with empty tag", func(t *testing.T) {
		n := newNodeElement(false, "")
		got, err := n.RenderString()
		if err != nil {
			t.Error(err)
		}
		if got != "" {
			t.Errorf("Expected empty string, got %q", got)
		}
	})

	t.Run("Element with nil child", func(t *testing.T) {
		tag := "div"
		children := []Node{
			nil,
			newNodeText("Content"),
		}
		expected := "<div>Content</div>"

		n := newNodeElement(false, tag, children...)
		got, err := n.RenderString()
		if err != nil {
			t.Error(err)
		}

		if got != expected {
			t.Errorf("Expected %q, got %q", expected, got)
		}
	})

	t.Run("Complex element with many nested levels", func(t *testing.T) {
		tag := "html"
		children := []Node{
			newNodeElement(false, "head",
				newNodeElement(false, "title", newNodeText("Hello, World!")),
				newNodeElement(true, "meta", newNodeAttribute("charset", "UTF-8")),
			),
			newNodeElement(false, "body",
				newNodeElement(false, "div",
					newNodeAttribute("class", "container"),
					newNodeElement(false, "h1", newNodeText("Hello, World!")),
					newNodeElement(false, "p", newNodeText("This is a paragraph.")),
				),
			),
		}
		expected := "<html><head><title>Hello, World!</title><meta charset=\"UTF-8\"/></head><body><div class=\"container\"><h1>Hello, World!</h1><p>This is a paragraph.</p></div></body></html>"

		n := newNodeElement(false, tag, children...)
		got, err := n.RenderString()
		if err != nil {
			t.Error(err)
		}

		if got != expected {
			t.Errorf("Expected %q, got %q", expected, got)
		}
	})

	t.Run("Element with mixed text and element children", func(t *testing.T) {
		tag := "div"
		children := []Node{
			newNodeText("Hello, "),
			newNodeElement(false, "strong", newNodeText("World")),
			newNodeText("!"),
		}
		expected := "<div>Hello, <strong>World</strong>!</div>"

		n := newNodeElement(false, tag, children...)
		got, err := n.RenderString()
		if err != nil {
			t.Error(err)
		}
		if got != expected {
			t.Errorf("Expected %q, got %q", expected, got)
		}
	})
}
