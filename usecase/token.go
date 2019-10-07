package usecase

import (
	"fmt"
	"log"
	"strconv"

	"github.com/arjunksofficial/token-pooling/infrastructure"
	"github.com/arjunksofficial/token-pooling/lru"
	"github.com/jinzhu/gorm"
)

type Token struct {
	TokenID   int    `gorm:"column:token_id"`
	TokenData string `gorm:"column:token_data"`
	Count     int    `gorm:"column:count"`
}

type UseCaseInterface interface {
	GetToken() string
	Stats() string
}

// UseCase ..
type UseCase struct {
	SQLHandler *infrastructure.SQL
	Cache      *lru.Cache
	LastKey    int
	FirstKey   int
	CurrentKey *int
}

// GetToken ..
func (u *UseCase) GetToken() string {
	if *u.CurrentKey == u.LastKey {
		*u.CurrentKey = u.FirstKey
	} else {
		*u.CurrentKey = *u.CurrentKey + 1
	}
	val, _ := u.Cache.Get(*u.CurrentKey)
	u.SQLHandler.Master.Table("tokens").Where("token_id = "+strconv.Itoa(*u.CurrentKey)).UpdateColumn("count", gorm.Expr("count + 1"))
	return fmt.Sprintf("%v", val)
}

// Stats ..
func (u *UseCase) Stats() string {
	tokens := []Token{}
	err := u.SQLHandler.Master.Table("tokens").Find(&tokens).Error
	if err != nil {
		log.Panic(err)
	}
	res := "token_id \t token\t count\n"
	for _, v := range tokens {
		res += fmt.Sprintf("%v\t\t\t%v\t\t%v\n", v.TokenID, v.TokenData, v.Count)
	}
	return res
}

// NewUsecase ..
func NewUsecase(db *infrastructure.SQL, cache *lru.Cache, currentkey *int) *UseCase {
	return &UseCase{
		SQLHandler: db,
		Cache:      cache,
		CurrentKey: currentkey,
		FirstKey:   1,
		LastKey:    12,
	}
}
