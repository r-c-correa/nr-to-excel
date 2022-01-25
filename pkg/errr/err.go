package errr

func PanicIfIsNotNull(err error) {
	if err != nil {
		panic(err)
	}
}
