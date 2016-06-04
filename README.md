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
