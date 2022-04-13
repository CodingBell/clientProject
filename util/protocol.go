package util

// BootNotificationRequest return the necessary bytes slices
func BootNotificationRequest() []byte {
	str := []string{
		"5431363431373335323131",
		"4A4F59534F4EFFFFFFFF",
		"4A4F59534F4EFFFFFFFF",
		"76312E302E31FFFF",
		"3839383630335959584D48484858585858585850",
		"303132333435363738393031323334",
		"00",
		"7F000001",
		"31323334353637383930",
	}

	result := GetBytes(str)

	result = addHeadAndTail(0x01, result)
	return result
}
