package constants

const (
	LoginRedisPrefix      = "login_"
	VerifyCodeRedisPrefix = "verify_code_"

	AssayLoginRedisPrefix      = LoginRedisPrefix + "assay_"
	AssayVerifyCodeRedisPrefix = VerifyCodeRedisPrefix + "assay_"
)
