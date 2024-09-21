import { Node } from './nodx.js'
import { escapeHtml } from './escape_html.js'

/**
 * NodeAttribute represents an HTML attribute.
 */
export class NodeAttribute implements Node {
  private readonly name: string
  private readonly value: string

  // Constructor to create a new HTML attribute.
  constructor (name: string, value: string) {
    this.name = name
    this.value = value
  }

  // Render returns the HTML attribute as a string.
  render (): string {
    if (this.name === '') return ''
    return ` ${this.name}="${escapeHtml(this.value)}"`
  }
}
