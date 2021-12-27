package redmine

type CustomField struct {
	Id          int         `json:"id"`
	Name        string      `json:"name,omitempty"`
	Description string      `json:"description,omitempty"`
	Multiple    bool        `json:"multiple,omitempty"`
	Value       interface{} `json:"value"`
}
