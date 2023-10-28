package User

type (
	SCard struct {
		Bank   string `json:"bank"`
		Secret int    `json:"secret"`
	}

	SWallet struct {
		Money int     `json:"money"`
		Cards []SCard `json:"cards"`
	}

	SUser struct {
		Id     int     `json:"id"`
		Age    int     `json:"age"`
		Name   string  `json:"name"`
		Wallet SWallet `json:"wallet"`
	}
)
