import { describe, it, expect } from 'vitest'
import { NodeAttribute } from './nodx_attribute.js'

describe('NodeAttribute', () => {
  it('Attribute render', () => {
    const key = 'class'
    const value = 'container'
    const expected = ' class="container"'

    const n = new NodeAttribute(key, value)
    const got = n.render()

    expect(got).toBe(expected)
  })

  it('Empty attribute render', () => {
    const key = 'class'
    const value = ''
    const expected = ' class=""'

    const n = new NodeAttribute(key, value)
    const got = n.render()

    expect(got).toBe(expected)
  })

  it('Empty key attribute render', () => {
    const key = ''
    const value = 'container'
    const expected = ''

    const n = new NodeAttribute(key, value)
    const got = n.render()

    expect(got).toBe(expected)
  })

  it('Empty key and value attribute render', () => {
    const key = ''
    const value = ''
    const expected = ''

    const n = new NodeAttribute(key, value)
    const got = n.render()

    expect(got).toBe(expected)
  })

  it('Double quote in value attribute render', () => {
    const key = 'class'
    const value = 'container"container\'container'
    const expected = ' class="container&#34;container&#39;container"'

    const n = new NodeAttribute(key, value)
    const got = n.render()

    expect(got).toBe(expected)
  })
})
