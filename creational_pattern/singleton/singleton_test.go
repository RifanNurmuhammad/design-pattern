package singleton

import "testing"

func TestInit(t *testing.T) {
	db := GetDBInstanceWithLock()
	db.GetConnection()
}
