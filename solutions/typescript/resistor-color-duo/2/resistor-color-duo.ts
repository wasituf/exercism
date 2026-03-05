enum Resistance {
  black,
  brown,
  red,
  orange,
  yellow,
  green,
  blue,
  violet,
  grey,
  white,
}

type COLORS = keyof typeof Resistance

export function decodedValue(colors: COLORS[]): number {
  const val1: Resistance = Resistance[colors[0]]
  const val2: Resistance = Resistance[colors[1]]

  return parseInt(val1.toString() + val2.toString())
}