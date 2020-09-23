package feed_test

import (
	"github.com/golang/mock/gomock"
	"github.com/mmcdole/gofeed"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"

	"NewsAppV2/feed"
	mock_feed "NewsAppV2/feed/mocks"
)

var _ = Describe("Feed", func() {
	var (
		mockCtrl   *gomock.Controller
		mockParser *mock_feed.MockParser
		mockGetter *mock_feed.MockGetter

		c feed.Client

		urls  []string
		feeds []*gofeed.Feed
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockParser = mock_feed.NewMockParser(mockCtrl)
		mockGetter = mock_feed.NewMockGetter(mockCtrl)

		c = feed.New(mockGetter, mockParser)

		urls = []string{
			"some-valid-url",
			"another-valid-url",
		}

		feeds = []*gofeed.Feed{
			{
				Title: "some-valid-title",
				Items: []*gofeed.Item{
					{},
				},
			},
			{
				Title: "another-valid-title",
				Items: []*gofeed.Item{
					{},
				},
			},
		}
	})

	Context("GetFeeds", func() {
		Context("When ParseUrl does not return an error", func() {
			It("should return a list of feeds", func() {
				mockGetter.
					EXPECT().
					GetUrls().
					Return(urls)

				gomock.InOrder(
					mockParser.EXPECT().ParseUrl(urls[0]).Return(feeds[0], nil),
					mockParser.EXPECT().ParseUrl(urls[1]).Return(feeds[1], nil),
				)

				actualFeeds, actualError := c.GetFeeds()

				Expect(actualError).ToNot(HaveOccurred())
				Expect(actualFeeds).To(Equal(feeds))
			})
		})

		Context("When ParseUrl does return an error", func() {
			It("returns a nil slice of feeds and an error", func() {
				err := errors.New("some-error")
				expectedError := errors.Wrap(err, "get feeds failed")

				mockGetter.
					EXPECT().
					GetUrls().
					Return(urls)

				gomock.InOrder(
					mockParser.EXPECT().ParseUrl(urls[0]).Return(feeds[0], nil),
					mockParser.EXPECT().ParseUrl(urls[1]).Return(nil, err),
				)

				actualFeeds, actualError := c.GetFeeds()

				Expect(actualError.Error()).To(Equal(expectedError.Error()))
				Expect(actualFeeds).To(BeNil())
			})
		})
	})
})
