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
| [mcloud-client-linux-arm.tar.gz](https://github.com/daqnext/meson-cloud-client/releases/download/v0.01/mcloud-client-linux-arm.tar.gz)     | dc67282c7251126cdd88064757a6c2c8c0250995c44136f2c65eb71049d9f807 | 
| [mcloud-client-linux-arm64.tar.gz](https://github.com/daqnext/meson-cloud-client/releases/download/v0.01/mcloud-client-linux-arm64.tar.gz)   | 0c56bff09e8323702fbe68519a0650b99783363b179e3bb01bec27d4253839a5 | 
| [mcloud-client-linux-386.tar.gz](https://github.com/daqnext/meson-cloud-client/releases/download/v0.01/mcloud-client-linux-386.tar.gz)     | ab95169c6a8202163eaa084f8f74a7acbd248ad05ae1e6f178e08d690745bdc3 | 
| [mcloud-client-linux-amd64.tar.gz](https://github.com/daqnext/meson-cloud-client/releases/download/v0.01/mcloud-client-linux-amd64.tar.gz)   | 86083b9133df984798ff3f04b2fd45739fc65cf67c691b62a485c0a8773cae60 | 
| [mcloud-client-darwin-arm64.tar.gz](https://github.com/daqnext/meson-cloud-client/releases/download/v0.01/mcloud-darwin-arm64.tar.gz)  | bd455456ebe0f3d758a4982bf6282ed7560ee581bfbb3214157aab845006d462 |
| [mcloud-client-darwin-amd64.tar.gz](https://github.com/daqnext/meson-cloud-client/releases/download/v0.01/mcloud-darwin-amd64.tar.gz)  | 49d6505f4d5678fd91bc2bb3fc3d0c3b5d29283f2829b28376582c624467065b | 
| [mcloud-client-windows-386.tar.gz](https://github.com/daqnext/meson-cloud-client/releases/download/v0.01/mcloud-windows-386.tar.gz)   | 0c5d3e23fa664a0a5aaaa0df521e1e051c48b52ce6c0e39945d7291f218ad1e0 |
| [mcloud-client-windows-amd64.tar.gz](https://github.com/daqnext/meson-cloud-client/releases/download/v0.01/mcloud-windows-amd64.tar.gz) | ab5d0f4bdd65b35e3fc20776b4f19a1cae6af95b0e2c237ecca2f77bae2a7d17 |

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

## License
- Apache License, Version 2.0, [LICENSE-APACHE](http://www.apache.org/licenses/LICENSE-2.0)