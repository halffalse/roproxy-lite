# roproxy-lite
A modified version of RoProxy made for self-hosting.

Setup is easy, simply change the options at the top of main.go and run. Alternatively, deploy to Heroku with one click:

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/halffalse/roproxy-lite)

Use with HTTPService as you normally would, just adding your proxy before the roblox url. E.g. https://roproxytest.herokuapp.com/https://games.roblox.com/

Note: This isn't actually restricted to roblox domains and requires no sort of authentication. If this is something you'd like to see, let me know or feel free to implement it yourself.
