

# fixtures
`import "github.com/coralproject/shelf/cmd/corald/fixtures"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [func Error(err error) web.Handler](#Error)
* [func Handler(name string, code int) web.Handler](#Handler)
* [func Load(name string, v interface{}) error](#Load)
* [func NoContent(c *web.Context) error](#NoContent)


#### <a name="pkg-files">Package files</a>
[fixtures.go](/src/github.com/coralproject/shelf/cmd/corald/fixtures/fixtures.go) 





## <a name="Error">func</a> [Error](/src/target/fixtures.go?s=753:786#L24)
``` go
func Error(err error) web.Handler
```
Error will simply return the error to the calling request stack.



## <a name="Handler">func</a> [Handler](/src/target/fixtures.go?s=1122:1169#L39)
``` go
func Handler(name string, code int) web.Handler
```
Handler will serve a JSON payload as the endpoint response.



## <a name="Load">func</a> [Load](/src/target/fixtures.go?s=334:377#L6)
``` go
func Load(name string, v interface{}) error
```
Load unmarshals the specified fixture into the provided
data value.



## <a name="NoContent">func</a> [NoContent](/src/target/fixtures.go?s=967:1003#L33)
``` go
func NoContent(c *web.Context) error
```
NoContent simply responds with a HTTP Status Code of 204.








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
