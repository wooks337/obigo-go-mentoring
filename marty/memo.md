# 정리 

## (Call by Value) vs (Call by Rerence)
- https://deveric.tistory.com/92  // Java 에서의 개념 (golang과 비교)


## Golang OOP
- https://golangkorea.github.io/post/go-start/object-oriented/


## go-callvis : go 패키지 시각화 
- https://github.com/ofabry/go-callvis
- https://codesk.tistory.com/120


### UML - goplantuml
- 설치
```
go get github.com/jfeliu007/goplantuml/parser 
go get github.com/jfeliu007/goplantuml/cmd/goplantuml
cd $GOPATH/src/github.com/jfeliu007/goplantuml 
go install ./...
```
- puml 생성 
```
goplantuml [-recursive] path/to/gofiles path/to/gofiles2 > outfile.puml
```




## Goland plugin 
- Theme
    - Atom material icons
    - One Dark them
- CSV (종류 많음 취향대로~)
- Git
    - Emoji commit log viewer   : git log 에서 이모지 아이콘 적용
    - Git Commit Template :  git conventional  템플릿 적용
    - Gitmoji Plus : Commit Button   : git emoji
    - GitToolBox    :  기본적인 git 기능 편리하게 인터페이스 적용
- Rainbow Brackets  :  {}, (), 개행 구분 플러그인 
- Toml  : toml 환경변수 invalid 체크
- Progress : 다양한 progress UI 제공 (취향대로 선택)