package types

type Message struct {
	A UserInfo
	B MessageContent
}

type UserInfo struct {
	A string // User IP
	B string // Room Info
}

type MessageContent struct {
	A string // Content
	B int64  // Time
}
