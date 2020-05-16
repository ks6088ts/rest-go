# References

* [Go for Visual Studio Code](https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go)
* [cobra](https://github.com/spf13/cobra)
* [gin](https://github.com/gin-gonic/gin)
* [gin/example](https://github.com/gin-gonic/examples)
* [go-gin-example](https://github.com/eddycjy/go-gin-example)
* [Docker で MySQL 8.0.4 を使う](https://qiita.com/yensaki/items/9e453b7320ca2d0461c7)
* [Go + docker で Mysqlを使う(multi-stage builds & docker-composeで)](https://qiita.com/t0w4/items/e886a514559cdb295600)
* [【go】golangでyamlを読み込んでstructに入れるメモ - gopkg.in/yaml.v2](https://www.tweeeety.blog/entry/2017/06/04/231043)
* [【はじめてのGO】gin + gormでシンプルなCRUDなREST APIを作成する](https://qiita.com/daitasu/items/9428220810816972b5d5#read)
* [TeamSQL もいいけど、 DBeaver もいいぞ](https://qiita.com/nanasess/items/609c7cda4adee344221c)

# MySQL

```bash
docker-compose exec mysql mysql -u root -p

# mysql> ---
show databases;
use db;
show tables;
desc products;
insert into products (id, name, description) values \
(1, "yamada", "I am yamada"), \
(2, "okada", "I am okada");
```
