// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/SpenserCai/sd-webui-discord/api/gen/restapi/operations/admin"
	"github.com/SpenserCai/sd-webui-discord/api/gen/restapi/operations/system"
	"github.com/SpenserCai/sd-webui-discord/api/gen/restapi/operations/user"
)

// NewAPIServiceAPI creates a new APIService instance
func NewAPIServiceAPI(spec *loads.Document) *APIServiceAPI {
	return &APIServiceAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		UserAuthHandler: user.AuthHandlerFunc(func(params user.AuthParams) middleware.Responder {
			return middleware.NotImplemented("operation user.Auth has not yet been implemented")
		}),
		SystemClusterHandler: system.ClusterHandlerFunc(func(params system.ClusterParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation system.Cluster has not yet been implemented")
		}),
		UserCommunityHistoryHandler: user.CommunityHistoryHandlerFunc(func(params user.CommunityHistoryParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.CommunityHistory has not yet been implemented")
		}),
		SystemDiscordServerHandler: system.DiscordServerHandlerFunc(func(params system.DiscordServerParams) middleware.Responder {
			return middleware.NotImplemented("operation system.DiscordServer has not yet been implemented")
		}),
		UserLoginHandler: user.LoginHandlerFunc(func(params user.LoginParams) middleware.Responder {
			return middleware.NotImplemented("operation user.Login has not yet been implemented")
		}),
		SystemOpenDiscordServerHandler: system.OpenDiscordServerHandlerFunc(func(params system.OpenDiscordServerParams) middleware.Responder {
			return middleware.NotImplemented("operation system.OpenDiscordServer has not yet been implemented")
		}),
		AdminSetUserEnableHandler: admin.SetUserEnableHandlerFunc(func(params admin.SetUserEnableParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation admin.SetUserEnable has not yet been implemented")
		}),
		AdminSetUserPrivateHandler: admin.SetUserPrivateHandlerFunc(func(params admin.SetUserPrivateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation admin.SetUserPrivate has not yet been implemented")
		}),
		UserUserHistoryHandler: user.UserHistoryHandlerFunc(func(params user.UserHistoryParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.UserHistory has not yet been implemented")
		}),
		UserUserInfoHandler: user.UserInfoHandlerFunc(func(params user.UserInfoParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.UserInfo has not yet been implemented")
		}),
		AdminUserListHandler: admin.UserListHandlerFunc(func(params admin.UserListParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation admin.UserList has not yet been implemented")
		}),

		// Applies when the "Authorization" header is set
		BearerAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (Bearer) Authorization from header param [Authorization] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*APIServiceAPI API */
type APIServiceAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// BearerAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Authorization provided in the header
	BearerAuth func(string) (interface{}, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// UserAuthHandler sets the operation handler for the auth operation
	UserAuthHandler user.AuthHandler
	// SystemClusterHandler sets the operation handler for the cluster operation
	SystemClusterHandler system.ClusterHandler
	// UserCommunityHistoryHandler sets the operation handler for the community history operation
	UserCommunityHistoryHandler user.CommunityHistoryHandler
	// SystemDiscordServerHandler sets the operation handler for the discord server operation
	SystemDiscordServerHandler system.DiscordServerHandler
	// UserLoginHandler sets the operation handler for the login operation
	UserLoginHandler user.LoginHandler
	// SystemOpenDiscordServerHandler sets the operation handler for the open discord server operation
	SystemOpenDiscordServerHandler system.OpenDiscordServerHandler
	// AdminSetUserEnableHandler sets the operation handler for the set user enable operation
	AdminSetUserEnableHandler admin.SetUserEnableHandler
	// AdminSetUserPrivateHandler sets the operation handler for the set user private operation
	AdminSetUserPrivateHandler admin.SetUserPrivateHandler
	// UserUserHistoryHandler sets the operation handler for the user history operation
	UserUserHistoryHandler user.UserHistoryHandler
	// UserUserInfoHandler sets the operation handler for the user info operation
	UserUserInfoHandler user.UserInfoHandler
	// AdminUserListHandler sets the operation handler for the user list operation
	AdminUserListHandler admin.UserListHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *APIServiceAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *APIServiceAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *APIServiceAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *APIServiceAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *APIServiceAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *APIServiceAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *APIServiceAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *APIServiceAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *APIServiceAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the APIServiceAPI
func (o *APIServiceAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.BearerAuth == nil {
		unregistered = append(unregistered, "AuthorizationAuth")
	}

	if o.UserAuthHandler == nil {
		unregistered = append(unregistered, "user.AuthHandler")
	}
	if o.SystemClusterHandler == nil {
		unregistered = append(unregistered, "system.ClusterHandler")
	}
	if o.UserCommunityHistoryHandler == nil {
		unregistered = append(unregistered, "user.CommunityHistoryHandler")
	}
	if o.SystemDiscordServerHandler == nil {
		unregistered = append(unregistered, "system.DiscordServerHandler")
	}
	if o.UserLoginHandler == nil {
		unregistered = append(unregistered, "user.LoginHandler")
	}
	if o.SystemOpenDiscordServerHandler == nil {
		unregistered = append(unregistered, "system.OpenDiscordServerHandler")
	}
	if o.AdminSetUserEnableHandler == nil {
		unregistered = append(unregistered, "admin.SetUserEnableHandler")
	}
	if o.AdminSetUserPrivateHandler == nil {
		unregistered = append(unregistered, "admin.SetUserPrivateHandler")
	}
	if o.UserUserHistoryHandler == nil {
		unregistered = append(unregistered, "user.UserHistoryHandler")
	}
	if o.UserUserInfoHandler == nil {
		unregistered = append(unregistered, "user.UserInfoHandler")
	}
	if o.AdminUserListHandler == nil {
		unregistered = append(unregistered, "admin.UserListHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *APIServiceAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *APIServiceAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "Bearer":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.BearerAuth)

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *APIServiceAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *APIServiceAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *APIServiceAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *APIServiceAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the API service API
func (o *APIServiceAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *APIServiceAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/auth"] = user.NewAuth(o.context, o.UserAuthHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/cluster"] = system.NewCluster(o.context, o.SystemClusterHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/community_history"] = user.NewCommunityHistory(o.context, o.UserCommunityHistoryHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/discord_server"] = system.NewDiscordServer(o.context, o.SystemDiscordServerHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/login"] = user.NewLogin(o.context, o.UserLoginHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/open_discord_server"] = system.NewOpenDiscordServer(o.context, o.SystemOpenDiscordServerHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/set_user_enable"] = admin.NewSetUserEnable(o.context, o.AdminSetUserEnableHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/set_user_private"] = admin.NewSetUserPrivate(o.context, o.AdminSetUserPrivateHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/user_history"] = user.NewUserHistory(o.context, o.UserUserHistoryHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user_info"] = user.NewUserInfo(o.context, o.UserUserInfoHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/user_list"] = admin.NewUserList(o.context, o.AdminUserListHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *APIServiceAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *APIServiceAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *APIServiceAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *APIServiceAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *APIServiceAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
