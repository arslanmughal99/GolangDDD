package dtos

type RegisterUser struct {
	Name     string `json:"name" form:"name" valid:"maxstringlength(100)~Name must be max 100 characters.,required~Name is required."`
	Email    string `json:"email" form:"email" valid:"email~Invalid email address provided.,required~Email is required."`
	Username string `json:"username" form:"username" valid:"minstringlength(5)~Username must be at least 5 characters long."`
	Password string `json:"password" form:"password" valid:"minstringlength(5)~Password must be at least 5 characters long.,maxstringlength(100)~Password cannot be more than 100 characters long."`
}

type LoginUser struct {
	Remember bool   `json:"remember" form:"remember"`
	Username string `json:"username" form:"username" valid:"minstringlength(5)~Invalid username."`
	Password string `json:"password" form:"password" valid:"minstringlength(5)~Invalid password.,maxstringlength(100)~Invalid password."`
}

type LoginUserResponse struct {
	Token   string `json:"token"`
	Refresh string `json:"refreshToken,omitempty"`
}

type RegisterUserResponse struct {
	Username string `json:"username"`
}
