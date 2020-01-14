package models
import "time"

type Good struct {
	ID	int	`gorm:"primary_key" json:"id"` //
	Name	string	`json:"name" binding:"required" ` //
	Price	float64	`json:"price" binding:"required"` //
	Remain	int	`json:"remain" binding:"required"` //
	Seller	string	`json:"seller" binding:"required"` //
	Pic	string	`gorm:"default":'www.baidu.com' json:"pic"` //
	LikeSum	int	`gorm:"default":0 json:"like_sum"` //
	BuySum	int	`gorm:"default":0 json:"buy_sum"` //
	PostAddress	string	`json:"post_address" binding:"required"` //
	CommentSum	string	`gorm:"default":0 json:"comment_sum"` //
	CreatedAt time.Time `json:"-"`
}

