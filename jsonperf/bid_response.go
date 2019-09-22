package jsonperf

import "github.com/bsm/openrtb"

func bidresponse() *openrtb.BidResponse {
	return &openrtb.BidResponse{
		ID:       "some_id",
		BidID:    "some_bid_id",
		Currency: "currency",
		SeatBid: []openrtb.SeatBid{
			{
				Seat:  "seat_1",
				Group: 1,
				Bid: []openrtb.Bid{
					{
						ID:         "bid_1_1",
						ImpID:      "imp_1_1",
						Price:      100,
						NURL:       "http://example165.com/win/12345-1?won=${AUCTION_PRICE}",
						BURL:       "http://example165.com/win/12345-1?won=${AUCTION_PRICE}",
						LURL:       "http://example165.com/win/12345-1?won=${AUCTION_PRICE}",
						CampaignID: "7771",
						AdMarkup:   "<html><title>some title</tile><body>some body</body></html>",
					},
					{
						ID:         "bid_1_2",
						ImpID:      "imp_1_2",
						Price:      200,
						NURL:       "http://example165.com/win/12345-1?won=${AUCTION_PRICE}",
						BURL:       "http://example165.com/win/12345-1?won=${AUCTION_PRICE}",
						LURL:       "http://example165.com/win/12345-1?won=${AUCTION_PRICE}",
						CampaignID: "7772",
						AdMarkup:   "<html><title>some title</tile><body>some body</body></html>",
					},
					{
						ID:         "bid_1_3",
						ImpID:      "imp_1_3",
						Price:      200,
						NURL:       "http://example165.com/win/12345-1?won=${AUCTION_PRICE}",
						BURL:       "http://example165.com/win/12345-1?won=${AUCTION_PRICE}",
						LURL:       "http://example165.com/win/12345-1?won=${AUCTION_PRICE}",
						CampaignID: "7773",
						AdMarkup:   "<html><title>some title</tile><body>some body</body></html>",
					},
				},
			},
			{
				Seat:  "seat_1",
				Group: 1,
				Bid: []openrtb.Bid{
					{
						ID:         "bid_1_1",
						ImpID:      "imp_1_1",
						Price:      100,
						NURL:       "http://example165.com/win/12345-1?won=${AUCTION_PRICE}",
						BURL:       "http://example165.com/win/12345-1?won=${AUCTION_PRICE}",
						LURL:       "http://example165.com/win/12345-1?won=${AUCTION_PRICE}",
						CampaignID: "7771",
						AdMarkup:   "<html><title>some title</tile><body>some body</body></html>",
					},
					{
						ID:         "bid_1_2",
						ImpID:      "imp_1_2",
						Price:      200,
						NURL:       "http://example165.com/win/12345-1?won=${AUCTION_PRICE}",
						BURL:       "http://example165.com/win/12345-1?won=${AUCTION_PRICE}",
						LURL:       "http://example165.com/win/12345-1?won=${AUCTION_PRICE}",
						CampaignID: "7772",
						AdMarkup:   "<html><title>some title</tile><body>some body</body></html>",
					},
					{
						ID:         "bid_1_3",
						ImpID:      "imp_1_3",
						Price:      200,
						NURL:       "http://example165.com/win/12345-1?won=${AUCTION_PRICE}",
						BURL:       "http://example165.com/win/12345-1?won=${AUCTION_PRICE}",
						LURL:       "http://example165.com/win/12345-1?won=${AUCTION_PRICE}",
						CampaignID: "7773",
						AdMarkup:   "<html><title>some title</tile><body>some body</body></html>",
					},
				},
			},
			{
				Seat:  "seat_1",
				Group: 1,
				Bid: []openrtb.Bid{
					{
						ID:         "bid_1_1",
						ImpID:      "imp_1_1",
						Price:      100,
						NURL:       "http://example165.com/win/12345-1?won=${AUCTION_PRICE}",
						BURL:       "http://example165.com/win/12345-1?won=${AUCTION_PRICE}",
						LURL:       "http://example165.com/win/12345-1?won=${AUCTION_PRICE}",
						CampaignID: "7771",
						AdMarkup:   "<html><title>some title</tile><body>some body</body></html>",
					},
					{
						ID:         "bid_1_2",
						ImpID:      "imp_1_2",
						Price:      200,
						NURL:       "http://example165.com/win/12345-1?won=${AUCTION_PRICE}",
						BURL:       "http://example165.com/win/12345-1?won=${AUCTION_PRICE}",
						LURL:       "http://example165.com/win/12345-1?won=${AUCTION_PRICE}",
						CampaignID: "7772",
						AdMarkup:   "<html><title>some title</tile><body>some body</body></html>",
					},
					{
						ID:         "bid_1_3",
						ImpID:      "imp_1_3",
						Price:      200,
						NURL:       "http://example165.com/win/12345-1?won=${AUCTION_PRICE}",
						BURL:       "http://example165.com/win/12345-1?won=${AUCTION_PRICE}",
						LURL:       "http://example165.com/win/12345-1?won=${AUCTION_PRICE}",
						CampaignID: "7773",
						AdMarkup:   "<html><title>some title</tile><body>some body</body></html>",
					},
				},
			},
		},
	}
}
