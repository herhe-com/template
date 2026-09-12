package basic

// DoLogin 登录响应
type DoLogin struct {
	SessionID       string `json:"session_id"`               // 登录会话ID；同一轮登录及后续令牌轮换保持不变，不作为请求凭证
	AccessToken     string `json:"access_token"`             // 访问令牌；仅用于 Authorization 请求头
	RefreshToken    string `json:"refresh_token"`            // 刷新令牌；仅在访问令牌失效后用于 Refresh-Token 请求头
	IssuedAt        int64  `json:"issued_at"`                // 签发时间；Unix 秒（UTC）
	AccessLifetime  int64  `json:"access_lifetime"`          // 访问令牌有效时长；单位：秒；过期时间为 issued_at + access_lifetime
	RefreshLifetime int64  `json:"refresh_lifetime"`         // 刷新令牌有效时长；单位：秒；过期时间为 issued_at + refresh_lifetime
	GraceLifetime   int64  `json:"grace_lifetime,omitempty"` // 并发刷新容错时长；单位：秒；登录响应通常省略，刷新令牌对时返回
}
