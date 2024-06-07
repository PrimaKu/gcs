type (
	GCSManagerInterface interface {
	}

	GCSManager struct {
		client *storage.Client
		ctx    context.Context
	}
)

func NewGCSManager(credentialsFile string) (*GCSManager, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialsFile))
	if err != nil {
		return nil, fmt.Errorf("failed to create GCS client: %v", err)
	}

	return &GCSManager{
		client: client,
		ctx:    ctx,
	}, nil
}
