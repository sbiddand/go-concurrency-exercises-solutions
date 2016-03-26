package main

type MockDB struct{}

// Get only returns an empty string, as this is only for demonstration purposes
func (*MockDB) Get(key string) (string, error) {
	return "", nil
}

func GetMockDB() *MockDB {
	return &MockDB{}
}
