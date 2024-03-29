<p align="center">
  <img src="./web/ene/ene.png" width="350"/>
</p>
<h1 align="center">Ene</h1>
<p align="center">A go bot framework</p>

### Summary
Ene is a twitch bot that functions as plugin framework for making fun chatbot plugins

### Features
Ene is divided into three parts:
* `/cmd/`: The core bot engine
* `/internal/adapters/`: Service Input and Output
* `/internal/plugins/`: dynamic logic for handling input

### running

```bash
# create your starter config
cp ./.waifu.example.toml ./.waifu
# edit your config file to add necessary details
vim ./.waifu
# run the bot
go run ./main.go  
```

### Goals

* adapters
  - [x] cli
  - [ ] twitch
  - [ ] openai
  - [ ] discord
  - [ ] opera.gx
  - [ ] twitter
  - [ ] tts
* plugins
  - [ ] trivia
  - [x] spam
  - [ ] quotes
  - [ ] tarot
  - [ ] gpt
  - [ ] mahjong
  - [ ] chinese dictionary
* lib 
  - [ ] database
  - [ ] chatgpt
  - [ ] emotes
    - [ ] ffz
    - [ ] bttv
    - [ ] 7tv