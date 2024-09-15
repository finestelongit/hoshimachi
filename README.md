# Hoshimachi

## ✩°｡ An experimental, Suisei-themed gRPC server, made with Go 💙

![Hoshimachi Suisei](/docs/hoshimachi-cover.png)

## Information
- Runs on localhost or server's IP address
- Runs at port :22318 (Suisei's debut date, 22/3/18)

## How to use
### Method 1: Using <a href="https://github.com/finestelongit/hoshiyomi">Hoshiyomi</a> (recommended)
See the documentation of Hoshiyomi for further explanation.

### Method 2: Using Postman
Maybe you want to view <a href="https://learning.postman.com/docs/sending-requests/grpc/grpc-client-overview/">this</a> link? :o

### Method 3: Build your own gRPC client
I'm assuming that you know the drill. So, just download the `.proto` file, and use it to build the gRPC client! ^o^

I'll recommend to use Hoshiyomi instead as your starting point!

## Setup 

### Normal setup

#### Clone the repository 📃
```bash
git clone https://github.com/finestelongit/hoshimachi
```

#### Build to make the binary file ⚒️
```bash
go build
```

#### And, run the program 📦
```bash
./hoshimachi
```

#### Or, directly run the program 📦
```bash
go run .
```

### (Optional) Compile `.proto` file manually

Coming soon! gRPC is kinda *"spain without the s"* to set up, so as the tutorial.

If you want to use the `.proto` file in other programming languages (let's say, building your own gRPC client to fetch the Hoshimachi server), read https://protobuf.dev/reference/other/




