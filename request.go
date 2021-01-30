package simpleclient

type Method string

const (
	GET  Method = "GET"
	POST        = "POST"
)

type Request struct {
	url         string
	method      Method
	response    interface{}
	payload     []byte
	headers     map[string]string
	queryParams map[string]string
}

func (r *Request) addHeader(key, value string) {
	r.headers[key] = value
}

type Option func(*Request)

func NewRequest(url string, method Method, response interface{}, options ...Option) *Request {
	request := Request{
		url:         url,
		method:      method,
		response:    response,
		payload:     nil,
		headers:     map[string]string{},
		queryParams: map[string]string{},
	}

	for _, option := range options {
		option(&request)
	}

	return &request
}

func WithHeader(key, value string) Option {
	return func(request *Request) {
		request.headers[key] = value
	}
}

func WithHeaders(headers map[string]string) Option {
	return func(request *Request) {
		for key, value := range headers {
			request.headers[key] = value
		}
	}
}

func WithQueryParam(key, value string) Option {
	return func(request *Request) {
		request.queryParams[key] = value
	}
}

func WithQueryParams(headers map[string]string) Option {
	return func(request *Request) {
		for key, value := range headers {
			request.queryParams[key] = value
		}
	}
}

func WithPayload(payload []byte) Option {
	return func(request *Request) {
		request.payload = payload
	}
}
