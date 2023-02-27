package proxy

const (
	baseServiceName = "base"
)

var BaseProxy *Proxy

func initBaseProxy() error {
	var err error
	BaseProxy, err = InitServiceByName(baseServiceName)
	return err
}
