# mdprev

A markdown preview service.


```
go build mdprev.go
./mdprev ~/mymarkdown/
```

![image](https://cloud.githubusercontent.com/assets/302375/15802829/9537f3a8-2a8e-11e6-99dc-0331529b3134.png)


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
