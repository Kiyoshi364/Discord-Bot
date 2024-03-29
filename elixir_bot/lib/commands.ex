defmodule ElixirBot.Commands do
  use Alchemy.Cogs
  # require ElixirBot.Commands.Dev

  alias Alchemy.{Client, Embed}
  # alias Alchemy.{Client, Cache, User, Embed}
  alias ElixirBot.Utils.{Time, Channel}

  def all_commands() do
    Cogs.all_commands()
  end

  @default_color 0x880085
  @default_color 0x410056

  Cogs.set_parser(:ping, &List.wrap/1)

  Cogs.def ping(rest \\ "") do
    IO.inspect("ping_command: {#{rest}}")

    old_time = message.timestamp

    task = Task.async(fn -> Cogs.say("pong!") end)
    {:ok, new_message} = Task.await(task)
    # {:ok, new_message} = Cogs.say "pong!"

    time = Time.diff(new_message.timestamp, old_time)
    Client.edit_message(new_message, new_message.content <> " (#{time} ms)")
  end

  Cogs.set_parser(:hello, &List.wrap/1)

  Cogs.def hello(rest \\ "") do
    rest =
      case rest do
        "" -> "#{message.author.username}"
        _ -> rest
      end

    Cogs.say("Hallo, #{message.author.username} #{rest}")
  end

  Cogs.set_parser(:hallo, &List.wrap/1)

  Cogs.def hallo(rest \\ "") do
    rest =
      case rest do
        "" -> "#{message.author.username}"
        _ -> rest
      end

    require Alchemy.Embed

    %Embed{}
    |> Embed.title("#{message.author.username}")
    |> Embed.description("Hallo, #{rest}.")
    |> Embed.color(@default_color)
    |> Embed.send()
  end

  Cogs.set_parser(:join, &List.wrap/1)

  Cogs.def join(rest \\ "") do
    alias Alchemy.Voice

    user_id = message.author.id
    {:ok, guild} = Cogs.guild()

    voice_channel =
      guild.channels
      |> Enum.filter(&Channel.voiceChannel?/1)
      |> hd

    IO.inspect(rest)
    IO.inspect(guild.id)
    IO.inspect(voice_channel.id)
    IO.inspect(voice_channel)

    Voice.join(guild.id, voice_channel.id)
    |> IO.inspect()

    Cogs.say("Join")
  end
end
