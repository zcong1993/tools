# tools [![Go Report Card](https://goreportcard.com/badge/github.com/zcong1993/tools)](https://goreportcard.com/report/github.com/zcong1993/tools) [![CircleCI branch](https://img.shields.io/circleci/project/github/zcong1993/tools/master.svg)](https://circleci.com/gh/zcong1993/tools/tree/master)

> simple tools by golang and aws lambda

## Usage

### dns resolver
```sh
$ curl https://tools.zcong.moe/dns?d=google.com
# cli
$ tools dns google.com
```

### ulid generator
```sh
$ curl https://tools.zcong.moe/ulid?n=10
# cli
$ tools ulid 100
```

### md5 hash
```sh
$ curl https://tools.zcong.moe/md5?s=string4md5
# cli
$ tools md5 string4md5
```

### qrcode generator

[https://tools.zcong.moe/qrcode?s=https://google.com](https://tools.zcong.moe/qrcode?s=https://google.com)

## Deploy

with awesome [now](https://zeit.co/now)

```bash
$ now zcong1993/tools -e GIN_MODE=release --alias your.custom.domain.com
```

## License

MIT &copy; zcong1993
