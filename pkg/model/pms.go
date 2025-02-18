package model

import "time"

const (
	TableNamePMSEquipmentFirstCategory = "ship_manager.pms_equipment_first_category"
)

// PMS设备第一大类
type PMSEquipmentFirstCategory struct {
	ID          int64     `gorm:"primaryKey;autoIncrement;column:id" json:"id"` // ID
	FirstCode   string    `gorm:"column:first_code" json:"first_code"`          // 一大类类别编码
	Description string    `gorm:"column:description" json:"description"`        // 名称描述
	Editor      string    `gorm:"column:editor" json:"editor"`                  // 编辑者
	Source      string    `gorm:"column:source" json:"source"`                  // 数据来源
	ModifyTime  time.Time `gorm:"column:modify_time" json:"modify_time"`        // 修改时间
	CreateTime  time.Time `gorm:"column:create_time" json:"create_time"`        // 创建时间
}

func (*PMSEquipmentFirstCategory) TableName() string {
	return TableNamePMSEquipmentFirstCategory
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	TableNamePMSEquipmentSecondCategory = "ship_manager.pms_equipment_second_category"
)

// PMS设备第二大类
type PMSEquipmentSecondCategory struct {
	ID          int64     `gorm:"primaryKey;autoIncrement;column:id" json:"id"` // ID
	FirstCode   string    `gorm:"column:first_code" json:"first_code"`          // 一大类类别编码
	SecondCode  string    `gorm:"column:second_code" json:"second_code"`        // 二大类类别编码
	Description string    `gorm:"column:description" json:"description"`        // 名称描述
	Editor      string    `gorm:"column:editor" json:"editor"`                  // 编辑者
	Source      string    `gorm:"column:source" json:"source"`                  // 数据来源
	ModifyTime  time.Time `gorm:"column:modify_time" json:"modify_time"`        // 修改时间
	CreateTime  time.Time `gorm:"column:create_time" json:"create_time"`        // 创建时间
}

func (*PMSEquipmentSecondCategory) TableName() string {
	return TableNamePMSEquipmentSecondCategory
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	TableNamePMSEquipmentThirdCategory = "ship_manager.pms_equipment_third_category"
)

// PMS设备第三大类
type PMSEquipmentThirdCategory struct {
	ID          int64     `gorm:"primaryKey;autoIncrement;column:id" json:"id"` // ID
	FirstCode   string    `gorm:"column:first_code" json:"first_code"`          // 一大类类别编码
	SecondCode  string    `gorm:"column:second_code" json:"second_code"`        // 二大类类别编码
	ThirdCode   string    `gorm:"column:third_code" json:"third_code"`          // 三大类类别编码
	Description string    `gorm:"column:description" json:"description"`        // 名称描述
	Editor      string    `gorm:"column:editor" json:"editor"`                  // 编辑者
	Source      string    `gorm:"column:source" json:"source"`                  // 数据来源
	ModifyTime  time.Time `gorm:"column:modify_time" json:"modify_time"`        // 修改时间
	CreateTime  time.Time `gorm:"column:create_time" json:"create_time"`        // 创建时间
}

func (*PMSEquipmentThirdCategory) TableName() string {
	return TableNamePMSEquipmentThirdCategory
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	TableNamePMSEquipmentFourthCategory = "ship_manager.pms_equipment_fourth_category"
)

// PMS 设备第四大类
type PMSEquipmentFourthCategory struct {
	ID                int64     `gorm:"primaryKey;autoIncrement;column:id" json:"id"`          // ID
	FirstCode         string    `gorm:"column:first_code" json:"first_code"`                   // 一大类类别编码
	SecondCode        string    `gorm:"column:second_code" json:"second_code"`                 // 二大类类别编码
	ThirdCode         string    `gorm:"column:third_code" json:"third_code"`                   // 三大类类别编码
	FourthCode        string    `gorm:"column:fourth_code" json:"fourth_code"`                 // 四大类类别编码
	Description       string    `gorm:"column:description" json:"description"`                 // 名称描述
	DailyRunningHours int       `gorm:"column:daily_running_hours" json:"daily_running_hours"` // 日均运行小时数
	Location          string    `gorm:"column:location" json:"location"`                       // 所在位置
	Editor            string    `gorm:"column:editor" json:"editor"`                           // 编辑者
	Source            string    `gorm:"column:source" json:"source"`                           // 数据来源
	ModifyTime        time.Time `gorm:"column:modify_time" json:"modify_time"`                 // 修改时间
	CreateTime        time.Time `gorm:"column:create_time" json:"create_time"`                 // 创建时间
}

