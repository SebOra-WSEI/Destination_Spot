package token

import "time"

var expireTime = (time.Hour * 1).Milliseconds()

func CreateToken() {
	expAt := time.Now().UnixMilli() + expireTime

}
