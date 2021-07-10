package xiaomi

import (
	"fmt"
	"testing"
)

func Test_CreateHandphone(t *testing.T) {
	miHandphone := Xiaomi{}
	handphone := miHandphone.CreateHandphone()

	fmt.Println(handphone.Name())
	fmt.Println(handphone.Price())
	fmt.Println(handphone.HasNFC())
	fmt.Println(handphone.HandphoneType())
}

func Test_CreateHeadset(t *testing.T) {
	miHeadset := Xiaomi{}
	headset := miHeadset.CreateHeadset()

	fmt.Println(headset.Price())
	fmt.Println(headset.Color())
	fmt.Println(headset.IsBluetoothType())
}
