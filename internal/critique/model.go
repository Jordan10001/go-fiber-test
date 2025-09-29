package critique

// Critique represents a single user feedback entry.
type Critique struct {
	ID        int64  `json:"id" db:"id"`
	Title     string `json:"title" db:"title"`
	Content   string `json:"content" db:"content"`
	CreatedAt int64  `json:"created_at" db:"created_at"`
}