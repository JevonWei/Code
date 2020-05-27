package models

import (
	"encoding/json"
	"time"

	"github.com/astaxie/beego/orm"
)

const (
	LOGResource = 0X0001
)

type Log struct {
	UUID string     `json:"uuid"`
	Type int        `json:"type"`
	Msg  string     `json:"msg"`
	Time *time.Time `json:"time"`
}

type LogManager struct{}

func NewLogManager() *LogManager {
	return &LogManager{}
}

func (m *LogManager) Create(log *Log) {
	switch log.Type {
	case LOGResource:
		resource := &Resource{}
		if err := json.Unmarshal([]byte(log.Msg), resource); err == nil {
			DefaultResourceManager.Create(log, resource)
		}
	}
}

type Resource struct {
	Id          int        `orm:"column(id);" json:"id"`
	UUID        string     `orm:"column(uuid);size(64);" json:"uuid"`
	Load        string     `orm:"column(load);size(1024);" json:"load"`
	CPUPrecent  float64    `orm:"column(cpu_precent);" json:"cpu_precent"`
	RAMPrecent  float64    `orm:"column(ram_precent);" json:"ram_precent"`
	DiskPrecent string     `orm:"column(disk_precent);size(4096);" json:"disk_precent"`
	Time        *time.Time `orm:"column(time);" json:"time"`
	CreatedTime *time.Time `orm:"column(created_time);auto_now_add" json:"created_time"`
	DeletedTime *time.Time `orm:"column(deleted_time);null;" json:"deleted_time"`
	AgentObject *Agent     `orm:"-" json:"agent_object"`
}

func (m *Resource) Patch() {
	if m.UUID != "" {
		ormer := orm.NewOrm()
		a := &Agent{UUID: m.UUID}
		if err := ormer.Read(a, "UUID"); err == nil {
			m.AgentObject = a
		}
	}
}

type ResourceManager struct{}

func NewResourceManager() *ResourceManager {
	return &ResourceManager{}
}

func (m *ResourceManager) Create(log *Log, resource *Resource) {
	resource.UUID = log.UUID
	resource.Time = log.Time

	orm.NewOrm().Insert(resource)
}

func (m *ResourceManager) Query(q string, start int64, length int) ([]*Resource, int64, int64) {
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&Resource{})

	condition := orm.NewCondition()
	condition = condition.And("deleted_time__isnull", true)

	total, _ := queryset.SetCond(condition).Count()
	qtotal := total

	if q != "" {
		query := orm.NewCondition()
		query = query.Or("uuid__icontains", q)
		query = query.Or("created_time__icontains", q)
		query = query.Or("load__icontains", q)
		condition = condition.AndCond(query)
		qtotal, _ = queryset.SetCond(condition).Count()
	}

	var result []*Resource
	queryset.SetCond(condition).Limit(length).Offset(start).All(&result)
	return result, total, qtotal

}

func (m *ResourceManager) DeleteById(pk int) error {
	orm.NewOrm().QueryTable(&Resource{}).Filter("id__exact", pk).Update(orm.Params{"deleted_time": time.Now()})
	return nil
}

var DefaultLogManager = NewLogManager()
var DefaultResourceManager = NewResourceManager()

func init() {
	orm.RegisterModel(new(Resource))
}
