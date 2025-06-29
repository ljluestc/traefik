package dynamic

// Middleware holds the HTTP middleware configuration.
type Middleware struct {
	// AddPrefix holds the AddPrefix middleware configuration.
	AddPrefix *AddPrefix `json:"addPrefix,omitempty" toml:"addPrefix,omitempty" yaml:"addPrefix,omitempty" export:"true"`

	// StripPrefix holds the StripPrefix middleware configuration.
	StripPrefix *StripPrefix `json:"stripPrefix,omitempty" toml:"stripPrefix,omitempty" yaml:"stripPrefix,omitempty" export:"true"`

	// StripPrefixRegex holds the StripPrefixRegex middleware configuration.
	StripPrefixRegex *StripPrefixRegex `json:"stripPrefixRegex,omitempty" toml:"stripPrefixRegex,omitempty" yaml:"stripPrefixRegex,omitempty" export:"true"`

	// ReplacePath holds the ReplacePath middleware configuration.
	ReplacePath *ReplacePath `json:"replacePath,omitempty" toml:"replacePath,omitempty" yaml:"replacePath,omitempty" export:"true"`

	// ReplacePathRegex holds the ReplacePathRegex middleware configuration.
	ReplacePathRegex *ReplacePathRegex `json:"replacePathRegex,omitempty" toml:"replacePathRegex,omitempty" yaml:"replacePathRegex,omitempty" export:"true"`

	// Chain holds the Chain middleware configuration.
	Chain *Chain `json:"chain,omitempty" toml:"chain,omitempty" yaml:"chain,omitempty" export:"true"`

	// IPAllowList holds the IPAllowList middleware configuration.
	IPAllowList *IPAllowList `json:"ipAllowList,omitempty" toml:"ipAllowList,omitempty" yaml:"ipAllowList,omitempty" export:"true"`

	// Headers holds the Headers middleware configuration.
	Headers *Headers `json:"headers,omitempty" toml:"headers,omitempty" yaml:"headers,omitempty" export:"true"`

	// Errors holds the Errors middleware configuration.
	Errors *Errors `json:"errors,omitempty" toml:"errors,omitempty" yaml:"errors,omitempty" export:"true"`

	// TraefikErrors holds the TraefikErrors middleware configuration.
	TraefikErrors *TraefikErrors `json:"traefikErrors,omitempty" toml:"traefikErrors,omitempty" yaml:"traefikErrors,omitempty" export:"true"`

	// RateLimit holds the RateLimit middleware configuration.
	RateLimit *RateLimit `json:"rateLimit,omitempty" toml:"rateLimit,omitempty" yaml:"rateLimit,omitempty" export:"true"`

	// RedirectRegex holds the RedirectRegex middleware configuration.
	RedirectRegex *RedirectRegex `json:"redirectRegex,omitempty" toml:"redirectRegex,omitempty" yaml:"redirectRegex,omitempty" export:"true"`

	// RedirectScheme holds the RedirectScheme middleware configuration.
	RedirectScheme *RedirectScheme `json:"redirectScheme,omitempty" toml:"redirectScheme,omitempty" yaml:"redirectScheme,omitempty" export:"true"`

	// BasicAuth holds the BasicAuth middleware configuration.
	BasicAuth *BasicAuth `json:"basicAuth,omitempty" toml:"basicAuth,omitempty" yaml:"basicAuth,omitempty" export:"true"`

	// DigestAuth holds the DigestAuth middleware configuration.
	DigestAuth *DigestAuth `json:"digestAuth,omitempty" toml:"digestAuth,omitempty" yaml:"digestAuth,omitempty" export:"true"`

	// ForwardAuth holds the ForwardAuth middleware configuration.
	ForwardAuth *ForwardAuth `json:"forwardAuth,omitempty" toml:"forwardAuth,omitempty" yaml:"forwardAuth,omitempty" export:"true"`

	// InFlightReq holds the InFlightReq middleware configuration.
	InFlightReq *InFlightReq `json:"inFlightReq,omitempty" toml:"inFlightReq,omitempty" yaml:"inFlightReq,omitempty" export:"true"`

	// Buffering holds the Buffering middleware configuration.
	Buffering *Buffering `json:"buffering,omitempty" toml:"buffering,omitempty" yaml:"buffering,omitempty" export:"true"`

	// CircuitBreaker holds the CircuitBreaker middleware configuration.
	CircuitBreaker *CircuitBreaker `json:"circuitBreaker,omitempty" toml:"circuitBreaker,omitempty" yaml:"circuitBreaker,omitempty" export:"true"`

	// Compress holds the Compress middleware configuration.
	Compress *Compress `json:"compress,omitempty" toml:"compress,omitempty" yaml:"compress,omitempty" export:"true"`

	// PassTLSClientCert holds the PassTLSClientCert middleware configuration.
	PassTLSClientCert *PassTLSClientCert `json:"passTLSClientCert,omitempty" toml:"passTLSClientCert,omitempty" yaml:"passTLSClientCert,omitempty" export:"true"`

	// Retry holds the Retry middleware configuration.
	Retry *Retry `json:"retry,omitempty" toml:"retry,omitempty" yaml:"retry,omitempty" export:"true"`

	// ContentType holds the ContentType middleware configuration.
	ContentType *ContentType `json:"contentType,omitempty" toml:"contentType,omitempty" yaml:"contentType,omitempty" export:"true"`

	// Plugin holds the Plugin middleware configuration.
	Plugin map[string]interface{} `json:"plugin,omitempty" toml:"plugin,omitempty" yaml:"plugin,omitempty" export:"true"`
}

