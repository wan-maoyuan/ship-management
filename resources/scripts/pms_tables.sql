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

--------------------------------------------------------------------------------------------------------------------------

create table ship_manager.pms_job_order
(
    id                              INTEGER identity constraint PK_pms_job_order primary key,
    vessel_name                     varchar(64),
    vessel_imo                      varchar(64),
    order_id                        varchar(64),
    task                            varchar(512),
    first_code                      varchar(64),
    second_code                     varchar(64),
    third_code                      varchar(64),
    fourth_code                     varchar(64),
    equipment                       varchar(64),
    equipment_component             varchar(64),
    remarks                         varchar(256),
    planned_start_time              datetime,
    create_time                     datetime,
    completed                       bit,    
    deleted                         bit
)
go

--------------------------------------------------------------------------------------------------------------------------

create table ship_manager.pms_work_done
(
    id                              INTEGER identity constraint PK_pms_work_done primary key,
    vessel_name                     varchar(64),
    vessel_imo                      varchar(64),
    order_id                        varchar(64),
    task                            varchar(512),
    first_code                      varchar(64),
    second_code                     varchar(64),
    third_code                      varchar(64),
    fourth_code                     varchar(64),
    equipment                       varchar(64),
    equipment_component             varchar(64),
    result                          varchar(256),
    remarks                         varchar(256),
    speed_hours                     int,
    planned_start_time                      datetime,
    end_time                        datetime
)
go


--------------------------------------------------------------------------------------------------------------------------

create table ship_manager.pms_over_due_order
(
    id                              INTEGER identity constraint PK_pms_over_due_order primary key,
    vessel_name                     varchar(64),
    vessel_imo                      varchar(64),
    order_id                        varchar(64),
    task                            varchar(512),
    first_code                      varchar(64),
    second_code                     varchar(64),
    third_code                      varchar(64),
    fourth_code                     varchar(64),
    equipment                       varchar(64),
    equipment_component             varchar(64),
    reason                          varchar(256),
    remarks                         varchar(256),
    speed_hours                     int,
    planned_start_time              datetime
)
go















































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

--------------------------------------------------------------------------------------------------------------------------

create table ship_manager.pms_defect_list
(
    id                  INTEGER identity constraint PK_pms_defect_list primary key,
    vessel_name         varchar(64),
    vessel_imo          varchar(64),
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



