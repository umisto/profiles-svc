package middlewares

import (
	"net/http"

	"github.com/netbill/ape"
	"github.com/netbill/ape/problems"
	"github.com/netbill/logium"
	"github.com/netbill/profiles-svc/internal/tokenmanager"
	"github.com/netbill/restkit/mdlv"
	"github.com/netbill/restkit/tokens"
)

type Service struct {
	accountAccessSK string
	uploadFilesSK   string

	log *logium.Logger
}

type Config struct {
	AccountAccessSK string
	UploadFilesSK   string
}

func New(
	log *logium.Logger,
	cfg Config,
) Service {
	return Service{
		accountAccessSK: cfg.AccountAccessSK,
		uploadFilesSK:   cfg.UploadFilesSK,
		log:             log,
	}
}

func (s Service) AccountAuth() func(next http.Handler) http.Handler {
	return mdlv.AccountAuth(s.log, accountDataCtxKey, s.accountAccessSK)
}

func (s Service) AccountRolesGrant(
	allowedRoles map[string]bool,
) func(http.Handler) http.Handler {
	return mdlv.AccountRoleGrant(s.log, accountDataCtxKey, allowedRoles)
}

func (s Service) UpdateOwnProfile() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			userData, ok := ctx.Value(accountDataCtxKey).(tokens.AccountJwtData)
			if !ok {
				s.log.Errorf("missing account data in context")
				ape.RenderErr(w, problems.Unauthorized("missing account data in context"))
				return
			}

			confirm := mdlv.ConfirmUploadFiles(
				s.log,
				uploadFilesCtxKey,
				s.uploadFilesSK,
				mdlv.ConfirmUploadFilesParams{
					Audience:   tokenmanager.ProfilesService,
					Resource:   tokenmanager.ProfileResource,
					ResourceID: userData.AccountID.String(),
				},
			)

			confirm(next).ServeHTTP(w, r)
		})
	}
}
