package mem

import (
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
	"net/http"
)

func init() {
	http.HandleFunc("/mem", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	item := &memcache.Item{
		Key:   "myKey",
		Value: []byte("FirstMemcache"),
	}

	if err := memcache.Add(ctx, item); err == memcache.ErrNotStored {
		fmt.Fprint(w, "Key already exist. myKes is ", string(item.Value))
		return
	} else if err != nil {
		fmt.Fprint(w, "memchache add error", item.Key)
		return
	}

	if item, err := memcache.Get(ctx, "myKey"); err == memcache.ErrCacheMiss {
		fmt.Fprint(w, "memchache not found myKey")
	} else if err != nil {
		fmt.Fprint(w, "memchache get error")
	} else {
		fmt.Fprint(w, "Value of myKey is ", string(item.Value))
	}
}
