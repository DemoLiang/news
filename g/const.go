package g

const (
	REQ__MSG_TYPE__TEXT                   = "text"
	REQ__MSG_TYPE__IMAGE                  = "image"
	REQ__MSG_TYPE__VOICE                  = "voice"
	REQ__MSG_TYPE__VIDEO                  = "video"
	REQ__MSG_TYPE__LOCATION               = "location"
	REQ__MSG_TYPE__LINK                   = "link"
	REQ__EVENT_TYPE__SUBSCRIBE            = "subscribe"
	REQ__EVENT_TYPE__ENTER_AGENT          = "enter_agent"
	REQ__EVENT_TYPE__LOCATION             = "LOCATION"
	REQ__EVENT_TYPE__BATCH_JOB_RESULT     = "batch_job_result"
	REQ__EVENT_TYPE__CHANGE_CONTACT       = "change_contact"
	REQ__EVENT_TYPE__MSGTYPE__CREATE_USER = "create_user"
	REQ__EVENT_TYPE__MSGTYPE__UPDATE_USER = "update_user"
	REQ__EVENT_TYPE__DELETE_USER          = "delete_user"
	REQ__EVENT_TYPE__CREATE_PARTY         = "create_party"
	REQ__EVENT_TYPE__UPDATE_PARTY         = "update_party"
	REQ__EVENT_TYPE__DELETE_PARTY         = "delete_party"
	REQ__EVENT_TYPE__UPDATE_TAG           = "update_tag"
	REQ__EVENT_TYPE__CLICK                = "click"
	REQ__EVENT_TYPE__VIEW                 = "view"
	REQ__EVENT_TYPE__SCANCODE_PUSH        = "scancode_push"
	REQ__EVENT_TYPE__SCANCODE_WAITMSG     = "scancode_waitmsg"
	REQ__EVENT_TYPE__PIC_SYSPHOTO         = "pic_sysphoto"
	REQ__EVENT_TYPE__PIC_PHOTO_OR_ALBUM   = "pic_photo_or_album"
	REQ__EVENT_TYPE__PIC_WEIXIN           = "pic_weixin"
	REQ__EVENT_TYPE__LOCATION_SELECT      = "location_select"
)
const (
	RESP_MESSAGE_TYPE__TEXT  = "text"
	RESP_MESSAGE_TYPE__MUSIC = "music"
	RESP_MESSAGE_TYPE__NEWS  = "news"

	REQ_MESSAGE_TYPE__TEXT              = "text"
	REQ_MESSAGE_TYPE__IMAGE             = "image"
	REQ_MESSAGE_TYPE__LINK              = "link"
	REQ_MESSAGE_TYPE__LOCATION          = "location"
	REQ_MESSAGE_TYPE__VOICE             = "voice"
	REQ_MESSAGE_TYPE__EVENT             = "event"
	REQ_MESSAGE_TYPE_EVENT__SUBSCRIBE   = "subscribe"
	REQ_MESSAGE_TYPE_EVENT__UNSUBSCRIBE = "unsubscribe"
	REQ_MESSAGE_TYPE_EVENT__CLICK       = "CLICK"
	REQ_MESSAGE_TYPE_EVENT__VIEW        = "VIEW"
)

const (
	WECHAT__MENU_BUTTON__BT0 = "bt0"
	WECHAT__MENU_BUTTON__BT1 = "bt1"
	WECHAT__MENU_BUTTON__BT2 = "bt2"
)

const (
	DEFAULT_SUBSCRIBE__MSG   = "感谢关注每日新闻集,今后我们将每日推送你想要的每日新闻集合喔～"
	DEFAULT_LATEST_NEWS_NULL = "暂无最新消息"
)

const (
	VERSION = "0.0.1"
)
