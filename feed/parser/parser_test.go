package parser_test

import (
	"context"
	"net/url"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/mmcdole/gofeed"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"

	mock_parser "NewsAppV2/feed/mocks/parser"
	"NewsAppV2/feed/parser"
)

type httpError struct {
	err     string
	timeout bool
}

func (e *httpError) Error() string   { return e.err }
func (e *httpError) Timeout() bool   { return e.timeout }
func (e *httpError) Temporary() bool { return true }

var _ = Describe("Parser", func() {
	const validUrl = "some-valid-url"
	const invalidUrl = "some-invalid-url"

	var (
		mockCtrl         *gomock.Controller
		mockGofeedParser *mock_parser.MockGofeedParser

		p parser.Parser

		feed *gofeed.Feed
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockGofeedParser = mock_parser.NewMockGofeedParser(mockCtrl)

		p = parser.New(mockGofeedParser)

		feed = &gofeed.Feed{
			Title: "some-title",
			Items: []*gofeed.Item{
				{},
			},
		}
	})

	Context("ParseUrl", func() {
		Context("When given a valid url", func() {
			It("returns the expected feed", func() {
				mockGofeedParser.
					EXPECT().
					ParseURLWithContext(validUrl, context.Background()).
					Return(feed, nil)

				actualFeed, actualError := p.ParseUrl(validUrl)

				Expect(actualError).ToNot(HaveOccurred())
				Expect(actualFeed.Title).To(Equal("some-title"))
				Expect(actualFeed.Items).ToNot(BeEmpty())
			})
		})

		Context("When given an invalid url", func() {
			It("returns a nil feed and an error", func() {
				err := errors.New("some-error")
				expectedError := errors.Wrap(err, "parsing URL failed")

				mockGofeedParser.
					EXPECT().
					ParseURLWithContext(invalidUrl, context.Background()).
					Return(nil, err)

				actualFeed, actualError := p.ParseUrl(invalidUrl)

				Expect(actualError.Error()).To(Equal(expectedError.Error()))
				Expect(actualFeed).To(BeNil())
			})
		})
	})

	Context("ParseUrlWithContext", func() {
		var (
			ctx    context.Context
			cancel context.CancelFunc
		)

		BeforeEach(func() {
			ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
		})

		AfterEach(func() {
			cancel()
		})

		Context("When given a valid url", func() {
			Context("and a timeout context", func() {
				It("returns the expected feed", func() {
					mockGofeedParser.
						EXPECT().
						ParseURLWithContext(validUrl, ctx).
						Return(feed, nil)

					actualFeed, actualError := p.ParseUrlWithContext(validUrl, ctx)

					Expect(actualError).ToNot(HaveOccurred())
					Expect(actualFeed.Title).To(Equal("some-title"))
					Expect(actualFeed.Items).ToNot(BeEmpty())
				})
			})
		})

		Context("When given an invalid url", func() {
			Context("and a timeout context", func() {
				It("returns a nil feed and an error", func() {
					err := &url.Error{
						Err: &httpError{
							err:     "some-error",
							timeout: true,
						},
					}
					expectedError := errors.Wrap(err, "parsing URL timed out")

					mockGofeedParser.
						EXPECT().
						ParseURLWithContext(invalidUrl, ctx).
						Return(nil, err)

					actualFeed, actualError := p.ParseUrlWithContext(invalidUrl, ctx)

					Expect(actualError.Error()).To(Equal(expectedError.Error()))
					Expect(actualFeed).To(BeNil())
				})
			})
		})
	})
})
