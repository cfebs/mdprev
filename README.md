# mdprev

A markdown preview service.


```
go build mdprev.go
./mdprev ~/mymarkdown/
```

### Endpoints

```
GET /
```
- lists files

```
GET /file/:file_path
```
- renders file

> Note: Makefile is catered to my GOPATH of ~/gopaths/mdprev
