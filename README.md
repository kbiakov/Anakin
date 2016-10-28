# Anakin

Codegeneration tool for isomorphic server and mobile Go apps backed by gRPC & Protobuf. Share code between your backend, Android & iOS app!

## Description
<b>Anakin</b> takes care about some routine tasks and helps to create shared code between client (mobile apps) & server app written in Go and backed by gRPC & Protobuf. So, how it works?

At the first stage, it just generates from your ```*.proto``` (with defined service, RPC-calls & allowed messages) main gRPC-file ```*.pb.go``` which may be used by any Go app. It uses ```protoc``` utility.

Next <b>Anakin</b> parses ```*.proto```-file to extract RPC-methods and messages, takes templates for client and server & generates similar Go code. After that it builds binaries for Android (```*.aar```) and iOS (```*.framework```) using ```gomobile```.

Mobile binaries build stage may fail by different reasons and you may want to go back again later, when environment will be ready. Also build stage may be ignored by you if want to make changes in generated code. Anyway <b>Anakin</b> has another ```anakin-build``` script for this purpose.

![Anakin plugin stages](http://i64.tinypic.com/1f4uh.png)

## Required
1. Go >=1.5.<br>

2. Xcode Command Line Tools (Mac OS X only, will be installed if needed).<br>

3. Android SDK (if Android build is needed).<br>

## Usage
```
anakin -P myrpc.proto [-O output_dir] [-h localhost] [-p 50051] [-android] [-ios]
```

Flags:
-P | --proto <proto>	(required) path to ```*.proto```-file
-O | --output <output>	(optional) path to output directory, default: ```/gen```
-h | --host <host>		(optional) server host, default: localhost
-p | --port <port>		(optional) server port, default: 50051
-android 				(optional) is Android build needed, default: false
-ios 					(optional) is iOS build needed, default: false

```
anakin-build [--android 1] [--ios 1]
```

Flags:
--android 1				(optional) is Android build needed, default: 1 (true)
--ios 1 				(optional) is iOS build needed, default: 1 (true)

## Author
### [Kirill Biakov](https://github.com/kbiakov)

## License
```
Copyright (c) 2016 Softwee

![GPL-3.0 License](https://github.com/Softwee/Anakin/blob/master/LICENSE)
```