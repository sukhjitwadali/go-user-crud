package controllers

import (
	"go-user-crud/models"
	"sync"
)

var (
	users    = make(map[int]models.User)
	nextID   = 1
	usersMux sync.Mutex
)
