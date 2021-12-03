package helix

type Tag struct {
	IsAuto                   bool              `json:"is_auto"`
	LocalizationDescriptions map[string]string `json:"localization_descriptions"`
	LocalizationNames        map[string]string `json:"localization_names"`
	TagID                    string            `json:"tag_id"`
}

type ManyTags struct {
	Tags       []Tag      `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type TagsResponse struct {
	ResponseCommon
	Data ManyTags
}

type TagsParams struct {
	After string `query:"after"`
	First int    `query:"first,20"` // Limit 100
}

func (c *Client) GetAllStreamTags(params *TagsParams) (*TagsResponse, error) {
	resp, err := c.get("/tags/streams", &ManyTags{}, params)
	if err != nil {
		return nil, err
	}

	streams := &TagsResponse{}
	resp.HydrateResponseCommon(&streams.ResponseCommon)
	streams.Data.Tags = resp.Data.(*ManyTags).Tags
	streams.Data.Pagination = resp.Data.(*ManyTags).Pagination

	return streams, nil
}
