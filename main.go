package EchoTestContext

import (
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"testing"

	"github.com/labstack/echo"
)

type EchoTestContext struct {
	t *testing.T

	AttachmentResponseCounter         int
	BindResponseCounter               int
	BlobResponseCounter               int
	CookieErrResponseCounter          int
	CookieMultiResponseCounter        int
	CookiesResponseCounter            int
	EchoResponseCounter               int
	FileResponseCounter               int
	FormFileErrResponseCounter        int
	FormFileMultiResponseCounter      int
	FormParamsErrResponseCounter      int
	FormParamsURLResponseCounter      int
	FormValueResponseCounter          int
	GetResponseCounter                int
	HandlerResponseCounter            int
	HTMLBlobResponseCounter           int
	HTMLResponseCounter               int
	InlineResponseCounter             int
	IsTLSResponseCounter              int
	IsWebSocketResponseCounter        int
	JSONBlobResponseCounter           int
	JSONPBlobResponseCounter          int
	JSONPResponseCounter              int
	JSONPrettyResponseCounter         int
	JSONResponseCounter               int
	LoggerResponseCounter             int
	MultipartFormErrResponseCounter   int
	MultipartFormMultiResponseCounter int
	NoContentResponseCounter          int
	ParamNamesResponseCounter         int
	ParamResponseCounter              int
	ParamValuesResponseCounter        int
	PathResponseCounter               int
	QueryParamResponseCounter         int
	QueryParamsResponseCounter        int
	QueryStringResponseCounter        int
	RealIPResponseCounter             int
	RedirectResponseCounter           int
	RenderResponseCounter             int
	RequestResponseCounter            int
	ResponseResponseCounter           int
	SchemeResponseCounter             int
	StreamResponseCounter             int
	StringResponseCounter             int
	ValidateResponseCounter           int
	XMLBlobResponseCounter            int
	XMLPrettyResponseCounter          int
	XMLResponseCounter                int

	AttachmentFileRequestValues    []string
	AttachmentNameRequestValues    []string
	BindRequestValues              []interface{}
	BlobBRequestValues             [][]byte
	BlobCodeRequestValues          []int
	BlobContentRequestValues       []string
	CookieRequestValues            []string
	ErrorRequestValues             []error
	FileRequestValues              []string
	FormFileRequestValues          []string
	FormValueRequestValues         []string
	GetRequestValues               []string
	HTMLBlobBRequestValues         [][]byte
	HTMLBlobCodeRequestValues      []int
	HTMLCodeRequestValues          []int
	HTMLHTMLRequestValues          []string
	InlineFileRequestValues        []string
	InlineNameRequestValues        []string
	JSONBlobBRequestValues         [][]byte
	JSONBlobCodeRequestValues      []int
	JSONCodeRequestValues          []int
	JSONIRequestValues             []interface{}
	JSONPBlobBRequestValues        [][]byte
	JSONPBlobCallbackRequestValues []string
	JSONPBlobCodeRequestValues     []int
	JSONPCallbackRequestValues     []string
	JSONPCodeRequestValues         []int
	JSONPIRequestValues            []interface{}
	JSONPrettyCodeRequestValues    []int
	JSONPrettyIndentRequestValues  []string
	JSONPrettyIRequestValues       []interface{}
	NoContentRequestValues         []int
	ParamRequestValues             []string
	QueryParamRequestValues        []string
	RedirectCodeRequestValues      []int
	RedirectURLRequestValues       []string
	RenderCodeRequestValues        []int
	RenderDataRequestValues        []interface{}
	RenderNameRequestValues        []string
	ResetRRequestValues            []*http.Request
	ResetWRequestValues            []http.ResponseWriter
	SetCookieRequestValues         []*http.Cookie
	SetHandlerRequestValues        []echo.HandlerFunc
	SetKeyRequestValues            []string
	SetLoggerRequestValues         []echo.Logger
	SetParamNamesRequestValues     []string
	SetParamValuesRequestValues    []string
	SetPathRequestValues           []string
	SetRequestValues               []*http.Request
	SetResponseValues              []*echo.Response
	SetValRequestValues            []interface{}
	StreamCodeRequestValues        []int
	StreamContentRequestValues     []string
	StreamRRequestValues           []io.Reader
	StringCodeRequestValues        []int
	StringSRequestValues           []string
	ValidateRequestValues          []interface{}
	XMLBlobBRequestValues          [][]byte
	XMLBlobCodeRequestValues       []int
	XMLCodeRequestValues           []int
	XMLIRequestValues              []interface{}
	XMLPrettyCodeRequestValues     []int
	XMLPrettyIndentRequestValues   []string
	XMLPrettyIRequestValues        []interface{}

	AttachmentResponse         []error
	BindResponse               []error
	BlobResponse               []error
	CookieErrResponse          []error
	CookieMultiResponse        []*http.Cookie
	CookiesResponse            [][]*http.Cookie
	EchoResponse               []*echo.Echo
	FileResponse               []error
	FormFileErrResponse        []error
	FormFileMultiResponse      []*multipart.FileHeader
	FormParamsErrResponse      []error
	FormParamsURLResponse      []url.Values
	FormValueResponse          []string
	GetResponse                []interface{}
	HandlerResponse            []echo.HandlerFunc
	HTMLBlobResponse           []error
	HTMLResponse               []error
	InlineResponse             []error
	IsTLSResponse              []bool
	IsWebSocketResponse        []bool
	JSONBlobResponse           []error
	JSONPBlobResponse          []error
	JSONPResponse              []error
	JSONPrettyResponse         []error
	JSONResponse               []error
	LoggerResponse             []echo.Logger
	MultipartFormErrResponse   []error
	MultipartFormMultiResponse []*multipart.Form
	NoContentResponse          []error
	ParamNamesResponse         [][]string
	ParamResponse              []string
	ParamValuesResponse        [][]string
	PathResponse               []string
	QueryParamResponse         []string
	QueryParamsResponse        []url.Values
	QueryStringResponse        []string
	RealIPResponse             []string
	RedirectResponse           []error
	RenderResponse             []error
	RequestResponse            []*http.Request
	ResponseResponse           []*echo.Response
	SchemeResponse             []string
	StreamResponse             []error
	StringResponse             []error
	ValidateResponse           []error
	XMLBlobResponse            []error
	XMLPrettyResponse          []error
	XMLResponse                []error
}

