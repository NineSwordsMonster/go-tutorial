# Go tour

```shell
cd hello
go mod edit -replace github.com/NineSwordsMonster/go-tutorial/greetings=../greetings  
go mod tidy 
```

```shell
cd ../greetings/
go test
go test -v

```