func (*PMSEquipmentFourthCategory) TableName() string {
	return TableNamePMSEquipmentFourthCategory
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	TableNamePMSCriticalSpare = "ship_manager.pms_critical_spare"
)

// PMS 关键性设备的备件清单
type PMSCriticalSpare struct {
	ID          int64     `gorm:"primaryKey;autoIncrement;column:id" json:"id"` // ID
	FirstCode   string    `gorm:"column:first_code" json:"first_code"`          // 一大类类别编码
	SecondCode  string    `gorm:"column:second_code" json:"second_code"`        // 二大类类别编码
	ThirdCode   string    `gorm:"column:third_code" json:"third_code"`          // 三大类类别编码
	FourthCode  string    `gorm:"column:fourth_code" json:"fourth_code"`        // 四大类类别编码
	Description string    `gorm:"column:description" json:"description"`        // 名称描述
	PartNumber  string    `gorm:"column:part_number" json:"part_number"`        // 设备编号
	Location    string    `gorm:"column:location" json:"location"`              // 所在位置
	Stock       int       `gorm:"column:stock" json:"stock"`                    // 库存数
	MinStock    int       `gorm:"column:min_stock" json:"min_stock"`            // 最低库存数
	Unit        string    `gorm:"column:unit" json:"unit"`                      // 备件单位
	Editor      string    `gorm:"column:editor" json:"editor"`                  // 编辑者
	Source      string    `gorm:"column:source" json:"source"`                  // 数据来源
	ModifyTime  time.Time `gorm:"column:modify_time" json:"modify_time"`        // 修改时间
	CreateTime  time.Time `gorm:"column:create_time" json:"create_time"`        // 创建时间
}

func (*PMSCriticalSpare) TableName() string {
	return TableNamePMSCriticalSpare
}

// ///////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	TableNamePMSDefectList = "ship_manager.pms_defect_list"
)

// PMS 设备损坏工单
type PMSDefectList struct {
	ID               int64     `gorm:"primaryKey;autoIncrement;column:id" json:"id"`          // ID
	SystemModule     string    `gorm:"column:system_module" json:"system_module"`             // 系统模块
	FirstCode        string    `gorm:"column:first_code" json:"first_code"`                   // 一大类类别编码
	SecondCode       string    `gorm:"column:second_code" json:"second_code"`                 // 二大类类别编码
	ThirdCode        string    `gorm:"column:third_code" json:"third_code"`                   // 三大类类别编码
	FourthCode       string    `gorm:"column:fourth_code" json:"fourth_code"`                 // 四大类类别编码
	Description      string    `gorm:"column:description" json:"description"`                 // 名称描述
	RepairMode       string    `gorm:"column:repair_mode" json:"repair_mode"`                 // 修复模式
	OutOfServiceDate time.Time `gorm:"column:out_of_service_date" json:"out_of_service_date"` // 设备停止服务日期
	RepairedDate     time.Time `gorm:"column:repaired_date" json:"repaired_date"`             // 设备修复日期
	BreakDownHours   int       `gorm:"column:break_down_hours" json:"break_down_hours"`       // 设备故障小时数
	RunningHours     int       `gorm:"column:running_hours" json:"running_hours"`             // 设备运行小时数
	Location         string    `gorm:"column:location" json:"location"`                       // 所在位置
	Environment      string    `gorm:"column:environment" json:"environment"`                 // 所在环境
	Remarks          string    `gorm:"column:remarks" json:"remarks"`                         // 评价、备注
	Editor           string    `gorm:"column:editor" json:"editor"`                           // 编辑者
	Source           string    `gorm:"column:source" json:"source"`                           // 数据来源
	ModifyTime       time.Time `gorm:"column:modify_time" json:"modify_time"`                 // 修改时间
	CreateTime       time.Time `gorm:"column:create_time" json:"create_time"`                 // 创建时间
}

