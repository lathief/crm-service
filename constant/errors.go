package constant

import (
	"fmt"
)

var (
	ErrAdminNotFound         = fmt.Errorf("admin not found")
	ErrCustomerNotFound      = fmt.Errorf("customer not found")
	ErrAdminUsernameExists   = fmt.Errorf("Username already in use")
	ErrRoleNotFound          = fmt.Errorf("role not found")
	ErrSuperAdminNotFound    = fmt.Errorf("super admin not found")
	ErrAdminPasswordNotMatch = fmt.Errorf("Password does not match")
	ErrAdminAccountNotActive = fmt.Errorf("Your account not active, please contact super admin")
	ErrApprovalNotFound      = fmt.Errorf("approval not found")
	ErrAdminNotApprove       = fmt.Errorf("approved first, to active/deactive account")
	ErrTokenInvalid          = fmt.Errorf("Token invalid")
	ErrAdminNeedLogin        = fmt.Errorf("login to proceed")
	ErrNotAllowedAccess      = fmt.Errorf("You are not allowed to access this data")
	ErrAdminNotActive        = fmt.Errorf("Admin not active and verified")
)
