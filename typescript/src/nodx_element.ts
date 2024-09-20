import { Node } from './nodx.js'
import { NodeAttribute } from './nodx_attribute.js'
import { NodeText } from './nodx_text.js'

/**
 * NodeElement represents an HTML element.
 */
export class NodeElement implements Node {
  isVoid: boolean
  name: string
  children: Node[]

  // Constructor to create a new HTML element.
  constructor (isVoid: boolean, name: string, ...children: Node[]) {
    this.isVoid = isVoid
    this.name = name
    this.children = children
  }

  // Render returns the HTML element and all its children as a string.
  render (): string {
    if (this.name === '') return ''

    let result = `<${this.name}`

    // Render only NodeAttribute instances
    for (const child of this.children) {
      if (child instanceof NodeAttribute) {
        result += child.render()
      }
    }

    if (this.isVoid) {
      result += '/>'
    }

    if (!this.isVoid) {
      result += '>'

      // Render only NodeElement and NodeText instances
      for (const child of this.children) {
        if (child instanceof NodeElement || child instanceof NodeText) {
          result += child.render()
        }
      }

      result += `</${this.name}>`
    }

    return result
  }
}
