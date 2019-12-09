FROM golang:latest

## 存在しない場合は、ディレクトリ作成される
WORKDIR /go/src/github.com/MasatoTokuse/todo

# WORKDIRからの相対パスの位置に配置する
# src配下が全て /go/src/github.com/MasatoTokuse/todo 以下にコピーされる
COPY ./src .

# 使用している外部ライブラリをダウンロード
RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/go-sql-driver/mysql

# go install で /go/bin 配下に実行ファイルがビルドされ、/go/pkg 配下に実行ファイル以外がビルドされる。
RUN go install github.com/MasatoTokuse/todo/model
RUN go install github.com/MasatoTokuse/todo/controller
RUN go install github.com/MasatoTokuse/todo

# image自体に最初からPATHに/go/binのパスが設定されているため、バイナリファイルを書くだけで動かすことができる
# TODO:ENTRYPOINTである必要はあるか？
ENTRYPOINT [ "todo" ]