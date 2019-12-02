defmodule Workshop do
  def generate(n) do
    case n do
      1 ->
        []

      _ ->
        find(n, 2, [])
    end
  end

  defp find(n, candidate, acc_candidates) do
    {n, candidate, acc_candidates} = next(n, candidate, acc_candidates)

    if n < candidate do
      acc_candidates
    else
      find(n, candidate, acc_candidates)
    end
  end

  defp next(n, candidate, acc_candidates) when rem(n, candidate) == 0 do
    n = div(n, candidate)
    acc_candidates = acc_candidates ++ [candidate]
    {n, candidate, acc_candidates}
  end

  defp next(n, candidate, acc_candidates) do
    {n, candidate + 1, acc_candidates}
  end
end
