package models

import (
	"database/sql"
	"time"
)


// Snippet represents a snippet of code
type Snippet struct {
    ID int
    Title string
    Content string
    Created time.Time
    Expires time.Time
}

// SnippetModel defines a model type which wraps a sql.DB connection pool.
type SnippetModel struct {
    DB *sql.DB
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
    // 创建一个sql语句,INTERVAL ? DAY 为了让expires天数后过期
    stmt := `INSERT INTO snippets (title, content, created, expires)
    VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))
    `

    // 使用Exec()方法执行sql语句
    result, err := m.DB.Exec(stmt, title, content, expires)
    if err != nil {
        return 0, err
    }

    // 调用LastInsertId()方法获取最后插入的ID
    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    // 返回ID
    return int(id), nil


}

func (m *SnippetModel) Get(id int) (*Snippet, error) {
    return nil, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
    return nil, nil
}
