package service

import "errors"

// Domain-level errors returned by application services during validation and business operations
var (
	ErrNotEnoughMoney = errors.New("the club cannot afford this operation: insufficient balance")
	ErrSalaryCapExceeded = errors.New("transaction rejected: this operation exceeds the league salary cap")
	ErrPlayerNotFound = errors.New("operation failed: the specified player does not belong to this team")
	ErrPlayerAlreadySigned = errors.New("operation failed: the player is already signed to this team")
)
