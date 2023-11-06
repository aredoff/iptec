package iptec

import (
	"sync"
)

func New() *App {
	return &App{
		plugins: map[string]Plugin{},
		cash:    NewCach(),
	}
}

type App struct {
	plugins map[string]Plugin
	cash    *cash
}

func (a *App) Use(p Plugin) {
	a.plugins[p.Name()] = p
	curatorMixin, ok := p.(СuratorMixinInterface)
	if ok {
		curatorMixin.curatorInitialization(a)
	}

	cachMixin, ok := p.(CashMixinInterface)
	if ok {
		cachMixin.cashInitialization(p.Name(), a.cash)
	}
}

func (a *App) Activate() {
	var wg sync.WaitGroup
	for _, v := range a.plugins {
		wg.Add(1)
		go pluginActivation(&wg, v)
	}
	wg.Wait()
}

func pluginActivation(wg *sync.WaitGroup, plugin Plugin) {
	defer wg.Done()
	plugin.Activate()
}

func (a *App) Collect() {
	// defer a.cash.Close()

	// for k, v := range a.plugins {
	// 	fmt.Printf("%s %s\n", k, v.Name())
	// }
}

type Plugin interface {
	Name() string
	Activate()
	// Deactivate() error
	// Execute() string
}

func (a *App) Close() {
	a.cash.Close()
}

// type App struct {
// 	mutex sync.Mutex

// 	plugins map[string]interface{}

// 	httpClient =
// 	// Route stack divided by HTTP methods
// 	stack [][]*Route
// 	// Route stack divided by HTTP methods and route prefixes
// 	treeStack []map[string][]*Route
// 	// contains the information if the route stack has been changed to build the optimized tree
// 	routesRefreshed bool
// 	// Amount of registered routes
// 	routesCount uint32
// 	// Amount of registered handlers
// 	handlersCount uint32
// 	// Ctx pool
// 	pool sync.Pool
// 	// Fasthttp server
// 	server *fasthttp.Server
// 	// App config
// 	config Config
// 	// Converts string to a byte slice
// 	getBytes func(s string) (b []byte)
// 	// Converts byte slice to a string
// 	getString func(b []byte) string
// 	// Hooks
// 	hooks *Hooks
// 	// Latest route & group
// 	latestRoute *Route
// 	// TLS handler
// 	tlsHandler *TLSHandler
// 	// Mount fields
// 	mountFields *mountFields
// 	// Indicates if the value was explicitly configured
// 	configured Config
// }

// type Config struct {
// 	// When set to true, this will spawn multiple Go processes listening on the same port.
// 	//
// 	// Default: false
// 	Prefork bool `json:"prefork"`

// 	// Enables the "Server: value" HTTP header.
// 	//
// 	// Default: ""
// 	ServerHeader string `json:"server_header"`

// 	// When set to true, the router treats "/foo" and "/foo/" as different.
// 	// By default this is disabled and both "/foo" and "/foo/" will execute the same handler.
// 	//
// 	// Default: false
// 	StrictRouting bool `json:"strict_routing"`

// 	// When set to true, enables case sensitive routing.
// 	// E.g. "/FoO" and "/foo" are treated as different routes.
// 	// By default this is disabled and both "/FoO" and "/foo" will execute the same handler.
// 	//
// 	// Default: false
// 	CaseSensitive bool `json:"case_sensitive"`

// 	// When set to true, this relinquishes the 0-allocation promise in certain
// 	// cases in order to access the handler values (e.g. request bodies) in an
// 	// immutable fashion so that these values are available even if you return
// 	// from handler.
// 	//
// 	// Default: false
// 	Immutable bool `json:"immutable"`

// 	// When set to true, converts all encoded characters in the route back
// 	// before setting the path for the context, so that the routing,
// 	// the returning of the current url from the context `ctx.Path()`
// 	// and the parameters `ctx.Params(%key%)` with decoded characters will work
// 	//
// 	// Default: false
// 	UnescapePath bool `json:"unescape_path"`

// 	// Enable or disable ETag header generation, since both weak and strong etags are generated
// 	// using the same hashing method (CRC-32). Weak ETags are the default when enabled.
// 	//
// 	// Default: false
// 	ETag bool `json:"etag"`

// 	// Max body size that the server accepts.
// 	// -1 will decline any body size
// 	//
// 	// Default: 4 * 1024 * 1024
// 	BodyLimit int `json:"body_limit"`

// 	// Maximum number of concurrent connections.
// 	//
// 	// Default: 256 * 1024
// 	Concurrency int `json:"concurrency"`

// 	// Views is the interface that wraps the Render function.
// 	//
// 	// Default: nil
// 	Views Views `json:"-"`

