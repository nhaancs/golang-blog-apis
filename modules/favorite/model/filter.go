package favoritemodel

import "nhaancs/common"

type Filter struct {
	FakePostId string `json:"post_id,omitempty" form:"post_id"`
	FakeUserId string `json:"user_id,omitempty" form:"user_id"`
	PostId     int    `json:"-"`
	UserId     int    `json:"-"`
}

func (f *Filter) Fulfill() {
	if len(f.FakePostId) > 0 {
		if postId, err := common.FromBase58(f.FakePostId); err == nil {
			f.PostId = int(postId.GetLocalID())
		}
	}
	if len(f.FakeUserId) > 0 {
		if userId, err := common.FromBase58(f.FakeUserId); err == nil {
			f.UserId = int(userId.GetLocalID())
		}
	}
}
