package models

// 项目状态
const (
	StatusSchoolVerify  string = "学院审核"
	StatusICenterVerify string = "网信中心审核"
	StatusMasterVerify  string = "专家论证"
	StatusVerifyProject string = "核定参数"
	StatusBudget        string = "招投标"
	StatusRunning       string = "执行中"
	StatusFinish        string = "已完成"
)

// 招标方式
const (
	InviteOT string = "公开招标"
	InviteIT string = "邀请招标"
	InviteCN string = "竞争性谈判"
	InviteSS string = "单一来源"
	InviteIS string = "询价采购"
)

// 专家审核项目状态
const (
	StatusWaitToAudit = "audit"
	StatusFinishAudit = "finish"
	ResultPass        = "pass"
	ResultFail        = "fail"
)
