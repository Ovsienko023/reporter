package dbmessage

// --------------------------------
//      	CREATE USER
// --------------------------------

type CreateUser struct {
	InvokerId   string `json:"invoker_id,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
}

type CreatedUser struct {
	UserId *string `json:"user_id,omitempty"`
}
