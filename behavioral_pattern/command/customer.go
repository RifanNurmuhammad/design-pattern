package command

func Execute() {
	jne := JNE{}

	shipper1 := Shipper{}

	nextDayDelivery := ShippingNextDayCommand{
		Courier: jne,
		Params: ShippingNextDayParams{
			Box:              2,
			ExtraBubbleWraps: true,
		},
	}
	shipper1.AddShipping(nextDayDelivery)

	regularDelivery := ShippingRegularCommand{
		Courier: jne,
		Params: ShippingRegularParams{
			Kargo: true,
		},
	}
	shipper1.AddShipping(regularDelivery)
	shipper1.BulkShipping()
}
