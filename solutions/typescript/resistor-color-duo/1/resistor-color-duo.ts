enum COLORS {
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

export function decodedValue(colors: string[]): number {
  const val1: COLORS = COLORS[colors[0] as keyof typeof COLORS]
  const val2: COLORS = COLORS[colors[1] as keyof typeof COLORS]

  return parseInt(val1.toString() + val2.toString())
}