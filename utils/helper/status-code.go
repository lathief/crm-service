package helper

import (
	"github.com/lathief/crm-service/constant"
	"net/http"
)

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case constant.ErrCustomerNotFound:
		fallthrough
	case constant.ErrAdminNotFound:
		fallthrough
	case constant.ErrRoleNotFound:
		fallthrough
	case constant.ErrSuperAdminNotFound:
		fallthrough
	case constant.ErrApprovalNotFound:
		return http.StatusNotFound
	case constant.ErrAdminUsernameExists:
		return http.StatusConflict
	case constant.ErrAdminPasswordNotMatch:
		fallthrough
	case constant.ErrAdminAccountNotActive:
		fallthrough
	case constant.ErrAdminNotApprove:
		fallthrough
	case constant.ErrTokenInvalid:
		fallthrough
	case constant.ErrAdminNeedLogin:
		fallthrough
	case constant.ErrNotAllowedAccess:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
