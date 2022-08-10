# llvmcalc

goyaccとllvmを使って計算機作る。

## 実行方法

```bash
# yacc実行
$ make yacc

# 計算機をビルド、llvmcalcができる
$ make build-go


# `1 + 2` を計算するプログラムを a.ll に出力
$ ./llvmcalc '1 + 2'

# llvmでコンパイル＆リンク（M1 MacOSのみ対応）
$ make llc
$ make lld

# 実行すると3が変えてくる
$ ./a.out; echo $?
3
```
