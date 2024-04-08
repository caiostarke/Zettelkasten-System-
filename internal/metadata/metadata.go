package metadata

type Metadata struct {
	Tags       []string
	Categories []string
}

func New(tags, categories []string) Metadata {
	return Metadata{
		Tags:       tags,
		Categories: categories,
	}
}
