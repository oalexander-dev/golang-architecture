package operations

import "github.com/oalexander-dev/golang-architecture/domain"

func NewOps(r domain.Repo) domain.Ops {
	return domain.Ops{
		User: newUserOps(r),
	}
}
