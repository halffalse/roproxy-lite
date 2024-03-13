# RoProxy Lite
A modified version of RoProxy made for self-hosting.

Setup is easy, simply deploy with the button below and configure environment variables. The KEY variable is optional, leave blank to not require an auth key.

[![Docker](https://img.shields.io/badge/docker-2496ED.svg?style=for-the-badge&logo=docker&logoColor=white)](https://hub.docker.com/rocord01/roproxy-docker)
[![Deploy](https://img.shields.io/badge/heroku-%23430098.svg?style=for-the-badge&logo=heroku&logoColor=white)](https://heroku.com/deploy?template=https://github.com/halffalse/roproxy-lite)
[![Deploy on Railway](https://img.shields.io/badge/Railway-0B0D0E.svg?style=for-the-badge&logo=railway&logoColor=white)](https://railway.app/new/template/fV9Lxm)
[![Deploy to Render](https://img.shields.io/badge/Render-8B05FF.svg?style=for-the-badge&logo=render&logoColor=white)](https://render.com/deploy)

When "KEY" environment variable is populated, a matching "PROXYKEY" header must be present. Requests must be made in the format /subdomain/path. E.g. https://games.roblox.com/docs -> https://roproxytest.heroku.com/games/docs
