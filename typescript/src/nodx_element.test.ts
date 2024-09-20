import { describe, it, expect } from 'vitest'
import { NodeText } from './nodx_text.js'
import { NodeElement } from './nodx_element.js'
import { NodeAttribute } from './nodx_attribute.js'

describe('NodeElement', () => {
  it('Basic element render', () => {
    const tag = 'div'
    const expected = '<div></div>'

    const n = new NodeElement(false, tag)
    const got = n.render()

    expect(got).toBe(expected)
  })

  it('Basic void element', () => {
    const tag = 'img'
    const expected = '<img/>'

    const n = new NodeElement(true, tag)
    const got = n.render()

    expect(got).toBe(expected)
  })

  it('Element with text', () => {
    const tag = 'p'
    const expected = '<p>Hello, World!</p>'

    const n = new NodeElement(false, tag, new NodeText('Hello, World!'))
    const got = n.render()

    expect(got).toBe(expected)
  })

  it('Void element with text', () => {
    const tag = 'link'
    const expected = '<link/>'

    const n = new NodeElement(true, tag, new NodeText('Hello, World!'))
    const got = n.render()

    expect(got).toBe(expected)
  })

  it('Element with attributes', () => {
    const tag = 'div'
    const expected = '<div class="container" id="main"></div>'

    const n = new NodeElement(
      false,
      tag,
      new NodeAttribute('class', 'container'),
      new NodeAttribute('id', 'main')
    )
    const got = n.render()

    expect(got).toBe(expected)
  })

  it('Void element with attributes', () => {
    const tag = 'img'
    const expected = '<img src="image.jpg" alt="Image"/>'

    const n = new NodeElement(
      true,
      tag,
      new NodeAttribute('src', 'image.jpg'),
      new NodeAttribute('alt', 'Image')
    )
    const got = n.render()

    expect(got).toBe(expected)
  })

  it('Element with text and attributes', () => {
    const tag = 'a'
    const expected = '<a href="#" class="btn">Click</a>'

    const n = new NodeElement(
      false,
      tag,
      new NodeAttribute('href', '#'),
      new NodeAttribute('class', 'btn'),
      new NodeText('Click')
    )
    const got = n.render()

    expect(got).toBe(expected)
  })

  it('Void element with text and attributes', () => {
    const tag = 'link'
    const expected = '<link rel="stylesheet" href="style.css"/>'

    const n = new NodeElement(
      true,
      tag,
      new NodeAttribute('rel', 'stylesheet'),
      new NodeAttribute('href', 'style.css'),
      new NodeText('This will be ignored')
    )
    const got = n.render()

    expect(got).toBe(expected)
  })

  it('Element with children in mixed order', () => {
    const tag = 'div'
    const expected = '<div class="container" id="main" style="color: red;">Hello, World! - nodx</div>'

    const n = new NodeElement(
      false,
      tag,
      new NodeAttribute('class', 'container'),
      new NodeText('Hello'),
      new NodeAttribute('id', 'main'),
      new NodeText(', World!'),
      new NodeAttribute('style', 'color: red;'),
      new NodeText(' - nodx')
    )
    const got = n.render()

    expect(got).toBe(expected)
  })

  it('Void element with children in mixed order', () => {
    const tag = 'img'
    const expected = '<img src="image.jpg" alt="Image"/>'

    const n = new NodeElement(
      true,
      tag,
      new NodeAttribute('src', 'image.jpg'),
      new NodeText('This will be ignored'),
      new NodeAttribute('alt', 'Image'),
      new NodeText('This will also be ignored')
    )
    const got = n.render()

    expect(got).toBe(expected)
  })

  it('Element with nested elements', () => {
    const tag = 'ul'
    const expected = '<ul><li>Item 1</li><li>Item 2</li></ul>'

    const n = new NodeElement(
      false,
      tag,
      new NodeElement(false, 'li', new NodeText('Item 1')),
      new NodeElement(false, 'li', new NodeText('Item 2'))
    )
    const got = n.render()

    expect(got).toBe(expected)
  })

  it('Element with duplicate attributes', () => {
    const tag = 'div'
    const expected = '<div class="container" class="main"></div>'

    const n = new NodeElement(
      false,
      tag,
      new NodeAttribute('class', 'container'),
      new NodeAttribute('class', 'main')
    )
    const got = n.render()

    expect(got).toBe(expected)
  })

  it('Element with special attributes', () => {
    const tag = 'button'
    const expected = '<button data-toggle="modal" aria-label="Open Modal"></button>'

    const n = new NodeElement(
      false,
      tag,
      new NodeAttribute('data-toggle', 'modal'),
      new NodeAttribute('aria-label', 'Open Modal')
    )
    const got = n.render()

    expect(got).toBe(expected)
  })

  it('Element with empty tag', () => {
    const n = new NodeElement(false, '')
    const got = n.render()

    expect(got).toBe('')
  })

  it('Element with null child', () => {
    const tag = 'div'
    const expected = '<div>Content</div>'

    const n = new NodeElement(false, tag, null as any, new NodeText('Content'))
    const got = n.render()

    expect(got).toBe(expected)
  })

  it('Complex element with many nested levels', () => {
    const tag = 'html'
    const expected =
      '<html><head><title>Hello, World!</title><meta charset="UTF-8"/></head><body><div class="container"><h1>Hello, World!</h1><p>This is a paragraph.</p></div></body></html>'

    const n = new NodeElement(
      false,
      tag,
      new NodeElement(
        false,
        'head',
        new NodeElement(false, 'title', new NodeText('Hello, World!')),
        new NodeElement(true, 'meta', new NodeAttribute('charset', 'UTF-8'))
      ),
      new NodeElement(
        false,
        'body',
        new NodeElement(
          false,
          'div',
          new NodeAttribute('class', 'container'),
          new NodeElement(false, 'h1', new NodeText('Hello, World!')),
          new NodeElement(false, 'p', new NodeText('This is a paragraph.'))
        )
      )
    )

    const got = n.render()

    expect(got).toBe(expected)
  })

  it('Element with mixed text and element children', () => {
    const tag = 'div'
    const expected = '<div>Hello, <strong>World</strong>!</div>'

    const n = new NodeElement(
      false,
      tag,
      new NodeText('Hello, '),
      new NodeElement(false, 'strong', new NodeText('World')),
      new NodeText('!')
    )
    const got = n.render()

    expect(got).toBe(expected)
  })

  it('Element with xss injection', () => {
    const tag = 'div'
    const expected = '<div style="alert(&#34;Hello, World!&#34;);">This is a paragraph.</div>'

    const n = new NodeElement(
      false,
      tag,
      new NodeAttribute('style', 'alert("Hello, World!");'),
      new NodeText('This is a paragraph.')
    )
    const got = n.render()

    expect(got).toBe(expected)
  })
})
