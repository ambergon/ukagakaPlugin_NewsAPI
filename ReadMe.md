# ukagakaPlugin_NewsAPI
このプラグインはNewsAPIを使い、BBCニュースに網を張るプラグインです。<br>


## 動作環境
SSP 2.6.48で動作確認しています。<br>


## 必要なモノ
NewsAPIの開発用API(無料)が必要です。<br>
[Register - News API](https://newsapi.org/register)<br>


## 注意
開発用APIは現在時刻から24時間以内のニュースは拾えないので、二日前の現在時刻から24時間分のニュースを拾うようにしています。<br>
APIの料金やルールは2023/07/20現在の情報なので、ご自分でご確認ください。<br>
<br>
Go 言語で作成されたDLLなので、プラグインのリロード(free library)を行うとフリーズします。<br>


## 設定
ディレクトリ直下にある、Config.jsonを編集します。<br>
CountはSSPが起動してから何分以降にニュースの読み上げを実行するか指定します。<br>
2だと二分以上経過していてかつ、ニュースの読み込みが終わってから実行されます。<br>
ゴーストの起動トークとのダブり解消用です。<br>
```
{
     "API"          : "YourAPIKey",
     "SearchWord"   : "検索ワードを英語で。" ,
     "Count"        : 2
}
```

SearchWordには[Everything - Documentation - News API](https://newsapi.org/docs/endpoints/everything)で使用できるクエリの書き方が使用できます。
```
//プーチンが含まれていてプリゴジンが含まれない検索
"putin -prigozhin"
//プーチンかトランプが含まれる検索
"trump OR putin"
//英単語を複数指定して検索したい。
"SearchWord"   : "\"3D print\" OR \"3D printer\" OR \"3D printed\"" ,
```



## 表示される内容

- `\0\b[2]`を使用して表示されます。
- クリックで記事リンクを開きます。


## 他
メニューからこのプラグインを呼び出すと設定の再読み込みと記事の再検索を行います。


## 制作動機
まじめな記事はあまり見たくありませんが、特定の記事をいつでも拾える状態を作りたくて制作しました。<br>
24時間以上経過した内容じゃないと拾えない欠点はありますが、必要な内容以外頭に入れないで済みます。<br>


## License
MIT

## Author
ambergon









