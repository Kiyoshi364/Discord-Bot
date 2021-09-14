defmodule ElixirBot.Utils.Channel do
  def channelCategory?(channel) do
    case channel do
      %{bitrate: _bitrate} -> false
      %{topic: _topic} -> false
      _ -> true
    end
  end

  def textChannel?(channel) do
    case channel do
      %{topic: _topic} -> true
      _ -> false
    end
  end

  def voiceChannel?(channel) do
    case channel do
      %{bitrate: _bitrate} -> true
      _ -> false
    end
  end
end
