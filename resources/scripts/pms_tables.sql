create table ship_manager.pms_equipment_first_category
(
    id              INTEGER identity constraint PK_pms_equipment_first_category primary key,
    first_code      varchar(64),
    description     varchar(256),
    editor          varchar(64),
    source          varchar(64),
    modify_time     datetime,
    create_time     datetime
)
go

exec sp_addextendedproperty 'MS_Description', N'主键', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_first_category', 'COLUMN', 'id'
go

exec sp_addextendedproperty 'MS_Description', N'设备第一大类编码', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_first_category', 'COLUMN', 'first_code'
go

exec sp_addextendedproperty 'MS_Description', N'名称描述', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_first_category', 'COLUMN', 'description'
go

exec sp_addextendedproperty 'MS_Description', N'编辑者', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_first_category', 'COLUMN', 'editor'
go

exec sp_addextendedproperty 'MS_Description', N'数据来源', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_first_category', 'COLUMN', 'source'
go

exec sp_addextendedproperty 'MS_Description', N'修改时间', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_first_category', 'COLUMN', 'modify_time'
go

exec sp_addextendedproperty 'MS_Description', N'创建时间', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_first_category', 'COLUMN', 'create_time'
go

--------------------------------------------------------------------------------------------------------------------------

create table ship_manager.pms_equipment_second_category
(
    id              INTEGER identity constraint PK_pms_equipment_second_category primary key,
    first_code      varchar(64),
    second_code     varchar(64),
    description     varchar(256),
    editor          varchar(64),
    source          varchar(64),
    modify_time     datetime,
    create_time     datetime
)
go

exec sp_addextendedproperty 'MS_Description', N'主键', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_second_category', 'COLUMN', 'id'
go

exec sp_addextendedproperty 'MS_Description', N'设备第一大类编码', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_second_category', 'COLUMN', 'first_code'
go

exec sp_addextendedproperty 'MS_Description', N'设备第二大类编码', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_second_category', 'COLUMN', 'second_code'
go

exec sp_addextendedproperty 'MS_Description', N'名称描述', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_second_category', 'COLUMN', 'description'
go

exec sp_addextendedproperty 'MS_Description', N'编辑者', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_second_category', 'COLUMN', 'editor'
go

exec sp_addextendedproperty 'MS_Description', N'数据来源', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_second_category', 'COLUMN', 'source'
go

exec sp_addextendedproperty 'MS_Description', N'修改时间', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_second_category', 'COLUMN', 'modify_time'
go

exec sp_addextendedproperty 'MS_Description', N'创建时间', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_second_category', 'COLUMN', 'create_time'
go

--------------------------------------------------------------------------------------------------------------------------

create table ship_manager.pms_equipment_third_category
(
    id              INTEGER identity constraint PK_pms_equipment_third_category primary key,
    first_code      varchar(64),
    second_code     varchar(64),
    third_code      varchar(64),
    description     varchar(256),
    editor          varchar(64),
    source          varchar(64),
    modify_time     datetime,
    create_time     datetime
)
go

exec sp_addextendedproperty 'MS_Description', N'主键', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_third_category', 'COLUMN', 'id'
go

exec sp_addextendedproperty 'MS_Description', N'设备第一大类编码', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_third_category', 'COLUMN', 'first_code'
go

exec sp_addextendedproperty 'MS_Description', N'设备第二大类编码', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_third_category', 'COLUMN', 'second_code'
go

exec sp_addextendedproperty 'MS_Description', N'设备第三大类编码', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_third_category', 'COLUMN', 'third_code'
go

exec sp_addextendedproperty 'MS_Description', N'名称描述', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_third_category', 'COLUMN', 'description'
go

exec sp_addextendedproperty 'MS_Description', N'编辑者', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_third_category', 'COLUMN', 'editor'
go

exec sp_addextendedproperty 'MS_Description', N'数据来源', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_third_category', 'COLUMN', 'source'
go

exec sp_addextendedproperty 'MS_Description', N'修改时间', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_third_category', 'COLUMN', 'modify_time'
go

