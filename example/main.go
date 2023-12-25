package example

import querycacher "github.com/seipan/query-cacher"

func main() {
	db := querycacher.NewCacher(&TestDB{})
	db.Get("key")
}
