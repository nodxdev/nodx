// escapeHTML escapes the input string to prevent XSS attacks.
export function escapeHtml (unsafe: string): string {
  return unsafe
    .replace(/&/g, '&#38;')
    .replace(/</g, '&#60;')
    .replace(/>/g, '&#62;')
    .replace(/"/g, '&#34;')
    .replace(/'/g, '&#39;')
}