exec sp_addextendedproperty 'MS_Description', N'创建时间', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_third_category', 'COLUMN', 'create_time'
go

--------------------------------------------------------------------------------------------------------------------------

create table ship_manager.pms_equipment_fourth_category
(
    id                  INTEGER identity constraint PK_pms_equipment_fourth_category primary key,
    first_code          varchar(64),
    second_code         varchar(64),
    third_code          varchar(64),
    fourth_code         varchar(64),
    description         varchar(256),
    daily_running_hours float,
    location            varchar(256),
    editor              varchar(64),
    source              varchar(64),
    modify_time         datetime,
    create_time         datetime
)
go

exec sp_addextendedproperty 'MS_Description', N'主键', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_fourth_category', 'COLUMN', 'id'
go

exec sp_addextendedproperty 'MS_Description', N'设备第一大类编码', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_fourth_category', 'COLUMN', 'first_code'
go

exec sp_addextendedproperty 'MS_Description', N'设备第二大类编码', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_fourth_category', 'COLUMN', 'second_code'
go

exec sp_addextendedproperty 'MS_Description', N'设备第三大类编码', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_fourth_category', 'COLUMN', 'third_code'
go

exec sp_addextendedproperty 'MS_Description', N'设备第四大类编码', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_fourth_category', 'COLUMN', 'fourth_code'
go

exec sp_addextendedproperty 'MS_Description', N'名称描述', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_fourth_category', 'COLUMN', 'description'
go

exec sp_addextendedproperty 'MS_Description', N'日均运行小时数', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_fourth_category', 'COLUMN', 'daily_running_hours'
go

exec sp_addextendedproperty 'MS_Description', N'组件所在位置', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_fourth_category', 'COLUMN', 'location'
go

exec sp_addextendedproperty 'MS_Description', N'编辑者', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_fourth_category', 'COLUMN', 'editor'
go

exec sp_addextendedproperty 'MS_Description', N'数据来源', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_fourth_category', 'COLUMN', 'source'
go

exec sp_addextendedproperty 'MS_Description', N'修改时间', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_fourth_category', 'COLUMN', 'modify_time'
go

exec sp_addextendedproperty 'MS_Description', N'创建时间', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_fourth_category', 'COLUMN', 'create_time'
go

--------------------------------------------------------------------------------------------------------------------------

create table ship_manager.pms_critical_spare
(
    id                  INTEGER identity constraint PK_pms_critical_spare primary key,
    first_code          varchar(64),
    second_code         varchar(64),
    third_code          varchar(64),
    fourth_code         varchar(64),
    description         varchar(256),
    part_number         varchar(64),
    location            varchar(256),
    stock               int,
    min_stock           int,
    unit                varchar(64),
    editor              varchar(64),
    source              varchar(64),
    modify_time         datetime,
    create_time         datetime
)
go

exec sp_addextendedproperty 'MS_Description', N'主键', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_critical_spare', 'COLUMN', 'id'
go

exec sp_addextendedproperty 'MS_Description', N'设备第一大类编码', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_critical_spare', 'COLUMN', 'first_code'
go

exec sp_addextendedproperty 'MS_Description', N'设备第二大类编码', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_critical_spare', 'COLUMN', 'second_code'
go

exec sp_addextendedproperty 'MS_Description', N'设备第三大类编码', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_critical_spare', 'COLUMN', 'third_code'
go

exec sp_addextendedproperty 'MS_Description', N'设备第四大类编码', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_critical_spare', 'COLUMN', 'fourth_code'
go

exec sp_addextendedproperty 'MS_Description', N'名称描述', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_critical_spare', 'COLUMN', 'description'
go

exec sp_addextendedproperty 'MS_Description', N'设备编号', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_critical_spare', 'COLUMN', 'part_number'
go

exec sp_addextendedproperty 'MS_Description', N'所在位置', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_critical_spare', 'COLUMN', 'location'
go

exec sp_addextendedproperty 'MS_Description', N'库存数', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_critical_spare', 'COLUMN', 'stock'
go

