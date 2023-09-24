package request

import "strings"

type GetDocumentsRequest struct {
	Tags string `json:"tags" query:"tags"`
}

func (r *GetDocumentsRequest) ParseTags() []string {
	tags := strings.Split(r.Tags, ",")

	return tags
}
