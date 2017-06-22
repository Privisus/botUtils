# Configuration

You'll need to configure these environment variables to ensure successful loading: `DARKSKY_TOKEN`, `LATLONGITUDE`, `CHANNEL_SECRET`, `CHANNEL_TOKEN`, `PORT`, `MOVIE_TOKEN`.

`DARKSKY_TOKEN` can be obtained from darksky.net/dev/.  
`LATLONGITUDE` is `"LATITUDE,LONGITUDE"`.  
`MOVIE_TOKEN` can be obtained from themoviedb.org/.  
`PORT` is without the `:` character.

All environment variables are in string.

You'll also need to install ImageMagick on your server.

---

# Setting Up

Initialize first by doing `botUtils.Initialize(randomNumberAsASeed)`, then you can use `botUtils.Call.SetReplyToken(replyToken)` and `handler.HandleCommand(message.Text)` accordingly.

You have to open a fileserver of the `downloaded` directory.

That's it.

---

# Commands

1. `!<a question>?` -> ball8
2. `!coin` -> flip a coin
3. `!dice` -> dice
4. `=<some math equation>` -> latex
5. `!film <film name>` -> movies
6. `!qod` -> quote of the day
7. `!weather` -> weather

An alternative word is also available (check commandHandler).

---

# Contributing

Don't hesitate to ask a question regarding contributing. I'll try to answer them if I can.
