package client

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	awscfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go/logging"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/digitalocean/godo"
	"github.com/rs/zerolog"
)

var defaultSpacesRegions = []string{"nyc3", "sfo3", "ams3", "sgp1", "fra1"}

const MaxItemsPerPage = 200

type Client struct {
	// This is a client that you need to create and initialize in Configure
	// It will be passed for each resource fetcher.
	logger           zerolog.Logger
	DoClient         *godo.Client
	Regions          []string
	SpacesRegion     string
	CredentialStatus DoCredentialStruct
	S3               *s3.Client
	Services         *Services
}

type DoCredentialStruct struct {
	Api    bool
	Spaces bool
}

type SpacesCredentialsProvider struct {
	SpacesAccessKey   string
	SpacesAccessKeyId string
}

type SpacesEndpointResolver struct{}

type DoLogger struct {
	l zerolog.Logger
}

func (s SpacesCredentialsProvider) Retrieve(ctx context.Context) (aws.Credentials, error) {
	return aws.Credentials{
		AccessKeyID:     s.SpacesAccessKeyId,
		SecretAccessKey: s.SpacesAccessKey,
		Source:          "digitalocean",
	}, nil
}

func (c *Client) WithSpacesRegion(region string) *Client {
	return &Client{
		logger:       c.Logger().With().Str("spaces_region", region).Logger(),
		DoClient:     c.DoClient,
		SpacesRegion: region,
		S3:           c.S3,
		Services:     initServices(c.DoClient),
	}
}

func (SpacesEndpointResolver) ResolveEndpoint(_, region string) (aws.Endpoint, error) {
	return aws.Endpoint{
		URL:    fmt.Sprintf("https://%s.digitaloceanspaces.com", region),
		Source: aws.EndpointSourceCustom,
	}, nil
}

func getTokenFromEnv() string {
	doToken := os.Getenv("DIGITALOCEAN_TOKEN")
	doAccessToken := os.Getenv("DIGITALOCEAN_ACCESS_TOKEN")
	if doToken != "" {
		return doToken
	}
	if doAccessToken != "" {
		return doAccessToken
	}
	return ""
}

func getSpacesTokenFromEnv() (string, string) {
	spacesAccessKey := os.Getenv("SPACES_ACCESS_KEY_ID")
	spacesSecretKey := os.Getenv("SPACES_SECRET_ACCESS_KEY")
	if spacesAccessKey == "" {
		return "", ""
	}
	if spacesSecretKey == "" {
		return "", ""
	}
	return spacesAccessKey, spacesSecretKey
}

type Services struct {
	Account        AccountService
	Cdn            CdnService
	BillingHistory BillingHistoryService
	Monitoring     MonitoringService
	Balance        BalanceService
	Certificates   CertificatesService
	Databases      DatabasesService
	Domains        DomainsService
	Droplets       DropletsService
	Firewalls      FirewallsService
	FloatingIps    FloatingIpsService
	Images         ImagesService
	Keys           KeysService
	LoadBalancers  LoadBalancersService
	Projects       ProjectsService
	Registry       RegistryService
	Sizes          SizesService
	Snapshots      SnapshotsService
	Storage        StorageService
	Vpcs           VpcsService
}

type ServicesRegionMap map[string]*Services

// ServicesManager will hold the entire map of region services
type ServicesManager struct {
	services ServicesRegionMap
}

func (s *ServicesManager) ServicesByRegion(region string) *Services {
	return s.services[region]
}

func New(ctx context.Context, logger zerolog.Logger, s specs.Source) (schema.ClientMeta, error) {
	// providerConfig := config.(*Config)
	var doSpec Spec
	if err := s.UnmarshalSpec(&doSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal gcp spec: %w", err)
	}

	if doSpec.Token == "" {
		doSpec.Token = getTokenFromEnv()
	}
	if doSpec.Token == "" {
		return nil, fmt.Errorf("missing API token")
	}

	credStatus := DoCredentialStruct{
		Api:    true,
		Spaces: true,
	}

	if doSpec.SpacesAccessKey == "" || doSpec.SpacesAccessKeyId == "" {
		doSpec.SpacesAccessKeyId, doSpec.SpacesAccessKey = getSpacesTokenFromEnv()
	}
	if doSpec.SpacesAccessKey == "" || doSpec.SpacesAccessKeyId == "" {
		credStatus.Spaces = false
	}

	awsCfg, err := awscfg.LoadDefaultConfig(context.Background(),
		awscfg.WithCredentialsProvider(SpacesCredentialsProvider{doSpec.SpacesAccessKey, doSpec.SpacesAccessKeyId}),
		awscfg.WithEndpointResolver(SpacesEndpointResolver{}),
	)

	if err != nil {
		return nil, diag.FromError(err, diag.INTERNAL)
	}

	if doSpec.SpacesDebugLogging {
		awsCfg.ClientLogMode = aws.LogRequest | aws.LogResponse | aws.LogRetries
		awsCfg.Logger = DoLogger{logger}
	}

	spacesRegions := defaultSpacesRegions
	if len(doSpec.SpacesRegions) > 0 {
		spacesRegions = doSpec.SpacesRegions
	}

	client := godo.NewFromToken(doSpec.Token)
	c := Client{
		logger:           logger,
		DoClient:         godo.NewFromToken(doSpec.Token),
		Regions:          spacesRegions,
		SpacesRegion:     "nyc3",
		S3:               s3.NewFromConfig(awsCfg),
		CredentialStatus: credStatus,
		Services:         initServices(client),
	}
	return &c, nil
}

func (a DoLogger) Logf(classification logging.Classification, format string, v ...interface{}) {
	if classification == logging.Warn {
		a.l.Warn().Msg(fmt.Sprintf(format, v...))
	} else {
		a.l.Debug().Msg(fmt.Sprintf(format, v...))
	}
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func initServices(doClient *godo.Client) *Services {
	return &Services{
		Account:        doClient.Account,
		Cdn:            doClient.CDNs,
		BillingHistory: doClient.BillingHistory,
		Monitoring:     doClient.Monitoring,
		Balance:        doClient.Balance,
		Certificates:   doClient.Certificates,
		Databases:      doClient.Databases,
		Domains:        doClient.Domains,
		Droplets:       doClient.Droplets,
		Firewalls:      doClient.Firewalls,
		FloatingIps:    doClient.FloatingIPs,
		Images:         doClient.Images,
		Keys:           doClient.Keys,
		LoadBalancers:  doClient.LoadBalancers,
		Projects:       doClient.Projects,
		Registry:       doClient.Registry,
		Sizes:          doClient.Sizes,
		Snapshots:      doClient.Snapshots,
		Storage:        doClient.Storage,
		Vpcs:           doClient.VPCs,
	}
}
