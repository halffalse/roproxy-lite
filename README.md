# roproxy-lite
A modified version of RoProxy made for self-hosting.

Setup is easy, simply deploy with the button below and configure environment variables. The KEY variable is optional, leave blank to not require an auth key.

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/halffalse/roproxy-lite)

When "KEY" environment variable is populated, a matching "PROXYKEY" header must be present. Use with HTTPService as you normally would, simply adding your proxy before the roblox url. E.g. https://roproxytest.herokuapp.com/https://games.roblox.com/

Note: This isn't actually restricted to roblox domains. If this is something you'd like to see, let me know or feel free to implement it yourself.
