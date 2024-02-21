## Discord Greeting Bot

This is a simple Discord bot that responds with a greeting message when someone sends "hi" in a Discord channel where the bot is present. The bot is written in Go and uses the DiscordGo library for interacting with the Discord API.

### Installation

Clone the repository:

   ```
   git clone https://github.com/MISHRA-TUSHAR/discord-greeting-bot.git
   ```

Navigate to the project directory:

```
cd discord-greeting-bot
```

Install dependencies:
```
go mod tidy
```

### Configuration


The bot requires a config.json file in the root directory to store the bot token and prefix. You can create a config.json file with the following structure:
```
{
  "Token": "YOUR_DISCORD_BOT_TOKEN",
  "BotPrefix": "!"
}
```

***Replace "YOUR_DISCORD_BOT_TOKEN" with your actual Discord bot token.***


### Usage


To start the bot, run the main.go file:
```
go run main.go
```


**Once the bot is running, it will listen for messages in Discord channels. When someone sends "hi", the bot will respond with "Hello, hope you are fine :)".**

