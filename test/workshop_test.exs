defmodule WorkshopTest do
  use ExUnit.Case

  test "calculate nth fibonacci number" do
    expectations = [
      {1, 1},
      {2, 1},
      {3, 2},
      {4, 3},
      {10, 55},
      {12, 144},
      {20, 6765}
    ]

    for {input, expected_output} <- expectations do
      assert Workshop.fib(input) == expected_output
    end
  end
end
