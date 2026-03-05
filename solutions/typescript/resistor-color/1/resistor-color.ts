const colorsArray: string[] = [
  'black',
  'brown',
  'red',
  'orange',
  'yellow',
  'green',
  'blue',
  'violet',
  'grey',
  'white',
]

export const colorCode = (color: string): number => {
  return colorsArray.indexOf(color.toLowerCase())
}

export const COLORS = colorsArray
