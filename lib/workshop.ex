defmodule Workshop do
  def to_roman(n) do
    n
    |> convert_to_roman()
    |> replace_patterns()
  end

  #

  @dec_to_roman %{
    1000 => "M",
    500 => "D",
    100 => "C",
    50 => "L",
    10 => "X",
    5 => "V",
    1 => "I"
  }

  #

  defp replace_patterns(roman) do
    roman
    |> replace_string("LXXXX", "XC")
    |> replace_string("VIIII", "IX")
    |> replace_string("DCCCC", "CM")
    |> replace_string("IIII", "IV")
    |> replace_string("XXXX", "XL")
    |> replace_string("CCCC", "CD")
  end

  defp replace_string(roman, long, short) do
    String.replace(roman, long, short, global: true)
  end

  #

  defp convert_to_roman(n) do
    n
    |> roman_string()
    |> (fn
          nil ->
            roman_ceiling(n)

          r ->
            r
        end).()
    |> (fn
          {ceiling, diff_to_ceiling} ->
            {ceiling, roman_string(diff_to_ceiling)}

          x ->
            {:continue, x}
        end).()
    |> (fn
          {:continue, x} ->
            x

          {_, nil} ->
            r_floor = roman_floor(n)
            divisor = div(n, r_floor)
            remainder = rem(n, r_floor)

            repeat(roman_string(r_floor), divisor) <> process_rem(remainder)

          {ceiling, r_diff_str} ->
            r_diff_str <> roman_string(ceiling)
        end).()
  end

  defp roman_string(rnum) do
    Map.get(@dec_to_roman, rnum)
  end

  defp process_rem(remainder) do
    if remainder > 0 do
      convert_to_roman(remainder)
    else
      ""
    end
  end

  defp repeat(roman, divisor) do
    if divisor > 0 do
      String.duplicate(roman, divisor)
    else
      ""
    end
  end

  defp roman_ceiling(n) do
    cond do
      n <= 1 ->
        {1, 1 - n}

      n <= 5 ->
        {5, 5 - n}

      n <= 10 ->
        {10, 10 - n}

      n <= 50 ->
        {50, 50 - n}

      n <= 100 ->
        {100, 100 - n}

      n <= 500 ->
        {500, 500 - n}

      n <= 1000 ->
        {1000, 1000 - n}

      true ->
        {:error, :too_big_to_be_roman}
    end
  end

  defp roman_floor(n) do
    cond do
      n > 1000 ->
        1000

      n > 500 ->
        500

      n > 100 ->
        100

      n > 50 ->
        50

      n > 10 ->
        10

      n > 5 ->
        5

      true ->
        1
    end
  end
end
