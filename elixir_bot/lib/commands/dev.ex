defmodule ElixirBot.Commands.Dev do
  use Alchemy.Cogs
  alias Alchemy.{Embed}
  alias ElixirBot.Utils.Channel

  @default_color 0x410056

  Cogs.group("dev")

  Cogs.set_parser(:tt, &[String.split(&1)])

  Cogs.def tt(_rest \\ "") do
    # alias Alchemy.Client
    _u = message
    # Client.create_channel(message.gui
  end

  Cogs.def message(_rest \\ "") do
    require Alchemy.Embed

    attachments =
      if message.attachments != [] do
        message.attachments
      else
        "[]"
      end

    edited_timestamp =
      if message.edited_timestamp do
        message.edited_timestamp
      else
        "nil"
      end

    embeds =
      if message.embeds != [] do
        message.embeds
      else
        "[]"
      end

    mention_roles =
      if message.mention_roles != [] do
        message.mention_roles
      else
        "[]"
      end

    mentions =
      if message.mentions != [] do
        message.mentions
      else
        "[]"
      end

    reactions =
      if message.reactions do
        message.reactions
      else
        "nil"
      end

    webhook_id =
      if message.webhook_id do
        message.webhook_id
      else
        "nil"
      end

    %Embed{}
    |> Embed.title("Message")
    # |> Embed.description("description")
    |> Embed.field("attachments", "#{attachments}")
    |> Embed.field(
      "author",
      "avatar: #{message.author.avatar}\nbot: #{message.author.bot}\ndiscriminator: #{message.author.discriminator}\nemail: #{message.author.email}\nid: #{message.author.id}\nusername: #{message.author.username}\nverified: #{message.author.verified}"
    )
    |> Embed.field("channel_id", "#{message.channel_id}")
    |> Embed.field("content", "#{message.content}")
    |> Embed.field("edited_timestamp", "#{edited_timestamp}")
    |> Embed.field("embeds", "#{embeds}")
    |> Embed.field("id", "#{message.id}")
    |> Embed.field("mention_everyone", "#{message.mention_everyone}")
    |> Embed.field("mention_roles", "#{mention_roles}")
    |> Embed.field("mentions", "#{mentions}")
    |> Embed.field("nonce", "#{message.nonce}")
    |> Embed.field("pinned", "#{message.pinned}")
    |> Embed.field("reactions", "#{reactions}")
    |> Embed.field("timestamp", "#{message.timestamp}")
    |> Embed.field("tts", "#{message.tts}")
    |> Embed.field("webhook_id", "#{webhook_id}")
    |> Embed.color(@default_color)
    |> Embed.send()

    Cogs.say("message")
  end

  Cogs.def embed(_a \\ "") do
    require Alchemy.Embed

    %Embed{}
    |> Embed.title("title")
    |> Embed.url("https://www.ecosia.org/search?q=url")
    |> Embed.author(
      name: "author_name",
      url: "https://www.ecosia.org/search?q=author+name",
      icon_url: "https://essayclick.net/static/img/regular/Famous-American-Authors.jpg"
    )
    |> Embed.thumbnail("https://louisem.com/wp-content/uploads/2017/06/youtube-thumbnail-FB.jpg")
    |> Embed.color(@default_color)
    |> Embed.description("description")
    |> Embed.field("my_field_name", "my_field_value")
    |> Embed.image(
      "http://www.fotos-imagens.net/wp-content/uploads/2011/08/Imagens-Bonitas-quadro-550x371.jpg"
    )
    |> Embed.footer(
      text: "footer_text",
      icon_url: "https://www.wirelesseducation.org/wp-content/uploads/2016/03/footer.png"
    )
    |> Embed.timestamp(DateTime.utc_now())
    |> Embed.send("send_content")

    Cogs.say("Embed")
  end

  Cogs.def guild(_a \\ "") do
    require Alchemy.Embed

    {:ok, guild} = Cogs.guild()

    %Embed{}
    |> Embed.title("Guild")
    |> Embed.color(@default_color)
    |> Embed.field("afk_channel_id", "#{guild.afk_channel_id}")
    |> Embed.field("afk_timeout", "#{guild.afk_timeout}")
    |> Embed.field("channels", "<> []ChannelCategory|TextChannel|VoiceChannel")
    |> Embed.field("default_message_notifications", "#{guild.default_message_notifications}")
    |> Embed.field("embed_enabled", "#{guild.embed_enabled} <end>")
    |> Embed.field("emojis", "#{guild.emojis} <end>")
    |> Embed.field("icon", "#{guild.icon} <end>")
    |> Embed.field("id", "#{guild.id}")
    |> Embed.field("joined_at", "#{guild.joined_at}")
    |> Embed.field("large", "#{guild.large}")
    |> Embed.field("member_count", "#{guild.member_count}")
    |> Embed.field("members", "<> []GuildMember")
    |> Embed.field("mfa_level", "#{guild.mfa_level}")
    |> Embed.field("name", "#{guild.name}")
    |> Embed.field("presences", "#{guild.presences} <end>")
    |> Embed.field("region", "#{guild.region}")
    |> Embed.field("roles", "<> []Role")
    |> Embed.field("splash", "#{guild.splash} <end>")
    |> Embed.field("unavailable", "#{guild.unavailable}")
    |> Embed.field("verification_level", "#{guild.verification_level}")
    |> Embed.field("voice_states", "#{guild.voice_states} <end>")
    |> Embed.send()

    Cogs.say("Guild")
  end

  Cogs.def channel(_a \\ "") do
    require Alchemy.Embed

    {:ok, guild} = Cogs.guild()
    channels = guild.channels

    channels
    |> Enum.filter(&Channel.channelCategory?/1)
    |> Enum.take(1)
    |> Enum.map(fn category ->
      %Embed{}
      |> Embed.title("Channel Category")
      |> Embed.color(@default_color)
      |> Embed.field("guild_id", "#{category.guild_id} <end>")
      |> Embed.field("id", "#{category.id}")
      |> Embed.field("name", "#{category.name}")
      |> Embed.field("nsfw", "#{category.nsfw}")
      |> Embed.field("permission_overwrites", "<> []OverWrite")
      |> Embed.field("position", "#{category.position}")
      |> Embed.send()
    end)

    channels
    |> Enum.filter(&Channel.textChannel?/1)
    |> Enum.take(1)
    |> Enum.map(fn text ->
      %Embed{}
      |> Embed.title("Text Channel")
      |> Embed.color(@default_color)
      |> Embed.field("guild_id", "#{text.guild_id} <end>")
      |> Embed.field("id", "#{text.id}")
      |> Embed.field("last_message_id", "#{text.last_message_id}")
      |> Embed.field("last_pin_timestamp", "#{text.last_pin_timestamp} <end>")
      |> Embed.field("name", "#{text.name}")
      |> Embed.field("nsfw", "#{text.nsfw}")
      |> Embed.field("parent_id", "#{text.parent_id}")
      |> Embed.field("permission_overwrites", "<> []OverWrite")
      |> Embed.field("position", "#{text.position}")
      |> Embed.field("topic", "#{text.topic} <end>")
      |> Embed.send()
    end)

    channels
    |> Enum.filter(&Channel.voiceChannel?/1)
    |> Enum.take(1)
    |> Enum.map(fn voice ->
      %Embed{}
      |> Embed.title("Voice Channel")
      |> Embed.color(@default_color)
      |> Embed.field("bitrate", "#{voice.bitrate}")
      |> Embed.field("guild_id", "#{voice.guild_id} <end>")
      |> Embed.field("id", "#{voice.id}")
      |> Embed.field("name", "#{voice.name}")
      |> Embed.field("nsfw", "#{voice.nsfw} <end>")
      |> Embed.field("parent_id", "#{voice.parent_id}")
      |> Embed.field("permission_overwrites", "<> []OverWrite")
      |> Embed.field("position", "#{voice.position}")
      |> Embed.field("user_limit", "#{voice.user_limit}")
      |> Embed.send()
    end)

    Cogs.say("Channel")
  end
end
