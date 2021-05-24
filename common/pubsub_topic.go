package common

import "nhaancs/pubsub"

const (
	TopicUserFavoritePost   pubsub.Topic = "TopicUserFavoritePost"
	TopicUserUnfavoritePost pubsub.Topic = "TopicUserUnfavoritePost"
	TopicPostDeleted        pubsub.Topic = "TopicPostDeleted"
	TopicCategoryDeleted    pubsub.Topic = "TopicCategoryDeleted"
	TopicCategoryDisabled   pubsub.Topic = "TopicCategoryDisabled"
)
