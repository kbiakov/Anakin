# Anakin [![GoDoc](https://godoc.org/github.com/Softwee/Anakin?status.svg)](https://godoc.org/github.com/Softwee/Anakin) [![Go Report Card](https://goreportcard.com/badge/github.com/Softwee/Anakin)](https://goreportcard.com/report/github.com/Softwee/Anakin) [![Maintainability](https://api.codeclimate.com/v1/badges/9f7e630f0ab78a73ccca/maintainability)](https://codeclimate.com/github/kbiakov/Anakin/maintainability) [![Build Status](https://travis-ci.org/Softwee/Anakin.svg?branch=master)](https://travis-ci.org/Softwee/Anakin) [![Android Arsenal](https://img.shields.io/badge/Android%20Arsenal-Anakin-blue.svg?style=true)](https://android-arsenal.com/details/1/4625)

Codegeneration tool for isomorphic server and mobile Go apps with gRPC & Protobuf. Share code between your backend, Android & iOS app!

## Description
<b>Anakin</b> takes care about some routine tasks and helps to create shared code between client (mobile apps) & server app written in Go and backed by gRPC & Protobuf. So, how it works?

At the first stage, it just generates from your ```*.proto``` (with defined service, RPC-calls & allowed messages) main gRPC-file ```*.pb.go``` which may be used by any Go app. It uses ```protoc``` utility.

Next <b>Anakin</b> parses ```*.proto```-file to extract RPC-methods and messages, takes templates for client and server & generates similar Go code. After that it builds binaries for Android (```*.aar```) and iOS (```*.framework```) using ```gomobile```.

Mobile binaries build stage may fail by different reasons and you may want to go back again later when environment will be ready. Also you can ignore build stage if you want to make some changes in generated code. Anyway <b>Anakin</b> has another ```anakin-build``` script inside for this purpose, which automatically copied for generated ```$YOUR_OUTPUT/client``` directory with other source files when you run original ```anakin``` script.

![Anakin plugin stages](http://i64.tinypic.com/1f4uh.png)

## Required
1. Go 1.5 and higher.<br>

2. Xcode Command Line Tools (Mac OS X only, will be installed if needed).<br>

3. Android SDK (if Android build is needed).<br>

## Usage
```
anakin -P myrpc.proto [-O output_dir] [-h localhost] [-p 50051] [-android] [-ios]
```

<b>Flags</b>:<br>
```-P | --proto <proto>``` *(required)* path to ```*.proto```-file<br>
```-O | --output <output>``` *(optional)* path to output directory, default: ```/gen```<br>
```-h | --host <host>``` *(optional)* server host, default: localhost<br>
```-p | --port <port>``` *(optional)* server port, default: 50051<br>
```-android``` *(optional)* is Android build needed, default: false<br>
```-ios``` *(optional)* is iOS build needed, default: false<br>

```
anakin-build [--android 1] [--ios 1]
```

<b>Flags</b>:<br>
```--android 1``` *(optional)* is Android build needed, default: 1 (true)<br>
```--ios 1``` *(optional)* is iOS build needed, default: 1 (true)<br>

## Author
### [Kirill Biakov](https://github.com/kbiakov)

## License
```
                     GNU GENERAL PUBLIC LICENSE
                       Version 3, 29 June 2007

 Copyright (C) 2007 Free Software Foundation, Inc. <http://fsf.org/>
 Everyone is permitted to copy and distribute verbatim copies
 of this license document, but changing it is not allowed.
```
