package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ultrazg/xyz/handlers"
)

func RegisterRouters(engine *gin.Engine) {
	engine.Static("/doc", "../docs")                             // 文档
	engine.POST("/sendCode", handlers.SendCode)                  // 发送验证码
	engine.POST("/login", handlers.Login)                        // 验证码登录
	engine.POST("/subscription", handlers.Subscription)          // 订阅列表
	engine.POST("/search", handlers.Search)                      // 搜索
	engine.POST("/refresh_token", handlers.RefreshToken)         // 刷新 token
	engine.POST("/episode_list", handlers.EpisodeList)           // 剧集列表
	engine.POST("/episode_detail", handlers.EpisodeDetail)       // 查询单集详情
	engine.POST("/podcast_detail", handlers.PodcastDetail)       // 查询节目详情
	engine.POST("/podcast_related", handlers.RelatedPodcastList) // 相关节目推荐
	engine.POST("/profile", handlers.Profile)                    // 查询我的信息
	engine.POST("/sticker", handlers.StickerList)                // 查询我的信息
}
