
# galleryfix
    import "github.com/coralproject/shelf/internal/ask/form/gallery/galleryfix"






## func Add
``` go
func Add(context interface{}, db *db.DB, gs []gallery.Gallery) error
```
Add inserts gallerys for testing.


## func Get
``` go
func Get(fixture string) ([]gallery.Gallery, error)
```
Get loads gallery data.


## func Remove
``` go
func Remove(context interface{}, db *db.DB, pattern string) error
```
Remove removes gallerys in Mongo that match a given pattern.









- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)