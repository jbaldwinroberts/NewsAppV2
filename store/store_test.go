package store_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"NewsAppV2/store"
)

var _ = Describe("Store", func() {
	Context("Urls", func() {
		Context("When called", func() {
			It("should return a list of urls", func() {
				expectedUrls := []string{
					"http://urls.bbci.co.uk/news/uk/rss.xml",
					"http://urls.bbci.co.uk/news/technology/rss.xml",
					"http://urls.reuters.com/reuters/UKdomesticNews?format=xml",
					"http://urls.reuters.com/reuters/technologyNews?format=xml",
				}

				s := store.New()
				actualUrls := s.Urls()

				Expect(actualUrls).To(Equal(expectedUrls))
			})
		})
	})
})
