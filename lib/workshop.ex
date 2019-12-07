defmodule Workshop do
  def sort(input) do
    input
    |> split_to_segments()
    |> sort_by_segments()
    |> Enum.map(fn {str, _segments} -> str end)
  end

  #

  defp sort_by_segments(segments) do
    segments
    |> Enum.sort_by(fn {_str, segments} -> segments end, fn segments1, segments2 ->
      less(segments1, segments2)
    end)
  end

  #

  defp split_to_segments(input) do
    input
    |> Enum.map(fn str ->
      segments = split_string_segments(str)

      {str, segments}
    end)
  end

  defp split_string_segments(str) do
    str
    |> String.codepoints()
    |> Enum.chunk_by(fn x ->
      "0" <= x and x <= "9"
    end)
    |> Enum.map(fn x ->
      segment = join_chars(x)

      first = Enum.at(x, 0)

      if "0" <= first and first <= "9" do
        String.to_integer(segment)
      else
        segment
      end
    end)
  end

  defp join_chars(chars) do
    chars
    |> Enum.reduce("", fn x, acc -> acc <> x end)
  end

  #

  defp less([], _) do
    true
  end

  defp less([h1 | _t1], [h2 | _t2]) when is_number(h1) and is_binary(h2) do
    true
  end

  defp less([h1 | _t1], [h2 | _t2]) when is_binary(h1) and is_number(h2) do
    false
  end

  defp less([h1 | t1], [h2 | t2]) do
    cond do
      h1 < h2 -> true
      h1 > h2 -> false
      true -> less(t1, t2)
    end
  end
end
