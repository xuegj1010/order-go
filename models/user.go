package models

type User struct {
	Uid         int    `gorm:"type:int;primaryKey;not null;autoIncrement;comment:用户Uid"`
	Nickname    string `gorm:"type:varchar(100);not null;default:'';comment:用户名"`
	Mobile      string `gorm:"type:varchar(20);not null;default:'';comment:手机号码"`
	Email       string `gorm:"type:varchar(100);not null;default:'';comment:邮箱地址"`
	Sex         int    `gorm:"type:tinyint(1);not null;default:0;comment:1：男 2：女 0：没填写"`
	Avatar      string `gorm:"type:varchar(64);not null;default:'';comment:头像"`
	LoginName   string `gorm:"type:varchar(20);uniqueIndex;not null;default:'';comment:登录用户名"`
	LoginPwd    string `gorm:"type:varchar(32);not null;default:'';comment:登录密码"`
	LoginSalt   string `gorm:"type:varchar(32);not null;default:'';comment:登录密码的随机加密密钥"`
	Status      int    `gorm:"type:tinyint(1);not null;default:1;comment:1：有效 0：无效"`
	UpdatedTime int    `gorm:"type:timestamp; not null;autoUpdateTime;comment:最后一次更新时间"`
	CreatedTime int    `gorm:"type:timestamp; not null;autoCreateTime;comment:插入时间"`
}

func (u User) TableName() string {
	return "user"
}
