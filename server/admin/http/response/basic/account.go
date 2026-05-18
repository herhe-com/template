package basic

type ToAccountOfInformation struct {
	Nickname string `json:"nickname"`
	Username string `json:"username,omitempty"`
	Mobile   string `json:"mobile,omitempty"`
	Email    string `json:"email,omitempty"`
	Avatar   string `json:"avatar"`
	Platform struct {
		Code uint16 `json:"code"`
		Name string `json:"name"`
	} `json:"platform"`
}

type ToAccountOfModules struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
