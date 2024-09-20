import { describe, it, expect } from 'vitest'
import { escapeHtml } from './escape_html.js'

describe('escapeHtml', () => {
  it('Escape HTML', () => {
    const input = "<script>alert('XSS')</script>"
    const expected = '&#60;script&#62;alert(&#39;XSS&#39;)&#60;/script&#62;'
    const got = escapeHtml(input)
    expect(got).toBe(expected)
  })

  it('All characters', () => {
    const input = "test-&<>'\"-test"
    const expected = 'test-&#38;&#60;&#62;&#39;&#34;-test'
    const got = escapeHtml(input)
    expect(got).toBe(expected)
  })

  it('No escape', () => {
    const input = 'test'
    const expected = 'test'
    const got = escapeHtml(input)
    expect(got).toBe(expected)
  })

  it('Empty string', () => {
    const input = ''
    const expected = ''
    const got = escapeHtml(input)
    expect(got).toBe(expected)
  })

  it('Escape with special characters', () => {
    const input = "test👌áéíóú-&<>'\"-👌test"
    const expected = 'test👌áéíóú-&#38;&#60;&#62;&#39;&#34;-👌test'
    const got = escapeHtml(input)
    expect(got).toBe(expected)
  })

  it('No escape with special characters', () => {
    const input = 'test👌áéíóú'
    const expected = 'test👌áéíóú'
    const got = escapeHtml(input)
    expect(got).toBe(expected)
  })

  it('Single special character', () => {
    const input = '👌'
    const expected = '👌'
    const got = escapeHtml(input)
    expect(got).toBe(expected)
  })
})
