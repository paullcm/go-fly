package models

type Message struct {
	Model
	KefuId    string `json:"kefu_id"`
	VisitorId string `json:"visitor_id"`
	Content   string `json:"content"`
	MesType   string `json:"mes_type"`
	Status    string `json:"status"`
}

func CreateMessage(kefu_id string, visitor_id string, content string, mes_type string) {
	v := &Message{
		KefuId:    kefu_id,
		VisitorId: visitor_id,
		Content:   content,
		MesType:   mes_type,
		Status:    "unread",
	}
	DB.Create(v)
}
func FindMessageByVisitorId(visitor_id string) []Message {
	var messages []Message
	DB.Where("visitor_id=?", visitor_id).Order("id asc").Find(&messages)
	return messages
}

//修改消息状态
func ReadMessageByVisitorId(visitor_id string) {
	message := &Message{
		Status: "read",
	}
	DB.Model(&message).Where("visitor_id=?", visitor_id).Update(message)
}

//获取未读数
func FindUnreadMessageNumByVisitorId(visitor_id string) uint {
	var count uint
	DB.Where("visitor_id=? and status=?", visitor_id, "unread").Count(&count)
	return count
}

//查询最后一条消息
func FindLastMessage(visitorIds []string) []Message {
	var messages []Message
	subQuery := DB.
		Table("message").
		Where(" visitor_id in (? )", visitorIds).
		Order("id desc").
		Limit(1024).
		SubQuery()
	DB.Raw("SELECT ANY_VALUE(visitor_id) visitor_id,ANY_VALUE(id) id,ANY_VALUE(content) content FROM ? message_alia GROUP BY visitor_id", subQuery).Scan(&messages)
	//DB.Select("ANY_VALUE(visitor_id) visitor_id,MAX(ANY_VALUE(id)) id,ANY_VALUE(content) content").Group("visitor_id").Find(&messages)
	return messages
}

//查询条数
func CountMessage() uint {
	var count uint
	DB.Model(&Message{}).Count(&count)
	return count
}
