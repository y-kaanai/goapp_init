# GoApp-Init

<!-- ![Badge Status](https://ci-as-a-service) -->

A initializer for Golang application

## Description

私がGoで新たなアプリケーションを作成する際に、<br>
参考にしている [Standard Go Project Layout](https://qiita.com/vengavengavnega/items/2235589445dd0effda05) に則り、<br>
1コマンドで必要なフォルダを一括で作成することで、<br>
毎回ディレクトリを作成しなければならない煩わしさを解消する。

## Features

- make directories

## Requirement

- Go(1.15.12) or more

## Installation

Paste the following commands at a Terminal prompt.

```shel
$ go install github.com/Kanai-Yuki/goapp_init
go: downloading github.com...
.
.
.
```

## Usage

### 1. Create setting JSON file

```shel
$ goapp gen-json
```

#### 1-2. Add dir-name to goappinit.json if necessary more directories

exmaple

```json
{
    "folders": [
        "build",
        "cmd",
        "configs",
        "developments",
        "docs",
        "internal",
        "pkg",
        "test",
        "tools",
        "gen",    // ← add
        ".github" // ← add
    ],
    "files": [
        "cmd/main.go",
        "Makefile",  // ← add
        "Dockerfile" // ← add
    ]
}
```

### 2. Create need folders

```shel
$ goapp init
```

### 3. See Your Project Folder

```shel
$ ls
```

OK if it looks like the following.

```
comming soon
```

### 4. goapp del json

```shel
goapp del-json
```

## Q&A

### cannot find module providing package ~: working directory is not part of a module

Try it ...

```shel
$ go mod init github.com/{{user_name}}/{{dir_name}}
go: creating new go.mod: module github.com/{{user_name}}/{{dir_name}}
```
