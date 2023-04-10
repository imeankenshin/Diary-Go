package auth

import (
	"context"
	"first_go/pkg/app"
	"first_go/pkg/db"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type ServerUser struct {
	ID        string    `bson:"_id,omitempty"`
	Name      string    `bson:"name"`
	Password  string    `bson:"password"`
	Mail      string    `bson:"mail"`
	Birthday  time.Time `bson:"birthday"`
	CreatedAt time.Time `bson:"created_at"`
	LastLogIn time.Time `bson:"last_login"`
}
type ClientUser struct {
	ID        string    `json:"_id"`
	Name      string    `json:"name"`
	Mail      string    `json:"mail"`
	Birthday  time.Time `json:"birthday"`
	CreatedAt time.Time `json:"created_at"`
	LastLogIn time.Time `json:"last_login"`
}
type Auth struct {
	Token string     `json:"token"`
	Model ClientUser `json:"model"`
}
type JWTCustomClaims struct {
	Name      string `json:"username"`
	Mail      string `json:"mail"`
	Birthday  string `json:"birthday"`
	CreatedAt string `json:"created_at"`
	jwt.StandardClaims
}

func CreateUser(c echo.Context) error {
	ctx := context.Background()
	// ヘッダーを取得
	req := c.Request().Header
	name := req.Get("Name")
	mail := req.Get("Mail")
	password := req.Get("Password")

	if name == "" || mail == "" || password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, app.ErrorMessage{Message: "name, mail, password can not be empty"})
	}

	//　データベースに接続する
	client, err := db.Connect()
	if err != nil {
		return app.ServerError(c, err)
	}

	// MongoDBから切断する
	defer client.Disconnect(context.Background())

	// ユザー
	usrs := client.Database("maindb").Collection("users")
	birthday := time.Date(2007, 1, 13, 0, 0, 0, 0, time.UTC)
	realPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return app.ServerError(c, err)
	}
	// 既に入力されたユーザーが存在している場合、エラーを出す
	filter := bson.M{"mail": mail}
	var userHasSentMail bson.M
	if err := usrs.FindOne(ctx, filter).Decode(&userHasSentMail); err == nil {
		return c.JSON(http.StatusBadRequest, app.ErrorMessage{Message: "user already exist"})
	}
	data := bson.M{
		"name":       name,
		"mail":       mail,
		"password":   string(realPass),
		"birthday":   primitive.NewDateTimeFromTime(birthday),
		"created_at": primitive.NewDateTimeFromTime(time.Now()),
	}
	res, err := usrs.InsertOne(ctx, data)
	if err != nil {
		return app.ServerError(c, err)
	}
	return c.JSON(http.StatusOK, res)
}

func Login(c echo.Context) error {
	// ヘッダーを取得
	mail := c.Request().Header.Get("Mail")
	password := c.Request().Header.Get("Password")
	if mail == "" || password == "" {
		return c.JSON(http.StatusBadRequest, app.ErrorMessage{Message: "mail or password is empty"})
	}
	ctx := context.Background()

	//　データベースに接続
	client, err := db.Connect()
	if err != nil {
		return app.ServerError(c, err)
	}

	// 処理が終わったら、MongoDBから切断する
	defer client.Disconnect(context.Background())

	// ユーザーの情報をデータベースから取得
	usrs := client.Database("maindb").Collection("users")
	var result *ServerUser
	err = usrs.FindOne(ctx, bson.M{"mail": mail}).Decode(&result)
	if err != nil {
		return app.ServerError(c, err)
	}

	// ユーザーがいなかった場合、エラーを返す
	if result == nil {
		return c.JSON(http.StatusBadRequest, "mail or password is wrong.")
	}

	// 入力されたメールが正しいか確認
	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password)); err != nil {
		return c.JSON(http.StatusBadRequest, app.ErrorMessage{Message: "your mail or password is not matched"})
	}
	// JWTをする
	decode := &JWTCustomClaims{
		Name: result.Name,
		Mail: result.Mail,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "https://localhost:1000",
			Subject:   "demo",
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 24 * 365).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, decode)

	// tokenを文字列に直す
	tokenString, err := token.SignedString([]byte("nothing"))
	if err != nil {
		return app.ServerError(c, err)
	}

	// ログインを更新する
	loginUpdate := bson.M{"$set": bson.M{"last_login": primitive.NewDateTimeFromTime(time.Now())}}
	_, err = usrs.UpdateOne(ctx, bson.M{"mail": mail}, loginUpdate)
	if err != nil {
		return app.ServerError(c, err)
	}

	return c.JSON(http.StatusOK, &Auth{
		Token: tokenString,
		Model: ClientUser{
			ID:        result.ID,
			Name:      result.Name,
			Mail:      result.Mail,
			Birthday:  result.Birthday,
			CreatedAt: result.CreatedAt,
			LastLogIn: result.LastLogIn,
		},
	})
}
