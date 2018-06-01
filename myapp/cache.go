package myapp

// UserCache wraps a UserService to provide an in-memory cache.
type UserCache struct {
	cache   map[int]*User
	service UserService
}

// NewUserCache returns a new read-through cache for service.
func NewUserCache(service UserService) *UserCache {
	return &UserCache{
		cache:   make(map[int]*User),
		service: service,
	}
}

// User returns a user for a given id.
// Returns the cached instance if available.
func (c *UserCache) User(id int) (*User, error) {
	// Check the local cache first.
	if u := c.cache[id]; u != nil {
		return u, nil
	}

	// Otherwise fetch from the underlying service.
	u, err := c.service.User(id)
	if err != nil {
		return nil, err
	} else if u != nil {
		c.cache[id] = u
	}
	return u, err
}
