package proxy

const (
	payServiceName = "pay"
)

var PayProxy *Proxy

func initPayProxy() error {
	var err error
	PayProxy, err = InitServiceByName(payServiceName)
	return err
}
