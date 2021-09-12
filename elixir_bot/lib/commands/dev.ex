defmodule ElixirBot.Commands.Dev do
  use Alchemy.Cogs
  alias Alchemy.{Embed}

  @default_color 0x410056

  Cogs.group("dev")

  Cogs.set_parser(:tt, &([String.split(&1)]))
  Cogs.def tt(rest \\ "") do
    alias Alchemy.Client
    # Client.create_channel(message.gui
  end

  Cogs.def message(_rest \\ "") do
    require Alchemy.Embed

    attachments =
    if message.attachments != [] do message.attachments else "[]" end

    edited_timestamp =
    if message.edited_timestamp do message.edited_timestamp else "nil" end

    embeds =
    if message.embeds != [] do message.embeds else "[]" end

    mention_roles =
    if message.mention_roles != [] do message.mention_roles else "[]" end

    mentions =
    if message.mentions != [] do message.mentions else "[]" end

    reactions =
    if message.reactions do message.reactions else "nil" end

    webhook_id =
    if message.webhook_id do message.webhook_id else "nil" end

    %Embed{}
    |> Embed.title("Message")
    # |> Embed.description("description")
    |> Embed.field("attachments", "#{attachments}")
    |> Embed.field("author", "avatar: #{message.author.avatar}\nbot: #{message.author.bot}\ndiscriminator: #{message.author.discriminator}\nemail: #{message.author.email}\nid: #{message.author.id}\nusername: #{message.author.username}\nverified: #{message.author.verified}")
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
    |> Embed.send

    Cogs.say "message"
  end

  Cogs.def embed(_a \\ "") do
    require Alchemy.Embed
    %Embed{}
    |> Embed.title("title")
    |> Embed.url("https://www.ecosia.org/search?q=url")
    |> Embed.author(
      name: "author_name",
      url: "https://www.ecosia.org/search?q=author+name",
      icon_url: "https://essayclick.net/static/img/regular/Famous-American-Authors.jpg")
    |> Embed.thumbnail("https://louisem.com/wp-content/uploads/2017/06/youtube-thumbnail-FB.jpg")
    |> Embed.color(@default_color)
    |> Embed.description("description")
    |> Embed.field("my_field_name", "my_field_value")
    |> Embed.image("http://www.fotos-imagens.net/wp-content/uploads/2011/08/Imagens-Bonitas-quadro-550x371.jpg")
    |> Embed.footer(
      text: "footer_text",
      icon_url: "https://www.wirelesseducation.org/wp-content/uploads/2016/03/footer.png")
    |> Embed.timestamp(DateTime.utc_now())
    |> Embed.send("send_content")

    Cogs.say "Embed"
  end

end
