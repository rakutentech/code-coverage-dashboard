/**
 * queryToString return string from query param
 * @param {string | string[] | undefined} val query param
 * @returns {string} query string
 */
export const queryToString = (val: string | string[] | undefined): string => {
  const type = typeof val
  switch (true) {
    case type === 'string':
      return val as string
    case Array.isArray(val):
      return (val as string[])[0] // return only first one
    default:
      return ''
  }
}

