package auto

type Role string

const (
	ADMIN     Role = "ADMIN"
	MODERATOR Role = "MODERATOR"
	ANALYST   Role = "ANALYST"
	SUPPORT   Role = "SUPPORT"
)

func IsValidRole(role Role) bool {
	switch role {
	case ADMIN, MODERATOR, ANALYST, SUPPORT:
		return true
	default:
		return false
	}
}

type User struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" gorm:"unique;not null" binding:"required"`
	Password string `json:"password,omitempty" gorm:"not null" binding:"required" json:"-"`

	Role Role `json:"role" gorm:"type:varchar(20);default:'SUPPORT'"`
}
