# discord-notify-cli

The CLI tool for discord notification

## Usage

```bash
$ echo "# test" | ./notify
Notification sent to Discord successfully.
```

![demo image](demo.png)

## Installation

```bash
$ go version
go version go1.22.3 darwin/arm64

$ go build -o notify notify.go

$ sudo install /usr/local/bin notify

$ echo "export DISCORD_WEBHOOK_URL='https://discord.com/api/webhooks/...'" | sudo tee -a /etc/profile
```
