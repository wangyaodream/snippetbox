package models

import (
	"database/sql"
	"errors"
	"time"
)

// Snippet represents a snippet of code
type Snippet struct {
	ID      int
	Title   string
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
	// 创建sql语句
	stmt := `SELECT id, title, content, created, expires
    FROM snippets
    WHERE expires > UTC_TIMESTAMP() AND id = ?`
	// 使用QueryRow()方法执行sql语句
	row := m.DB.QueryRow(stmt, id)
	// 创建一个Snippet对象
	s := &Snippet{}
	// 使用Scan()方法将查询结果赋值给Snippet对象
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}
	// 返回Snippet对象
	return s, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	stmt := `SELECT id, title, content, created, expires
	FROM snippets
	WHERE expires > UTC_TIMESTAMP()
	ORDER BY created DESC
	LIMIT 10`
	// 使用Query()方法执行sql语句获取多个结果
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	// 确保在函数返回时关闭rows
	defer rows.Close()
	snippets := []*Snippet{}
	// 遍历结果集
	for rows.Next() {
		s := &Snippet{}
		// 使用Scan()方法将查询结果赋值给Snippet对象
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		// 将Snippet对象添加到切片中
		snippets = append(snippets, s)
	}
	// 检查遍历结果集时是否发生错误
	if err = rows.Err(); err != nil {
		return nil, err
	}
	// 返回切片
	return snippets, nil
}