func New(t *testing.T) *EchoTestContext {
	return &EchoTestContext{
		t:                              t,
		AttachmentFileRequestValues:    make([]string, 0),
		AttachmentNameRequestValues:    make([]string, 0),
		BindRequestValues:              make([]interface{}, 0),
		BlobBRequestValues:             make([][]byte, 0),
		BlobCodeRequestValues:          make([]int, 0),
		BlobContentRequestValues:       make([]string, 0),
		CookieRequestValues:            make([]string, 0),
		ErrorRequestValues:             make([]error, 0),
		FileRequestValues:              make([]string, 0),
		FormFileRequestValues:          make([]string, 0),
		FormValueRequestValues:         make([]string, 0),
		GetRequestValues:               make([]string, 0),
		HTMLBlobBRequestValues:         make([][]byte, 0),
		HTMLBlobCodeRequestValues:      make([]int, 0),
		HTMLCodeRequestValues:          make([]int, 0),
		HTMLHTMLRequestValues:          make([]string, 0),
		InlineFileRequestValues:        make([]string, 0),
		InlineNameRequestValues:        make([]string, 0),
		JSONBlobBRequestValues:         make([][]byte, 0),
		JSONBlobCodeRequestValues:      make([]int, 0),
		JSONCodeRequestValues:          make([]int, 0),
		JSONIRequestValues:             make([]interface{}, 0),
		JSONPBlobBRequestValues:        make([][]byte, 0),
		JSONPBlobCallbackRequestValues: make([]string, 0),
		JSONPBlobCodeRequestValues:     make([]int, 0),
		JSONPCallbackRequestValues:     make([]string, 0),
		JSONPCodeRequestValues:         make([]int, 0),
		JSONPIRequestValues:            make([]interface{}, 0),
		JSONPrettyCodeRequestValues:    make([]int, 0),
		JSONPrettyIndentRequestValues:  make([]string, 0),
		JSONPrettyIRequestValues:       make([]interface{}, 0),
		NoContentRequestValues:         make([]int, 0),
		ParamRequestValues:             make([]string, 0),
		QueryParamRequestValues:        make([]string, 0),
		RedirectCodeRequestValues:      make([]int, 0),
		RedirectURLRequestValues:       make([]string, 0),
		RenderCodeRequestValues:        make([]int, 0),
		RenderDataRequestValues:        make([]interface{}, 0),
		RenderNameRequestValues:        make([]string, 0),
		ResetRRequestValues:            make([]*http.Request, 0),
		ResetWRequestValues:            make([]http.ResponseWriter, 0),
		SetCookieRequestValues:         make([]*http.Cookie, 0),
		SetHandlerRequestValues:        make([]echo.HandlerFunc, 0),
		SetKeyRequestValues:            make([]string, 0),
		SetLoggerRequestValues:         make([]echo.Logger, 0),
		SetParamNamesRequestValues:     make([]string, 0),
		SetParamValuesRequestValues:    make([]string, 0),
		SetPathRequestValues:           make([]string, 0),
		SetRequestValues:               make([]*http.Request, 0),
		SetResponseValues:              make([]*echo.Response, 0),
		SetValRequestValues:            make([]interface{}, 0),
		StreamCodeRequestValues:        make([]int, 0),
		StreamContentRequestValues:     make([]string, 0),
		StreamRRequestValues:           make([]io.Reader, 0),
		StringCodeRequestValues:        make([]int, 0),
		StringSRequestValues:           make([]string, 0),
		ValidateRequestValues:          make([]interface{}, 0),
		XMLBlobBRequestValues:          make([][]byte, 0),
		XMLBlobCodeRequestValues:       make([]int, 0),
		XMLCodeRequestValues:           make([]int, 0),
		XMLIRequestValues:              make([]interface{}, 0),
		XMLPrettyCodeRequestValues:     make([]int, 0),
		XMLPrettyIndentRequestValues:   make([]string, 0),
		XMLPrettyIRequestValues:        make([]interface{}, 0),

		AttachmentResponse:         make([]error, 0),
		BindResponse:               make([]error, 0),
		BlobResponse:               make([]error, 0),
		CookieErrResponse:          make([]error, 0),
		CookieMultiResponse:        make([]*http.Cookie, 0),
		CookiesResponse:            make([][]*http.Cookie, 0),
		EchoResponse:               make([]*echo.Echo, 0),
		FileResponse:               make([]error, 0),
		FormFileErrResponse:        make([]error, 0),
		FormFileMultiResponse:      make([]*multipart.FileHeader, 0),
		FormParamsErrResponse:      make([]error, 0),
		FormParamsURLResponse:      make([]url.Values, 0),
		FormValueResponse:          make([]string, 0),
		GetResponse:                make([]interface{}, 0),
		HandlerResponse:            make([]echo.HandlerFunc, 0),
		HTMLBlobResponse:           make([]error, 0),
		HTMLResponse:               make([]error, 0),
		InlineResponse:             make([]error, 0),
		IsTLSResponse:              make([]bool, 0),
		IsWebSocketResponse:        make([]bool, 0),
		JSONBlobResponse:           make([]error, 0),
		JSONPBlobResponse:          make([]error, 0),
		JSONPResponse:              make([]error, 0),
		JSONPrettyResponse:         make([]error, 0),
		JSONResponse:               make([]error, 0),
		LoggerResponse:             make([]echo.Logger, 0),
		MultipartFormErrResponse:   make([]error, 0),
		MultipartFormMultiResponse: make([]*multipart.Form, 0),
		NoContentResponse:          make([]error, 0),
		ParamNamesResponse:         make([][]string, 0),
		ParamResponse:              make([]string, 0),
		ParamValuesResponse:        make([][]string, 0),
		PathResponse:               make([]string, 0),
		QueryParamResponse:         make([]string, 0),
		QueryParamsResponse:        make([]url.Values, 0),
		QueryStringResponse:        make([]string, 0),
		RealIPResponse:             make([]string, 0),
		RedirectResponse:           make([]error, 0),
		RenderResponse:             make([]error, 0),
		RequestResponse:            make([]*http.Request, 0),
		ResponseResponse:           make([]*echo.Response, 0),
		SchemeResponse:             make([]string, 0),
		StreamResponse:             make([]error, 0),
		StringResponse:             make([]error, 0),
		ValidateResponse:           make([]error, 0),
		XMLBlobResponse:            make([]error, 0),
		XMLPrettyResponse:          make([]error, 0),
		XMLResponse:                make([]error, 0),
	}
}

