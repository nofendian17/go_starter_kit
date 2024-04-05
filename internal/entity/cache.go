package entity

type Cache struct {
	Async AsyncCache `json:"async,omitempty"`
}

type AsyncCache struct {
	Identifier  string `json:"identifier"`
	CallbackURL string `json:"callback_url"`
}
