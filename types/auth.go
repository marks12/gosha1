package types

type Auth struct {

    Email     string    `gorm:"many2many:PersonAccount;association_foreignkey:idAccount;foreignkey:idPerson"`
    Password  string
    Token     string
    //Auth remove this line for disable generator functionality
}

func (auth *Auth) Validate()  {
//Validate remove this line for disable generator functionality
}