// Request returns `*http.Request`.
func (e *EchoTestContext) Request() *http.Request {
	v := e.RequestResponse[e.RequestResponseCounter]
	e.RequestResponseCounter++
	return v
}

// SetRequest sets `*http.Request`.
func (e *EchoTestContext) SetRequest(r *http.Request) {
	e.SetRequestValues = append(e.SetRequestValues, r)
}

// SetResponse sets `*Response`.
func (e *EchoTestContext) SetResponse(r *echo.Response) {
	e.SetResponseValues = append(e.SetResponseValues, r)
}

// Response returns `*Response`.
func (e *EchoTestContext) Response() *echo.Response {
	v := e.ResponseResponse[e.ResponseResponseCounter]
	e.ResponseResponseCounter++
	return v
}

// IsTLS returns true if HTTP connection is TLS otherwise false.
func (e *EchoTestContext) IsTLS() bool {
	v := e.IsTLSResponse[e.IsTLSResponseCounter]
	e.IsTLSResponseCounter++
	return v
}

// IsWebSocket returns true if HTTP connection is WebSocket otherwise false.
func (e *EchoTestContext) IsWebSocket() bool {
	v := e.IsWebSocketResponse[e.IsWebSocketResponseCounter]
	e.IsWebSocketResponseCounter++
	return v
}

