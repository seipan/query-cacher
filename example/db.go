package example

type TestDB struct{}

func (db *TestDB) Get(key interface{}) (interface{}, error) {
	return nil, nil
}

func (db *TestDB) Set(key interface{}, value interface{}) error {
	return nil
}
