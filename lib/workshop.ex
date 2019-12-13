defmodule Workshop do
  alias Workshop.Items
  alias Workshop.Taxes

  @spec taxes(String.t()) :: String.t()
  def taxes(_input)

  def taxes(input) do
    items =
      input
      |> Items.parse()

    Taxes.summary(items)
  end
end

defmodule Workshop.Taxes do
  alias Workshop.Amount

  def summary(items) do
    items_with_taxes = apply(items)
    taxes_total = total(items_with_taxes)

    %{sales_taxes: taxes, total: total_amount} = taxes_total

    items_lines =
      items_with_taxes
      |> Enum.map(fn %{count: count, name: name, total: total} ->
        "#{count} #{name}: #{Amount.to_string(total)}"
      end)

    text_lines =
      items_lines ++
        ["Sales Taxes: #{Amount.to_string(taxes)}", "Total: #{Amount.to_string(total_amount)}"]

    text_lines
    |> Enum.join("\n")
    |> String.trim()
  end

  @spec total([map]) :: %{sales_taxes: float, total: any}
  def total(input) do
    {total_amount, taxes} =
      input
      |> Enum.reduce({0, 0}, fn %{count: count, name: _, price: price, total: total},
                                {total_amount, taxes} ->
        total_amount = total_amount + total
        taxes = taxes + (total - count * price)
        {total_amount, taxes}
      end)

    taxes = Amount.round_up_tax(taxes)

    %{sales_taxes: taxes, total: total_amount}
  end

  @spec apply([map]) :: [map]
  def apply(input) do
    input
    |> Enum.map(fn %{count: count, name: name, price: price} ->
      total_price = price * count

      tax = calculate_tax(name, total_price)

      total =
        (total_price + tax)
        |> Amount.round_decimals()

      %{
        count: count,
        name: name,
        price: price,
        total: total
      }
    end)
  end

  defp calculate_tax(name, price) do
    (calculate_tax_10(name, price) +
       calculate_tax_5(name, price))
    |> Amount.round_up_tax()
  end

  defp calculate_tax_5(name, price) do
    apply? =
      cond do
        String.match?(name, ~r/imported/) -> true
        true -> false
      end

    if apply? do
      5.0 * price / 100.0
    else
      0.0
    end
  end

  defp calculate_tax_10(name, price) do
    apply? =
      cond do
        String.match?(name, ~r/perfume/) -> true
        String.match?(name, ~r/music/) -> true
        true -> false
      end

    if apply? do
      10.0 * price / 100.0
    else
      0.0
    end
  end
end

defmodule Workshop.Items do
  @spec parse(String.t()) :: [map]
  def parse(items_input_text) do
    items_input_text
    |> String.split("\n")
    |> Enum.map(fn str ->
      Regex.named_captures(~r/(?<count>\d+) (?<name>.+) at (?<price>.+)/, str)
    end)
    |> Enum.map(fn %{"count" => count, "name" => name, "price" => price} ->
      %{count: String.to_integer(count), name: name, price: String.to_float(price)}
    end)
  end
end

defmodule Workshop.Amount do
  @spec round_up_tax(float) :: float
  def round_up_tax(n) do
    ceil(n * 20) / 20
  end

  @spec to_string(float) :: String.t()
  def to_string(amount) do
    :erlang.float_to_binary(amount, decimals: 2)
  end

  @spec round_decimals(float) :: float
  def round_decimals(n) do
    round(n * 100) / 100
  end
end
