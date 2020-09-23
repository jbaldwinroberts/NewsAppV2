package store_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	store2 "NewsAppV2/feed/store"
)

var _ = Describe("Store", func() {
	Context("GetUrls", func() {
		Context("When called", func() {
			It("should return a list of urls", func() {
				expectedUrls := []string{
					"http://urls.bbci.co.uk/news/uk/rss.xml",
					"http://urls.bbci.co.uk/news/technology/rss.xml",
					"http://urls.reuters.com/reuters/UKdomesticNews?format=xml",
					"http://urls.reuters.com/reuters/technologyNews?format=xml",
				}

				s := store2.New()
				actualUrls := s.GetUrls()

				Expect(actualUrls).To(Equal(expectedUrls))
			})
		})
	})
})
