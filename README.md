# tools [![Go Report Card](https://goreportcard.com/badge/github.com/zcong1993/tools)](https://goreportcard.com/report/github.com/zcong1993/tools) [![CircleCI branch](https://img.shields.io/circleci/project/github/zcong1993/tools/master.svg)](https://circleci.com/gh/zcong1993/tools/tree/master)

> simple tools by golang and aws lambda

## Usage

### dns resolver
```sh
$ curl https://f48s21atqi.execute-api.ap-southeast-1.amazonaws.com/staging/dns?d=google.com
# cli
$ tools dns google.com
```

### ulid generator
```sh
$ curl https://f48s21atqi.execute-api.ap-southeast-1.amazonaws.com/staging/ulid?n=10
# cli
$ tools ulid 100
```

### md5 hash
```sh
$ curl https://f48s21atqi.execute-api.ap-southeast-1.amazonaws.com/staging/md5?s=string4md5
# cli
$ tools md5 string4md5
```

### qrcode generator

[https://f48s21atqi.execute-api.ap-southeast-1.amazonaws.com/staging/qrcode?s=https://google.com](https://f48s21atqi.execute-api.ap-southeast-1.amazonaws.com/staging/qrcode?s=https://google.com)

## License

MIT &copy; zcong1993
