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

type Color = keyof typeof Resistance

export function decodedResistorValue([first, second, third]: Color[]): string {
  const firstTwoDecoded: number = parseInt(
    Resistance[first].toString() + Resistance[second].toString(),
  )
  const thirdDecoded: number = parseInt(Resistance[third].toString())
  const decodedValue: number = firstTwoDecoded * Math.pow(10, thirdDecoded)

  let result: string = 'omh shakalaka'

  if (decodedValue < 1000) {
    result = decodedValue.toString() + ' ohms'
  } else if (decodedValue < 10000) {
    result = decodedValue.toString().substring(0, 1) + ' kiloohms'
  } else if (decodedValue < 100000) {
    result = decodedValue.toString().substring(0, 2) + ' kiloohms'
  } else if (decodedValue < 1000000) {
    result = decodedValue.toString().substring(0, 3) + ' kiloohms'
  } else if (decodedValue < 10000000) {
    result = decodedValue.toString().substring(0, 1) + ' megaohms'
  } else if (decodedValue < 100000000) {
    result = decodedValue.toString().substring(0, 2) + ' megaohms'
  } else if (decodedValue < 1000000000) {
    result = decodedValue.toString().substring(0, 3) + ' megaohms'
  } else if (decodedValue < 10000000000) {
    result = decodedValue.toString().substring(0, 1) + ' gigaohms'
  } else if (decodedValue < 100000000000) {
    result = decodedValue.toString().substring(0, 2) + ' gigaohms'
  }

  return result
}
