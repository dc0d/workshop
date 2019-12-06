defmodule WorkshopTest do
  use ExUnit.Case

  test "convert decimal to roman" do
    expectations = [
      {1, "I"},
      {5, "V"},
      {10, "X"},
      {50, "L"},
      {100, "C"},
      {500, "D"},
      {1000, "M"},
      {7, "VII"},
      {14, "XIV"},
      {15, "XV"},
      {99, "IC"},
      {2006, "MMVI"},
      {1944, "MCMXLIV"},
      {3497, "MMMCDXCVII"},
      {1999, "MIM"}
    ]

    for {input, expected_output} <- expectations do
      assert Workshop.to_roman(input) == expected_output
    end
  end
end
