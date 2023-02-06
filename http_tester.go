package taste

import (
	"net/http"
	"testing"
)

// ServerTester core http server tester
type ServerTester struct {
	ip      string
	t       *testing.T
	handler http.Handler
	store   map[string][]byte
}

// NewHTTPServerTester creates new http server tester
// after you can set mock ip with 'SetIP' method
func NewHTTPServerTester(t *testing.T, hadler http.Handler) *ServerTester {
	return &ServerTester{
		t:       t,
		handler: hadler,
		store:   make(map[string][]byte)}
}

// SetIP set mock ip to tester
func (tt *ServerTester) SetIP(ip string) {
	tt.ip = ip
}

// Run run the test cases
func (tt *ServerTester) Run(scenario []*HTTPTestCase) {
	index := tt.hasOnlyRunMe(scenario)
	if index > 0 {
		c := scenario[index]
		tt.t.Logf("RUN [ONLY]\t%s\n", c.Name)
		tt.t.Run(c.Name, func(_ *testing.T) {
			tasteIt(tt, c)
		})
		return
	}

	for _, c := range scenario {
		tt.t.Run(c.Name, func(_ *testing.T) {
			tasteIt(tt, c)
		})
	}
}

// ResetStore clear store
func (tt *ServerTester) ResetStore() {
	tt.store = make(map[string][]byte)
}

// Store store byte slice value with a key
func (tt *ServerTester) Store(key string, body []byte) {
	tt.store[key] = body
}

// StoreKeyValue store key value
func (tt *ServerTester) StoreKeyValue(key, value string) {
	tt.Store(key, []byte(value))
}