// 	// Views Layout is the global layout for all template render until override on Render function.
// 	//
// 	// Default: ""
// 	ViewsLayout string `json:"views_layout"`

// 	// PassLocalsToViews Enables passing of the locals set on a fiber.Ctx to the template engine
// 	//
// 	// Default: false
// 	PassLocalsToViews bool `json:"pass_locals_to_views"`

// 	// The amount of time allowed to read the full request including body.
// 	// It is reset after the request handler has returned.
// 	// The connection's read deadline is reset when the connection opens.
// 	//
// 	// Default: unlimited
// 	ReadTimeout time.Duration `json:"read_timeout"`

// 	// The maximum duration before timing out writes of the response.
// 	// It is reset after the request handler has returned.
// 	//
// 	// Default: unlimited
// 	WriteTimeout time.Duration `json:"write_timeout"`

// 	// The maximum amount of time to wait for the next request when keep-alive is enabled.
// 	// If IdleTimeout is zero, the value of ReadTimeout is used.
// 	//
// 	// Default: unlimited
// 	IdleTimeout time.Duration `json:"idle_timeout"`

// 	// Per-connection buffer size for requests' reading.
// 	// This also limits the maximum header size.
// 	// Increase this buffer if your clients send multi-KB RequestURIs
// 	// and/or multi-KB headers (for example, BIG cookies).
// 	//
// 	// Default: 4096
// 	ReadBufferSize int `json:"read_buffer_size"`

// 	// Per-connection buffer size for responses' writing.
// 	//
// 	// Default: 4096
// 	WriteBufferSize int `json:"write_buffer_size"`

// 	// CompressedFileSuffix adds suffix to the original file name and
// 	// tries saving the resulting compressed file under the new file name.
// 	//
// 	// Default: ".fiber.gz"
// 	CompressedFileSuffix string `json:"compressed_file_suffix"`

// 	// ProxyHeader will enable c.IP() to return the value of the given header key
// 	// By default c.IP() will return the Remote IP from the TCP connection
// 	// This property can be useful if you are behind a load balancer: X-Forwarded-*
// 	// NOTE: headers are easily spoofed and the detected IP addresses are unreliable.
// 	//
// 	// Default: ""
// 	ProxyHeader string `json:"proxy_header"`

// 	// GETOnly rejects all non-GET requests if set to true.
// 	// This option is useful as anti-DoS protection for servers
// 	// accepting only GET requests. The request size is limited
// 	// by ReadBufferSize if GETOnly is set.
// 	//
// 	// Default: false
// 	GETOnly bool `json:"get_only"`

// 	// ErrorHandler is executed when an error is returned from fiber.Handler.
// 	//
// 	// Default: DefaultErrorHandler
// 	ErrorHandler ErrorHandler `json:"-"`

// 	// When set to true, disables keep-alive connections.
// 	// The server will close incoming connections after sending the first response to client.
// 	//
// 	// Default: false
// 	DisableKeepalive bool `json:"disable_keepalive"`

// 	// When set to true, causes the default date header to be excluded from the response.
// 	//
// 	// Default: false
// 	DisableDefaultDate bool `json:"disable_default_date"`

// 	// When set to true, causes the default Content-Type header to be excluded from the response.
// 	//
// 	// Default: false
// 	DisableDefaultContentType bool `json:"disable_default_content_type"`

// 	// When set to true, disables header normalization.
// 	// By default all header names are normalized: conteNT-tYPE -> Content-Type.
// 	//
// 	// Default: false
// 	DisableHeaderNormalizing bool `json:"disable_header_normalizing"`

// 	// When set to true, it will not print out the «Fiber» ASCII art and listening address.
// 	//
// 	// Default: false
// 	DisableStartupMessage bool `json:"disable_startup_message"`

// 	// This function allows to setup app name for the app
// 	//
// 	// Default: nil
// 	AppName string `json:"app_name"`

// 	// StreamRequestBody enables request body streaming,
// 	// and calls the handler sooner when given body is
// 	// larger then the current limit.
// 	StreamRequestBody bool

// 	// Will not pre parse Multipart Form data if set to true.
// 	//
// 	// This option is useful for servers that desire to treat
// 	// multipart form data as a binary blob, or choose when to parse the data.
// 	//
// 	// Server pre parses multipart form data by default.
// 	DisablePreParseMultipartForm bool

// 	// Aggressively reduces memory usage at the cost of higher CPU usage
// 	// if set to true.
// 	//
// 	// Try enabling this option only if the server consumes too much memory
// 	// serving mostly idle keep-alive connections. This may reduce memory
// 	// usage by more than 50%.
// 	//
// 	// Default: false
// 	ReduceMemoryUsage bool `json:"reduce_memory_usage"`