// AddPrefix holds the AddPrefix middleware configuration.
type AddPrefix struct {
	// Prefix is the string to add before the request path.
	Prefix string `json:"prefix,omitempty" toml:"prefix,omitempty" yaml:"prefix,omitempty" export:"true"`
}

// StripPrefix holds the StripPrefix middleware configuration.
type StripPrefix struct {
	// Prefixes defines the path prefixes to strip from the request URL.
	Prefixes []string `json:"prefixes,omitempty" toml:"prefixes,omitempty" yaml:"prefixes,omitempty" export:"true"`
}

// StripPrefixRegex holds the StripPrefixRegex middleware configuration.
type StripPrefixRegex struct {
	// Regex defines the regular expressions to match the path.
	Regex []string `json:"regex,omitempty" toml:"regex,omitempty" yaml:"regex,omitempty" export:"true"`
}

// ReplacePath holds the ReplacePath middleware configuration.
type ReplacePath struct {
	// Path defines the path to use as replacement.
	Path string `json:"path,omitempty" toml:"path,omitempty" yaml:"path,omitempty" export:"true"`
}

// ReplacePathRegex holds the ReplacePathRegex middleware configuration.
type ReplacePathRegex struct {
	// Regex defines the regular expression to match the path.
	Regex string `json:"regex,omitempty" toml:"regex,omitempty" yaml:"regex,omitempty" export:"true"`

	// Replacement defines the replacement path.
	Replacement string `json:"replacement,omitempty" toml:"replacement,omitempty" yaml:"replacement,omitempty" export:"true"`
}

// Chain holds the Chain middleware configuration.
type Chain struct {
	// Middlewares defines the list of middleware names.
	Middlewares []string `json:"middlewares,omitempty" toml:"middlewares,omitempty" yaml:"middlewares,omitempty" export:"true"`
}

// IPAllowList holds the IPAllowList middleware configuration.
type IPAllowList struct {
	// SourceRange defines the allowed IPs (CIDR notation).
	SourceRange []string `json:"sourceRange,omitempty" toml:"sourceRange,omitempty" yaml:"sourceRange,omitempty" export:"true"`
}

// Headers holds the Headers middleware configuration.
type Headers struct {
	// CustomRequestHeaders defines custom headers to be set on request.
	CustomRequestHeaders map[string]string `json:"customRequestHeaders,omitempty" toml:"customRequestHeaders,omitempty" yaml:"customRequestHeaders,omitempty" export:"true"`

	// CustomResponseHeaders defines custom headers to be set on response.
	CustomResponseHeaders map[string]string `json:"customResponseHeaders,omitempty" toml:"customResponseHeaders,omitempty" yaml:"customResponseHeaders,omitempty" export:"true"`
}

// Errors holds the Errors middleware configuration.
type Errors struct {
	// Status defines the HTTP status to match.
	Status []string `json:"status,omitempty" toml:"status,omitempty" yaml:"status,omitempty" export:"true"`

	// Service defines the service name to use.
	Service string `json:"service,omitempty" toml:"service,omitempty" yaml:"service,omitempty" export:"true"`

	// Query defines the query to use.
	Query string `json:"query,omitempty" toml:"query,omitempty" yaml:"query,omitempty" export:"true"`
}

// TraefikErrors holds the TraefikErrors middleware configuration.
type TraefikErrors struct {
	// HomePage defines the path to a custom home page.
	HomePage string `json:"homePage,omitempty" toml:"homePage,omitempty" yaml:"homePage,omitempty" export:"true"`
}

