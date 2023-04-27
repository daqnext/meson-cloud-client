# ☁️ Meson Cloud Client

![](/static/nyancat.svg)

## What is Meson Cloud Client?

Meson Cloud Client is a local client for Gateway X that helps users make better use of their local unlimited storage space. Currently, it perfectly supports IPFS and Filecoin, and it won't be long before the client also supports Arweave.

### Files and folders tree

```
├── .gitignore
├── LICENSE
├── README.md
├── api
│   ├── api.go
│   ├── api_test.go
│   └── request.go
├── config.templates.yml
├── config.yml
├── daemon
│   ├── ipfs.go
│   └── ipfs_test.go
├── go.mod
├── go.sum
├── ipfs
├── logger
│   └── log.go
├── main.go
└── static
```

## Getting Started

### Download Meson Cloud Client

Download the latest Meson Cloud Client for use.

| Filename | SHA256 | 
|---------------------------|------------------------------------------------------------------|
| mcloud-client-linux-arm.tar.gz     | 19e3482e301dcfe25456b80a1fd7d54201068ad956f8a87f35dde48c258f221d | 
| mcloud-client-linux-arm64.tar.gz   | 5c2916650ea1365159f5deceb1d4574dd3b0d4b579243483bfe72adf217d26e6 | 
| mcloud-client-linux-386.tar.gz     | 9eb17e78f49f7a87f291e9785b3fb67b5535ed1b3b07b6a35a615ecfc79d2f4e | 
| mcloud-client-linux-amd64.tar.gz   | 44cf16df823ccf969c6cb1cbfe8543fcbb8efdbb8418af3487c5fce120bfe277 | 
| mcloud-client-darwin-arm64.tar.gz  | e33b8f57de7395ab185f7202186d244be5d6ed7e0613e4045a0823f0ac40bfbd |
| mcloud-client-darwin-amd64.tar.gz  | 90b033d60d5c1fbee551a67b5500e70b858740f793b2573cfade745cf5e992cf | 
| mcloud-client-windows-386.tar.gz   | a6b4116c64e76841bdd0f3b22a659b1e005ca4fc9f5ef47a8201bf834afd68d3 |
| mcloud-client-windows-amd64.tar.gz | 9a85f4d6233bd963ef8f08004f5b3abfef880689ea628e651d63364e71b1e058 |

You can also update Meson Cloud Client by downloading the latest version of [Meson IPFS](https://github.com/daqnext/kubo/releases/) and extracting it to the current directory.

### Configure `config.yml`

Simply obtain the exclusive Meson platform token from the [Meson Dashboard](https://dashboard.meson.network/user/account), insert it into the `token` field in `config.yml`, and save it.

```
version: 1.0

token: "your token"
queryUrl: "http://127.0.0.1:3535"
logLevel: "dev"

ipfs:
  ipfsCmd: "./ipfs"
  ipfsDataRoot: "~/.ipfs"
```

### Run Client

Go language is required to run. If it is not available, please download and install Go language from the following website: https://go.dev/dl/.

```
go run main.go
```

After running the command, you can click on the `Upload to Local` button or open the link http://127.0.0.1:5001/webui/ to access it.

```
http://127.0.0.1:5001/webui/
```

### Develop
- [Dev Reference](DEVELOP.md)

## License
- Apache License, Version 2.0, [LICENSE-APACHE](http://www.apache.org/licenses/LICENSE-2.0)