// Scheme returns the HTTP protocol scheme, `http` or `https`.
func (e *EchoTestContext) Scheme() string {
	v := e.SchemeResponse[e.SchemeResponseCounter]
	e.SchemeResponseCounter++
	return v
}

// RealIP returns the client's network address based on `X-Forwarded-For`
// or `X-Real-IP` request header.
// The behavior can be configured using `Echo#IPExtractor`.
func (e *EchoTestContext) RealIP() string {
	v := e.RealIPResponse[e.RealIPResponseCounter]
	e.RealIPResponseCounter++
	return v
}

// Path returns the registered path for the handler.
func (e *EchoTestContext) Path() string {
	v := e.PathResponse[e.PathResponseCounter]
	e.PathResponseCounter++
	return v
}

// SetPath sets the registered path for the handler.
func (e *EchoTestContext) SetPath(p string) {
	e.SetPathRequestValues = append(e.SetPathRequestValues, p)

}

// Param returns path parameter by name.
func (e *EchoTestContext) Param(name string) string {
	e.ParamRequestValues = append(e.ParamRequestValues, name)

	v := e.ParamResponse[e.ParamResponseCounter]
	e.ParamResponseCounter++
	return v
}

// ParamNames returns path parameter names.
func (e *EchoTestContext) ParamNames() []string {
	v := e.ParamNamesResponse[e.ParamNamesResponseCounter]
	e.ParamNamesResponseCounter++
	return v
}

