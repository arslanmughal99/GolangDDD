package models

import "massivleads/exceptions"

type Result struct {
	Result    interface{}
	Exception *exceptions.BaseException
}
