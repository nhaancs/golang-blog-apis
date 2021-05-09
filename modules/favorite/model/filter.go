package favoritemodel

import "nhaancs/common"

type Filter struct {
	FakePostId string `json:"post_id,omitempty"`
	PostId     int    `json:"-"`
	UserId     int    `json:"-"`
}

func (f *Filter) Fulfill() {
	if len(f.FakePostId) > 0 {
		if postId, err := common.FromBase58(f.FakePostId); err == nil {
			f.PostId = int(postId.GetLocalID())
		}
	}
}
