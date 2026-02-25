package model

type Task struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Done        bool    `json:"done"`
	IsDaily     bool    `json:"isDaily"`
	IsExpired 	bool `json:"isExpired"`
	Lifetime    *string `json:"lifetime,omitempty"`
	CategoryID  string  `json:"categoryId"`
}
