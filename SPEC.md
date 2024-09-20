# Specification for "nodx" Cross-Language Template Engine

## Overview

**nodx** is a cross-language HTML template engine that enables developers to construct HTML DOM structures using native functions or methods in various programming languages. In **nodx**, **everything is a node**—both HTML elements and attributes are treated as nodes. This unified approach simplifies the programmatic creation and manipulation of HTML structures.

The DOM structure is built using functions or methods that represent HTML elements and attributes. Once the node tree is constructed, a `render` method/function is used to produce the final rendered template, typically as an HTML string.

## Core Concepts

### Nodes

- **Element Nodes**: Represent HTML elements (e.g., `div`, `img`).
- **Attribute Nodes**: Represent HTML attributes (e.g., `class`, `src`).
- **Text Nodes**: Represent text content within elements.

All nodes can render themselves into a string representation that forms part of the final HTML output.

## API

### Node Creation

#### Element Nodes

- Created using functions or constructors named after HTML elements.
- Accept a variable list of attributes and child nodes.

**Syntax Example:**

```pseudo
Element(attributes..., children...) -> ElementNode
```

#### Attribute Nodes

- Created using functions or constructors named after HTML attributes.
- Accept a value parameter.

**Syntax Example:**

```pseudo
Attribute(value) -> AttributeNode
```

#### Text Nodes

- Created using a specific function for text content.

**Syntax Example:**

```pseudo
Text(content) -> TextNode
```

### Rendering

- When possible each node has a `render` that produces its string representation.
- Rendering is recursive: element nodes render their attributes and children.

## Usage Example

### Creating a Component

Create an avatar component:

**Pseudo-code:**

```pseudo
avatar = Div(
    Class("avatar-container"),
    Img(
        Class("avatar-image"),
        Alt("Avatar"),
        Src("avatar.png")
    )
)
```

### Rendering the Structure

Render the node tree to produce the final output:

```pseudo
output = avatar.render()
```

**Expected Output:**

```html
<div class="avatar-container">
    <img class="avatar-image" alt="Avatar" src="avatar.png" />
</div>
```

## Advantages of the Simplified Approach

- **Simplicity**: Fewer classes and interfaces make it easier to understand and implement.
- **Language-Agnostic**: The design can be implemented in any programming language.
- **Flexibility**: Developers can easily extend, customize the functionality and build component libraries.
- **No special syntax**: No need for special syntax or templating language. This allows to use the full power of the programming language to create and manipulate the templates.

## Final Considerations

- **Void Elements**: Handle HTML elements that do not require closing tags (e.g., `img`, `br`) appropriately.
- **HTML Escaping**: Always escape attribute values and text content to prevent security vulnerabilities.
- **Customization**: Developers can create additional helper functions or nodes for custom components.