// SetParamNames sets path parameter names.
func (e *EchoTestContext) SetParamNames(names ...string) {
	for _, v := range names {
		e.SetParamNamesRequestValues = append(e.SetParamNamesRequestValues, v)
	}

}

// ParamValues returns path parameter values.
func (e *EchoTestContext) ParamValues() []string {
	v := e.ParamValuesResponse[e.ParamValuesResponseCounter]
	e.ParamValuesResponseCounter++
	return v
}

// SetParamValues sets path parameter values.
func (e *EchoTestContext) SetParamValues(values ...string) {
	for _, v := range values {
		e.SetParamValuesRequestValues = append(e.SetParamValuesRequestValues, v)
	}

}

// QueryParam returns the query param for the provided name.
func (e *EchoTestContext) QueryParam(name string) string {
	e.QueryParamRequestValues = append(e.QueryParamRequestValues, name)

	v := e.QueryParamResponse[e.QueryParamResponseCounter]
	e.QueryParamResponseCounter++
	return v
}

// QueryParams returns the query parameters as `url.Values`.
func (e *EchoTestContext) QueryParams() url.Values {
	v := e.QueryParamsResponse[e.QueryParamsResponseCounter]
	e.QueryParamsResponseCounter++
	return v
}

// QueryString returns the URL query string.
func (e *EchoTestContext) QueryString() string {
	v := e.QueryStringResponse[e.QueryStringResponseCounter]
	e.QueryStringResponseCounter++
	return v
}

// FormValue returns the form field value for the provided name.
func (e *EchoTestContext) FormValue(name string) string {
	e.FormValueRequestValues = append(e.FormValueRequestValues, name)

	v := e.FormValueResponse[e.FormValueResponseCounter]
	e.FormValueResponseCounter++
	return v
}

// FormParams returns the form parameters as `url.Values`.
func (e *EchoTestContext) FormParams() (url.Values, error) {
	v := e.FormParamsURLResponse[e.FormParamsURLResponseCounter]
	e.FormParamsURLResponseCounter++
	err := e.FormParamsErrResponse[e.FormParamsErrResponseCounter]
	e.FormParamsErrResponseCounter++
	return v, err
}

// FormFile returns the multipart form file for the provided name.
func (e *EchoTestContext) FormFile(name string) (*multipart.FileHeader, error) {
	e.FormFileRequestValues = append(e.FormFileRequestValues, name)

	m := e.FormFileMultiResponse[e.FormFileMultiResponseCounter]
	e.FormFileMultiResponseCounter++
	v := e.FormFileErrResponse[e.FormFileErrResponseCounter]
	e.FormFileErrResponseCounter++
	return m, v
}

// MultipartForm returns the multipart form.
func (e *EchoTestContext) MultipartForm() (*multipart.Form, error) {
	m := e.MultipartFormMultiResponse[e.MultipartFormMultiResponseCounter]
	e.MultipartFormMultiResponseCounter++
	v := e.MultipartFormErrResponse[e.MultipartFormErrResponseCounter]
	e.MultipartFormErrResponseCounter++
	return m, v
}

// Cookie returns the named cookie provided in the request.
func (e *EchoTestContext) Cookie(name string) (*http.Cookie, error) {
	e.CookieRequestValues = append(e.CookieRequestValues, name)

	m := e.CookieMultiResponse[e.CookieMultiResponseCounter]
	e.CookieMultiResponseCounter++
	v := e.CookieErrResponse[e.CookieErrResponseCounter]
	e.CookieErrResponseCounter++
	return m, v
}

// SetCookie adds a `Set-Cookie` header in HTTP response.
func (e *EchoTestContext) SetCookie(cookie *http.Cookie) {
	e.SetCookieRequestValues = append(e.SetCookieRequestValues, cookie)

}

