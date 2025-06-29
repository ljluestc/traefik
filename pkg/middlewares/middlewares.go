package middlewares

// RegisterFactories registers the auth middlewares factories.
func RegisterFactories(factoriesRegistry *middlewares.FactoriesRegistry) {
	factoriesRegistry.Register(addprefix.New)
	factoriesRegistry.Register(auth.NewBasic)
	factoriesRegistry.Register(auth.NewDigest)
	factoriesRegistry.Register(auth.NewForward)
	factoriesRegistry.Register(buffering.New)
	factoriesRegistry.Register(chain.New)
	factoriesRegistry.Register(circuitbreaker.New)
	factoriesRegistry.Register(compress.New)
	factoriesRegistry.Register(contenttype.New)
	factoriesRegistry.Register(customerrors.New)
	factoriesRegistry.Register(headers.New)
	factoriesRegistry.Register(inflightreq.New)
	factoriesRegistry.Register(ipallowlist.New)
	factoriesRegistry.Register(passtlsclientcert.New)
	factoriesRegistry.Register(ratelimiter.New)
	factoriesRegistry.Register(redirect.NewRedirectRegex)
	factoriesRegistry.Register(redirect.NewRedirectScheme)
	factoriesRegistry.Register(replacepath.New)
	factoriesRegistry.Register(replacepathregex.New)
	factoriesRegistry.Register(retry.New)
	factoriesRegistry.Register(stripprefix.New)
	factoriesRegistry.Register(stripprefixregex.New)
	factoriesRegistry.Register(traefik_errors.New)
}
