defmodule WorkshopTest do
  use ExUnit.Case

  test "sort by natural order" do
    expectations = [
      %{
        input: ["Alpha 100", "Alpha 1"],
        expected_output: ["Alpha 1", "Alpha 100"]
      },
      %{
        input: ["Alpha 100", "Alpha 2"],
        expected_output: ["Alpha 2", "Alpha 100"]
      },
      %{
        input: ["Alpha 2A-8000", "Alpha 100"],
        expected_output: ["Alpha 2A-8000", "Alpha 100"]
      },
      %{
        input: [
          "Alpha 100",
          "Alpha 2",
          "Alpha 200",
          "Alpha 2A",
          "Alpha 2A-8000",
          "Alpha 2A-900"
        ],
        expected_output: [
          "Alpha 2",
          "Alpha 2A",
          "Alpha 2A-900",
          "Alpha 2A-8000",
          "Alpha 100",
          "Alpha 200"
        ]
      }
    ]

    for %{input: input, expected_output: expected_output} <- expectations do
      assert Workshop.sort(input) == expected_output
    end
  end
end
