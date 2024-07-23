package skill

import (
	"database/sql"
	"testing"

	"github.com/lib/pq"
	_ "modernc.org/sqlite"
)

func newMockDB() *sql.DB {
	db, _ := sql.Open("sqlite", "file:skill?mode=memory&cache=shared")
	q := `
CREATE TABLE IF NOT EXISTS skill (
    key TEXT PRIMARY KEY,
    name TEXT NOT NULL DEFAULT '',
    description TEXT NOT NULL DEFAULT '',
    logo TEXT NOT NULL DEFAULT '',
    tags TEXT [] NOT NULL DEFAULT '{}'
);
`
	db.Exec(q)
	return db
}

func getCount(db *sql.DB) int {
	var count int
	db.QueryRow("SELECT COUNT(*) FROM skill").Scan(&count)
	return count
}

func getData(db *sql.DB, key string) Skill {
	rows := db.QueryRow("SELECT key, name, description, logo, tags FROM skill WHERE key = $1", key)
	var skill Skill
	rows.Scan(&skill.Key, &skill.Name, &skill.Description, &skill.Logo, pq.Array(&skill.Tags))
	return skill
}

func TestStorageCreate(t *testing.T) {
	// Arrange
	db := newMockDB()
	defer db.Close()

	storage := NewSkillStorage(db)
	give := CreateSkillRequest{
		Key:         "go",
		Name:        "Go",
		Description: "Golang",
		Logo:        "https://golang.org/doc/gopher/frontpage.png",
		Tags:        []string{"go", "golang"},
	}

	// Act
	err := storage.CreateSkill(give)

	// Assert
	if err != nil {
		t.Fatal(err)
	}

	if getCount(db) != 1 {
		t.Errorf("getCount() = %d, want 1", getCount(db))
	}
}

func TestStorageUpdate(t *testing.T) {
	// Arrange
	db := newMockDB()
	defer db.Close()
	db.Exec("INSERT INTO skill (key, name, description, logo, tags) VALUES ('go', 'Go', 'Golang', 'https://golang.org/doc/gopher/frontpage.png', '{go, golang}')")

	storage := NewSkillStorage(db)
	give := UpdateSkillRequest{
		Name:        "Golang Intensive Course",
		Description: "Go programming language",
		Logo:        "https://golang.org/doc/gopher/frontpage.png",
		Tags:        []string{"go", "golang"},
	}

	// Act
	err := storage.UpdateSkill("go", give)

	// Assert
	if err != nil {
		t.Fatal(err)
	}

	data := getData(db, "go")
	if data.Name != "Golang Intensive Course" {
		t.Errorf("got.Name = %s, want Golang", data.Name)
	}
}

func TestStorageUpdateName(t *testing.T) {
	// Arrange
	db := newMockDB()
	defer db.Close()
	db.Exec("INSERT INTO skill (key, name, description, logo, tags) VALUES ('go', 'Go', 'Golang', 'https://golang.org/doc/gopher/frontpage.png', '{go, golang}')")

	storage := NewSkillStorage(db)

	// Act
	err := storage.UpdateName("go", "Golang Intensive Course")

	// Assert
	if err != nil {
		t.Fatal(err)
	}

	data := getData(db, "go")
	if data.Name != "Golang Intensive Course" {
		t.Errorf("got.Name = %s, want Golang", data.Name)
	}
}

func TestStorageUpdateDescription(t *testing.T) {
	// Arrange
	db := newMockDB()
	defer db.Close()
	db.Exec("INSERT INTO skill (key, name, description, logo, tags) VALUES ('go', 'Go', 'Golang', 'https://golang.org/doc/gopher/frontpage.png', '{go, golang}')")

	storage := NewSkillStorage(db)

	// Act
	err := storage.UpdateDescription("go", "Go programming language")

	// Assert
	if err != nil {
		t.Fatal(err)
	}

	data := getData(db, "go")
	if data.Description != "Go programming language" {
		t.Errorf("got.Description = %s, want Golang", data.Description)
	}
}

func TestStorageUpdateLogo(t *testing.T) {
	// Arrange
	db := newMockDB()
	defer db.Close()
	db.Exec("INSERT INTO skill (key, name, description, logo, tags) VALUES ('go', 'Go', 'Golang', 'https://golang.org/doc/gopher/frontpage.png', '{go, golang}')")

	storage := NewSkillStorage(db)

	// Act
	err := storage.UpdateLogo("go", "https://golang.org/doc/gopher/frontpage.png")

	// Assert
	if err != nil {
		t.Fatal(err)
	}

	data := getData(db, "go")
	if data.Logo != "https://golang.org/doc/gopher/frontpage.png" {
		t.Errorf("got.Logo = %s, want https://golang.org/doc/gopher/frontpage.png", data.Logo)
	}
}

func TestStorageUpdateTags(t *testing.T) {
	// Arrange
	db := newMockDB()
	defer db.Close()
	db.Exec("INSERT INTO skill (key, name, description, logo, tags) VALUES ('go', 'Go', 'Golang', 'https://golang.org/doc/gopher/frontpage.png', '{go, golang}')")

	storage := NewSkillStorage(db)

	// Act
	err := storage.UpdateTags("go", []string{"go", "golang", "programming"})

	// Assert
	if err != nil {
		t.Fatal(err)
	}

	data := getData(db, "go")
	if len(data.Tags) != 3 {
		t.Errorf("got.Tags = %v, want [go, golang, programming]", data.Tags)
	}
}

func TestStorageDeleteSkill(t *testing.T) {
	// Arrange
	db := newMockDB()
	defer db.Close()
	db.Exec("INSERT INTO skill (key, name, description, logo, tags) VALUES ('go', 'Go', 'Golang', 'https://golang.org/doc/gopher/frontpage.png', '{go, golang}')")

	storage := NewSkillStorage(db)

	// Act
	err := storage.DeleteSkill("go")

	// Assert
	if err != nil {
		t.Fatal(err)
	}

	if getCount(db) != 0 {
		t.Errorf("getCount() = %d, want 0", getCount(db))
	}
}
