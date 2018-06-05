# ImpactBot
A simple bot for the Impact Discord

## What does it do?
It answers to some common questions and provides 2 simple commands

## Running

Run using `npm start`. The `TOKEN` environment variable must be set or the bot will fail to connect (see [Configuration](#Configuration)).

_You sholdn't use `node index.js` because we pass some custom arguments. See `package.json`._

## Configuration

Configuration is managed using environment variables. You can set variables specific to this project in `.env` or you can use your standard shell/os method of configuring environment variables.

Using `.env` has the advantage of keeping your config seperate from your system environment and specific to this project. It is ignored by git, so it is safe to include secrets such as `TOKEN` withoug fear of accidentally committing them.

| Variable   | Description
|:---        |:---
| `CLIENT_ID` | The Discord API Client ID associated with the bot.
| `TOKEN`    | The OAuth token to use for authenticating with Discord's API.
