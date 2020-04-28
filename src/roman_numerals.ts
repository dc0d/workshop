const NOT_FOUND = "NOT FOUND";

function decToRoman(num: number): string {
  switch (num) {
    case 1000:
      return "M";
    case 500:
      return "D";
    case 100:
      return "C";
    case 50:
      return "L";
    case 10:
      return "X";
    case 5:
      return "V";
    case 1:
      return "I";
    default:
      return NOT_FOUND;
  }
}

function romanFloor(n: number): number {
  if (n > 1000) return 1000;
  if (n > 500) return 500;
  if (n > 100) return 100;
  if (n > 50) return 50;
  if (n > 10) return 10;
  if (n > 5) return 5;

  return 1;
}

function romanCeiling(n: number): [number, number] {
  let ceiling = -1;
  let diffToCeiling = -1;

  if (n <= 1) [ceiling, diffToCeiling] = [1, 1 - n];
  else if (n <= 5) [ceiling, diffToCeiling] = [5, 5 - n];
  else if (n <= 10) [ceiling, diffToCeiling] = [10, 10 - n];
  else if (n <= 50) [ceiling, diffToCeiling] = [50, 50 - n];
  else if (n <= 100) [ceiling, diffToCeiling] = [100, 100 - n];
  else if (n <= 500) [ceiling, diffToCeiling] = [500, 500 - n];
  else [ceiling, diffToCeiling] = [1000, 1000 - n];

  return [ceiling, diffToCeiling];
}

function div(n: number, roman: number): [number, number] {
  const q: number = Math.floor(n / roman);
  const r: number = n % roman;
  return [q, r];
}

export function convert(num: number): string {
  let roman = decToRoman(num);
  if (roman != NOT_FOUND) {
    return roman;
  }

  const [ceiling, diff] = romanCeiling(num);
  const ceilingNumeral = decToRoman(diff);
  if (ceilingNumeral != NOT_FOUND) {
    return ceilingNumeral + decToRoman(ceiling);
  }

  const floor = romanFloor(num);
  const [divisor, remainder] = div(num, floor);

  let result = decToRoman(floor).repeat(divisor);
  if (remainder > 0) {
    result += convert(remainder);
  }

  return result;
}

export function toRomanNumerals(num: number): string {
  let result = convert(num);

  result = result.replace("LXXXX", "XC");
  result = result.replace("VIIII", "IX");
  result = result.replace("DCCCC", "CM");
  result = result.replace("IIII", "IV");
  result = result.replace("XXXX", "XL");
  result = result.replace("CCCC", "CD");

  return result;
}
