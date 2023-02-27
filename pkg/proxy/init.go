package proxy

func InitProxy() {
	err := initBaseProxy()
	if err != nil {
		panic(err)
	}
	err = initPayProxy()
	if err != nil {
		panic(err)
	}
	err = initStoreProxy()
	if err != nil {
		panic(err)
	}
}
