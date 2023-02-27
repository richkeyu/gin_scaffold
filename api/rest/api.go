package rest

import (
	"net/http"

	"gateway/api/rest/hello"
	"gateway/pkg/middleware/auth"
	"gateway/pkg/proxy"
)

func registerRoute() {
	route.GET("/hello", hello.Greeter)
	route.GET("/base/hello", proxy.BaseProxy.Handler(http.MethodGet, "/v1/hello"))
	route.GET("/pay/hello", proxy.PayProxy.Handler(http.MethodGet, "/v1/hello"))
	route.GET("/store/hello", proxy.StoreProxy.Handler(http.MethodGet, "/v1/hello"))

	// 商户后台
	merchant := route.Group("/merchant/")
	merchant.GET("transactions/status-pairs", proxy.PayProxy.Handler(http.MethodGet, "/v1/transactions/status-pairs"))
	merchant.Use(middleware.MerchantAuthMiddleWare())
	merchant.GET("pay/payment-intents", proxy.PayProxy.Handler(http.MethodGet, "/v1/transactions"))
	merchant.GET("pay/payment-intents/export", proxy.PayProxy.Handler(http.MethodGet, "/v1/transactions/export"))
	merchant.GET("pay/payment-intents/get-export-status", proxy.PayProxy.Handler(http.MethodGet, "/v1/transactions/get-export-status"))
	merchant.GET("pay/payment-intent/:pi_id", proxy.PayProxy.Handler(http.MethodGet, "/v1/transaction/:pi_id"))
	merchant.POST("pay/payment-intents/refund", proxy.PayProxy.Handler(http.MethodPost, "/v1/transaction/merchant-refund"))
	merchant.POST("pay/date-line", proxy.PayProxy.Handler(http.MethodPost, "/v1/transaction/date-line"))
	merchant.POST("bi/merchant", proxy.BaseProxy.Handler(http.MethodPost, "/v1/bi/merchant"))
	merchant.GET("bi/conversion/rate", proxy.BaseProxy.Handler(http.MethodGet, "/v1/bi/conversion/rate"))
	merchant.GET("payment-method/country-pairs", proxy.PayProxy.Handler(http.MethodGet, "/v1/payment-method/country-pairs"))
	merchant.GET("payment-method/type-pairs", proxy.PayProxy.Handler(http.MethodGet, "/v1/payment-method/type-pairs"))
	merchant.GET("payment-method/project-list", proxy.PayProxy.Handler(http.MethodGet, "/v1/payment-method/project-list"))
	merchant.GET("payment-method/project-export", proxy.PayProxy.Handler(http.MethodGet, "/v1/payment-method/project-export"))
	merchant.POST("payment-method/project-update", proxy.PayProxy.Handler(http.MethodPost, "/v1/payment-method/project-update"))
	merchant.POST("payment-method/project-batch-update", proxy.PayProxy.Handler(http.MethodPost, "/v1/payment-method/project-batch-update"))

	//virtual
	merchant.POST("virtual-currency/create", proxy.StoreProxy.Handler(http.MethodPost, "/v1/virtual-currency/create"))
	merchant.POST("virtual-currency/update", proxy.StoreProxy.Handler(http.MethodPost, "/v1/virtual-currency/update"))
	merchant.POST("virtual-currency/delete", proxy.StoreProxy.Handler(http.MethodPost, "/v1/virtual-currency/delete"))
	merchant.GET("virtual-currency/get", proxy.StoreProxy.Handler(http.MethodGet, "/v1/virtual-currency/get"))
	merchant.POST("tool/image/upload", proxy.StoreProxy.Handler(http.MethodPost, "/v1/tool/image/upload")) //公共上传

	siteGroup := merchant.Group("site/")
	siteGroup.GET("language/export", proxy.StoreProxy.Handler(http.MethodGet, "/v1/site/language/export"))
	siteGroup.POST("language/import", proxy.StoreProxy.Handler(http.MethodPost, "/v1/site/language/import"))

	siteBuilder := merchant.Group("site/builder/")
	siteBuilder.GET("desc/export", proxy.StoreProxy.Handler(http.MethodGet, "/v1/site/builder/desc/export"))
	siteBuilder.POST("desc/import", proxy.StoreProxy.Handler(http.MethodPost, "/v1/site/builder/desc/import"))
	siteBuilder.POST("edit-base-info", proxy.StoreProxy.Handler(http.MethodPost, "/v1/site/builder/edit-base-info")) // 修改商城基础信息
	siteBuilder.POST("edit-user-info", proxy.StoreProxy.Handler(http.MethodPost, "/v1/site/builder/edit-user-info")) // 修改商城用户配置信息
	siteBuilder.GET("info", proxy.StoreProxy.Handler(http.MethodGet, "/v1/site/builder/info"))                       // 获取商城详细信息
	siteBuilder.POST("image/add", proxy.StoreProxy.Handler(http.MethodPost, "/v1/site/build/image/add"))             // 建站器素材添加
	siteBuilder.DELETE("image/del", proxy.StoreProxy.Handler(http.MethodDelete, "/v1/site/build/image/del"))         //建站器素材删除
	siteBuilder.GET("image/list", proxy.StoreProxy.Handler(http.MethodGet, "/v1/site/build/image/list"))             //建站器素材列表
	siteBuilder.GET("default-conf", proxy.StoreProxy.Handler(http.MethodGet, "/v1/site/builder/default-conf"))
	siteBuilder.GET("list", proxy.StoreProxy.Handler(http.MethodGet, "/v1/site/builder/list"))
	siteBuilder.POST("add", proxy.StoreProxy.Handler(http.MethodPost, "/v1/site/builder/add"))
	siteBuilder.POST("edit", proxy.StoreProxy.Handler(http.MethodPost, "/v1/site/builder/edit"))
	siteBuilder.POST("copy", proxy.StoreProxy.Handler(http.MethodPost, "/v1/site/builder/copy"))
	siteBuilder.POST("delete", proxy.StoreProxy.Handler(http.MethodPost, "/v1/site/builder/delete"))
	siteBuilder.POST("publish", proxy.StoreProxy.Handler(http.MethodPost, "/v1/site/builder/publish"))
	siteBuilder.GET("domain-check", proxy.StoreProxy.Handler(http.MethodGet, "/v1/site/builder/domain-check"))

	//语言
	siteBuilder.GET("language/list", proxy.StoreProxy.Handler(http.MethodGet, "/v1/site/builder/language/list"))
	siteBuilder.POST("language/delete", proxy.StoreProxy.Handler(http.MethodPost, "/v1/site/builder/language/delete"))
	siteBuilder.POST("language/change-default", proxy.StoreProxy.Handler(http.MethodPost, "/v1/site/builder/language/change-default"))
	siteBuilder.POST("language/edit", proxy.StoreProxy.Handler(http.MethodPost, "/v1/site/builder/language/edit"))

	//self-domain
	siteBuilder.GET("self-domain/list", proxy.StoreProxy.Handler(http.MethodGet, "/v1/site/builder/self-domain/list"))
	siteBuilder.POST("self-domain/delete", proxy.StoreProxy.Handler(http.MethodPost, "/v1/site/builder/self-domain/delete"))
	siteBuilder.POST("self-domain/add", proxy.StoreProxy.Handler(http.MethodPost, "/v1/site/builder/self-domain/add"))

	//页面
	siteBuilder.GET("page/list", proxy.StoreProxy.Handler(http.MethodGet, "/v1/site/builder/page/list"))
	siteBuilder.POST("page/add", proxy.StoreProxy.Handler(http.MethodPost, "/v1/site/builder/page/add"))
	siteBuilder.POST("page/edit", proxy.StoreProxy.Handler(http.MethodPost, "/v1/site/builder/page/edit"))
	siteBuilder.POST("page/delete", proxy.StoreProxy.Handler(http.MethodPost, "/v1/site/builder/page/delete"))

	siteBuilder.GET("desc/info", proxy.StoreProxy.Handler(http.MethodGet, "/v1/site/builder/desc/info"))
	siteBuilder.POST("desc/edit", proxy.StoreProxy.Handler(http.MethodPost, "/v1/site/builder/desc/edit"))
	siteBuilder.POST("language/publish", proxy.StoreProxy.Handler(http.MethodPost, "/v1/site/builder/language/publish"))

	merchant.GET("order/list", proxy.StoreProxy.Handler(http.MethodGet, "/v1/order/list"))
	merchant.GET("order/export", proxy.StoreProxy.Handler(http.MethodGet, "/v1/order/export"))
	merchant.GET("order/export-status", proxy.StoreProxy.Handler(http.MethodGet, "/v1/order/export-status"))
	merchant.GET("order/detail", proxy.StoreProxy.Handler(http.MethodGet, "/v1/order/detail"))
	merchant.POST("order/refund", proxy.StoreProxy.Handler(http.MethodPost, "/v1/order/refund"))

	merchant.POST("virtual-group/create", proxy.StoreProxy.Handler(http.MethodPost, "/v1/virtual-group/create"))
	merchant.POST("virtual-group/update", proxy.StoreProxy.Handler(http.MethodPost, "/v1/virtual-group/update"))
	merchant.POST("virtual-group/delete", proxy.StoreProxy.Handler(http.MethodPost, "/v1/virtual-group/delete"))
	merchant.GET("virtual-group/get", proxy.StoreProxy.Handler(http.MethodGet, "/v1/virtual-group/get"))

	merchant.POST("virtual-item/create", proxy.StoreProxy.Handler(http.MethodPost, "/v1/virtual-item/create"))
	merchant.POST("virtual-item/update", proxy.StoreProxy.Handler(http.MethodPost, "/v1/virtual-item/update"))
	merchant.POST("virtual-item/delete", proxy.StoreProxy.Handler(http.MethodPost, "/v1/virtual-item/delete"))
	merchant.GET("virtual-item/get", proxy.StoreProxy.Handler(http.MethodGet, "/v1/virtual-item/get"))
	merchant.GET("virtual-item/detail", proxy.StoreProxy.Handler(http.MethodGet, "/v1/virtual-item/detail"))

	merchant.POST("price-template/create", proxy.StoreProxy.Handler(http.MethodPost, "/v1/price-template/create"))
	merchant.POST("price-template/update", proxy.StoreProxy.Handler(http.MethodPost, "/v1/price-template/update"))
	merchant.POST("price-template/update-many", proxy.StoreProxy.Handler(http.MethodPost, "/v1/price-template/update-many"))
	merchant.POST("price-template/delete", proxy.StoreProxy.Handler(http.MethodPost, "/v1/price-template/delete"))
	merchant.GET("price-template/get", proxy.StoreProxy.Handler(http.MethodGet, "/v1/price-template/get"))
	merchant.GET("price-template/info", proxy.StoreProxy.Handler(http.MethodGet, "/v1/price-template/detail"))
	merchant.GET("price-template/config/country2currency", proxy.StoreProxy.Handler(http.MethodGet, "/v1/price-template/config/country2currency"))

	merchant.POST("price-template/associate-item/associate", proxy.StoreProxy.Handler(http.MethodPost, "/v1/price-template/associate-item/associate"))
	merchant.POST("price-template/associate-item/disassociate", proxy.StoreProxy.Handler(http.MethodPost, "/v1/price-template/associate-item/disassociate"))
	merchant.GET("price-template/associate-item/get", proxy.StoreProxy.Handler(http.MethodGet, "/v1/price-template/associate-item/get"))
	merchant.POST("site/builder/image/upload", proxy.StoreProxy.Handler(http.MethodPost, "/v1/tool/image/upload"))
	merchant.GET("sale/biz/add/credit/details", proxy.StoreProxy.Handler(http.MethodGet, "/v1/sale/biz/add/credit/details"))

	merchant.GET("exchange-rate/get-currencies", proxy.StoreProxy.Handler(http.MethodGet, "/v1/exchange-rate/get-currencies"))
	merchant.POST("exchange-rate/update", proxy.StoreProxy.Handler(http.MethodPost, "/v1/exchange-rate/update"))
	merchant.POST("exchange-rate/exec-task", proxy.StoreProxy.Handler(http.MethodPost, "/v1/exchange-rate/exec-task"))

	merchant.GET("statement/merchant-list", proxy.PayProxy.Handler(http.MethodGet, "/v1/statement/merchant-list"))
	merchant.GET("statement/export-detail", proxy.PayProxy.Handler(http.MethodGet, "/v1/statement/export-detail"))

	merchant.GET("customer-group/options", proxy.BaseProxy.Handler(http.MethodGet, "/v1/customer-group/options"))
	merchant.GET("customer-group/list", proxy.BaseProxy.Handler(http.MethodGet, "/v1/customer-group/list"))
	merchant.POST("customer-group/create", proxy.BaseProxy.Handler(http.MethodPost, "/v1/customer-group/create"))
	merchant.POST("customer-group/update", proxy.BaseProxy.Handler(http.MethodPost, "/v1/customer-group/update"))
	merchant.GET("customer-group/delete", proxy.BaseProxy.Handler(http.MethodGet, "/v1/customer-group/delete"))
	merchant.GET("customer-group/items", proxy.BaseProxy.Handler(http.MethodGet, "/v1/customer-group/items"))
	merchant.GET("customer-group/export", proxy.BaseProxy.Handler(http.MethodGet, "/v1/customer-group/export"))
	merchant.POST("customer-group/items-delete", proxy.BaseProxy.Handler(http.MethodPost, "/v1/customer-group/items-delete"))
	merchant.POST("customer-group/items-add", proxy.BaseProxy.Handler(http.MethodPost, "/v1/customer-group/items-add"))
	merchant.GET("customer-group/list-pairs", proxy.BaseProxy.Handler(http.MethodGet, "/v1/customer-group/list-pairs"))
	merchant.GET("customer-group/risk-items", proxy.BaseProxy.Handler(http.MethodGet, "/v1/customer-group/risk-items"))
	merchant.POST("customer-group/items-add-one", proxy.BaseProxy.Handler(http.MethodPost, "/v1/customer-group/items-add-one"))
	merchant.POST("customer-group/items-update-one", proxy.BaseProxy.Handler(http.MethodPost, "/v1/customer-group/items-update-one"))
	merchant.POST("customer-group/items-delete-one", proxy.BaseProxy.Handler(http.MethodPost, "/v1/customer-group/items-delete-one"))

	// admin运营后台
	admin := route.Group("/admin/")
	admin.Use(middleware.IpAuthMiddleware())
	admin.GET("pay/payment-intents", proxy.PayProxy.Handler(http.MethodGet, "/v1/transactions-admin"))
	admin.GET("pay/payment-intents/export", proxy.PayProxy.Handler(http.MethodGet, "/v1/transactions-admin/export"))
	admin.GET("pay/payment-intents/detail", proxy.PayProxy.Handler(http.MethodGet, "/v1/transactions-admin/detail"))
	admin.GET("pay/payment-intents/get-export-status", proxy.PayProxy.Handler(http.MethodGet, "/v1/transactions/get-export-status"))
	admin.GET("projects/list", proxy.BaseProxy.Handler(http.MethodGet, "/v1/projects/list"))
	admin.GET("site/builder/pairs", proxy.StoreProxy.Handler(http.MethodGet, "/v1/site/build/pairs"))
	admin.GET("merchant/list", proxy.BaseProxy.Handler(http.MethodGet, "/v1/merchant/list"))
	admin.POST("tool/image/upload", proxy.StoreProxy.Handler(http.MethodPost, "/v1/tool/image/upload")) //公共上传

	admin.GET("sale/biz/list", proxy.StoreProxy.Handler(http.MethodGet, "/v1/admin/sale/biz/list"))
	admin.POST("sale/biz/create", proxy.StoreProxy.Handler(http.MethodPost, "/v1/admin/sale/biz/create"))
	admin.POST("sale/biz/update", proxy.StoreProxy.Handler(http.MethodPost, "/v1/admin/sale/biz/update"))
	admin.POST("sale/biz/status/edit", proxy.StoreProxy.Handler(http.MethodPost, "/v1/admin/sale/biz/status/edit"))
	admin.GET("sale/filter/goods", proxy.StoreProxy.Handler(http.MethodGet, "/v1/admin/sale/filter/goods"))

	admin.GET("account/diff-reasons", proxy.PayProxy.Handler(http.MethodGet, "/v1/account/diff-reasons"))
	admin.GET("account/third-platforms", proxy.PayProxy.Handler(http.MethodGet, "/v1/account/third-platforms"))
	admin.POST("account/third-bills/upload", proxy.PayProxy.Handler(http.MethodPost, "/v1/account/third-bills/upload"))
	admin.GET("account/third-bills/upload-status", proxy.PayProxy.Handler(http.MethodGet, "/v1/account/third-bills/upload-status"))
	admin.POST("account/third-final-amount/save", proxy.PayProxy.Handler(http.MethodPost, "/v1/account/third-final-amount/save"))
	admin.GET("account/reconciliations", proxy.PayProxy.Handler(http.MethodGet, "/v1/account/reconciliations"))
	admin.GET("account/reconciliations-status", proxy.PayProxy.Handler(http.MethodGet, "/v1/account/reconciliations-status"))

	admin.GET("statement/list", proxy.PayProxy.Handler(http.MethodGet, "/v1/statement/list"))
	admin.GET("statement/export-detail", proxy.PayProxy.Handler(http.MethodGet, "/v1/statement/export-detail"))
	admin.POST("statement/edit", proxy.PayProxy.Handler(http.MethodPost, "/v1/statement/edit"))

	admin.GET("payment-method/country-pairs", proxy.PayProxy.Handler(http.MethodGet, "/v1/payment-method/country-pairs"))
	admin.GET("payment-method/type-pairs", proxy.PayProxy.Handler(http.MethodGet, "/v1/payment-method/type-pairs"))
	admin.GET("payment-method/merchant-list", proxy.PayProxy.Handler(http.MethodGet, "/v1/payment-method/merchant-list"))
	admin.GET("payment-method/merchant-export", proxy.PayProxy.Handler(http.MethodGet, "/v1/payment-method/merchant-export"))
	admin.POST("payment-method/merchant-update", proxy.PayProxy.Handler(http.MethodPost, "/v1/payment-method/merchant-update"))

	// 无需验证的基础接口
	base := route.Group("/base/")
	// i18n
	base.GET("category/tran", proxy.BaseProxy.Handler(http.MethodGet, "/v1/i18n/category/tran")) // category language tran
	base.GET("get/tran", proxy.BaseProxy.Handler(http.MethodGet, "/v1/i18n/get/tran"))
	base.POST("batch/tran", proxy.BaseProxy.Handler(http.MethodPost, "/v1/i18n/batch/tran"))
	// region
	base.GET("countries/complete/all", proxy.BaseProxy.Handler(http.MethodGet, "/v1/countries/complete/all"))
	base.GET("countries/all", proxy.BaseProxy.Handler(http.MethodGet, "/v1/countries/all"))
	base.GET("countries/currency/pairs", proxy.BaseProxy.Handler(http.MethodGet, "/v1/countries/currency/pairs"))
	base.GET("states/provinces", proxy.BaseProxy.Handler(http.MethodGet, "/v1/states/provinces"))
	base.GET("city/cities", proxy.BaseProxy.Handler(http.MethodGet, "/v1/city/cities"))
	// language
	base.GET("language/all", proxy.BaseProxy.Handler(http.MethodGet, "/v1/language/all"))
	// 状态
	base.GET("trans/status/pairs", proxy.PayProxy.Handler(http.MethodGet, "/v1/transactions/status/pairs"))

	// 旧交易流程接口 后期可废弃
	oldPay := route.Group("/")
	oldPay.POST("/transaction/create-payment-ui", proxy.PayProxy.Handler(http.MethodPost, "/v1/transaction/create-payment-ui"))
	oldPay.GET("/transaction/payment/token", proxy.PayProxy.Handler(http.MethodGet, "/v1/transaction/payment/token"))
	oldPay.POST("/transaction/payment-ui-token/inner", proxy.PayProxy.Handler(http.MethodPost, "/v1/transaction/payment-ui-token/inner"))
	oldPay.POST("/transaction/pay", proxy.PayProxy.Handler(http.MethodPost, "/v1/transaction/pay"))
	oldPay.PATCH("/transaction/payment/token", proxy.PayProxy.Handler(http.MethodPost, "/v1/transaction/payment/information/save"))
	oldPay.GET("/transaction/payment/page/config", proxy.PayProxy.Handler(http.MethodGet, "/v1/transaction/payment/page/config"))
	oldPay.GET("/transaction/payment/cards", proxy.PayProxy.Handler(http.MethodGet, "/v1/transaction/payment/get/card"))
	oldPay.DELETE("/transaction/payment/cards", proxy.PayProxy.Handler(http.MethodPost, "/v1/transaction/payment/delete/card"))

	// 支付前端 新支付流程接口 通过接口参数进行验证
	pay := route.Group("/pay")
	pay.GET("/trans/info/get", proxy.PayProxy.Handler(http.MethodGet, "/v1/transaction/payment/info/get"))                    // 查询付款详情接口
	pay.POST("/trans/fail", proxy.PayProxy.Handler(http.MethodPost, "/v1/transaction/payment/fail"))                          // 将交易状态更新为失败接口
	pay.GET("/trans/config", proxy.PayProxy.Handler(http.MethodGet, "/v1/transaction/payment/config"))                        // 查询付款配置接口
	pay.GET("/trans/cards", proxy.PayProxy.Handler(http.MethodGet, "/v1/transaction/payment/get/card"))                       // 查询card接口
	pay.DELETE("/trans/cards", proxy.PayProxy.Handler(http.MethodPost, "/v1/transaction/payment/delete/card"))                // 删除card接口
	pay.POST("/trans/payment/info/save", proxy.PayProxy.Handler(http.MethodPost, "/v1/transaction/payment/information/save")) // 保存付款信息接口
	pay.POST("/trans/payment/confirm", proxy.PayProxy.Handler(http.MethodPost, "/v1/transaction/pay"))                        // dlocal确认付款接口
	pay.POST("/trans/payment/token", proxy.PayProxy.Handler(http.MethodPost, "/v1/transaction/platform/token"))               // 平台token接口
	pay.POST("/customers/password/code", proxy.PayProxy.Handler(http.MethodPost, "/v1/customers/password/code"))              // 忘记密码，发送邮箱验证码
	pay.POST("/customers/password/change", proxy.PayProxy.Handler(http.MethodPost, "/v1/customers/password/change"))          // 重置密码

	// 回调
	callback := route.Group("/sdk-callback")
	callback.POST("/xsolla/webhook", proxy.PayProxy.Handler(http.MethodPost, "/v1/callback/xsolla/webhook"))                       //xsolla
	callback.POST("/airwallex/webhook/live", proxy.PayProxy.Handler(http.MethodPost, "/v1/callback/airwallex/webhook/live"))       //airwallex-live
	callback.POST("/airwallex/webhook/sandbox", proxy.PayProxy.Handler(http.MethodPost, "/v1/callback/airwallex/webhook/sandbox")) //airwallex-sandbox
	callback.POST("/dlocal/payment", proxy.PayProxy.Handler(http.MethodPost, "/v1/notify/dlocal/payment"))
	callback.POST("/dlocal/refund", proxy.PayProxy.Handler(http.MethodPost, "/v1/notify/dlocal/refund"))
	//callback.POST("/dlocal/chargeback", proxy.PayProxy.Handler(http.MethodPost, "/v1/notify/dlocal/chargeback"))
	callback.POST("/mycard/feedback", proxy.PayProxy.Handler(http.MethodPost, "/v1/notify/mycard/feedback"))
	callback.POST("/mycard/notify", proxy.PayProxy.Handler(http.MethodPost, "/v1/notify/mycard/notify"))
	callback.POST("/mycard/bill", proxy.PayProxy.Handler(http.MethodPost, "/v1/notify/mycard/bill"))

	callback.POST("/wechatpay/payment", proxy.PayProxy.Handler(http.MethodPost, "/v1/notify/wechatpay/payment"))
	callback.POST("/wechatpay/refund", proxy.PayProxy.Handler(http.MethodPost, "/v1/notify/wechatpay/refund"))
	callback.POST("/alipay/payment", proxy.PayProxy.Handler(http.MethodPost, "/v1/callback/alipay/webhook"))

	callback.POST("/paypal/payment", proxy.PayProxy.Handler(http.MethodPost, "/v1/notify/paypal/payment"))

	//外部其他平台调用
	external := route.Group("/external")
	external.POST("/transaction/create-payment-ui", proxy.PayProxy.Handler(http.MethodPost, "/v1/external/transaction/create-payment-ui"))
	external.GET("/store/get-cost", proxy.StoreProxy.Handler(http.MethodGet, "/v1/external/virtual-currency/get-cost"))
	external.POST("/price-template/compute-price", proxy.StoreProxy.Handler(http.MethodPost, "/v1/external/price-template/compute-price"))

	// 商城公共接口 无需登录信息
	storePublic := route.Group("/store")
	storePublic.POST("/register", proxy.StoreProxy.Handler(http.MethodPost, "/v1/register"))
	storePublic.POST("/login", proxy.StoreProxy.Handler(http.MethodPost, "/v1/login"))
	storePublic.POST("/login/player", proxy.StoreProxy.Handler(http.MethodPost, "/v1/login/player"))
	storePublic.GET("/player/:projectId/:playerId", proxy.StoreProxy.Handler(http.MethodGet, "/v1/player/:projectId/:playerId"))
	storePublic.GET("/virtual-group/get", proxy.StoreProxy.Handler(http.MethodGet, "/v1/virtual-group/get"))

	// 商城未登录及登录都可使用的接口 登录后才进行登录用户信息解析
	storeNoAuth := route.Group("/store")
	storeNoAuth.Use(middleware.StoreWithOutAuthMiddleWare())
	storeNoAuth.GET("/virtual-item/get", proxy.StoreProxy.Handler(http.MethodGet, "/v1/virtual-item/list-store"))
	storeNoAuth.GET("/virtual-item/detail", proxy.StoreProxy.Handler(http.MethodGet, "/v1/virtual-item/detail"))
	storeNoAuth.GET("/site/builder/page/list", proxy.StoreProxy.Handler(http.MethodGet, "/v1/site/builder/page/list"))
	storeNoAuth.GET("site/builder/desc/info", proxy.StoreProxy.Handler(http.MethodGet, "/v1/site/builder/desc/info"))
	storeNoAuth.GET("/site/builder/info", proxy.StoreProxy.Handler(http.MethodGet, "/v1/site/builder/info"))
	storeNoAuth.GET("/site/builder/language/list", proxy.StoreProxy.Handler(http.MethodGet, "/v1/site/builder/language/list"))
	storeNoAuth.GET("/tool/country-currency", proxy.StoreProxy.Handler(http.MethodGet, "/v1/tool/country-currency"))
	storeNoAuth.POST("/point/reporting", proxy.StoreProxy.Handler(http.MethodPost, "/v1/data/point/reporting")) // 数据上报接口
	storeNoAuth.GET("/sale/biz/add/credit/details", proxy.StoreProxy.Handler(http.MethodGet, "/v1/sale/biz/add/credit/details"))

	// 商城需登录接口
	store := route.Group("/store")
	store.Use(middleware.StoreAuthMiddleWare())
	store.POST("/order/create", proxy.StoreProxy.Handler(http.MethodPost, "/v1/order/create"))
	store.POST("/order/gift/create", proxy.StoreProxy.Handler(http.MethodPost, "/v1/order/gift/create"))
	store.GET("/order/info", proxy.StoreProxy.Handler(http.MethodGet, "/v1/order/info"))
	store.GET("/customer/info", proxy.StoreProxy.Handler(http.MethodGet, "/v1/customer/info"))
	store.POST("/customer/update", proxy.StoreProxy.Handler(http.MethodPost, "/v1/customer/update"))
	store.POST("/customer/bind/player", proxy.StoreProxy.Handler(http.MethodPost, "/v1/customer/bind/player"))
}