exec sp_addextendedproperty 'MS_Description', N'最低库存数', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_critical_spare', 'COLUMN', 'min_stock'
go

exec sp_addextendedproperty 'MS_Description', N'备件单位', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_critical_spare', 'COLUMN', 'unit'
go

exec sp_addextendedproperty 'MS_Description', N'编辑者', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_critical_spare', 'COLUMN', 'editor'
go

exec sp_addextendedproperty 'MS_Description', N'数据来源', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_critical_spare', 'COLUMN', 'source'
go

exec sp_addextendedproperty 'MS_Description', N'修改时间', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_critical_spare', 'COLUMN', 'modify_time'
go

exec sp_addextendedproperty 'MS_Description', N'创建时间', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_critical_spare', 'COLUMN', 'create_time'
go

--------------------------------------------------------------------------------------------------------------------------

create table ship_manager.pms_defect_list
(
    id                  INTEGER identity constraint PK_pms_defect_list primary key,
    system_module       varchar(64),
    first_code          varchar(64),
    second_code         varchar(64),
    third_code          varchar(64),
    fourth_code         varchar(64),
    description         varchar(256),
    repair_mode         varchar(64),
    out_of_service_date datetime,
    repaired_date       datetime,
    break_down_hours    int,
    running_hours       int,
    location            varchar(256),
    environment         varchar(256),
    remarks             varchar(512),
    editor              varchar(64),
    source              varchar(64),
    modify_time         datetime,
    create_time         datetime
)
go

exec sp_addextendedproperty 'MS_Description', N'主键', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_defect_list', 'COLUMN', 'id'
go

exec sp_addextendedproperty 'MS_Description', N'系统模块', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_defect_list', 'COLUMN', 'system_module'
go

exec sp_addextendedproperty 'MS_Description', N'设备第一大类编码', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_defect_list', 'COLUMN', 'first_code'
go

exec sp_addextendedproperty 'MS_Description', N'设备第二大类编码', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_defect_list', 'COLUMN', 'second_code'
go

exec sp_addextendedproperty 'MS_Description', N'设备第三大类编码', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_defect_list', 'COLUMN', 'third_code'
go

exec sp_addextendedproperty 'MS_Description', N'设备第四大类编码', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_defect_list', 'COLUMN', 'fourth_code'
go

exec sp_addextendedproperty 'MS_Description', N'修复模式', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_defect_list', 'COLUMN', 'description'
go

exec sp_addextendedproperty 'MS_Description', N'设备编号', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_defect_list', 'COLUMN', 'repair_mode'
go

exec sp_addextendedproperty 'MS_Description', N'设备停止服务日期', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_defect_list', 'COLUMN', 'out_of_service_date'
go

exec sp_addextendedproperty 'MS_Description', N'设备修复日期', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_defect_list', 'COLUMN', 'repaired_date'
go

exec sp_addextendedproperty 'MS_Description', N'设备故障小时数', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_defect_list', 'COLUMN', 'break_down_hours'
go

exec sp_addextendedproperty 'MS_Description', N'设备运行小时数', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_defect_list', 'COLUMN', 'running_hours'
go

exec sp_addextendedproperty 'MS_Description', N'所在位置', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_defect_list', 'COLUMN', 'location'
go

exec sp_addextendedproperty 'MS_Description', N'所在环境', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_defect_list', 'COLUMN', 'environment'
go

exec sp_addextendedproperty 'MS_Description', N'评价、备注', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_defect_list', 'COLUMN', 'remarks'
go

exec sp_addextendedproperty 'MS_Description', N'编辑者', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_defect_list', 'COLUMN', 'editor'
go

exec sp_addextendedproperty 'MS_Description', N'数据来源', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_defect_list', 'COLUMN', 'source'
go

exec sp_addextendedproperty 'MS_Description', N'修改时间', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_defect_list', 'COLUMN', 'modify_time'
go

exec sp_addextendedproperty 'MS_Description', N'创建时间', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_defect_list', 'COLUMN', 'create_time'
go

--------------------------------------------------------------------------------------------------------------------------

