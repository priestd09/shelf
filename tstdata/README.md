

# tstdata
`import "github.com/coralproject/shelf/tstdata"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func Docs() ([]map[string]interface{}, error)](#Docs)
* [func Drop(db *db.DB)](#Drop)
* [func Generate(db *db.DB) error](#Generate)


#### <a name="pkg-files">Package files</a>
[tstdata.go](/src/github.com/coralproject/shelf/tstdata/tstdata.go) 


## <a name="pkg-constants">Constants</a>
``` go
const CollectionExecTest = "test_xenia_data"
```
CollectionExecTest contains the name of the collection that is
going to store the xenia test data.




## <a name="Docs">func</a> [Docs](/src/target/tstdata.go?s=624:669#L18)
``` go
func Docs() ([]map[string]interface{}, error)
```
Docs reads the fixture and returns the documents.



## <a name="Drop">func</a> [Drop](/src/target/tstdata.go?s=1607:1627#L69)
``` go
func Drop(db *db.DB)
```
Drop drops the temp collection.



## <a name="Generate">func</a> [Generate](/src/target/tstdata.go?s=1163:1193#L45)
``` go
func Generate(db *db.DB) error
```
Generate creates a temp collection with data
that can be used for testing things.








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
