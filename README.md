# GoApp-Init

<!-- ![Badge Status](https://ci-as-a-service) -->

A initializer for Golang application

## Description

私がGoで新たなアプリケーションを作成する際に、参考にしている [Standard Go Project Layout](https://qiita.com/vengavengavnega/items/2235589445dd0effda05) に則り、  
1コマンドで必要なフォルダを一括で作成することで、毎回ディレクトリを作成しなければならない煩わしさを解消する。

## Features

- Awesome function
- Awesome UI
- ...

For more information, see `awesome-tool --help`.

## Requirement

- Go(1.15.12) or more

## Installation

Paste the following commands at a Terminal prompt.

```shel
$ go get github.com/Kanai-Yuki/goapp_init
```

## Usage

### 1. get cli tool

```shel
$ go get github.com/Kanai-Yuki/goapp_init
```

### 2. Create setting JSON file

```shel
$ goapp gen-json
```

### 3. Create need folders

```shel
$ goapp init
```

### 4. See Your Project Folder

```shel
$ ls
```

OK if it looks like the following.

```
comming soon
```

### 5. goapp del json

```shel
goapp del-json
```

## Author

comming soon ...

## Q&A

### cannot find module providing package ~: working directory is not part of a module

Try it ...

```shel
$ go mod init
go: creating new go.mod: module github.com/user_name/dir_name
```
