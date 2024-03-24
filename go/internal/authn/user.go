package authn

// User описывает доступ к данным абстрактного пользователя.
type User interface {
	ID() string
	Type() UserType
}

// UserType - тип пользователя.
type UserType int

const (
	UserTypeUnknown UserType = iota
	UserTypeAnonymous
	UserTypeAuthenticated
)

// AnonymousUser - структура данных анонимного пользователя.
type AnonymousUser struct{}

func (AnonymousUser) ID() string {
	panic("anonymous user")
}

func (AnonymousUser) Type() UserType {
	return UserTypeAnonymous
}

// AuthenticatedUser - структура данных аутенифицированного пользователя.
type AuthenticatedUser struct {
	id string
}

func (user AuthenticatedUser) ID() string {
	return user.id
}

func (user AuthenticatedUser) Type() UserType {
	return UserTypeAuthenticated
}
