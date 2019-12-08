defmodule Workshop do
  def generate(n) do
    case n do
      1 ->
        []

      _ ->
        find(n)
    end
  end

  defp find(n) do
    {:done, result} =
      Stream.iterate({n, 2, []}, fn
        {n, candidate, acc_candidates} ->
          {n, candidate, acc_candidates} = next(n, candidate, acc_candidates)

          if n < candidate do
            {:done, acc_candidates}
          else
            {n, candidate, acc_candidates}
          end
      end)
      |> Stream.drop_while(fn
        {:done, _} ->
          false

        _ ->
          true
      end)
      |> Enum.at(0)

    result
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