create table ship_manager.pms_down_time
(
    id                  INTEGER identity constraint PK_pms_down_time primary key,
    system_module       varchar(64),
    first_code          varchar(64),
    second_code         varchar(64),
    third_code          varchar(64),
    fourth_code         varchar(64),
    description         varchar(256),
    reason              varchar(512),
    start_time          datetime,
    end_time            datetime,
    down_hours          int,
    running_hours       int,
    location            varchar(256),
    environment         varchar(256),
    remarks             varchar(512),
    editor              varchar(64),
    source              varchar(64),
    modify_time         datetime,
    create_time         datetime
)
go

exec sp_addextendedproperty 'MS_Description', N'主键', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_down_time', 'COLUMN', 'id'
go

exec sp_addextendedproperty 'MS_Description', N'系统模块', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_down_time', 'COLUMN', 'system_module'
go

exec sp_addextendedproperty 'MS_Description', N'设备第一大类编码', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_down_time', 'COLUMN', 'first_code'
go

exec sp_addextendedproperty 'MS_Description', N'设备第二大类编码', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_down_time', 'COLUMN', 'second_code'
go

exec sp_addextendedproperty 'MS_Description', N'设备第三大类编码', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_down_time', 'COLUMN', 'third_code'
go

exec sp_addextendedproperty 'MS_Description', N'设备第四大类编码', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_down_time', 'COLUMN', 'fourth_code'
go

exec sp_addextendedproperty 'MS_Description', N'名称描述', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_down_time', 'COLUMN', 'description'
go

exec sp_addextendedproperty 'MS_Description', N'损坏原因', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_down_time', 'COLUMN', 'reason'
go

exec sp_addextendedproperty 'MS_Description', N'故障开始时间', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_down_time', 'COLUMN', 'start_time'
go

exec sp_addextendedproperty 'MS_Description', N'故障结束时间', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_down_time', 'COLUMN', 'end_time'
go

exec sp_addextendedproperty 'MS_Description', N'故障小时数', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_down_time', 'COLUMN', 'down_hours'
go

exec sp_addextendedproperty 'MS_Description', N'设备运行小时数', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_down_time', 'COLUMN', 'running_hours'
go

exec sp_addextendedproperty 'MS_Description', N'所在位置', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_down_time', 'COLUMN', 'location'
go

exec sp_addextendedproperty 'MS_Description', N'所在环境', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_down_time', 'COLUMN', 'environment'
go

exec sp_addextendedproperty 'MS_Description', N'评价、备注', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_down_time', 'COLUMN', 'remarks'
go

exec sp_addextendedproperty 'MS_Description', N'编辑者', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_down_time', 'COLUMN', 'editor'
go

exec sp_addextendedproperty 'MS_Description', N'数据来源', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_down_time', 'COLUMN', 'source'
go

exec sp_addextendedproperty 'MS_Description', N'修改时间', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_down_time', 'COLUMN', 'modify_time'
go

exec sp_addextendedproperty 'MS_Description', N'创建时间', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_down_time', 'COLUMN', 'create_time'
go

--------------------------------------------------------------------------------------------------------------------------

create table ship_manager.pms_operator_work
(
    id                  INTEGER identity constraint PK_pms_operator_work primary key,
    first_code          varchar(64),
    second_code         varchar(64),
    third_code          varchar(64),
    fourth_code         varchar(64),
    description         varchar(256),
    operator            varchar(64),
    remarks             varchar(512),
    editor              varchar(64),
    source              varchar(64),
    modify_time         datetime,
    create_time         datetime
)
go

exec sp_addextendedproperty 'MS_Description', N'主键', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_operator_work', 'COLUMN', 'id'
go

exec sp_addextendedproperty 'MS_Description', N'设备第一大类编码', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_operator_work', 'COLUMN', 'first_code'
go

exec sp_addextendedproperty 'MS_Description', N'设备第二大类编码', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_operator_work', 'COLUMN', 'second_code'
go

exec sp_addextendedproperty 'MS_Description', N'设备第三大类编码', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_operator_work', 'COLUMN', 'third_code'
go

