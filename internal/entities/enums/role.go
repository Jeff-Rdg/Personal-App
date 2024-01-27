package enums

type Role string

const (
	ADMIN   Role = "admin"
	COACH   Role = "coach"
	STUDENT Role = "student"
)

func IsValidRole(role string) bool {
	isValid := map[Role]bool{
		ADMIN:   true,
		COACH:   true,
		STUDENT: true,
	}

	return isValid[Role(role)]
}