func (*PMSDefectList) TableName() string {
	return TableNamePMSDefectList
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	TableNamePMSDownTime = "ship_manager.pms_down_time"
)

// PMS 设备损坏后果
type PMSDownTime struct {
	ID           int64     `gorm:"primaryKey;autoIncrement;column:id" json:"id"` // ID
	SystemModule string    `gorm:"column:system_module" json:"system_module"`    // 系统模块
	FirstCode    string    `gorm:"column:first_code" json:"first_code"`          // 一大类类别编码
	SecondCode   string    `gorm:"column:second_code" json:"second_code"`        // 二大类类别编码
	ThirdCode    string    `gorm:"column:third_code" json:"third_code"`          // 三大类类别编码
	FourthCode   string    `gorm:"column:fourth_code" json:"fourth_code"`        // 四大类类别编码
	Description  string    `gorm:"column:description" json:"description"`        // 名称描述
	Reason       string    `gorm:"column:reason" json:"reason"`                  // 损坏原因
	StartTime    time.Time `gorm:"column:start_time" json:"start_time"`          // 故障开始时间
	EndTime      time.Time `gorm:"column:end_time" json:"end_time"`              // 故障结束时间
	DownHours    int       `gorm:"column:down_hours" json:"down_hours"`          // 故障小时数
	RunningHours int       `gorm:"column:running_hours" json:"running_hours"`    // 设备运行小时数
	Location     string    `gorm:"column:location" json:"location"`              // 所在位置
	Environment  string    `gorm:"column:environment" json:"environment"`        // 所在环境
	Remarks      string    `gorm:"column:remarks" json:"remarks"`                // 评价、备注
	Editor       string    `gorm:"column:editor" json:"editor"`                  // 编辑者
	Source       string    `gorm:"column:source" json:"source"`                  // 数据来源
	ModifyTime   time.Time `gorm:"column:modify_time" json:"modify_time"`        // 修改时间
	CreateTime   time.Time `gorm:"column:create_time" json:"create_time"`        // 创建时间
}

func (*PMSDownTime) TableName() string {
	return TableNamePMSDownTime
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	TableNamePMSOperatorWork = "ship_manager.pms_operator_work"
)

// PMS 设备维护人员分工
type PMSOperatorWork struct {
	ID          int64     `gorm:"primaryKey;autoIncrement;column:id" json:"id"` // ID
	FirstCode   string    `gorm:"column:first_code" json:"first_code"`          // 一大类类别编码
	SecondCode  string    `gorm:"column:second_code" json:"second_code"`        // 二大类类别编码
	ThirdCode   string    `gorm:"column:third_code" json:"third_code"`          // 三大类类别编码
	FourthCode  string    `gorm:"column:fourth_code" json:"fourth_code"`        // 四大类类别编码
	Description string    `gorm:"column:description" json:"description"`        // 名称描述
	Operator    string    `gorm:"column:operator" json:"operator"`              // 维护人员
	Remarks     string    `gorm:"column:remarks" json:"remarks"`                // 评价、备注
	Editor      string    `gorm:"column:editor" json:"editor"`                  // 编辑者
	Source      string    `gorm:"column:source" json:"source"`                  // 数据来源
	ModifyTime  time.Time `gorm:"column:modify_time" json:"modify_time"`        // 修改时间
	CreateTime  time.Time `gorm:"column:create_time" json:"create_time"`        // 创建时间
}

func (*PMSOperatorWork) TableName() string {
	return TableNamePMSOperatorWork
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	TableNamePMSEquipmentFactory = "ship_manager.pms_equipment_factory"
)