exec sp_addextendedproperty 'MS_Description', N'设备第四大类编码', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_operator_work', 'COLUMN', 'fourth_code'
go

exec sp_addextendedproperty 'MS_Description', N'名称描述', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_operator_work', 'COLUMN', 'description'
go

exec sp_addextendedproperty 'MS_Description', N'维护人员', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_operator_work', 'COLUMN', 'operator'
go

exec sp_addextendedproperty 'MS_Description', N'评价、备注', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_operator_work', 'COLUMN', 'remarks'
go

exec sp_addextendedproperty 'MS_Description', N'编辑者', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_operator_work', 'COLUMN', 'editor'
go

exec sp_addextendedproperty 'MS_Description', N'数据来源', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_operator_work', 'COLUMN', 'source'
go

exec sp_addextendedproperty 'MS_Description', N'修改时间', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_operator_work', 'COLUMN', 'modify_time'
go

exec sp_addextendedproperty 'MS_Description', N'创建时间', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_operator_work', 'COLUMN', 'create_time'
go

--------------------------------------------------------------------------------------------------------------------------

create table ship_manager.pms_equipment_factory
(
    id                              INTEGER identity constraint PK_pms_equipment_factory primary key,
    first_code                      varchar(64),
    second_code                     varchar(64),
    third_code                      varchar(64),
    fourth_code                     varchar(64),
    description                     varchar(256),
    machinery_id                    varchar(64),
    ship_yard_code                  varchar(64),
    maker                           varchar(64),
    contact                         varchar(64),
    rating                          varchar(64),
    remarks                         varchar(512),
    model                           varchar(64),
    installed_date                  datetime,
    designed_sea_temperature        float,
    designed_ambient_temperature    float,
    score                           float,
    editor                          varchar(64),
    source                          varchar(64),
    modify_time                     datetime,
    create_time                     datetime
)
go

exec sp_addextendedproperty 'MS_Description', N'主键', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_factory', 'COLUMN', 'id'
go

exec sp_addextendedproperty 'MS_Description', N'设备第一大类编码', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_factory', 'COLUMN', 'first_code'
go

exec sp_addextendedproperty 'MS_Description', N'设备第二大类编码', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_factory', 'COLUMN', 'second_code'
go

exec sp_addextendedproperty 'MS_Description', N'设备第三大类编码', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_factory', 'COLUMN', 'third_code'
go

exec sp_addextendedproperty 'MS_Description', N'设备第四大类编码', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_factory', 'COLUMN', 'fourth_code'
go

exec sp_addextendedproperty 'MS_Description', N'名称描述', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_factory', 'COLUMN', 'description'
go

exec sp_addextendedproperty 'MS_Description', N'机器ID', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_factory', 'COLUMN', 'machinery_id'
go

exec sp_addextendedproperty 'MS_Description', N'船厂代码', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_factory', 'COLUMN', 'ship_yard_code'
go

exec sp_addextendedproperty 'MS_Description', N'制造商', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_factory', 'COLUMN', 'maker'
go

exec sp_addextendedproperty 'MS_Description', N'联系方式', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_factory', 'COLUMN', 'contact'
go

exec sp_addextendedproperty 'MS_Description', N'设备等级', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_factory', 'COLUMN', 'rating'
go

exec sp_addextendedproperty 'MS_Description', N'评价、备注', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_factory', 'COLUMN', 'remarks'
go

exec sp_addextendedproperty 'MS_Description', N'型号', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_factory', 'COLUMN', 'model'
go

exec sp_addextendedproperty 'MS_Description', N'设备安装日期', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_factory', 'COLUMN', 'installed_date'
go

exec sp_addextendedproperty 'MS_Description', N'设计海上使用温度', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_factory', 'COLUMN', 'designed_sea_temperature'
go

exec sp_addextendedproperty 'MS_Description', N'设计环境温度', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_factory', 'COLUMN', 'designed_ambient_temperature'
go

exec sp_addextendedproperty 'MS_Description', N'设备评分', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_factory', 'COLUMN', 'score'
go

