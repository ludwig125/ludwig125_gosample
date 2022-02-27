参考
https://docs.google.com/presentation/d/10ru3LdbofJqgdmD8pprZuZyWbGvOFC8rKxb6q5Q46Xc/mobilepresent?slide=id.gcf4887a11e_0_2

実行方法

```
$ls -l
-rw-r--r-- 1 ludwig125 ludwig125      62  4月 27 07:00 a.s
-rw-r--r-- 1 ludwig125 ludwig125      26  4月 27 06:58 main.go
```

以下のファイルがある状態で、`go build .`


```
$ cat a.s
TEXT    main·main(SB),0,$0
        MOVQ    $231, AX
        MOVQ    $1, DI
        SYSCALL
```

```
$cat main.go
package main

func main()
```

実行ファイルはディレクトリ名が自動で割り当てられる
`[~/go/src/github.com/ludwig125/ludwig125_gosample/assembly/helloWorld]`

```
$./helloWorld
$echo $?
1
```
`MOVQ    $1, DI`となっているので、終了ステータスが１になる