// PMS 设备厂家
type PMSEquipmentFactory struct {
	ID                         int64     `gorm:"primaryKey;autoIncrement;column:id" json:"id"`                            // ID
	FirstCode                  string    `gorm:"column:first_code" json:"first_code"`                                     // 一大类类别编码
	SecondCode                 string    `gorm:"column:second_code" json:"second_code"`                                   // 二大类类别编码
	ThirdCode                  string    `gorm:"column:third_code" json:"third_code"`                                     // 三大类类别编码
	FourthCode                 string    `gorm:"column:fourth_code" json:"fourth_code"`                                   // 四大类类别编码
	Description                string    `gorm:"column:description" json:"description"`                                   // 名称描述
	MachineryID                string    `gorm:"column:machinery_id" json:"machinery_id"`                                 // 机器ID
	ShipYardCode               string    `gorm:"column:ship_yard_code" json:"ship_yard_code"`                             // 船厂代码
	Maker                      string    `gorm:"column:maker" json:"maker"`                                               // 制造商
	Contact                    string    `gorm:"column:contact" json:"contact"`                                           // 联系方式
	Rating                     string    `gorm:"column:rating" json:"rating"`                                             // 设备等级
	Remarks                    string    `gorm:"column:remarks" json:"remarks"`                                           // 评价、备注
	Model                      string    `gorm:"column:model" json:"model"`                                               // 型号
	InstalledDate              time.Time `gorm:"column:installed_date" json:"installed_date"`                             // 设备安装日期
	DesignedSeaTemperature     float32   `gorm:"column:designed_sea_temperature" json:"designed_sea_temperature"`         // 设计海上使用温度
	DesignedAmbientTemperature float32   `gorm:"column:designed_ambient_temperature" json:"designed_ambient_temperature"` // 设计环境温度
	Score                      float32   `gorm:"column:score" json:"score"`                                               // 设备评分
	Editor                     string    `gorm:"column:editor" json:"editor"`                                             // 编辑者
	Source                     string    `gorm:"column:source" json:"source"`                                             // 数据来源
	ModifyTime                 time.Time `gorm:"column:modify_time" json:"modify_time"`                                   // 修改时间
	CreateTime                 time.Time `gorm:"column:create_time" json:"create_time"`                                   // 创建时间
}

func (*PMSEquipmentFactory) TableName() string {
	return TableNamePMSEquipmentFactory
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	TableNamePMSOverhaulStatus = "ship_manager.pms_overhaul_status"
)

// 设备大修的状态
type PMSOverhaulStatus struct {
	ID                       int64     `gorm:"primaryKey;autoIncrement;column:id" json:"id"`                        // ID
	FirstCode                string    `gorm:"column:first_code" json:"first_code"`                                 // 一大类类别编码
	SecondCode               string    `gorm:"column:second_code" json:"second_code"`                               // 二大类类别编码
	ThirdCode                string    `gorm:"column:third_code" json:"third_code"`                                 // 三大类类别编码
	FourthCode               string    `gorm:"column:fourth_code" json:"fourth_code"`                               // 四大类类别编码
	Description              string    `gorm:"column:description" json:"description"`                               // 名称描述
	Task                     string    `gorm:"column:task" json:"task"`                                             // 维修任务
	OfficeRef                int       `gorm:"column:office_ref" json:"office_ref"`                                 // 标准运行小时数
	OriginalDueDate          time.Time `gorm:"column:original_due_date" json:"original_due_date"`                   // 原定维修日期
	LastActionDate           time.Time `gorm:"column:last_action_date" json:"last_action_date"`                     // 上次维修日期
	ActionResult             string    `gorm:"column:action_result" json:"action_result"`                           // 维修的结果
	WorkDoneDetail           string    `gorm:"column:work_done_detail" json:"work_done_detail"`                     // 维修细节
	ActionCompleteHours      int       `gorm:"column:action_complete_hours" json:"action_complete_hours"`           // 维修完成小时数
	MaintenanceIntervelHours int       `gorm:"column:maintenance_intervel_hours" json:"maintenance_intervel_hours"` // 根据运行小时数进行维护干预
	MaintenanceIntervelDays  int       `gorm:"column:maintenance_intervel_days" json:"maintenance_intervel_days"`   // 根据运行天数进行维护干预
	DueDate                  time.Time `gorm:"column:due_date" json:"due_date"`                                     // 计划维修完成的日期
	Editor                   string    `gorm:"column:editor" json:"editor"`                                         // 编辑者
	Source                   string    `gorm:"column:source" json:"source"`                                         // 数据来源
	ModifyTime               time.Time `gorm:"column:modify_time" json:"modify_time"`                               // 修改时间
	CreateTime               time.Time `gorm:"column:create_time" json:"create_time"`                               // 创建时间
}

