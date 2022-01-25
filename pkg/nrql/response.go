package nrql

type Response struct {
	Data *Data `json:"data"`
}

type Data struct {
	Actor *Actor `json:"actor"`
}

type Actor struct {
	Account *Account `json:"account"`
}

type Account struct {
	NRQL *NRQL `json:"nrql"`
}

type NRQL struct {
	Results []map[string]interface{}
}
