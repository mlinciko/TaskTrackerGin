package models

// Registry содержит все модели, которые нужно мигрировать
func Registry() []interface{} {
	return []interface{}{
		&User{},
		&Task{},
		&Status{},
		&Priority{},
	}
}
