package models

type UserProfile struct {
	ID     		string 		`json:"id"`
	UserID 		string 		`json:"user_id"`
}

type User struct {
	ID 			string 	    `json:"id"`
	Email 		string 	    `json:"email"`
	Password    password      `json:"-"`
    FirstName   string      `json:"first_name"`
    LastName    string      `json:"last_name"`
	Profile     UserProfile `json:"profile"`
}

type password struct {
    plaintext  *string
    hash       string
}
