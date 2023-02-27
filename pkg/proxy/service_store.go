package proxy

const (
	storeServiceName = "store"
)

var StoreProxy *Proxy

func initStoreProxy() error {
	var err error
	StoreProxy, err = InitServiceByName(storeServiceName)
	return err
}