exec sp_addextendedproperty 'MS_Description', N'编辑者', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_factory', 'COLUMN', 'editor'
go

exec sp_addextendedproperty 'MS_Description', N'数据来源', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_factory', 'COLUMN', 'source'
go

exec sp_addextendedproperty 'MS_Description', N'修改时间', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_factory', 'COLUMN', 'modify_time'
go

exec sp_addextendedproperty 'MS_Description', N'创建时间', 'SCHEMA', 'ship_manager', 'TABLE',
     'pms_equipment_factory', 'COLUMN', 'create_time'
go

--------------------------------------------------------------------------------------------------------------------------

create table ship_manager.pms_overhaul_status
(
    id                              INTEGER identity constraint PK_pms_overhaul_status primary key,
    first_code                      varchar(64),
    second_code                     varchar(64),
    third_code                      varchar(64),
    fourth_code                     varchar(64),
    description                     varchar(256),
    task                            varchar(64),
    office_ref                      int,
    original_due_date               datetime,
    last_action_date                datetime,
    action_result                   varchar(64),
    work_done_detail                varchar(64),
    action_complete_hours           int,
    maintenance_intervel_hours      int,
    maintenance_intervel_days       int,
    due_date                        datetime,
    editor                          varchar(64),
    source                          varchar(64),
    modify_time                     datetime,
    create_time                     datetime
)
go

--------------------------------------------------------------------------------------------------------------------------

create table ship_manager.pms_running_hours
(
    id                              INTEGER identity constraint PK_pms_running_hours primary key,
    first_code                      varchar(64),
    second_code                     varchar(64),
    third_code                      varchar(64),
    fourth_code                     varchar(64),
    description                     varchar(256),
    office_ref                      int,
    running_date                    datetime,
    total_running_hours             int,
    setup                           bit,
    avg_daily_running_hours         int,
    editor                          varchar(64),
    source                          varchar(64),
    modify_time                     datetime,
    create_time                     datetime
)
go

--------------------------------------------------------------------------------------------------------------------------

create table ship_manager.pms_task_template
(
    id                              INTEGER identity constraint PK_pms_task_template primary key,
    first_code                      varchar(64),
    second_code                     varchar(64),
    third_code                      varchar(64),
    fourth_code                     varchar(64),
    description                     varchar(256),
    task                            varchar(256),
    full_description                varchar(256),
    office_ref                      int,
    maintenance_intervel_hours      int,
    maintenance_intervel_days       int,
    total_running_hours             int,
    setup                           bit,
    avg_daily_running_hours         int,
    editor                          varchar(64),
    source                          varchar(64),
    modify_time                     datetime,
    create_time                     datetime
)
go

--------------------------------------------------------------------------------------------------------------------------

create table ship_manager.pms_work_done
(
    id                              INTEGER identity constraint PK_pms_work_done primary key,
    first_code                      varchar(64),
    second_code                     varchar(64),
    third_code                      varchar(64),
    fourth_code                     varchar(64),
    description                     varchar(256),
    task                            varchar(256),
    office_ref                      int,
    original_due_date               datetime,
    last_action_date                datetime,
    action_result                   varchar(256),
    detail                          varchar(256),
    speed_hours                     int,
    completed                       bit,
    editor                          varchar(64),
    source                          varchar(64),
    modify_time                     datetime,
    create_time                     datetime
)
go

--------------------------------------------------------------------------------------------------------------------------

create table ship_manager.pms_work_order
(
    id                              INTEGER identity constraint PK_pms_work_order primary key,
    task_id                         int,
    date                            datetime,
    first_code                      varchar(64),
    second_code                     varchar(64),
    third_code                      varchar(64),
    fourth_code                     varchar(64),
    description                     varchar(256),
    task                            varchar(256),
    office_ref                      int,
    original_due_date               datetime,
    last_action_date                datetime,
    action_result                   varchar(256),
    detail                          varchar(256),
    speed_hours                     int,
    completed                       bit,
    editor                          varchar(64),
    source                          varchar(64),
    modify_time                     datetime,
    create_time                     datetime
)
go

