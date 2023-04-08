package auth

import (
	"context"
	"errors"
	"first_go/db"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func CreateUser(name string, mail string, password string) (*mongo.InsertOneResult, error) {
	ctx := context.Background()
	client, err := db.Connect()
	if err != nil {
		return nil, err
	}
	// MongoDBから切断する
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	usrs := client.Database("maindb").Collection("users")
	birthday := time.Date(2007, 1, 13, 0, 0, 0, 0, time.UTC)
	realPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	// 既に入力されたユーザーが存在している場合、エラーを出す
	filter := bson.M{"mail": mail}
	var userHasSentMail bson.M
	if err := usrs.FindOne(ctx, filter).Decode(&userHasSentMail); err == nil {
		return nil, errors.New("user already exists")
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
		return nil, err
	}
	return res, err
}

func Login(mail string, password string) (*Auth, error) {
	ctx := context.Background()
	client, err := db.Connect()
	if err != nil {
		return nil, err
	}
	// ユーザーの情報をデータベースから取得
	usrs := client.Database("maindb").Collection("users")
	var result *ServerUser
	err = usrs.FindOne(ctx, bson.M{"mail": mail}).Decode(&result)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, errors.New("user not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password)); err != nil {
		return nil, errors.New("password is not matched")
	}
	// JWTをする
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &JWTCustomClaims{
		Name: result.Name,
		Mail: result.Mail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 365).Unix(),
		},
	})
	tokenString, err := token.SignedString([]byte("nothing"))
	if err != nil {
		return nil, err
	}

	// MongoDBから切断する
	if client.Disconnect(context.Background()) != nil {
		return nil, err
	}
	return &Auth{
		Token: tokenString,
		Model: ClientUser{
			ID:        result.ID,
			Name:      result.Name,
			Mail:      result.Mail,
			Birthday:  result.Birthday,
			CreatedAt: result.CreatedAt,
			LastLogIn: result.LastLogIn},
	}, nil
}