func (*PMSOverhaulStatus) TableName() string {
	return TableNamePMSOverhaulStatus
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	TableNamePMSRunningHours = "ship_manager.pms_running_hours"
)

// 设备运行时间
type PMSRunningHours struct {
	ID                   int64     `gorm:"primaryKey;autoIncrement;column:id" json:"id"`                  // ID
	FirstCode            string    `gorm:"column:first_code" json:"first_code"`                           // 一大类类别编码
	SecondCode           string    `gorm:"column:second_code" json:"second_code"`                         // 二大类类别编码
	ThirdCode            string    `gorm:"column:third_code" json:"third_code"`                           // 三大类类别编码
	FourthCode           string    `gorm:"column:fourth_code" json:"fourth_code"`                         // 四大类类别编码
	Description          string    `gorm:"column:description" json:"description"`                         // 名称描述
	OfficeRef            int       `gorm:"column:office_ref" json:"office_ref"`                           // 标准运行小时数
	RunningDate          time.Time `gorm:"column:running_date" json:"running_date"`                       // 开始运行的日期
	TotalRunningHours    int       `gorm:"column:total_running_hours" json:"total_running_hours"`         // 总的运行小时数
	Setup                bool      `gorm:"column:setup" json:"setup"`                                     // 是否启动
	AvgDailyRunningHours int       `gorm:"column:avg_daily_running_hours" json:"avg_daily_running_hours"` // 日均运行小时数
	Editor               string    `gorm:"column:editor" json:"editor"`                                   // 编辑者
	Source               string    `gorm:"column:source" json:"source"`                                   // 数据来源
	ModifyTime           time.Time `gorm:"column:modify_time" json:"modify_time"`                         // 修改时间
	CreateTime           time.Time `gorm:"column:create_time" json:"create_time"`                         // 创建时间
}

func (*PMSRunningHours) TableName() string {
	return TableNamePMSRunningHours
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	TableNamePMSTaskTemplate = "ship_manager.pms_task_template"
)

// 系统默认生成的维修保养工单清单模板
type PMSTaskTemplate struct {
	ID                       int64     `gorm:"primaryKey;autoIncrement;column:id" json:"id"`                        // ID
	FirstCode                string    `gorm:"column:first_code" json:"first_code"`                                 // 一大类类别编码
	SecondCode               string    `gorm:"column:second_code" json:"second_code"`                               // 二大类类别编码
	ThirdCode                string    `gorm:"column:third_code" json:"third_code"`                                 // 三大类类别编码
	FourthCode               string    `gorm:"column:fourth_code" json:"fourth_code"`                               // 四大类类别编码
	Description              string    `gorm:"column:description" json:"description"`                               // 名称描述
	Task                     string    `gorm:"column:task" json:"task"`                                             // 维修任务
	FullDescription          string    `gorm:"column:full_description" json:"full_description"`                     // 完成的设备描述
	OfficeRef                int       `gorm:"column:office_ref" json:"office_ref"`                                 // 标准运行小时数
	MaintenanceIntervelHours int       `gorm:"column:maintenance_intervel_hours" json:"maintenance_intervel_hours"` // 根据运行小时数进行维护干预
	MaintenanceIntervelDays  int       `gorm:"column:maintenance_intervel_days" json:"maintenance_intervel_days"`   // 根据运行天数进行维护干预
	TotalRunningHours        int       `gorm:"column:total_running_hours" json:"total_running_hours"`               // 总的运行小时数
	Setup                    bool      `gorm:"column:setup" json:"setup"`                                           // 是否启动
	AvgDailyRunningHours     int       `gorm:"column:avg_daily_running_hours" json:"avg_daily_running_hours"`       // 日均运行小时数
	Editor                   string    `gorm:"column:editor" json:"editor"`                                         // 编辑者
	Source                   string    `gorm:"column:source" json:"source"`                                         // 数据来源
	ModifyTime               time.Time `gorm:"column:modify_time" json:"modify_time"`                               // 修改时间
	CreateTime               time.Time `gorm:"column:create_time" json:"create_time"`                               // 创建时间
}

func (*PMSTaskTemplate) TableName() string {
	return TableNamePMSTaskTemplate
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	TableNamePMSWorkDone = "ship_manager.pms_work_done"
)

