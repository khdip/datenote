package storage

type Event struct {
	ID       int64  `db:"id"`
	Name     string `db:"name"`
	Date     string `db:"date"`
	Info     string `db:"info"`
	Category string `db:"category"`
}

type Category struct {
	ID    int64  `db:"id"`
	Title string `db:"title"`
}
