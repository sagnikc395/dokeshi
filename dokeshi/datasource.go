package dokeshi

// data source fetching interface
type DataSource interface {
	Fetch(cfg *Config) ([]string, error)
}

// create a new data source directly from git as a source of truth
func NewDataSource() DataSource {
	return &GitDataSource{}
}