// 已经完成的维修保养工单
type PMSWorkDone struct {
	ID              int64     `gorm:"primaryKey;autoIncrement;column:id" json:"id"`      // ID
	FirstCode       string    `gorm:"column:first_code" json:"first_code"`               // 一大类类别编码
	SecondCode      string    `gorm:"column:second_code" json:"second_code"`             // 二大类类别编码
	ThirdCode       string    `gorm:"column:third_code" json:"third_code"`               // 三大类类别编码
	FourthCode      string    `gorm:"column:fourth_code" json:"fourth_code"`             // 四大类类别编码
	Description     string    `gorm:"column:description" json:"description"`             // 名称描述
	Task            string    `gorm:"column:task" json:"task"`                           // 维修任务
	OfficeRef       int       `gorm:"column:office_ref" json:"office_ref"`               // 标准运行小时数
	OriginalDueDate time.Time `gorm:"column:original_due_date" json:"original_due_date"` // 计划维修日期
	LastActionDate  time.Time `gorm:"column:last_action_date" json:"last_action_date"`   // 上次维修日期
	ActionResult    string    `gorm:"column:action_result" json:"action_result"`         // 维修的结果
	Detail          string    `gorm:"column:detail" json:"detail"`                       // 已完成工作和所用零件的详细信息
	SpeedHours      int       `gorm:"column:speed_hours" json:"speed_hours"`             // 维修所花费的小时数
	Completed       bool      `gorm:"column:completed" json:"completed"`                 // 是否已经完成
	Editor          string    `gorm:"column:editor" json:"editor"`                       // 编辑者
	Source          string    `gorm:"column:source" json:"source"`                       // 数据来源
	ModifyTime      time.Time `gorm:"column:modify_time" json:"modify_time"`             // 修改时间
	CreateTime      time.Time `gorm:"column:create_time" json:"create_time"`             // 创建时间
}

func (*PMSWorkDone) TableName() string {
	return TableNamePMSWorkDone
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	TableNamePMSWorkOrder = "ship_manager.pms_work_order"
)

// 需要做的工单号（当月需要完成的工单）
type PMSWorkOrder struct {
	ID              int64     `gorm:"primaryKey;autoIncrement;column:id" json:"id"`      // ID
	TaskID          int       `gorm:"column:task_id" json:"task_id"`                     // 任务ID
	Date            time.Time `gorm:"column:date" json:"date"`                           // 日期
	FirstCode       string    `gorm:"column:first_code" json:"first_code"`               // 一大类类别编码
	SecondCode      string    `gorm:"column:second_code" json:"second_code"`             // 二大类类别编码
	ThirdCode       string    `gorm:"column:third_code" json:"third_code"`               // 三大类类别编码
	FourthCode      string    `gorm:"column:fourth_code" json:"fourth_code"`             // 四大类类别编码
	Description     string    `gorm:"column:description" json:"description"`             // 名称描述
	Task            string    `gorm:"column:task" json:"task"`                           // 维修任务
	OfficeRef       int       `gorm:"column:office_ref" json:"office_ref"`               // 标准运行小时数
	OriginalDueDate time.Time `gorm:"column:original_due_date" json:"original_due_date"` // 计划维修日期
	LastActionDate  time.Time `gorm:"column:last_action_date" json:"last_action_date"`   // 上次维修日期
	ActionResult    string    `gorm:"column:action_result" json:"action_result"`         // 维修的结果
	Detail          string    `gorm:"column:detail" json:"detail"`                       // 已完成工作和所用零件的详细信息
	SpeedHours      int       `gorm:"column:speed_hours" json:"speed_hours"`             // 维修所花费的小时数
	Completed       bool      `gorm:"column:completed" json:"completed"`                 // 是否已经完成
	Editor          string    `gorm:"column:editor" json:"editor"`                       // 编辑者
	Source          string    `gorm:"column:source" json:"source"`                       // 数据来源
	ModifyTime      time.Time `gorm:"column:modify_time" json:"modify_time"`             // 修改时间
	CreateTime      time.Time `gorm:"column:create_time" json:"create_time"`             // 创建时间
}

func (*PMSWorkOrder) TableName() string {
	return TableNamePMSWorkOrder
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////