// Cookies returns the HTTP cookies sent with the request.
func (e *EchoTestContext) Cookies() []*http.Cookie {
	v := e.CookiesResponse[e.CookiesResponseCounter]
	e.CookiesResponseCounter++
	return v
}

// Get retrieves data from the context.
func (e *EchoTestContext) Get(key string) interface{} {
	e.GetRequestValues = append(e.GetRequestValues, key)

	v := e.GetResponse[e.GetResponseCounter]
	e.GetResponseCounter++
	return v
}

// Set saves data in the context.
func (e *EchoTestContext) Set(key string, val interface{}) {
	e.SetKeyRequestValues = append(e.SetKeyRequestValues, key)
	e.SetValRequestValues = append(e.SetValRequestValues, val)

}

// Bind binds the request body into provided type `i`. The default binder
// does it based on Content-Type header.
func (e *EchoTestContext) Bind(i interface{}) error {
	e.BindRequestValues = append(e.BindRequestValues, i)

	v := e.BindResponse[e.BindResponseCounter]
	e.BindResponseCounter++
	return v
}

// Validate validates provided `i`. It is usually called after `Context#Bind()`.
// Validator must be registered using `Echo#Validator`.
func (e *EchoTestContext) Validate(i interface{}) error {
	e.ValidateRequestValues = append(e.ValidateRequestValues, i)

	v := e.ValidateResponse[e.ValidateResponseCounter]
	e.ValidateResponseCounter++
	return v
}

// Render renders a template with data and sends a text/html response with status
// code. Renderer must be registered using `Echo.Renderer`.
func (e *EchoTestContext) Render(code int, name string, data interface{}) error {
	e.RenderCodeRequestValues = append(e.RenderCodeRequestValues, code)
	e.RenderNameRequestValues = append(e.RenderNameRequestValues, name)
	e.RenderDataRequestValues = append(e.RenderDataRequestValues, data)

	v := e.RenderResponse[e.RenderResponseCounter]
	e.RenderResponseCounter++
	return v
}

// HTML sends an HTTP response with status code.
func (e *EchoTestContext) HTML(code int, html string) error {
	e.HTMLCodeRequestValues = append(e.HTMLCodeRequestValues, code)
	e.HTMLHTMLRequestValues = append(e.HTMLHTMLRequestValues, html)

	v := e.HTMLResponse[e.HTMLResponseCounter]
	e.HTMLResponseCounter++
	return v
}

// HTMLBlob sends an HTTP blob response with status code.
func (e *EchoTestContext) HTMLBlob(code int, b []byte) error {
	e.HTMLBlobCodeRequestValues = append(e.HTMLBlobCodeRequestValues, code)
	e.HTMLBlobBRequestValues = append(e.HTMLBlobBRequestValues, b)

	v := e.HTMLBlobResponse[e.HTMLBlobResponseCounter]
	e.HTMLBlobResponseCounter++
	return v
}

// String sends a string response with status code.
func (e *EchoTestContext) String(code int, s string) error {
	e.StringCodeRequestValues = append(e.StringCodeRequestValues, code)
	e.StringSRequestValues = append(e.StringSRequestValues, s)

	v := e.StringResponse[e.StringResponseCounter]
	e.StringResponseCounter++
	return v
}

// JSON sends a JSON response with status code.
func (e *EchoTestContext) JSON(code int, i interface{}) error {
	e.JSONCodeRequestValues = append(e.JSONCodeRequestValues, code)
	e.JSONIRequestValues = append(e.JSONIRequestValues, i)

	v := e.JSONResponse[e.JSONResponseCounter]
	e.JSONResponseCounter++
	return v
}

// JSONPretty sends a pretty-print JSON with status code.
func (e *EchoTestContext) JSONPretty(code int, i interface{}, indent string) error {
	e.JSONPrettyCodeRequestValues = append(e.JSONPrettyCodeRequestValues, code)
	e.JSONPrettyIRequestValues = append(e.JSONPrettyIRequestValues, i)
	e.JSONPrettyIndentRequestValues = append(e.JSONPrettyIndentRequestValues, indent)

	v := e.JSONPrettyResponse[e.JSONPrettyResponseCounter]
	e.JSONPrettyResponseCounter++
	return v
}

