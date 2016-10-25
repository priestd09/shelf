
# sponge
    import "github.com/coralproject/shelf/internal/sponge"

Package sponge provides support for item importing.






## func Import
``` go
func Import(context interface{}, db *db.DB, graph *cayley.Handle, itm *item.Item) error
```
Import imports an item into the items collections and into the graph database.


## func Remove
``` go
func Remove(context interface{}, db *db.DB, graph *cayley.Handle, itemID string) error
```
Remove removes an item into the items collection and remove any
corresponding quads from the graph database.









- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)