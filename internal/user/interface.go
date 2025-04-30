package user

type UserRepositoryInterface interface {
	CreateUser(userModel UserModel) (*UserModel, error)
}

type UserServiceInterface interface{}

type UserControllerInterface interface{}
