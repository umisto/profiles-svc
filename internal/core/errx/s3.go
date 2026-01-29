package errx

import "github.com/netbill/ape"

var (
	ErrorProfileAvatarContentFormatIsNotAllowed = ape.DeclareError("PROFILE_AVATAR_CONTENT_FORMAT_IS_NOT_ALLOWED")
	ErrorProfileAvatarContentTypeIsNotAllowed   = ape.DeclareError("PROFILE_AVATAR_CONTENT_TYPE_IS_NOT_ALLOWED")
	ErrorProfileAvatarTooLarge                  = ape.DeclareError("PROFILE_AVATAR_TOO_LARGE")

	ErrorNoContentUploaded = ape.DeclareError("NO_CONTENT_UPLOADED")
)
