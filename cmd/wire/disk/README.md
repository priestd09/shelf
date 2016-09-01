

# disk
`import "github.com/coralproject/shelf/cmd/wire/disk"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [func LoadDir(dir string, loader func(string) error) error](#LoadDir)
* [func LoadQuadParams(context interface{}, path string) ([]wire.QuadParams, error)](#LoadQuadParams)


#### <a name="pkg-files">Package files</a>
[disk.go](/src/github.com/coralproject/shelf/cmd/wire/disk/disk.go) 





## <a name="LoadDir">func</a> [LoadDir](/src/target/disk.go?s=926:983#L26)
``` go
func LoadDir(dir string, loader func(string) error) error
```
LoadDir loadsup a given directory, calling a load function for each valid
json file found.



## <a name="LoadQuadParams">func</a> [LoadQuadParams](/src/target/disk.go?s=303:383#L5)
``` go
func LoadQuadParams(context interface{}, path string) ([]wire.QuadParams, error)
```
LoadQuadParams serializes the content of a set of QuadParams from a file using the
given file path. Returns the serialized QuadParams value.








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)