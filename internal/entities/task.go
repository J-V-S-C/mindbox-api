package entity

import "time"

type Task struct {
	ID          string
	Name        string
	Description string
	Done        bool
	IsDaily     bool
	IsExpired 	bool
	Lifetime    string
	CategoryID  string
}

func (t *Task) CheckExpired() bool {
    if t.Lifetime == "" {
        return false
    }
    // Lógica de parsing e comparação
    expiration, err := time.Parse(time.RFC3339, t.Lifetime)
    if err != nil {
        return false
    }
    return time.Now().After(expiration)
}