# tgsend

> 一个 Telegram 推送的小工具，用于调用 Bot API 发送告警等；Bot 以及 Token 请自行查阅相关文章创建。

## 安装

Release 页已经提供了一些预编译版本，下载后增加可执行权限即可使用。

## 使用

```sh
➜  ~ tgsend --help
NAME:
   tgsend - Telegram message send tool

USAGE:
   tgsend [global options] command [command options] [arguments...]

AUTHOR:
   mritd <mritd@linux.com>

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --api value                             Telegram api address (default: "https://api.telegram.org") [$TELEGRAM_ADDRESS]
   --token value                           Telegram api token (default: "") [$TELEGRAM_TOKEN]
   --id value, -i value                    Telegram user or group ID (default: 0) [$TELEGRAM_SEND_ID]
   --message value, -m value               Telegram message to be sent [$TELEGRAM_MESSAGE]
   --file value, -f value                  Telegram file to be sent [$TELEGRAM_FILE]
   --image value, --photo value, -p value  Telegram image to be sent [$TELEGRAM_IMAGE]
   --markdown                              Set the message format to markdown (default: false) [$TELEGRAM_MARKDOWN]
   --help, -h
```

**每个选项皆可通过环境变量设置，环境变量以 `TELEGRAM_` 开头，具体可查看 `--help` 文档**
