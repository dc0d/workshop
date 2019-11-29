defmodule Workshop do
  def fib(n) do
    {nthfib, _} =
      {0, 1}
      |> Stream.iterate(fn {a, b} -> {b, a + b} end)
      |> Enum.at(n)

    nthfib
  end
end
