package querycacher

import (
	"strconv"

	"github.com/bits-and-blooms/bloom/v3"
)

type Handler interface {
	Get(any) (any, error)
	Set(any, any) error
}

// Cacher sets the content in the Bloom Filter before any Set or Get operations on any cache or database,
// thereby preventing unnecessary access to the database in advance.
type Cacher struct {
	Handler

	// use : https://github.com/bits-and-blooms/bloom
	cache *bloom.BloomFilter
}

// Bloom filter capable of receiving 1 million elements with a false-positive rate of 1%.
func NewCacher(handler Handler) *Cacher {
	return &Cacher{
		Handler: handler,
		cache:   bloom.NewWithEstimates(1000000, 0.01),
	}
}

func (c *Cacher) Get(key any) (any, error) {
	switch key.(type) {
	case string:
		if !c.cache.TestString(key.(string)) {
			return nil, nil
		}
		res, err := c.Get(key)
		if err != nil {
			return nil, err
		}
		return res, nil
	case []byte:
		if !c.cache.Test(key.([]byte)) {
			return nil, nil
		}
		res, err := c.Get(key)
		if err != nil {
			return nil, err
		}
		return res, nil
	case int:
		if !c.cache.TestString(strconv.Itoa(key.(int))) {
			return nil, nil
		}
		res, err := c.Get(key)
		if err != nil {
			return nil, err
		}
		return res, nil
	case int64:
		if !c.cache.TestString(strconv.FormatInt(key.(int64), 10)) {
			return nil, nil
		}
		res, err := c.Get(key)
		if err != nil {
			return nil, err
		}
		return res, nil
	case uint64:
		if !c.cache.TestString(strconv.FormatUint(key.(uint64), 10)) {
			return nil, nil
		}
		res, err := c.Get(key)
		if err != nil {
			return nil, err
		}
		return res, nil
	case float64:
		if !c.cache.TestString(strconv.FormatFloat(key.(float64), 'f', -1, 64)) {
			return nil, nil
		}
		res, err := c.Get(key)
		if err != nil {
			return nil, err
		}
		return res, nil
	case float32:
		if !c.cache.TestString(strconv.FormatFloat(float64(key.(float32)), 'f', -1, 32)) {
			return nil, nil
		}
		res, err := c.Get(key)
		if err != nil {
			return nil, err
		}
		return res, nil
	default:
		return c.Get(key)
	}
}

func (c *Cacher) Set(key any, value any) error {
	switch key.(type) {
	case string:
		c.cache.AddString(key.(string))
		return c.Set(key, value)
	case []byte:
		c.cache.Add(key.([]byte))
		return c.Set(key, value)
	case int:
		c.cache.AddString(strconv.Itoa(key.(int)))
		return c.Set(key, value)
	case int64:
		c.cache.AddString(strconv.FormatInt(key.(int64), 10))
		return c.Set(key, value)
	case uint64:
		c.cache.AddString(strconv.FormatUint(key.(uint64), 10))
		return c.Set(key, value)
	case float64:
		c.cache.AddString(strconv.FormatFloat(key.(float64), 'f', -1, 64))
		return c.Set(key, value)
	case float32:
		c.cache.AddString(strconv.FormatFloat(float64(key.(float32)), 'f', -1, 32))
		return c.Set(key, value)
	default:
		return c.Set(key, value)
	}
}
