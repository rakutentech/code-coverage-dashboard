/**
 * Generate a function to debounce processing
 * @param {(arg: T) => void} fn callback function
 * @param {number} intervalMsec debounce interval msec
 * @returns {(arg: T, force?: boolean) => void} debounce function (immediate execute when force is true))
 */
export const debounce = <T>(fn: (arg: T) => void, intervalMsec: number): (arg: T, force?: boolean) => void => {
  var timer: NodeJS.Timeout
  return (arg, force = false) => {
    if (force) {
      clearTimeout(timer)
      return fn(arg)
    }
    console.log('clear timer: ', timer)
    clearTimeout(timer)
    timer = setTimeout(() => fn(arg), intervalMsec)
  }
}
