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

### Download Meson IPFS Core

Download the latest version of Meson IPFS and extract it in the current directory.

https://github.com/daqnext/kubo/releases/

| Filename | SHA256 | 
|---------------------------|------------------------------------------------------------------|
| ipfs-linux-arm.tar.gz     | 25a1903f385698fe692dc61ccf2760b5ce2efafeb456fa50a6ee428948bad4ce | 
| ipfs-linux-arm64.tar.gz   | 93cf7ee632eacdfb97fd56159050105e694e7c17e915482f7e6236db366d0620 | 
| ipfs-linux-386.tar.gz     | b8a56a777173566d5fd2579c1b25c8018c90c565d559a354968dd692d5b3ae8c | 
| ipfs-linux-amd64.tar.gz   | 4875e3fc14d29f484afaca2b03771b7809d73a206456e28a31c3dba346874461 | 
| ipfs-darwin-arm64.tar.gz  | 9ab79feb8c4f8a701cc577c218b9155052cdf861e288607e45526df33ff158c5 |
| ipfs-darwin-amd64.tar.gz  | f4490f51e39c624ef9b6fbde051788c0da503164b3e8858283eed2c2e985d5e4 | 
| ipfs-windows-386.tar.gz   | e3f818f9801fe0b180abf0dfefc5fe9882c464d37279f23849f0a8d6fab72f39 |
| ipfs-windows-amd64.tar.gz | e3fba1d0477bc9e3083d2f70167b0b996d2d179a8f9862c14938ba0e53033a19 |

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

## License
- Apache License, Version 2.0, [LICENSE-APACHE](http://www.apache.org/licenses/LICENSE-2.0)