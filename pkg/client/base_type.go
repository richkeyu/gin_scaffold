package client

import "time"

type Project struct {
	ID                     int64             `json:"id"`
	Website                string            `json:"website"`
	Names                  map[string]string `json:"names"`
	Webhook                string            `json:"webhook"`
	CreateTime             string            `json:"create_time"`
	Languages              []string          `json:"languages"`
	WebhookSignatureSecret string            `json:"webhook_signature_secret"`
	Icon                   string            `json:"icon"`
	Description            string            `json:"description"`
	WebhookStatus          int               `json:"webhook_status"`
	MerchantID             int64             `json:"merchant_id"`
	DefaultLanguage        string            `json:"default_language"`
	UpdateTime             string            `json:"update_time"`
	Deleted                int               `json:"deleted"`
}

func (p Project) GetName(language string) string {
	name := p.Names[language]
	if len(name) == 0 {
		name = p.Names[p.DefaultLanguage]
	}
	// 如果默认语言没有取第一个配置的语言 与前端逻辑保持一致
	if len(name) == 0 && len(p.Languages) > 0 {
		name = p.Names[p.Languages[0]]
	}
	// 如果还没有取第一个有值的
	if len(name) == 0 {
		for _, s := range p.Names {
			if len(s) > 0 {
				name = s
				break
			}
		}
	}
	return name
}

type GetProjectById struct {
	Data Project `json:"data"`
}

type GetProjects struct {
	Data []Project `json:"data"`
}

type Merchant struct {
	Id              int64     `json:"id,omitempty" form:"id"`
	Name            string    `json:"name,omitempty" form:"name"`
	Country         string    `json:"country,omitempty" form:"country"`
	Industry        string    `json:"industry,omitempty" form:"industry"`
	Size            string    `json:"size,omitempty" form:"size"`
	PostalCode      string    `json:"postal_code,omitempty" form:"postal_code"`
	PostalAddress   string    `json:"postal_address,omitempty" form:"postal_address"`
	Website         string    `json:"website,omitempty" form:"website"`
	PayFee          float64   `json:"pay_fee,omitempty" form:"im_30_pay_fee"`
	DefaultCurrency string    `json:"default_currency,omitempty" form:"default_currency"`
	Status          int       `json:"status,omitempty" form:"status"`
	WithTax         bool      `json:"with_tax,omitempty" form:"with_tax"`
	Formal          int       `json:"formal,omitempty" form:"formal"`
	ApiKey          string    `json:"api_key,omitempty" form:"api_key"`
	CreateTime      time.Time `json:"create_time" form:"create_time"`
	UpdateTime      time.Time `json:"update_time" form:"update_time"`
}

type GetMerchantResponse struct {
	Merchant
}

type GetMerchantAllResponse struct {
	Data []Merchant `json:"data"`
}

type GetUser struct {
	User
}

type User struct {
	IsSuperuser     int    `json:"is_superuser,omitempty"`
	CreateTime      string `json:"create_time,omitempty"`
	Level           int    `json:"level,omitempty"`
	MerchantID      int64  `json:"merchant_id,"`
	Login           int    `json:"login,omitempty"`
	FinishInit      int    `json:"finish_init,omitempty"`
	Password        string `json:"password,omitempty"`
	UpdateTime      string `json:"update_time,omitempty"`
	Phone           string `json:"phone,omitempty"`
	EmailVerifyCode string `json:"email_verify_code,omitempty"`
	Name            string `json:"name,omitempty"`
	EmailVerifyExp  int    `json:"email_verify_exp,omitempty"`
	ID              int    `json:"id,omitempty"`
	Email           string `json:"email,omitempty"`
}

type CurrencyAll struct {
	List map[string]CurrencyItem `json:"data"`
}
type CurrencyItem struct {
	CurrencyCode   string `json:"currency_code"`
	CurrencySymbol string `json:"currency_symbol"`
	CurrencyName   string `json:"currency_name"`
}

type CountriesAll []CountriesItem

type CountriesItem struct {
	Label        string            `json:"label"` // name
	Value        string            `json:"value"` // iso2
	Translations map[string]string `json:"translations"`
}
