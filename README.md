# mdprev

A markdown preview service.


```
go build mdprev.go
./mdprev ~/mymarkdown/
```

![image](https://cloud.githubusercontent.com/assets/302375/15802815/526d01da-2a8e-11e6-8ee5-bb851e0023f8.png)


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
