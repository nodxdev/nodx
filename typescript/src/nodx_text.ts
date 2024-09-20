import { Node } from './nodx.js'

/**
 * NodeText represents a text node.
 *
 * IMPORTANT: This node renders the text as is, without escaping it. It is the
 * responsibility of the caller to ensure that the text is properly escaped.
 */
export class NodeText implements Node {
  text: string

  // Constructor to create a new text node.
  constructor (text: string) {
    this.text = text
  }

  // Render returns the text as a string.
  render (): string {
    return this.text
  }
}
