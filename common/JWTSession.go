package common

type JWTSession struct {
	ExpireTime    int64                  `json:"expireTime"`
	Account       string                 `json:"account"`
	Role          string                 `json:"role"`
	AccountStatus string                 `json:"accountStatus"`
	AccountType   string                 `json:"accountType"`
	AccountID     string                 `json:"accountID"`
	AccountName   string                 `json:"accountName"`
	Token         string                 `json:"token"`
	Key           string                 `json:"key"`
	IsAdmin       bool                   `json:"isAdmin"`
	IsSuperAdmin  bool                   `json:"isSuperAdmin"`
	IsLogin       bool                   `json:"isLogin"`
	IsLogout      bool                   `json:"isLogout"`
	LoginTime     int64                  `json:"loginTime"`
	LoginIP       string                 `json:"loginIP"`
	LoginType     string                 `json:"loginType"`
	LoginDevice   string                 `json:"loginDevice"`
	LoginDeviceID string                 `json:"loginDeviceID"`
	OtherInfo     map[string]interface{} `json:"otherInfo"`
	LoginID       string                 `json:"loginID"`
}
