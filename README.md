![logo](https://socialify.git.ci/zkweb3/go-montgomery-reduce/image?description=1&language=1&name=1&pattern=Floating%20Cogs&theme=Light)

### Building
```bash
docker build . -t montgomery:latest --no-cache
```
### Testing
```bash
go test -timeout 5s -run . go-montgomery-reduce -v -count=1
```