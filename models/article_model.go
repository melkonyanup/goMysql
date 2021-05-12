package models

import (
	"database/sql"

	"../entities"
)

type ArticleModel struct {
	Db *sql.DB
}

func (articleModel ArticleModel) ViewAll() (articel []entities.Article, err error) {
	rows, err := articleModel.Db.Query("select * from tbl_article")

	if err != nil {
		return nil, err
	} else {
		var articles []entities.Article
		for rows.Next() {
			var ID int64
			var Title string
			var Description string
			var Content string
			err2 := rows.Scan(&ID, &Title, &Description, &Content)

			if err2 != nil {
				return nil, err2
			} else {
				article := entities.Article{
					ID:          ID,
					Title:       Title,
					Description: Description,
					Content:     Content,
				}

				articles = append(articles, article)
			}
		}

		return articles, nil
	}
}
func (articleModel ArticleModel) ViewArticle(id int64) (article []entities.Article, err error) {
	result, err := articleModel.Db.Query("select * from tbl_article where ID = ?", id)

	if err != nil {
		return nil, err
	} else {
		var articleX []entities.Article
		for result.Next() {
			var ID int64
			var Title string
			var Description string
			var Content string
			err2 := result.Scan(&ID, &Title, &Description, &Content)

			if err2 != nil {
				return nil, err2
			} else {
				article := entities.Article{
					ID:          ID,
					Title:       Title,
					Description: Description,
					Content:     Content,
				}

				articleX = append(articleX, article)
			}
		}

		return articleX, nil
	}
}
func (articleModel ArticleModel) CreateArticle(articel *entities.Article) (err error) {
	result, err := articleModel.Db.Exec("insert into tbl_article(Title, Description, Content) values(?,?,?)", articel.Title, articel.Description, articel.Content)

	if err != nil {
		return err
	} else {
		articel.ID, _ = result.LastInsertId()
		return nil
	}
}

func (articleModel ArticleModel) UpdateArticle(articel *entities.Article) (int64, error) {
	result, err := articleModel.Db.Exec("update tbl_article set Title = ?, Description = ?, Content = ? where ID = ?", articel.Title, articel.Description, articel.Content, articel.ID)

	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}

func (articleModel ArticleModel) DeleteArticle(id int64) (int64, error) {
	result, err := articleModel.Db.Exec("delete from tbl_article  where ID = ?", id)

	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}
