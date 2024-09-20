import { describe, it, expect } from 'vitest'
import { NodeText } from './nodx_text.js'

describe('NodeText', () => {
  it('Text render', () => {
    const text = 'Hello, World!'
    const n = new NodeText(text)
    const renderedText = n.render()
    expect(renderedText).toBe(text)
  })

  it('Empty text render', () => {
    const n = new NodeText('')
    const renderedText = n.render()
    expect(renderedText).toBe('')
  })
})
