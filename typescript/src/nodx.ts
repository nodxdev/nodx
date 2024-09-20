/**
 * Node is the interface that wraps the basic render methods used in the
 * text, attribute, and element nodes.
 */
export interface Node {
  render: () => string
}
