# query-casher
A library for caching using Bloom Filter on any database
![Screenshot from 2023-12-26 00-56-20](https://github.com/seipan/query-cacher/assets/88176012/26f44baf-d45f-432d-b251-1ebcc6d1cad5)
## Usage
```go
package example

import querycacher "github.com/seipan/query-cacher"

type TestDB struct{}

func (db *TestDB) Get(key interface{}) (interface{}, error) {
	return nil, nil
}

func (db *TestDB) Set(key interface{}, value interface{}) error {
	return nil
}

func main() {
	db := querycacher.NewCacher(&TestDB{})
	db.Get("key")
}

```