// 	// FEATURE: v2.3.x
// 	// The router executes the same handler by default if StrictRouting or CaseSensitive is disabled.
// 	// Enabling RedirectFixedPath will change this behavior into a client redirect to the original route path.
// 	// Using the status code 301 for GET requests and 308 for all other request methods.
// 	//
// 	// Default: false
// 	// RedirectFixedPath bool

// 	// When set by an external client of Fiber it will use the provided implementation of a
// 	// JSONMarshal
// 	//
// 	// Allowing for flexibility in using another json library for encoding
// 	// Default: json.Marshal
// 	JSONEncoder utils.JSONMarshal `json:"-"`

// 	// When set by an external client of Fiber it will use the provided implementation of a
// 	// JSONUnmarshal
// 	//
// 	// Allowing for flexibility in using another json library for decoding
// 	// Default: json.Unmarshal
// 	JSONDecoder utils.JSONUnmarshal `json:"-"`

// 	// XMLEncoder set by an external client of Fiber it will use the provided implementation of a
// 	// XMLMarshal
// 	//
// 	// Allowing for flexibility in using another XML library for encoding
// 	// Default: xml.Marshal
// 	XMLEncoder utils.XMLMarshal `json:"-"`

// 	// Known networks are "tcp", "tcp4" (IPv4-only), "tcp6" (IPv6-only)
// 	// WARNING: When prefork is set to true, only "tcp4" and "tcp6" can be chose.
// 	//
// 	// Default: NetworkTCP4
// 	Network string

// 	// If you find yourself behind some sort of proxy, like a load balancer,
// 	// then certain header information may be sent to you using special X-Forwarded-* headers or the Forwarded header.
// 	// For example, the Host HTTP header is usually used to return the requested host.
// 	// But when you’re behind a proxy, the actual host may be stored in an X-Forwarded-Host header.
// 	//
// 	// If you are behind a proxy, you should enable TrustedProxyCheck to prevent header spoofing.
// 	// If you enable EnableTrustedProxyCheck and leave TrustedProxies empty Fiber will skip
// 	// all headers that could be spoofed.
// 	// If request ip in TrustedProxies whitelist then:
// 	//   1. c.Protocol() get value from X-Forwarded-Proto, X-Forwarded-Protocol, X-Forwarded-Ssl or X-Url-Scheme header
// 	//   2. c.IP() get value from ProxyHeader header.
// 	//   3. c.Hostname() get value from X-Forwarded-Host header
// 	// But if request ip NOT in Trusted Proxies whitelist then:
// 	//   1. c.Protocol() WON't get value from X-Forwarded-Proto, X-Forwarded-Protocol, X-Forwarded-Ssl or X-Url-Scheme header,
// 	//    will return https in case when tls connection is handled by the app, of http otherwise
// 	//   2. c.IP() WON'T get value from ProxyHeader header, will return RemoteIP() from fasthttp context
// 	//   3. c.Hostname() WON'T get value from X-Forwarded-Host header, fasthttp.Request.URI().Host()
// 	//    will be used to get the hostname.
// 	//
// 	// Default: false
// 	EnableTrustedProxyCheck bool `json:"enable_trusted_proxy_check"`

// 	// Read EnableTrustedProxyCheck doc.
// 	//
// 	// Default: []string
// 	TrustedProxies     []string `json:"trusted_proxies"`
// 	trustedProxiesMap  map[string]struct{}
// 	trustedProxyRanges []*net.IPNet

// 	// If set to true, c.IP() and c.IPs() will validate IP addresses before returning them.
// 	// Also, c.IP() will return only the first valid IP rather than just the raw header
// 	// WARNING: this has a performance cost associated with it.
// 	//
// 	// Default: false
// 	EnableIPValidation bool `json:"enable_ip_validation"`

// 	// If set to true, will print all routes with their method, path and handler.
// 	// Default: false
// 	EnablePrintRoutes bool `json:"enable_print_routes"`

// 	// You can define custom color scheme. They'll be used for startup message, route list and some middlewares.
// 	//
// 	// Optional. Default: DefaultColors
// 	ColorScheme Colors `json:"color_scheme"`

// 	// RequestMethods provides customizibility for HTTP methods. You can add/remove methods as you wish.
// 	//
// 	// Optional. Default: DefaultMethods
// 	RequestMethods []string

// 	// EnableSplittingOnParsers splits the query/body/header parameters by comma when it's true.
// 	// For example, you can use it to parse multiple values from a query parameter like this:
// 	//   /api?foo=bar,baz == foo[]=bar&foo[]=baz
// 	//
// 	// Optional. Default: false
// 	EnableSplittingOnParsers bool `json:"enable_splitting_on_parsers"`
// }
