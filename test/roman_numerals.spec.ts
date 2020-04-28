import { toRomanNumerals } from "../src/roman_numerals";

import { expect } from "chai";
import "mocha";

function testCase(input: number, expectedRomannumber: string) {
  it(`should return roman number ${expectedRomannumber} for number ${input}`, () => {
    const result = toRomanNumerals(input);

    expect(result).to.equal(expectedRomannumber);
  });
}

function expectations() {
  return new Array<[number, string]>(
    [1, "I"],
    [5, "V"],
    [10, "X"],
    [50, "L"],
    [100, "C"],
    [500, "D"],
    [1000, "M"],
    [7, "VII"],
    [14, "XIV"],
    [15, "XV"],
    [99, "IC"],
    [2006, "MMVI"],
    [1944, "MCMXLIV"],
    [3497, "MMMCDXCVII"],
    [1999, "MIM"],
    [2020, "MMXX"],
    [509, "DIX"]
  );
}

describe("toRomanNumerals", () => {
  expectations().forEach((expectation) => {
    const [input, expectedRomannumber] = expectation;
    testCase(input, expectedRomannumber);
  });
});
