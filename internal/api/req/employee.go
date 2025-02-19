package req

type LoginDTO struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterDTO struct {
	//Id       uint64 `json:"id"`                          //员工id
	IdNumber string `json:"idNumber" binding:"required"` //身份证
	Name     string `json:"name" binding:"required"`     //姓名
	Phone    string `json:"phone" binding:"required"`    //手机号
	Sex      string `json:"sex" binding:"required"`      //性别
	UserName string `json:"username" binding:"required"` //用户名
}

type UpdateDTO struct {
	IdNumber string `json:"idNumber" ` //身份证
	Name     string `json:"name" `     //姓名
	Phone    string `json:"phone" `    //手机号
	Sex      string `json:"sex" `      //性别
	UserName string `json:"username" ` //用户名
}

type EmployeeEditPassword struct {
	//EmpId       uint64 `json:"empId"`
	NewPassword string `json:"newPassword" binding:"required"`
	OldPassword string `json:"oldPassword" binding:"required"`
}
