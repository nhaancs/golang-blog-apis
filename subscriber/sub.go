package subscriber

import (
	"context"
	"log"
	"nhaancs/common"
	"nhaancs/component"
	"nhaancs/component/asyncjob"
	"nhaancs/pubsub"
	"nhaancs/socketengine"
)

type subscribedJob struct {
	Title   string
	Handler func(ctx context.Context, message *pubsub.Message) error
}

func NewEngine(appContext component.AppContext, rtEngine socketengine.RealtimeEngine) *subscriberEngine {
	return &subscriberEngine{appCtx: appContext, rtEngine: rtEngine}
}

type subscriberEngine struct {
	appCtx   component.AppContext
	rtEngine socketengine.RealtimeEngine
}

func (engine *subscriberEngine) Start() error {
	engine.subscribeToATopic(
		common.TopicUserFavoritePost,
		true,
		RunIncreaseFavoriteCountAfterUserFavoritesAPost(engine.appCtx),
		EmitRealtimeAfterUserFavoritesAPost(engine.appCtx, engine.rtEngine),
	)

	engine.subscribeToATopic(
		common.TopicUserUnfavoritePost,
		true,
		RunDecreaseUnfavoriteCountAfterUserFavoritesAPost(engine.appCtx),
	)

	engine.subscribeToATopic(
		common.TopicCategoryDisabled,
		true,
		RunDisableAllPostsInACategoryGetDisabled(engine.appCtx),
	)

	engine.subscribeToATopic(
		common.TopicCategoryDeleted,
		true,
		RunDeleteAllPostsInACategoryGetDeleted(engine.appCtx),
	)

	return nil
}

func (engine *subscriberEngine) subscribeToATopic(topic pubsub.Topic, isConcurrent bool, subscribedJobs ...subscribedJob) error {
	// Subscribe to a topic and get back a channel to listen
	c, _ := engine.appCtx.GetPubsub().Subscribe(context.Background(), topic)
	// Helper function: convert a subscribedJob + pubsub.Message into an asyncjob JobHandler
	getJobHandler := func(job *subscribedJob, message *pubsub.Message) asyncjob.JobHandler {
		return func(ctx context.Context) error {
			log.Println("Running job: ", job.Title, ". Value: ", message.Data())
			return job.Handler(ctx, message)
		}
	}

	go func() {
		for { // listen to a returned channel forever
			// each time the topic publish a new message, we will get a new message from return channel
			msg := <-c
			asyncJobs := make([]asyncjob.Job, len(subscribedJobs))
			for i := range subscribedJobs {
				// combine all subscribed jobs with the new message
				// converted into asyncjob.JobHandler
				jobHandler := getJobHandler(&subscribedJobs[i], msg)
				asyncJobs[i] = asyncjob.NewJob(jobHandler)
			}

			group := asyncjob.NewGroup(isConcurrent, asyncJobs...)
			if err := group.Run(context.Background()); err != nil {
				log.Println(err)
			}
		}
	}()

	return nil
}
