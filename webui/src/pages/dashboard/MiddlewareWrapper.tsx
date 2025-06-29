  getMiddlewareTypeDescription = (type?: string): string => {
    if (!type) return ''

    switch (type) {
      case MiddlewareType.AddPrefix:
        return 'Adds a path prefix to the request URL'
      case MiddlewareType.BasicAuth:
        return 'Defines basic authentication mechanism'
      case MiddlewareType.Buffering:
        return 'Buffers the request/response'
      case MiddlewareType.Chain:
        return 'Combines multiple pieces of middleware'
      case MiddlewareType.CircuitBreaker:
        return 'Prevents calling unhealthy services'
      case MiddlewareType.Compress:
        return 'Compresses responses'
      case MiddlewareType.ContentType:
        return 'Auto detects the content type'
      case MiddlewareType.DigestAuth:
        return 'Defines digest authentication mechanism'
      case MiddlewareType.Errors:
        return 'Defines custom error pages'
      case MiddlewareType.TraefikErrors:
        return 'Defines custom error pages for Traefik-generated errors'
      case MiddlewareType.ForwardAuth:
        return 'Delegates request authentication to a Service'
      case MiddlewareType.Headers:
        return 'Adds/updates headers to the request/response'
      case MiddlewareType.IPAllowList:
        return 'Limits allowed client IPs'
      case MiddlewareType.InFlightReq:
        return 'Limits the number of simultaneous connections'
      case MiddlewareType.PassTLSClientCert:
        return 'Adds Client Certificates in a Header'
      case MiddlewareType.RateLimit:
        return 'Limits the call frequency'
      case MiddlewareType.RedirectRegex:
        return 'Redirects request using regex matching and replacement'
      case MiddlewareType.RedirectScheme:
        return 'Redirects request if scheme doesn\'t match'
      case MiddlewareType.ReplacePath:
        return 'Updates the path of the request URL'
      case MiddlewareType.ReplacePathRegex:
        return 'Updates the path of the request URL using regex'
      case MiddlewareType.Retry:
        return 'Automatically retries in case of error'
      case MiddlewareType.StripPrefix:
        return 'Removes a path prefix from the request URL'
      case MiddlewareType.StripPrefixRegex:
        return 'Removes a path prefix from the request URL using regex'
      default:
        if (type.startsWith('plugin-')) {
          // Plugins can have any type name
          return 'Custom plugin middleware'
        }
        return 'Plugin middleware'
    }
  }
