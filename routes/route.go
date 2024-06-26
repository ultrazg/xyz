package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ultrazg/xyz/handlers"
	"github.com/ultrazg/xyz/pkg/utils"
)

func RegisterRouters(engine *gin.Engine) {
	engine.Static("/doc", "../docs") // 文档
	engine.GET("/ping", handlers.Pong)
	engine.POST("/sendCode", handlers.SendCode)                                                              // 发送验证码
	engine.POST("/login", handlers.Login)                                                                    // 验证码登录
	engine.POST("/subscription", utils.CheckAccessToken(), handlers.Subscription)                            // 订阅列表
	engine.POST("/subscription_star", utils.CheckAccessToken(), handlers.StarSubscription)                   // 星标订阅
	engine.POST("/subscription_non_starred", utils.CheckAccessToken(), handlers.NonStarredSubscription)      // 未加星标订阅
	engine.POST("/subscription_update", utils.CheckAccessToken(), handlers.UpdateStarSubscription)           // 更新星标订阅
	engine.POST("/search", utils.CheckAccessToken(), handlers.Search)                                        // 搜索
	engine.POST("/search_preset", utils.CheckAccessToken(), handlers.SearchPreset)                           // 「你可能想搜的内容」
	engine.POST("/refresh_token", handlers.RefreshToken)                                                     // 刷新 token
	engine.POST("/episode_list", utils.CheckAccessToken(), handlers.EpisodeList)                             // 剧集列表
	engine.POST("/episode_detail", utils.CheckAccessToken(), handlers.EpisodeDetail)                         // 查询单集详情
	engine.POST("/podcast_detail", utils.CheckAccessToken(), handlers.PodcastDetail)                         // 查询节目详情
	engine.POST("/podcast_related", utils.CheckAccessToken(), handlers.RelatedPodcastList)                   // 相关节目推荐
	engine.POST("/profile", utils.CheckAccessToken(), handlers.Profile)                                      // 查询我的信息
	engine.POST("/sticker", utils.CheckAccessToken(), handlers.StickerList)                                  // 查询我的贴纸
	engine.POST("/sticker_board", utils.CheckAccessToken(), handlers.StickerBoard)                           // 查询我的贴纸墙
	engine.POST("/episode_play_progress", utils.CheckAccessToken(), handlers.PlaybackProgress)               // 查询单集播放进度
	engine.POST("/comment_primary", utils.CheckAccessToken(), handlers.CommentPrimary)                       // 查询单集的评论
	engine.POST("/comment_thread", utils.CheckAccessToken(), handlers.CommentThread)                         // 查询回复评论
	engine.POST("/comment_collect_create", utils.CheckAccessToken(), handlers.CreateCommentCollect)          // 收藏评论
	engine.POST("/comment_collect_remove", utils.CheckAccessToken(), handlers.RemoveCommentCollect)          // 取消收藏评论
	engine.POST("/comment_collect_list", utils.CheckAccessToken(), handlers.CommentCollectList)              // 获取收藏评论列表
	engine.POST("/discovery", utils.CheckAccessToken(), handlers.Discovery)                                  // 首页榜单、精选节目、推荐等
	engine.POST("/episode_live_count", utils.CheckAccessToken(), handlers.Live)                              // 正在收听的人数
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
}
