package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	docs "github.com/ultrazg/xyz/doc"
	"github.com/ultrazg/xyz/handlers"
	"github.com/ultrazg/xyz/utils"
)

func RegisterRouters(engine *gin.Engine) {
	engine.GET("/docs/*filepath", func(context *gin.Context) {
		server := http.FileServer(http.FS(docs.Fs))
		server.ServeHTTP(context.Writer, context.Request)
	})
	engine.GET("/ping", handlers.Pong)
	engine.POST("/sendCode", handlers.SendCode)                                                              // 发送验证码
	engine.POST("/login", handlers.Login)                                                                    // 验证码登录
	engine.POST("/subscription", utils.CheckAccessToken(), handlers.Subscription)                            // 订阅列表
	engine.POST("/subscription_update", utils.CheckAccessToken(), handlers.SubscriptionUpdate)               // 更新订阅
	engine.POST("/subscription_star", utils.CheckAccessToken(), handlers.StarSubscription)                   // 星标订阅
	engine.POST("/subscription_non_starred", utils.CheckAccessToken(), handlers.NonStarredSubscription)      // 未加星标订阅
	engine.POST("/subscription_star_update", utils.CheckAccessToken(), handlers.UpdateStarSubscription)      // 更新星标订阅
	engine.POST("/search", utils.CheckAccessToken(), handlers.Search)                                        // 搜索
	engine.POST("/search_preset", utils.CheckAccessToken(), handlers.SearchPreset)                           // 「你可能想搜的内容」
	engine.POST("/refresh_token", handlers.RefreshToken)                                                     // 刷新 token
	engine.POST("/episode_list", utils.CheckAccessToken(), handlers.EpisodeList)                             // 剧集列表
	engine.POST("/episode_list_by_filter", utils.CheckAccessToken(), handlers.EpisodeListByFilter)           // 节目内「最受欢迎」单集列表
	engine.POST("/episode_detail", utils.CheckAccessToken(), handlers.EpisodeDetail)                         // 查询单集详情
	engine.POST("/podcast_detail", utils.CheckAccessToken(), handlers.PodcastDetail)                         // 查询节目详情
	engine.POST("/podcast_get_info", utils.CheckAccessToken(), handlers.PodcastGetInfo)                      // 获取节目主体信息
	engine.POST("/podcast_honor_list", utils.CheckAccessToken(), handlers.PodcastHonorList)                  // 获取节目荣誉墙
	engine.POST("/podcast_related", utils.CheckAccessToken(), handlers.RelatedPodcastList)                   // 相关节目推荐
	engine.POST("/podcast_bulletin", utils.CheckAccessToken(), handlers.PodcastBulletin)                     // 获取节目公告
	engine.POST("/profile", utils.CheckAccessToken(), handlers.Profile)                                      // 根据 uid 查询用户信息
	engine.POST("/sticker", utils.CheckAccessToken(), handlers.StickerList)                                  // 根据 uid 查询已获得的贴纸
	engine.POST("/sticker_board", utils.CheckAccessToken(), handlers.StickerBoard)                           // 查询我的贴纸墙
	engine.POST("/episode_play_progress", utils.CheckAccessToken(), handlers.PlaybackProgress)               // 查询单集播放进度
	engine.POST("/episode_play_progress_update", utils.CheckAccessToken(), handlers.UpdatePlaybackProgress)  // 更新单集播放进度
	engine.POST("/comment_primary", utils.CheckAccessToken(), handlers.CommentPrimary)                       // 查询单集的评论
	engine.POST("/comment_thread", utils.CheckAccessToken(), handlers.CommentThread)                         // 查询回复评论
	engine.POST("/comment_collect_create", utils.CheckAccessToken(), handlers.CreateCommentCollect)          // 收藏评论
	engine.POST("/comment_collect_remove", utils.CheckAccessToken(), handlers.RemoveCommentCollect)          // 取消收藏评论
	engine.POST("/comment_collect_list", utils.CheckAccessToken(), handlers.CommentCollectList)              // 获取收藏评论列表
	engine.POST("/comment_like_update", utils.CheckAccessToken(), handlers.CommentLikeUpdate)                // 点赞/取消点赞评论
	engine.POST("/discovery", utils.CheckAccessToken(), handlers.Discovery)                                  // 首页榜单、精选节目、推荐等
	engine.POST("/refresh_episode_recommend", utils.CheckAccessToken(), handlers.RefreshEpisodeRecommend)    // 首页大家都在听-刷新推荐
	engine.POST("/episode_live_count", utils.CheckAccessToken(), handlers.Live)                              // 正在收听的人数
	engine.POST("/live_stats_report", utils.CheckAccessToken(), handlers.LiveStatsReport)                    // 上报播放状态
	engine.POST("/episode_clap", utils.CheckAccessToken(), handlers.Clap)                                    // 精彩时间点
	engine.POST("/episode_clap_create", utils.CheckAccessToken(), handlers.CreateClap)                       // 标记精彩时间点
	engine.POST("/inbox_list", utils.CheckAccessToken(), handlers.InboxList)                                 // 订阅更新列表
	engine.POST("/category_list", utils.CheckAccessToken(), handlers.CategoryList)                           // 全部分类
	engine.POST("/category_list_tab", utils.CheckAccessToken(), handlers.CategoryListTabById)                // 获取分类下的标签
	engine.POST("/category_podcast_list", utils.CheckAccessToken(), handlers.CategoryPodcastListByTab)       // 根据标签获取分类下的节目列表
	engine.POST("/favorite_episode_update", utils.CheckAccessToken(), handlers.UpdateEpisodeFavorite)        // 更新收藏单集
	engine.POST("/favorite_episode_list", utils.CheckAccessToken(), handlers.FavoriteEpisodeList)            // 获取收藏单集列表
	engine.POST("/episode_played_history_list", utils.CheckAccessToken(), handlers.EpisodePlayedHistoryList) // 收听历史
	engine.POST("/unread_count", utils.CheckAccessToken(), handlers.UnreadCount)                             // 未读消息
	engine.POST("/user_stats", utils.CheckAccessToken(), handlers.GetUserStats)                              // 用户统计数据
	engine.POST("/get_profile", utils.CheckAccessToken(), handlers.GetProfileByUid)                          // 根据 uid 查询用户信息
	engine.POST("/mileage_get", utils.CheckAccessToken(), handlers.GetMileage)                               // 获取收听数据概览
	engine.POST("/mileage_list", utils.CheckAccessToken(), handlers.GetMileageList)                          // 获取收听排行
	engine.POST("/mileage_update", utils.CheckAccessToken(), handlers.UpdateMileage)                         // 更新收听数据概览
	engine.POST("/played_list", utils.CheckAccessToken(), handlers.PlayedList)                               // 获取收听历史记录
	engine.POST("/pick_list_recent", utils.CheckAccessToken(), handlers.PickListRecent)                      // 获取「用户的喜欢」部分片段
	engine.POST("/pick_list_history", utils.CheckAccessToken(), handlers.PickListHistory)                    // 获取「用户的喜欢」全部内容
	engine.POST("/owned_podcasts", utils.CheckAccessToken(), handlers.OwnedPodcastsList)                     // 获取用户创建的播客
	engine.POST("/top_list", utils.CheckAccessToken(), handlers.GetTopList)                                  // 获取榜单
	engine.POST("/following_list", utils.CheckAccessToken(), handlers.FollowingList)                         // 获取「我」关注的人
	engine.POST("/follower_list", utils.CheckAccessToken(), handlers.FollowerList)                           // 获取关注「我」的人
	engine.POST("/blocked_user_lists", utils.CheckAccessToken(), handlers.BlockedUserLists)                  // 查询黑名单列表
	engine.POST("/blocked_user_create", utils.CheckAccessToken(), handlers.BlockedUserCreate)                // 将用户加入黑名单
	engine.POST("/blocked_user_remove", utils.CheckAccessToken(), handlers.BlockedUserRemove)                // 将用户移出黑名单
	engine.POST("/user_preference_get", utils.CheckAccessToken(), handlers.UserPreferenceGet)                // 获取用户偏好设置
	engine.POST("/user_preference_update", utils.CheckAccessToken(), handlers.UserPreferenceUpdate)          // 更新用户偏好设置
	engine.POST("/relation_update", utils.CheckAccessToken(), handlers.RelationUpdate)                       // 关注/取关用户
}
