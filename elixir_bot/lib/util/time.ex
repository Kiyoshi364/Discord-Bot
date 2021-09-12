defmodule ElixirBot.Utils.Time do
  def diff(time1, time2, unit \\ :millisecond) do
    from = fn
      %NaiveDateTime{} = x -> x
      x -> NaiveDateTime.from_iso8601!(x)
    end
    {time1, time2} = {from.(time1), from.(time2)}
    NaiveDateTime.diff(time1, time2, unit)
  end
end
