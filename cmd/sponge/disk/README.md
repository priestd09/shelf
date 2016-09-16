

# disk
`import "github.com/coralproject/shelf/cmd/sponge/disk"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [func LoadDir(dir string, loader func(string) error) error](#LoadDir)
* [func LoadItem(context interface{}, path string) (item.Item, error)](#LoadItem)


#### <a name="pkg-files">Package files</a>
[disk.go](/src/github.com/coralproject/shelf/cmd/sponge/disk/disk.go) 





## <a name="LoadDir">func</a> [LoadDir](/src/target/disk.go?s=873:930#L27)
``` go
func LoadDir(dir string, loader func(string) error) error
```
LoadDir loadsup a given directory, calling a load function for each valid
json file found.



## <a name="LoadItem">func</a> [LoadItem](/src/target/disk.go?s=284:350#L5)
``` go
func LoadItem(context interface{}, path string) (item.Item, error)
```
LoadItem serializes the content of Items from a file using the
given file path. Returns the serialized Item value.








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)