defmodule WorkshopTest do
  use ExUnit.Case

  test "prime factors" do
    expectations = [
      {1, []},
      {2, [2]},
      {3, [3]},
      {4, [2, 2]},
      {6, [2, 3]},
      {7, [7]},
      {8, [2, 2, 2]},
      {9, [3, 3]},
      {4620, [2, 2, 3, 5, 7, 11]}
    ]

    for {input, expected_output} <- expectations do
      assert Workshop.generate(input) == expected_output
    end
  end
end