// JSONBlob sends a JSON blob response with status code.
func (e *EchoTestContext) JSONBlob(code int, b []byte) error {
	e.JSONBlobCodeRequestValues = append(e.JSONBlobCodeRequestValues, code)
	e.JSONBlobBRequestValues = append(e.JSONBlobBRequestValues, b)

	v := e.JSONBlobResponse[e.JSONBlobResponseCounter]
	e.JSONBlobResponseCounter++
	return v
}

// JSONP sends a JSONP response with status code. It uses `callback` to construct
// the JSONP payload.
func (e *EchoTestContext) JSONP(code int, callback string, i interface{}) error {
	e.JSONPCodeRequestValues = append(e.JSONPCodeRequestValues, code)
	e.JSONPCallbackRequestValues = append(e.JSONPCallbackRequestValues, callback)
	e.JSONPIRequestValues = append(e.JSONPIRequestValues, i)

	v := e.JSONPResponse[e.JSONPResponseCounter]
	e.JSONPResponseCounter++
	return v
}

// JSONPBlob sends a JSONP blob response with status code. It uses `callback`
// to construct the JSONP payload.
func (e *EchoTestContext) JSONPBlob(code int, callback string, b []byte) error {
	e.JSONPBlobCodeRequestValues = append(e.JSONPBlobCodeRequestValues, code)
	e.JSONPBlobCallbackRequestValues = append(e.JSONPBlobCallbackRequestValues, callback)
	e.JSONPBlobBRequestValues = append(e.JSONPBlobBRequestValues, b)

	v := e.JSONPBlobResponse[e.JSONPBlobResponseCounter]
	e.JSONPBlobResponseCounter++
	return v
}

// XML sends an XML response with status code.
func (e *EchoTestContext) XML(code int, i interface{}) error {
	e.XMLCodeRequestValues = append(e.XMLCodeRequestValues, code)
	e.XMLIRequestValues = append(e.XMLIRequestValues, i)

	v := e.XMLResponse[e.XMLResponseCounter]
	e.XMLResponseCounter++
	return v
}

// XMLPretty sends a pretty-print XML with status code.
func (e *EchoTestContext) XMLPretty(code int, i interface{}, indent string) error {
	e.XMLPrettyCodeRequestValues = append(e.XMLPrettyCodeRequestValues, code)
	e.XMLPrettyIRequestValues = append(e.XMLPrettyIRequestValues, i)
	e.XMLPrettyIndentRequestValues = append(e.XMLPrettyIndentRequestValues, indent)

	v := e.XMLPrettyResponse[e.XMLPrettyResponseCounter]
	e.XMLPrettyResponseCounter++
	return v
}

// XMLBlob sends an XML blob response with status code.
func (e *EchoTestContext) XMLBlob(code int, b []byte) error {
	e.XMLBlobCodeRequestValues = append(e.XMLBlobCodeRequestValues, code)
	e.XMLBlobBRequestValues = append(e.XMLBlobBRequestValues, b)

	v := e.XMLBlobResponse[e.XMLBlobResponseCounter]
	e.XMLBlobResponseCounter++
	return v
}

// Blob sends a blob response with status code and content type.
func (e *EchoTestContext) Blob(code int, contentType string, b []byte) error {
	e.BlobCodeRequestValues = append(e.BlobCodeRequestValues, code)
	e.BlobContentRequestValues = append(e.BlobContentRequestValues, contentType)
	e.BlobBRequestValues = append(e.BlobBRequestValues, b)

	v := e.BlobResponse[e.BlobResponseCounter]
	e.BlobResponseCounter++
	return v
}

