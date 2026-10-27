package middleware

import (
	"github.com/gin-gonic/gin"
)

// ClientHints 在响应中声明需要的 User-Agent Client Hints，
// 以便浏览器在后续请求（含评论提交）中带上 Sec-CH-UA-Platform-Version。
//
// 关键点：Windows 10 与 Windows 11 的 User-Agent 字符串完全相同
// （均为 Windows NT 10.0），仅靠 UA 无法区分两者。评论区通过
// Sec-CH-UA-Platform-Version（Win11 为 13.0.0+，Win10 为 10.0.0）
// 来识别 Windows 11。但该提示属于“高熵”客户端提示，浏览器默认不会发送，
// 必须由服务器通过 Accept-CH 主动申请；否则评论 UA 中不会携带 |PV: 后缀，
// 前端 parseUserAgent 只能回退到 "Windows 10"。
//
// Critical-CH 确保浏览器在“首次看到该声明”的请求上即重发并附带提示，
// 无需等待下一次请求即可生效。
func ClientHints() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Accept-CH", "Sec-CH-UA-Platform-Version")
		c.Header("Critical-CH", "Sec-CH-UA-Platform-Version")
		// 允许跨域请求携带该提示（同源部署可忽略）；若仅同源使用，可改为 (self)
		c.Header("Permissions-Policy", "ch-ua-platform-version=*")
		c.Next()
	}
}
