package singleton

// Singleton is a creational design pattern that lets you ensure that a class has only one instance,
// while providing a global access point to this instance.
import "sync"

var (
	instance *db

	// mutual exclusion lock
	lock = &sync.Mutex{}

	once sync.Once
)

type db struct{}

func (d db) GetConnection() string {
	return "connecting"
}

var instanceBody = func() {
	instance = &db{}
}

// GetDBInstanceWithOnce singleton
func GetDBInstanceWithOnce() *db {
	// execution only once
	once.Do(instanceBody)
	return instance
}

// GetDBInstanceWithLock singleton
func GetDBInstanceWithLock() *db {
	if instance == nil {
		// accuire lock for set function with exclusion lock
		// another call process must have to queue
		lock.Lock()
		// unlock when function done
		defer lock.Unlock()
		instanceBody()
	}
	return instance
}
