# golang-module
Custom golang module for storing all functions needed for developement


Commands:
```
go mod init github.com/abhishekghoshh/golang-module

git add .
git commit -m <commit-message>
git push /  git push --set-upstream origin master

git tag v0.0.1
git push origin v0.0.1


go get github.com/abhishekghoshh/golang-module


 git describe --tags
 git describe --tags $(git rev-list --tags --max-count=1)
```