export function reverseMapping(o: { [x: string]: number; }) {
  return Object.keys(o).reduce((r: { [x: number]: [string] }, k) =>
    Object.assign(r, {[o[k]]: (r[o[k]] || []).concat(k)}), {})
}

export function getRandomInt(max: number) {
  return Math.floor(Math.random() * max);
}
