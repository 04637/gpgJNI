### 动态库编译
> windows:
>
`go build -buildmode=c-shared -o wingpg.dll ./encrypt.go`

> linux:
>
`go build -buildmode=c-shared -o liblinuxgpg.so ./encrypt.go`
