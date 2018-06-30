// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"log"
	"net/http"
	"sync"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	graceful "github.com/tylerb/graceful"

	"github.com/VivaLaPanda/isle-api/restapi/operations"
	"github.com/VivaLaPanda/isle-api/restapi/operations/comments"
	"github.com/VivaLaPanda/isle-api/restapi/operations/posts"
	"github.com/VivaLaPanda/isle-api/restapi/operations/users"

	models "github.com/VivaLaPanda/isle-api/models"
)

//go:generate swagger generate server --target .. --name IsleApi --spec ../swagger.yml --principal models.Principal

func configureFlags(api *operations.IsleAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.IsleAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.OauthSecurityAuth = func(token string, scopes []string) (*models.Principal, error) {
		// This handler is called by the runtime whenever a route needs authentication
		// against the 'OAuthSecurity' scheme.
		// It is passed a token extracted from the Authentication Bearer header, and
		// the list of scopes mentioned by the spec for this route.

		// NOTE: in this simple implementation, we do not check scopes against
		// the signed claims in the JWT token.
		// So whatever the required scope (passed a parameter by the runtime),
		// this will succeed provided we get a valid token.

		// authenticated validates a JWT token at userInfoURL
		ok, err := authenticated(token)
		if err != nil {
			return nil, errors.New(401, "error authenticate")
		}
		if !ok {
			return nil, errors.New(401, "invalid token")
		}

		// returns the authenticated principal (here just filled in with its token)
		prin := models.Principal(token)
		return &prin, nil
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	api.GetAuthCallbackHandler = operations.GetAuthCallbackHandlerFunc(func(params operations.GetAuthCallbackParams) middleware.Responder {
		// implements the callback operation
		token, err := callback(params.HTTPRequest)
		if err != nil {
			return middleware.NotImplemented("operation .GetAuthCallback error")
		}
		log.Println("Token", token)
		return operations.NewGetAuthCallbackDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(token)})
	})
	api.GetLoginHandler = operations.GetLoginHandlerFunc(func(params operations.GetLoginParams) middleware.Responder {
		// implements the login operation
		login(params.HTTPRequest)
		// TODO: Take a closer look at this
		return middleware.NotImplemented("operation .GetLogin has not yet been implemented")
	})
	api.GetPingHandler = operations.GetPingHandlerFunc(func(params operations.GetPingParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetPing has not yet been implemented")
	})
	api.CommentsGetCommentsHandler = comments.GetCommentsHandlerFunc(func(params comments.GetCommentsParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation comments.GetComments has not yet been implemented")
	})
	api.CommentsGetCommentsByIDHandler = comments.GetCommentsByIDHandlerFunc(func(params comments.GetCommentsByIDParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation comments.GetCommentsByID has not yet been implemented")
	})
	api.PostsGetPostByIDHandler = posts.GetPostByIDHandlerFunc(func(params posts.GetPostByIDParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation posts.GetPostByID has not yet been implemented")
	})
	api.PostsGetPostsHandler = posts.GetPostsHandlerFunc(func(params posts.GetPostsParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation posts.GetPosts has not yet been implemented")
	})
	api.UsersGetUserByIDHandler = users.GetUserByIDHandlerFunc(func(params users.GetUserByIDParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation users.GetUserByID has not yet been implemented")
	})
	api.UsersGetUsersHandler = users.GetUsersHandlerFunc(func(params users.GetUsersParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation users.GetUsers has not yet been implemented")
	})
	api.CommentsNewCommentHandler = comments.NewCommentHandlerFunc(func(params comments.NewCommentParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation comments.NewComment has not yet been implemented")
	})
	api.PostsNewPostHandler = posts.NewPostHandlerFunc(func(params posts.NewPostParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation posts.NewPost has not yet been implemented")
	})
	api.UsersNewUserHandler = users.NewUserHandlerFunc(func(params users.NewUserParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation users.NewUser has not yet been implemented")
	})
	api.CommentsUpdateCommentHandler = comments.UpdateCommentHandlerFunc(func(params comments.UpdateCommentParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation comments.UpdateComment has not yet been implemented")
	})
	api.PostsUpdatePostHandler = posts.UpdatePostHandlerFunc(func(params posts.UpdatePostParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation posts.UpdatePost has not yet been implemented")
	})
	api.UsersUpdateUserHandler = users.UpdateUserHandlerFunc(func(params users.UpdateUserParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation users.UpdateUser has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
}

var wG http.ResponseWriter
var wGLock sync.Mutex

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	ourFunc := func(w http.ResponseWriter, r *http.Request) {
		wGLock = sync.Mutex{}
		wG = w
		handler.ServeHTTP(w, r)
	}
	return http.HandlerFunc(ourFunc)
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
