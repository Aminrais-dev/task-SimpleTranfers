package delivery

import "task/simpleTranfers/features/users"

type Request struct {
	Name    string  `json:"name" form:"name"`
	Balance float64 `json:"balance" form:"balance"`
}

func (req *Request) toCore() users.UserCore {
	return users.UserCore{
		Name:    req.Name,
		Balance: req.Balance,
	}
}