// Stream sends a streaming response with status code and content type.
func (e *EchoTestContext) Stream(code int, contentType string, r io.Reader) error {
	e.StreamCodeRequestValues = append(e.StreamCodeRequestValues, code)
	e.StreamContentRequestValues = append(e.StreamContentRequestValues, contentType)
	e.StreamRRequestValues = append(e.StreamRRequestValues, r)

	v := e.StreamResponse[e.StreamResponseCounter]
	e.StreamResponseCounter++
	return v
}

// File sends a response with the content of the file.
func (e *EchoTestContext) File(file string) error {
	e.FileRequestValues = append(e.FileRequestValues, file)

	v := e.FileResponse[e.FileResponseCounter]
	e.FileResponseCounter++
	return v
}

// Attachment sends a response as attachment, prompting client to save the
// file.
func (e *EchoTestContext) Attachment(file string, name string) error {
	e.AttachmentFileRequestValues = append(e.AttachmentFileRequestValues, file)
	e.AttachmentNameRequestValues = append(e.AttachmentNameRequestValues, name)

	v := e.AttachmentResponse[e.AttachmentResponseCounter]
	e.AttachmentResponseCounter++
	return v
}

// Inline sends a response as inline, opening the file in the browser.
func (e *EchoTestContext) Inline(file string, name string) error {
	e.InlineFileRequestValues = append(e.InlineFileRequestValues, file)
	e.InlineNameRequestValues = append(e.InlineNameRequestValues, name)

	v := e.InlineResponse[e.InlineResponseCounter]
	e.InlineResponseCounter++
	return v
}

// NoContent sends a response with no body and a status code.
func (e *EchoTestContext) NoContent(code int) error {
	e.NoContentRequestValues = append(e.NoContentRequestValues, code)

	v := e.NoContentResponse[e.NoContentResponseCounter]
	e.NoContentResponseCounter++
	return v
}

// Redirect redirects the request to a provided URL with status code.
func (e *EchoTestContext) Redirect(code int, url string) error {
	e.RedirectCodeRequestValues = append(e.RedirectCodeRequestValues, code)
	e.RedirectURLRequestValues = append(e.RedirectURLRequestValues, url)

	v := e.RedirectResponse[e.RedirectResponseCounter]
	e.RedirectResponseCounter++
	return v
}

// Error invokes the registered HTTP error handler. Generally used by middleware.
func (e *EchoTestContext) Error(err error) {
	e.ErrorRequestValues = append(e.ErrorRequestValues, err)

}

// Handler returns the matched handler by router.
func (e *EchoTestContext) Handler() echo.HandlerFunc {
	v := e.HandlerResponse[e.HandlerResponseCounter]
	e.HandlerResponseCounter++
	return v
}

// SetHandler sets the matched handler by router.
func (e *EchoTestContext) SetHandler(h echo.HandlerFunc) {
	e.SetHandlerRequestValues = append(e.SetHandlerRequestValues, h)

}

// Logger returns the `Logger` instance.
func (e *EchoTestContext) Logger() echo.Logger {
	v := e.LoggerResponse[e.LoggerResponseCounter]
	e.LoggerResponseCounter++
	return v
}

// Set the logger
func (e *EchoTestContext) SetLogger(l echo.Logger) {
	e.SetLoggerRequestValues = append(e.SetLoggerRequestValues, l)

}

// Echo returns the `Echo` instance.
func (e *EchoTestContext) Echo() *echo.Echo {
	v := e.EchoResponse[e.EchoResponseCounter]
	e.EchoResponseCounter++
	return v
}

// Reset resets the context after request completes. It must be called along
// with `Echo#AcquireContext()` and `Echo#ReleaseContext()`.
// See `Echo#ServeHTTP()`
func (e *EchoTestContext) Reset(r *http.Request, w http.ResponseWriter) {
	e.ResetWRequestValues = append(e.ResetWRequestValues, w)
	e.ResetRRequestValues = append(e.ResetRRequestValues, r)

}
