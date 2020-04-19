# How to Setup Golang

## Installation Golang
まずはじめにgoenvをインストールしたほうが良いでしょう。

## Installation goenv

### for Mac
Homebrewを使ってgoenvをインストールすることができます。簡単です。
```
brew install goenv
```

## Retry Installation Golang 
goenvのインストールができたらGo言語をあらためてインストールします。

### Check what versions golang we can install
インストールできるバージョンとインストール済みのバージョンを確認します。

```
# available list
goenv install -l

# installed list
goenv versions
```

### Install specific version of golang.
特定のバージョンをインストールします。
勉強用であれば基本的に最新のバージョンをインストールするか、テキストと同じ内容のバージョンをインストールするのがいい気がします。

```
# rcファイルにパスを通すコマンドを追記する
vim ~/.zshrc

## 以下の内容を追記
export GOENV_ROOT=$HOME/.goenv
export PATH=$GOENV_ROOT/bin:$PATH
eval "$(goenv init -)"

# 追記した内容を読み込んで反映
source ~/.zshrc

# install 
goenv install 1.11.4

# specify current version
goenv global 1.11.4

# refresh caches
goenv rehash

# check current version
go version
```

### Set GOPATH
Go言語を利用するためには GOPATH という環境変数の設定が必要です。

```
vim ~/.zshprofile

# 以下を追加
GOPATH="$HOME/go"
export GOPATH
```