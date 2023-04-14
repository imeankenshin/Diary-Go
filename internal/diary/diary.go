package diary

import (
	"context"
	"first_go/pkg/app"
	"first_go/pkg/db"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Diary struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title   string             `json:"title" bson:"title"`
	Content string             `json:"content" bson:"content"`
	Date    time.Time          `json:"date" bson:"date"`
	Owner   string             `json:"owner" bson:"owner"`
}

/*データベースから、ユーザーの日記を取得し、クライアントに送信する*/
func GetDiary(c echo.Context) error {
	ctx := context.Background()
	var (
		header = c.Request().Header
		id     = header.Get("User-ID")
	)

	// IDが空なら、エラーを出す
	if id == "" {
		return c.JSON(http.StatusBadRequest, app.ErrorMessage{
			Message: "User-ID is empty",
		})
	}

	// データベースに接続する
	client, err := db.Connect()
	if err != nil {
		return app.ServerError(c, err)
	}

	// ユーザーの日記のデータを取得する
	cur, err := client.Database("maindb").Collection("diary").Find(ctx, bson.M{"owner": id})
	if err != nil {
		return app.ServerError(c, err)
	}
	defer cur.Close(ctx)

	// diaryスライスに日記を格納する
	var diary []Diary
	for cur.Next(ctx) {
		var d Diary
		err := cur.Decode(&d)
		if err != nil {
			return app.ServerError(c, err)
		}
		diary = append(diary, d)
	}
	if err := cur.Err(); err != nil {
		return app.ServerError(c, err)
	}

	return c.JSON(http.StatusOK, diary)
}

func PostDiary(c echo.Context) error {
	ctx := context.Background()

	// get header
	var (
		header  = c.Request().Header
		title   = header.Get("Diary-Title")
		content = header.Get("Diary-Content")
		id      = header.Get("User-ID")
	)
	if title == "" || content == "" || id == "" {
		return c.JSON(http.StatusBadRequest, app.ErrorMessage{Message: "title, content, date and id are required"})
	}

	// データベースに接続
	client, err := db.Connect()
	defer client.Disconnect(ctx)
	if err != nil {
		return app.ServerError(c, err)
	}

	// データを保存
	diary := &Diary{
		Title:   title,
		Content: content,
		Date:    time.Now(),
		Owner:   id,
	}
	res, err := client.Database("maindb").Collection("diary").InsertOne(ctx, &diary)
	if err != nil {
		return app.ServerError(c, err)
	}
	return c.JSON(http.StatusOK, &res)
}
