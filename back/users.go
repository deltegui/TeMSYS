package temsys

import "time"

// Role defines what kind of things can do someone in the system.
type Role string

// Defined roles.
const (
	AdminRole Role = "admin"
	UserRole  Role = "user"
)

// User is the representation of anyone inside the system.
type User struct {
	Name     string
	Password string
	Role     Role
}

// Token reprents a thing that you can interchange with a
// user session.
type Token struct {
	Value   string    `json:"value"`
	Expires time.Time `json:"expires"`
	Owner   string    `json:"owner"`
	Role    Role      `json:"role"`
}

// UserRepository is the abstraction of a persistence mechanism to store users.
type UserRepository interface {
	Save(User) error
	GetByName(name string) (User, error)
	ExistsWithName(name string) bool
	Delete(name string) error
	GetAll() []User
}

// PasswordHasher is the abstraction of anything that can hash or check users password.
type PasswordHasher interface {
	Hash(password string) string
	CheckHashPassword(hash, password string) bool
}

// Tokenizer is the abstraction of anything that can take a user and create a token from it.
type Tokenizer interface {
	Tokenize(user User) Token
	Decode(raw string) (Token, error)
}

// LoginRequest have the information needed to do a login handled by LoginCase.
type LoginRequest struct {
	Name     string `json:"name" db:"name" validate:"required,min=3,max=255"`
	Password string `json:"password" db:"password" validate:"required,min=3,max=255"`
}

// LoginResponse returns your login user, your role and a token.
type LoginResponse struct {
	Name  string `json:"name"`
	Role  Role   `json:"role"`
	Token Token  `json:"token"`
}

// LoginCase handles a user login.
type LoginCase struct {
	userRepository UserRepository
	hasher         PasswordHasher
	tokenizer      Tokenizer
}

// NewLoginCase creates a ready to go LoginCase.
func NewLoginCase(val Validator, userRepo UserRepository, hasher PasswordHasher, tokenizer Tokenizer) UseCase {
	return Validate(LoginCase{
		userRepository: userRepo,
		hasher:         hasher,
		tokenizer:      tokenizer,
	}, val)
}

// Exec the Login use case. Expects the request to be already validated.
func (login LoginCase) Exec(presenter Presenter, raw UseCaseRequest) {
	req := raw.(LoginRequest)
	if !login.userRepository.ExistsWithName(req.Name) {
		presenter.PresentError(UserNotFoundErr)
		return
	}
	user, _ := login.userRepository.GetByName(req.Name)
	if login.hasher.CheckHashPassword(user.Password, req.Password) {
		presenter.PresentError(InvalidPasswordErr)
		return
	}
	token := login.tokenizer.Tokenize(user)
	presenter.Present(LoginResponse{
		Name:  user.Name,
		Role:  user.Role,
		Token: token,
	})
}

// CreateUserRequest stores all data needed by CreateUserCase.
type CreateUserRequest struct {
	Name     string `validate:"required,min=3,max=255"`
	Password string `validate:"required,min=3,max=255"`
}

// UserResponse is the generic user response.
type UserResponse struct {
	Name string
	Role Role
}

// CreateUserCase handles the creation of a not admin user.
type CreateUserCase struct {
	userRepository UserRepository
	hasher         PasswordHasher
}

// NewCreateUserCase creates a ready to go CreateUserCase.
func NewCreateUserCase(val Validator, userRepo UserRepository, hasher PasswordHasher) UseCase {
	return Validate(CreateUserCase{
		userRepository: userRepo,
		hasher:         hasher,
	}, val)
}

// Exec the CreateUser use case. Expects the request to be already validated.
func (create CreateUserCase) Exec(presenter Presenter, raw UseCaseRequest) {
	req := raw.(CreateUserRequest)
	if create.userRepository.ExistsWithName(req.Name) {
		presenter.PresentError(UserAlreadyExitsErr)
		return
	}
	hashed := create.hasher.Hash(req.Password)
	user := User{
		Name:     req.Name,
		Password: hashed,
		Role:     UserRole,
	}
	create.userRepository.Save(user)
	presenter.Present(UserResponse{
		Name: req.Name,
		Role: UserRole,
	})
}

// DeleteUserResponse is the data structure returned
// by DeleteUserCase
type DeleteUserResponse struct {
	Name string
}

// DeleteUserCase handles the creation of a not admin user.
type DeleteUserCase struct {
	userRepository UserRepository
}

// NewDeleteUserCase creates a ready to go DeleteUserCase.
func NewDeleteUserCase(userRepo UserRepository) UseCase {
	return DeleteUserCase{
		userRepository: userRepo,
	}
}

// Exec the CreateUser use case. Expects the request to be already validated.
func (create DeleteUserCase) Exec(presenter Presenter, raw UseCaseRequest) {
	name := raw.(string)
	if !create.userRepository.ExistsWithName(name) {
		presenter.PresentError(UserNotFoundErr)
		return
	}
	create.userRepository.Delete(name)
	presenter.Present(DeleteUserResponse{
		Name: name,
	})
}

// GetAllUserCase handles the creation of a not admin user.
type GetAllUserCase struct {
	userRepository UserRepository
}

// NewGetAllUserCase creates a ready to go GetAllUserCase.
func NewGetAllUserCase(userRepo UserRepository) UseCase {
	return GetAllUserCase{
		userRepository: userRepo,
	}
}

// Exec the GetAll user case.
func (create GetAllUserCase) Exec(presenter Presenter, raw UseCaseRequest) {
	users := create.userRepository.GetAll()
	var result []UserResponse
	for _, user := range users {
		result = append(result, UserResponse{
			Name: user.Name,
			Role: user.Role,
		})
	}
	presenter.Present(result)
}
