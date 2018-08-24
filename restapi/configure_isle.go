// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"

	"cloud.google.com/go/logging"
	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	raven "github.com/getsentry/raven-go"
	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"
	"google.golang.org/grpc"

	"github.com/VivaLaPanda/isle-api/restapi/operations"
	"github.com/VivaLaPanda/isle-api/restapi/operations/comments"
	"github.com/VivaLaPanda/isle-api/restapi/operations/posts"
	"github.com/VivaLaPanda/isle-api/restapi/operations/roles"
	"github.com/VivaLaPanda/isle-api/restapi/operations/tags"
	"github.com/VivaLaPanda/isle-api/restapi/operations/users"

	models "github.com/VivaLaPanda/isle-api/models"
)

//go:generate swagger generate server --target .. --name IsleApi --spec ../swagger.yml --principal models.User

// Global threadsafe database object
var db *dgo.Dgraph

func configureFlags(api *operations.IsleAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.IsleAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Setting up stackdriver logging
	ctx := context.Background()
	logClient, err := logging.NewClient(ctx, "isle-network")
	if err != nil {
		panic("Failed to initialize Stackdriver")
	}
	// Sets the name of the log to write to.
	logName := "isle-server"
	// Selects the log to write to.
	logger := logClient.Logger(logName)
	stdLogger := func(f string, args ...interface{}) {
		stdlg := logger.StandardLogger(logging.Info)
		stdlg.Printf(f, args)
	}
	log.Println("API logging is now being handed over to Stackdriver.")
	//api.Logger = stdLogger

	// Init raven DSN
	raven.SetDSN("https://3db2a3653d054e29a65c9d2e1fba710e:e2b5990690eb4dabb38b977d7f79af7b@sentry.isle.network/2")

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Will validate the Firebase JWT, get the token data. Data contains uid and claims
	// We validate the claims against the scope, if we succeed we use the firebase uid
	api.HasRoleAuth = func(token string, scopes []string) (*models.User, error) {
		return nil, errors.NotImplemented("oauth2 bearer auth (hasRole) has not yet been implemented")
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	api.GetPingHandler = operations.GetPingHandlerFunc(func(params operations.GetPingParams) middleware.Responder {
		return operations.NewGetPingOK()
	})
	api.CommentsGetCommentsHandler = comments.GetCommentsHandlerFunc(func(params comments.GetCommentsParams, principal *models.User) middleware.Responder {
		return middleware.NotImplemented("operation comments.GetComments has not yet been implemented")
	})
	api.CommentsGetCommentsByIDHandler = comments.GetCommentsByIDHandlerFunc(func(params comments.GetCommentsByIDParams, principal *models.User) middleware.Responder {
		return middleware.NotImplemented("operation comments.GetCommentsByID has not yet been implemented")
	})
	api.PostsGetPostByIDHandler = posts.GetPostByIDHandlerFunc(func(params posts.GetPostByIDParams, principal *models.User) middleware.Responder {
		return middleware.NotImplemented("operation posts.GetPostByID has not yet been implemented")
	})
	api.PostsGetPostsHandler = posts.GetPostsHandlerFunc(func(params posts.GetPostsParams, principal *models.User) middleware.Responder {
		return middleware.NotImplemented("operation posts.GetPosts has not yet been implemented")
	})
	api.RolesGetRolesHandler = roles.GetRolesHandlerFunc(func(params roles.GetRolesParams, principal *models.User) middleware.Responder {
		return middleware.NotImplemented("operation roles.GetRoles has not yet been implemented")
	})
	api.TagsGetTagsHandler = tags.GetTagsHandlerFunc(func(params tags.GetTagsParams, principal *models.User) middleware.Responder {
		return middleware.NotImplemented("operation tags.GetTags has not yet been implemented")
	})
	api.UsersGetUserByIDHandler = users.GetUserByIDHandlerFunc(func(params users.GetUserByIDParams, principal *models.User) middleware.Responder {
		return middleware.NotImplemented("operation users.GetUserByID has not yet been implemented")
	})
	api.UsersGetUsersHandler = users.GetUsersHandlerFunc(func(params users.GetUsersParams, principal *models.User) middleware.Responder {
		return middleware.NotImplemented("operation users.GetUsers has not yet been implemented")
	})
	api.CommentsNewCommentHandler = comments.NewCommentHandlerFunc(func(params comments.NewCommentParams, principal *models.User) middleware.Responder {
		return middleware.NotImplemented("operation comments.NewComment has not yet been implemented")
	})
	api.PostsNewPostHandler = posts.NewPostHandlerFunc(func(params posts.NewPostParams, principal *models.User) middleware.Responder {
		return middleware.NotImplemented("operation posts.NewPost has not yet been implemented")
	})
	api.TagsNewTagHandler = tags.NewTagHandlerFunc(func(params tags.NewTagParams, principal *models.User) middleware.Responder {
		return middleware.NotImplemented("operation tags.NewTag has not yet been implemented")
	})
	api.UsersNewUserHandler = users.NewUserHandlerFunc(func(params users.NewUserParams, principal *models.User) middleware.Responder {
		return middleware.NotImplemented("operation users.NewUser has not yet been implemented")
	})
	api.CommentsUpdateCommentHandler = comments.UpdateCommentHandlerFunc(func(params comments.UpdateCommentParams, principal *models.User) middleware.Responder {
		return middleware.NotImplemented("operation comments.UpdateComment has not yet been implemented")
	})
	api.PostsUpdatePostHandler = posts.UpdatePostHandlerFunc(func(params posts.UpdatePostParams, principal *models.User) middleware.Responder {
		return middleware.NotImplemented("operation posts.UpdatePost has not yet been implemented")
	})
	api.TagsUpdateTagHandler = tags.UpdateTagHandlerFunc(func(params tags.UpdateTagParams, principal *models.User) middleware.Responder {
		return middleware.NotImplemented("operation tags.UpdateTag has not yet been implemented")
	})
	api.UsersUpdateUserHandler = users.UpdateUserHandlerFunc(func(params users.UpdateUserParams, principal *models.User) middleware.Responder {
		return middleware.NotImplemented("operation users.UpdateUser has not yet been implemented")
	})

	api.ServerShutdown = func() {

	}

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
	// Set up our DGraph server
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	d, err := grpc.Dial("dgraph.isle.network:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	db = dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)
	log.Println("DB Client created. Writing schema.")
	writeSchema()
	log.Println("Schema successfully written.")
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(raven.RecoveryHandler(handler.ServeHTTP))
}
