package constants

const (
	FEATURE_FLAG = "FEATURE_FLAG"
	USER_AUTH    = "USER_AUTH"
	NOT_FOUND    = "NOT_FOUND"

	INVALID_DATETIME_LAYOUT = "INVALID_DATETIME_LAYOUT"
)

const (
	DATE_FORMAT          = "2006-01-02"
	DATE_TIME_FORMAT     = "2006-01-02 15:04:05"
	DATE_TIME_UTC_FORMAT = "2006-01-02T15:04:05Z"
	THAI_LOCATION        = "Asia/Bangkok"
	ROW_DUPLICATE        = "Row duplicate"
	DATE_DUPLICATE       = "Date duplicate"
	WRONG_DATA           = "Wrong data"
	WRONG_FORMAT         = "Wrong format"
	REQUIRED_FIELD       = "Required field"
	FIND_NOT_FOUND       = "Not found"
	PRICE_WITH_VAT       = 1.07
)

const (
	DEFAULT_PAGE      = uint16(1)
	DEFAULT_PAGE_SIZE = int(50)
)

var SORT_FIELDS = map[string]string{
	"itemId":     "item_id",
	"supplierId": "supplier_id",
	"startDate":  "start_date",
	"endDate":    "end_date",
	"updatedAt":  "updated_at",
	"createdAt":  "created_at",
	"id":         "id",
	"itemName":   "item_name",
}

var SORT_ORDERS = map[string]string{
	"asc":  "asc",
	"desc": "desc",
}

var DEFAULT_SORT = []string{"item_id asc"}
