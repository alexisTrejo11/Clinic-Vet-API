package controller

type AuthController struct {
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (c *AuthController) Signup() error {
	return nil
}

func (c *AuthController) Login() error {
	return nil
}

func (c *AuthController) Logout() error {
	return nil
}

func (c *AuthController) LogoutAll() error {
	return nil
}
