package parser_test

import (
	"github.com/golang/mock/gomock"
	"github.com/mmcdole/gofeed"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"

	mock_parser "NewsAppV2/article/mocks"
	"NewsAppV2/article/parser"
)

var _ = Describe("Parser", func() {
	const validUrl = "some-valid-url"
	const invalidUrl = "some-invalid-url"

	var (
		mockCtrl         *gomock.Controller
		mockGofeedParser *mock_parser.MockGofeedParser

		p parser.Parser
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockGofeedParser = mock_parser.NewMockGofeedParser(mockCtrl)

		p = parser.New(mockGofeedParser)
	})

	Context("ParseUrl", func() {
		Context("When given a valid url", func() {
			It("returns the expected feed", func() {
				feed := &gofeed.Feed{
					Title: "some-title",
					Items: []*gofeed.Item{
						{},
					},
				}

				mockGofeedParser.EXPECT().ParseURL(validUrl).Return(feed, nil)

				actualFeed, err := p.ParseUrl(validUrl)

				Expect(err).ToNot(HaveOccurred())
				Expect(actualFeed.Title).To(Equal("some-title"))
				Expect(actualFeed.Items).ToNot(BeEmpty())
			})
		})

		Context("When given an invalid url", func() {
			It("returns a nil feed and an error", func() {
				error := errors.New("some-error")
				expectedError := errors.Wrap(error, "parsing failed")

				mockGofeedParser.EXPECT().ParseURL(invalidUrl).Return(nil, error)

				actualFeed, actualError := p.ParseUrl(invalidUrl)

				Expect(actualError.Error()).To(Equal(expectedError.Error()))
				Expect(actualFeed).To(BeNil())
			})
		})
	})
})
