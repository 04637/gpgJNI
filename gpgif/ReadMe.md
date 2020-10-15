### 动态库编译
> windows:
>
`go build -buildmode=c-shared -o gpg.dll .\encrypt.go`

> linux:
>
`go build -buildmode=c-shared -o gpg.so .\encrypt.go`
