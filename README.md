# RoProxy Lite
A modified version of RoProxy made for self-hosting.

Setup is easy, simply deploy with the button below and configure environment variables. The KEY variable is optional, leave blank to not require an auth key.

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/halffalse/roproxy-lite)
[![Deploy on Railway](https://railway.app/button.svg)](https://railway.app/new/template?template=https%3A%2F%2Fgithub.com%2Fhalffalse%2Froproxy-lite&envs=KEY%2CTIMEOUT%2CRETRIES&optionalEnvs=KEY&KEYDesc=The+key+used+to+access+proxy.&TIMEOUTDesc=The+number+of+seconds+before+a+request+times+out+and+is+retried.&RETRIESDesc=The+maximum+number+of+times+to+attempt+a+request.+Minimum%3A+1&TIMEOUTDefault=5&RETRIESDefault=5)
[![Deploy to Render](https://render.com/images/deploy-to-render-button.svg)](https://render.com/deploy)

When "KEY" environment variable is populated, a matching "PROXYKEY" header must be present. Requests must be made in the format /subdomain/path. E.g. https://games.roblox.com/docs -> https://roproxytest.heroku.com/games/docs
