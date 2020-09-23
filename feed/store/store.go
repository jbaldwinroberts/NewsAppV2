package store

type Store struct {
	urls []string
}

func New() Store {
	return Store{
		urls: []string{
			"http://urls.bbci.co.uk/news/uk/rss.xml",
			"http://urls.bbci.co.uk/news/technology/rss.xml",
			"http://urls.reuters.com/reuters/UKdomesticNews?format=xml",
			"http://urls.reuters.com/reuters/technologyNews?format=xml",
		},
	}
}

func (s Store) GetUrls() []string {
	return s.urls
}
