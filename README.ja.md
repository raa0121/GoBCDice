# GoBCDice

[![Build Status](https://travis-ci.org/raa0121/GoBCDice.svg?branch=master)](https://travis-ci.org/raa0121/GoBCDice)
[![Build status](https://ci.appveyor.com/api/projects/status/4gl47493rao9t4b8/branch/master?svg=true)](https://ci.appveyor.com/project/raa0121/gobcdice/branch/master)
[![codecov](https://codecov.io/gh/raa0121/GoBCDice/branch/master/graph/badge.svg)](https://codecov.io/gh/raa0121/GoBCDice)

GoBCDiceは、多くのゲームシステムに対応するTRPG向けダイスボット「[ボーンズ＆カーズ（BCDice）](https://github.com/bcdice/BCDice)」のGoによる実装です。中核となるダイスローラー（ダイス表記の構文解析器および評価器）および多くのゲームシステム固有ダイスボットで構成されます。

## 使い方

ビルド要件：[Go](https://golang.org/dl/) &ge; 1.12

現在は動作確認のためのREPLのみビルド、実行できます。

```bash
cd cmd/GoBCDiceREPL

# REPLをビルドする
go build

# REPLを実行する
./GoBCDiceREPL
```

ダイス表記の構文解析器（pkg/core/parser/parser.go）を修正し、ビルドするには、さらに[modernc.org/goyacc](https://godoc.org/modernc.org/goyacc)をインストールする必要があります：

```bash
GO111MODULE=off go get -u modernc.org/goyacc
```

そのうえで、以下のコマンドによってビルドします：

```bash
cd pkg/core/parser
make
```

## ダイスローラー

GoBCDiceのダイスローラーは、一般的なダイスロール機能を提供します。

BCDiceは以下のダイス表記に対応しています（詳細については、GitHub上の[bcdice/BCDice/docs/README.txt](https://github.com/bcdice/BCDice/tree/master/docs)を参照してください）。現在GoBCDiceが対応しているダイス表記にはチェックがついています。

* [x] 加算ロール（D）：`xDn`
    * [x] 成功判定つき：`xDn>=y` など
* [x] バラバラロール（B）：`nBx`
    * [ ] 成功判定つき：`xBn>=y` など
* [ ] 個数振り足しロール（R）：`xRn>=y` など
* [ ] 上方無限ロール（U）：`xUn[t]`
    * [ ] 成功判定つき：`xUn[t]>=y` など

x：ダイス数、n：ダイスの面数、y：目標値、t：振り足しの閾値

追加の構文は以下のとおりです：

* [x] ランダム数値埋め込み：`[最小値...最大値]`
* [ ] シークレットロール：`SxDn` など

ダイスローラーは以下のコマンドにも対応しています：

* [x] 計算（四則演算、C）：`C(1+2-3*4/5)` など
* [ ] ランダム選択：`CHOICE[A,B,C]` など

### 演算子

ダイスロールや計算で使用可能な演算子を示します。

#### 算術演算子

算術演算では、数値は整数として扱われます。

* 単項演算子
    * 単項プラス（何もしない）`+`
    * 単項マイナス（符号反転）`-`
* 二項演算子
    * 加算 `+`
    * 減算 `-`
    * 乗算 `*`
    * 除算
        * 端数を切り捨てる除算 `/`
        * 端数を四捨五入する除算 `/R`
        * 端数を切り上げる除算 `/U`

#### 比較演算子

比較演算子は、成功判定において使用します。

* 等価 `=`
* 非等価 `<>`
* 未満 `<`
* 以下 `<=`
* 超過 `>`
* 以上 `>=`

## 作者

[raa0121](https://twitter.com/raa0121)

移植元のBCDiceは、[Faceless氏](https://twitter.com/Faceless192x)および[たいたい竹流氏](https://twitter.com/torgtaitai)によって製作されました。
