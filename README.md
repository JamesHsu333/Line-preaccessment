# Preassessment for LINE
[![GoDoc](https://godoc.org/github.com/JamesHsu333/Line-preaccessment?status.svg)](https://godoc.org/github.com/JamesHsu333/Line-preaccessment) [![goreportcard.com](https://goreportcard.com/badge/github.com/JamesHsu333/Line-preaccessment)](https://goreportcard.com/report/github.com/JamesHsu333/Line-preaccessment)

Assignment for Line, which promote myself through chatbot.

## How to operate
![QR Code](./static/image/qrcode.png)

1. Scan this QR code or search ```@458pivlz``` to add LINE Friend .

2. Follow the instruction of dialog or richmenu.

3. Input any string to test the different reaction.

## Structure

![process](./static/image/process.png)

## How to develop

First, you can install the project by cloning the repo:

```zsh
git clone https://github.com/JamesHsu333/Line-preaccessment.git
```
Then, create ```.env``` and complete it with the ```CHANNEL_TOKEN``` and ```CHANNEL_SECRET``` from LINE Developers console:


```zsh
CHANNEL_SECRET=YOUR CHANNEL SECRET
CHANNEL_TOKEN=YOUR CHANNEL TOKEN
PORT=PORT YOU PREFER
```

Thus, install the dependencies and start the app:
```zsh
go install .
go run server.go message.go constant.go
```

Use ngrok or proxy to redirect localhost with https, then update the ```Webhook URL``` from LINE Developers console with:
```
https://(YOUR IP ADDRESS FROM ngrok OR proxy)/callback
```

## How to deploy

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://dashboard.heroku.com/new?template=https%3A%2F%2Fgithub.com%2FJamesHsu333%2FLine-preaccessment)

## CI/CD with GitHub/GitLab

![CICD](https://miro.medium.com/max/1400/1*pmCYFW7bPmEaysG2dl27IA.png)

You can use GitHub or GitLab as your git repo

Heroku can automatically build and release (if the build is successful) pushes to your GitHub repo

For more details, please read the [instruction of Heroku](https://devcenter.heroku.com/articles/github-integration)