// RateLimit holds the RateLimit middleware configuration.
type RateLimit struct {
	// Average defines the maximum rate, by default in requests/s.
	Average int64 `json:"average,omitempty" toml:"average,omitempty" yaml:"average,omitempty" export:"true"`

	// Burst defines the maximum allowed burst size.
	Burst int64 `json:"burst,omitempty" toml:"burst,omitempty" yaml:"burst,omitempty" export:"true"`
}

// RedirectRegex holds the RedirectRegex middleware configuration.
type RedirectRegex struct {
	// Regex defines the regular expression to match the request path.
	Regex string `json:"regex,omitempty" toml:"regex,omitempty" yaml:"regex,omitempty" export:"true"`

	// Replacement defines the replacement path.
	Replacement string `json:"replacement,omitempty" toml:"replacement,omitempty" yaml:"replacement,omitempty" export:"true"`

	// Permanent defines whether the redirect is permanent (301).
	Permanent bool `json:"permanent,omitempty" toml:"permanent,omitempty" yaml:"permanent,omitempty" export:"true"`
}

// RedirectScheme holds the RedirectScheme middleware configuration.
type RedirectScheme struct {
	// Scheme defines the scheme to use in the redirect.
	Scheme string `json:"scheme,omitempty" toml:"scheme,omitempty" yaml:"scheme,omitempty" export:"true"`

	// Port defines the port to use in the redirect.
	Port string `json:"port,omitempty" toml:"port,omitempty" yaml:"port,omitempty" export:"true"`

	// Permanent defines whether the redirect is permanent (301).
	Permanent bool `json:"permanent,omitempty" toml:"permanent,omitempty" yaml:"permanent,omitempty" export:"true"`
}

// BasicAuth holds the BasicAuth middleware configuration.
type BasicAuth struct {
	// Users defines the authorized users.
	Users []string `json:"users,omitempty" toml:"users,omitempty" yaml:"users,omitempty" export:"true"`

	// Realm defines the realm to display.
	Realm string `json:"realm,omitempty" toml:"realm,omitempty" yaml:"realm,omitempty" export:"true"`
}

// DigestAuth holds the DigestAuth middleware configuration.
type DigestAuth struct {
	// Users defines the authorized users.
	Users []string `json:"users,omitempty" toml:"users,omitempty" yaml:"users,omitempty" export:"true"`

	// Realm defines the realm to display.
	Realm string `json:"realm,omitempty" toml:"realm,omitempty" yaml:"realm,omitempty" export:"true"`
}

// ForwardAuth holds the ForwardAuth middleware configuration.
type ForwardAuth struct {
	// Address defines the authentication server address.
	Address string `json:"address,omitempty" toml:"address,omitempty" yaml:"address,omitempty" export:"true"`
}

// InFlightReq holds the InFlightReq middleware configuration.
type InFlightReq struct {
	// Amount defines the maximum amount of allowed simultaneous in-flight requests.
	Amount int64 `json:"amount,omitempty" toml:"amount,omitempty" yaml:"amount,omitempty" export:"true"`
}

// Buffering holds the Buffering middleware configuration.
type Buffering struct {
	// MaxRequestBodyBytes defines the maximum allowed body size.
	MaxRequestBodyBytes int64 `json:"maxRequestBodyBytes,omitempty" toml:"maxRequestBodyBytes,omitempty" yaml:"maxRequestBodyBytes,omitempty" export:"true"`
}

// CircuitBreaker holds the CircuitBreaker middleware configuration.
type CircuitBreaker struct {
	// Expression defines the expression to determine when the circuit trips.
	Expression string `json:"expression,omitempty" toml:"expression,omitempty" yaml:"expression,omitempty" export:"true"`
}

// Compress holds the Compress middleware configuration.
type Compress struct {
	// Exclude defines the list of content types to exclude from compression.
	Exclude []string `json:"exclude,omitempty" toml:"exclude,omitempty" yaml:"exclude,omitempty" export:"true"`
}

// PassTLSClientCert holds the PassTLSClientCert middleware configuration.
type PassTLSClientCert struct {
	// PEM defines whether to pass the certificate in PEM format.
	PEM bool `json:"pem,omitempty" toml:"pem,omitempty" yaml:"pem,omitempty" export:"true"`
}

// Retry holds the Retry middleware configuration.
type Retry struct {
	// Attempts defines the number of attempts.
	Attempts int `json:"attempts,omitempty" toml:"attempts,omitempty" yaml:"attempts,omitempty" export:"true"`
}

// ContentType holds the ContentType middleware configuration.
type ContentType struct {
	// AutoDetect defines whether to automatically detect the content type.
	AutoDetect bool `json:"autoDetect,omitempty" toml:"autoDetect,omitempty" yaml:"autoDetect,omitempty" export:"true"`
}

// HTTPMiddleware is an alias for Middleware
type HTTPMiddleware = Middleware
