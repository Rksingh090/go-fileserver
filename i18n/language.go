package i18n

const (
	// status
	StatusSuccess = "SUCCESS"
	StatusError   = "ERROR"
	StatusWarning = "WARNING"

	// db
	InsertInDbError = "INSERT_INTO_DB_ERROR"

	// http status
	StatusUnauthenticated = "UNAUTHENTICATED"
	StatusUnauthorized    = "UNAUTHORIZED"

	// general
	AllFieldsRequired  = "ALL_FIELDS_REQUIRED"
	InteralServerError = "INTERNAL_SERVER_ERROR"

	//permissions
	ErrorGettingPermission = "ERROR_GETTING_PERMISSIONS"
	PermissionFound        = "PERMISSION_FOUND"
	PermissionNotFound     = "PERMISSION_NOT_FOUND"

	// auth
	UserNotFound        = "USER_NOT_FOUND"
	GenerateOtpFirst    = "GENERATE_OTP_FIRST"
	OtpExpired          = "OTP_EXPIRED"
	UserAlreadExists    = "USER_ALREADY_EXISTS"
	LoginSuccess        = "LOGIN_SUCCESS"
	UserNotRegistered   = "USER_NOT_REGISTERED"
	InvalidMobileNo     = "INVALID_MOBILE_NO"
	InvalidOTP          = "INVALID_OTP"
	OtpVerified         = "OTP_VERIFIED"
	UserCreated         = "USER_CREATED"
	OtpGenerated        = "OTP_GENERATED"
	OtpGenerationFailed = "OTP_GENERATION_FAILED"

	// users
	ErrorFetchingUsers = "ERROR_FETCHING_USERS"
	UsersFetched       = "USER_FETCHED"

	// job roles
	JobRoleCreated  = "JOB_ROLE_CREATED"
	JobRolesFetched = "JOB_ROLES_FETCHED"

	// file server
	ErrorUploadFile            = "ERROR_UPLOADING_FILE"
	FileUploadedSuccess        = "FILE_UPLOADED_SUCCESSFULLY"
	PermErrorCreatingDirectory = "PERMISSION_ERROR_CREATE_DIR"
)
