package database

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	users_table = `CREATE TABLE IF NOT EXISTS users (
		uuid TEXT PRIMARY KEY NOT NULL,
		name CHAR(50) NOT NULL,
		username CHAR(50) NOT NULL UNIQUE,
		email CHAR(50) NOT NULL UNIQUE, 
		password TEXT NOT NULL,
	);`
	post_table = `CREATE TABLE IF NOT EXISTS post (
		id SERIAL NOT NULL PRIMARY KEY,
		uuid TEXT NOT NULL,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		author VARCHAR(50) NOT NULL,
		createdat VARCHAR(50) NOT NULL,
		categories VARCHAR(50) NOT NULL,
		likeField INTEGER,
		dislikeField INTEGER,
		FOREIGN KEY (uuid) REFERENCES users(uuid) ON DELETE CASCADE,
		FOREIGN KEY (author) REFERENCES users(username) ON DELETE CASCADE
	);`
	comments_table = `CREATE TABLE IF NOT EXISTS comments (
		id SERIAL NOT NULL PRIMARY KEY,
		postID INT DEFAULT 0,
		content TEXT NOT NULL,
		author VARCHAR(50) NOT NULL,
		likeField INT DEFAULT 0,
		dislikeField INT DEFAULT 0,
		createdat VARCHAR(50) NOT NULL,
		FOREIGN KEY (postID) REFERENCES post(id) ON DELETE CASCADE,
		FOREIGN KEY (author) REFERENCES users(username) ON DELETE CASCADE
	);`
	likePostTable = `CREATE TABLE IF NOT EXISTS likePost (
		id SERIAL NOT NULL PRIMARY KEY,
		userID TEXT,
		postID INT DEFAULT 0,
		status INT DEFAULT 0,
		FOREIGN KEY (userID) REFERENCES users(uuid) ON DELETE CASCADE,
		FOREIGN KEY (postID) REFERENCES post(id) ON DELETE CASCADE
		);`
	likeCommentsTable = `CREATE TABLE IF NOT EXISTS likeComments(
		id SERIAL NOT NULL PRIMARY KEY,
		userID TEXT,
		commentsID INT DEFAULT 0,
		status INT DEFAULT 0,
		FOREIGN KEY (userID) REFERENCES users(uuid) ON DELETE CASCADE,
		FOREIGN KEY (commentsID) REFERENCES comments(id) ON DELETE CASCADE
		);`
)

// Создание таблицы пользователя
func CreatTables(db *sql.DB) error {
	allTables := []string{users_table, post_table, comments_table, likePostTable, likeCommentsTable}
	for _, v := range allTables {
		stmt, err := db.Prepare(v)
		if err != nil {
			fmt.Println("1", err)
			return fmt.Errorf("Create table: %w", err)
		}
		_, err = stmt.Exec()
		if err != nil {
			fmt.Println("2", err)
			return fmt.Errorf("Create table: %w", err)
		}
	}
	log.Println("All table created successfully!")
	return nil
}
