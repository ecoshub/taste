package taste

import (
	"net/http"
)

const (
	// defaultMockIP default mock ip that can be access by using 'request.RemoteAddr()'
	defaultMockIP string = "127.0.0.1"
)

// MockServer mock http server
type MockServer struct {
	mux *http.ServeMux
	ip  string
}

// NewMockServer new mock server that can easily return its handler to 'tester'
func NewMockServer(mockIPOptional ...string) *MockServer {
	mockIP := defaultMockIP
	if len(mockIPOptional) == 1 {
		mockIP = mockIPOptional[0]
	}
	return &MockServer{
		mux: http.NewServeMux(),
		ip:  mockIP,
	}
}

// Handler get mock servers handler
func (s *MockServer) Handler() http.Handler {
	return s.mux
}

// Handle add handle function to mock server
func (s *MockServer) Handle(method, path string, handlerFunc http.HandlerFunc) {
	s.mux.HandleFunc(path, methodMiddleware(method, handlerFunc))
}

// methodMiddleware http method middleware.
func methodMiddleware(method string, next http.HandlerFunc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		next(w, r)
	}
}
