defmodule ElixirBot do
  use Application
  alias Alchemy.Client

  def start(_type, _args) do
    [token, prefix] = read_config()

    run = Client.start(token)
    Alchemy.Cogs.set_prefix(prefix)

    use ElixirBot.Commands
    use ElixirBot.Commands.Dev
    run
  end

  defp read_config() do
    {:ok, file} = File.read("config")

    file
    |> String.split("\n")
    |> Enum.take(2)
  end
end
