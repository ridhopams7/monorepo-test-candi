package warehoouse

import (
	"net/http"
	"time"

	"github.com/golangid/candi/candiutils"
)

type warehoouseRESTImpl struct {
	host    string
	authKey string
	httpReq candiutils.HTTPRequest
}

// NewWarehoouseServiceREST constructor
func NewWarehoouseServiceREST(host string, authKey string) Warehoouse {

	return &warehoouseRESTImpl{
		host:    host,
		authKey: authKey,
		httpReq: candiutils.NewHTTPRequest(
			candiutils.HTTPRequestSetRetries(5),
			candiutils.HTTPRequestSetSleepBetweenRetry(500*time.Millisecond),
			candiutils.HTTPRequestSetHTTPErrorCodeThreshold(http.StatusBadRequest),
			candiutils.HTTPRequestSetBreakerName("warehoouse"),
		),
	}
}
