defmodule WorkshopTest do
  use ExUnit.Case

  test "generate summary" do
    expectations = [
      {
        """
        2 book at 12.49
        1 music CD at 14.99
        1 chocolate bar at 0.85
        """
        |> String.trim(),
        """
        2 book: 24.98
        1 music CD: 16.49
        1 chocolate bar: 0.85
        Sales Taxes: 1.50
        Total: 42.32
        """
        |> String.trim()
      },
      {
        """
        1 imported box of chocolates at 10.00
        1 imported bottle of perfume at 47.50
        """
        |> String.trim(),
        """
        1 imported box of chocolates: 10.50
        1 imported bottle of perfume: 54.65
        Sales Taxes: 7.65
        Total: 65.15
        """
        |> String.trim()
      }
      # {
      #   """
      #   1 imported bottle of perfume at 27.99
      #   1 bottle of perfume at 18.99
      #   1 packet of headache pills at 9.75
      #   3 imported box of chocolates at 11.25
      #   """
      #   |> String.trim(),
      #   """
      #   1 imported bottle of perfume: 32.19
      #   1 bottle of perfume: 20.89
      #   1 packet of headache pills: 9.75
      #   3 imported box of chocolates: 35.55
      #   Sales Taxes: 7.90
      #   Total: 98.38
      #   """
      #   |> String.trim()
      # }
    ]

    for {input, expected_output} <- expectations do
      assert Workshop.taxes(input) == expected_output
    end
  end
end

defmodule Workshop.TaxesTest do
  use ExUnit.Case

  alias Workshop.Taxes

  test "summary" do
    expectations = [
      {
        [
          %{
            count: 2,
            name: "book",
            price: 12.49,
            total: 24.98
          },
          %{
            count: 1,
            name: "music CD",
            price: 14.99,
            total: 16.49
          },
          %{
            count: 1,
            name: "chocolate bar",
            price: 0.85,
            total: 0.85
          }
        ],
        """
        2 book: 24.98
        1 music CD: 16.49
        1 chocolate bar: 0.85
        Sales Taxes: 1.50
        Total: 42.32
        """
        |> String.trim()
      }
    ]

    for {input, expected_output} <- expectations do
      assert Taxes.summary(input) == expected_output
    end
  end

  test "total" do
    expectations = [
      {
        [
          %{
            count: 2,
            name: "book",
            price: 12.49,
            total: 24.98
          },
          %{
            count: 1,
            name: "music CD",
            price: 14.99,
            total: 16.49
          },
          %{
            count: 1,
            name: "chocolate bar",
            price: 0.85,
            total: 0.85
          }
        ],
        %{
          sales_taxes: 1.50,
          total: 42.32
        }
      },
      {
        [
          %{
            count: 1,
            name: "imported box of chocolates",
            price: 10.00,
            total: 10.50
          },
          %{
            count: 1,
            name: "imported bottle of perfume",
            price: 47.50,
            total: 54.63
          }
        ],
        %{
          sales_taxes: 7.65,
          total: 65.13
        }
      }
    ]

    for {input, expected_output} <- expectations do
      assert Taxes.total(input) == expected_output
    end
  end

  test "apply" do
    expectations = [
      {
        [
          %{
            count: 2,
            name: "book",
            price: 12.49
          },
          %{
            count: 1,
            name: "music CD",
            price: 14.99
          },
          %{
            count: 1,
            name: "chocolate bar",
            price: 0.85
          }
        ],
        [
          %{
            count: 2,
            name: "book",
            price: 12.49,
            total: 24.98
          },
          %{
            count: 1,
            name: "music CD",
            price: 14.99,
            total: 16.49
          },
          %{
            count: 1,
            name: "chocolate bar",
            price: 0.85,
            total: 0.85
          }
        ]
      },
      {
        [
          %{
            count: 1,
            name: "imported box of chocolates",
            price: 10.00
          },
          %{
            count: 1,
            name: "imported bottle of perfume",
            price: 47.50
          }
        ],
        [
          %{
            count: 1,
            name: "imported box of chocolates",
            price: 10.00,
            total: 10.50
          },
          %{
            count: 1,
            name: "imported bottle of perfume",
            price: 47.50,
            total: 54.65
          }
        ]
      }
    ]

    for {input, expected_output} <- expectations do
      assert Taxes.apply(input) == expected_output
    end
  end
end

defmodule Workshop.ItemsTest do
  use ExUnit.Case

  alias Workshop.Items

  test "parse" do
    expectations = [
      {
        """
        2 book at 12.49
        1 music CD at 14.99
        1 chocolate bar at 0.85
        """
        |> String.trim(),
        [
          %{
            count: 2,
            name: "book",
            price: 12.49
          },
          %{
            count: 1,
            name: "music CD",
            price: 14.99
          },
          %{
            count: 1,
            name: "chocolate bar",
            price: 0.85
          }
        ]
      },
      {
        """
        1 imported box of chocolates at 10.00
        1 imported bottle of perfume at 47.50
        """
        |> String.trim(),
        [
          %{
            count: 1,
            name: "imported box of chocolates",
            price: 10.00
          },
          %{
            count: 1,
            name: "imported bottle of perfume",
            price: 47.50
          }
        ]
      },
      {
        """
        1 imported bottle of perfume at 27.99
        1 bottle of perfume at 18.99
        1 packet of headache pills at 9.75
        3 box of imported chocolates at 11.25
        """
        |> String.trim(),
        [
          %{
            count: 1,
            name: "imported bottle of perfume",
            price: 27.99
          },
          %{
            count: 1,
            name: "bottle of perfume",
            price: 18.99
          },
          %{
            count: 1,
            name: "packet of headache pills",
            price: 9.75
          },
          %{
            count: 3,
            name: "box of imported chocolates",
            price: 11.25
          }
        ]
      }
    ]

    for {input, expected_output} <- expectations do
      assert Items.parse(input) == expected_output
    end
  end
end

defmodule Workshop.AmountTest do
  use ExUnit.Case

  alias Workshop.Amount

  test "round_up" do
    expectations = [
      {16.489, 16.5},
      {16.625, 16.65}
    ]

    for {input, expected_output} <- expectations do
      assert Amount.round_up_tax(input) == expected_output
    end
  end

  test "round_decimals" do
    expectations = [
      {16.490000000000002, 16.49}
    ]

    for {input, expected_output} <- expectations do
      assert Amount.round_decimals(input) == expected_output
    end
  end

  test "to_string" do
    expectations = [
      {1.499999999, "1.50"}
    ]

    for {input, expected_output} <- expectations do
      assert Amount.to_string(input) == expected_output
    end
  end
end
