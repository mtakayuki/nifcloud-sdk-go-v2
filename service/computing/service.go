// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"github.com/alice02/nifcloud-sdk-go-v2/nifcloud"
	"github.com/alice02/nifcloud-sdk-go-v2/nifcloud/signer/v2"
	"github.com/alice02/nifcloud-sdk-go-v2/private/protocol/computing"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/defaults"
)

// Computing provides the API operation methods for making requests to
// NIFCLOUD Computing. See this package's package overview docs
// for details on the service.
//
// Computing methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type Computing struct {
	*aws.Client
}

// Used for custom client initialization logic
var initClient func(*Computing)

// Used for custom request initialization logic
var initRequest func(*Computing, *aws.Request)

// Service information constants
const (
	ServiceName = "computing" // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the Computing client with a config.
//
// Example:
//     // Create a Computing client from just a config.
//     svc := computing.New(myConfig)
func New(config nifcloud.Config) *Computing {
	var signingName string
	signingRegion := config.Region

	svc := &Computing{
		Client: aws.NewClient(
			config.AWSConfig(),
			aws.Metadata{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "2016-11-15",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v2.SignRequestHandler)
	svc.Handlers.Sign.PushBackNamed(defaults.BuildContentLengthHandler)
	svc.Handlers.Build.PushBackNamed(computing.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(computing.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(computing.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(computing.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc)
	}

	return svc
}

// newRequest creates a new request for a Computing operation and runs any
// custom request initialization.
func (c *Computing) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(c, req)
	}

	return req
}
