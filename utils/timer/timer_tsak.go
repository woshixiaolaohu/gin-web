package timer

import (
	"github.com/robfig/cron/v3"
	"sync"
)

type Timer interface {
	// FindCronList 寻找所有 cron
	FindCronList() map[string]*taskManager
	// AddTaskByFuncWithSeconds 添加 Task 方法形式，以秒的形式加入
	AddTaskByFuncWithSeconds(cronName string, spec string, fun func(), taskName string, option ...cron.Option) (cron.EntryID, error)
	// AddTaskByJobWithSeconds 添加 Task 接口形式，以秒的形式传入
	AddTaskByJobWithSeconds(cronName string, spec string, job interface{ Run() }, taskName string, option ...cron.Option) (cron.EntryID, error)
	// AddTaskByFunc 通过函数的方法添加任务
	AddTaskByFunc(cronName string, spec string, fun func(), taskName string, option ...cron.Option) (cron.EntryID, error)
	// AddTaskByJob 通过接口的方法添加任务 要实现一个带有 Run 方法的接口触发
	AddTaskByJob(cronName string, spec string, job interface{ Run() }, taskName string, option ...cron.Option) (cron.EntryID, error)
	// FindCron 获取对应 taskName 的 cron 可能为空
	FindCron(cronName string) (*taskManager, bool)
	// StartCron 指定 cron 开始执行
	StartCron(cronName string)
	// StopCron 执行 cron 停止执行
	StopCron(cronName string)
	// FindTask 查找指定 cron 下指定的 task
	FindTask(cronName string, taskName string) (*task, bool)
	// RemoveTaskByID 根据 ID 删除指定 cron 下指定 task
	RemoveTaskByID(cronName string, id int)
	// RemoveTaskByName 根据 taskName 删除指定 cron 下指定 task
	RemoveTaskByName(cronName string, tasName string)
	// Clear 清理掉指定 cronName
	Clear(cronName string)
	// Close 停止所有 cron
	Close()
}

type task struct {
	EntryID  cron.EntryID
	Spec     string
	TaskName string
}

type taskManager struct {
	cron  *cron.Cron
	tasks map[cron.EntryID]*task
}

// timer 定时任务管理
type timer struct {
	cronList map[string]*taskManager
	sync.Mutex
}

// FindCronList 获取所有任务列表
func (t *timer) FindCronList() map[string]*taskManager {
	t.Lock()
	defer t.Unlock()
	return t.cronList
}

// AddTaskByFuncWithSeconds 通过函数的方法使用WithSeconds添加任务
func (t *timer) AddTaskByFuncWithSeconds(cronName string, spec string, fun func(), taskName string, option ...cron.Option) (cron.EntryID, error) {
	t.Lock()
	defer t.Unlock()
	option = append(option, cron.WithSeconds())
	if _, ok := t.cronList[cronName]; !ok {
		tasks := make(map[cron.EntryID]*task)
		t.cronList[cronName] = &taskManager{
			cron:  cron.New(option...),
			tasks: tasks,
		}
	}
	id, err := t.cronList[cronName].cron.AddFunc(spec, fun)
	t.cronList[cronName].cron.Start()
	t.cronList[cronName].tasks[id] = &task{
		EntryID:  id,
		Spec:     spec,
		TaskName: taskName,
	}
	return id, err
}

// AddTaskByJobWithSeconds 通过接口的方法添加任务
func (t *timer) AddTaskByJobWithSeconds(cronName string, spec string, job interface{ Run() }, taskName string, option ...cron.Option) (cron.EntryID, error) {
	t.Lock()
	defer t.Unlock()
	option = append(option, cron.WithSeconds())
	if _, ok := t.cronList[cronName]; !ok {
		tasks := make(map[cron.EntryID]*task)
		t.cronList[cronName] = &taskManager{
			cron:  cron.New(option...),
			tasks: tasks,
		}
	}
	id, err := t.cronList[cronName].cron.AddJob(spec, job)
	t.cronList[cronName].cron.Start()
	t.cronList[cronName].tasks[id] = &task{
		EntryID:  id,
		Spec:     spec,
		TaskName: taskName,
	}
	return id, err
}

// AddTaskByFunc 通过函数的方法添加任务
func (t *timer) AddTaskByFunc(cronName string, spec string, fun func(), taskName string, option ...cron.Option) (cron.EntryID, error) {
	t.Lock()
	defer t.Unlock()
	if _, ok := t.cronList[cronName]; !ok {
		tasks := make(map[cron.EntryID]*task)
		t.cronList[cronName] = &taskManager{
			cron:  cron.New(option...),
			tasks: tasks,
		}
	}
	id, err := t.cronList[cronName].cron.AddFunc(spec, fun)
	t.cronList[cronName].cron.Start()
	t.cronList[cronName].tasks[id] = &task{
		EntryID:  id,
		Spec:     spec,
		TaskName: taskName,
	}
	return id, err
}

// AddTaskByJob 通过接口的方法添加任务
func (t *timer) AddTaskByJob(cronName string, spec string, job interface{ Run() }, taskName string, option ...cron.Option) (cron.EntryID, error) {
	t.Lock()
	defer t.Unlock()
	if _, ok := t.cronList[cronName]; !ok {
		tasks := make(map[cron.EntryID]*task)
		t.cronList[cronName] = &taskManager{
			cron:  cron.New(option...),
			tasks: tasks,
		}
	}
	id, err := t.cronList[cronName].cron.AddJob(spec, job)
	t.cronList[cronName].cron.Start()
	t.cronList[cronName].tasks[id] = &task{
		EntryID:  id,
		Spec:     spec,
		TaskName: taskName,
	}
	return id, err
}

// FindCron 获取对应 cronName 的 cron 可能会为空
func (t *timer) FindCron(cronName string) (*taskManager, bool) {
	t.Lock()
	defer t.Unlock()
	v, ok := t.cronList[cronName]
	return v, ok
}

// StartCron 开始任务
func (t *timer) StartCron(cronName string) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.cronList[cronName]; ok {
		v.cron.Start()
	}
}

// StopCron 停止任务
func (t *timer) StopCron(cronName string) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.cronList[cronName]; ok {
		v.cron.Stop()
	}
}

// FindTask 获取对应 cronName 的 cron 可能会为空
func (t *timer) FindTask(cronName string, taskName string) (*task, bool) {
	t.Lock()
	defer t.Unlock()
	v, ok := t.cronList[cronName]
	if !ok {
		return nil, ok
	}
	for _, t2 := range v.tasks {
		if t2.TaskName == taskName {
			return t2, true
		}
	}
	return nil, false
}

// RemoveTaskByID 根据 ID 从 cronName 删除指定任务
func (t *timer) RemoveTaskByID(cronName string, id int) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.cronList[cronName]; ok {
		v.cron.Remove(cron.EntryID(id))
		delete(v.tasks, cron.EntryID(id))
	}
}

// RemoveTaskByName 从 cronName 使用 taskName 删除指定任务
func (t *timer) RemoveTaskByName(cronName string, taskName string) {
	fTask, ok := t.FindTask(cronName, taskName)
	if !ok {
		return
	}
	t.RemoveTaskByID(cronName, int(fTask.EntryID))
}

// Clear 清除任务
func (t *timer) Clear(cronName string) {
	t.Lock()
	defer t.Unlock()
	if v, ok := t.cronList[cronName]; ok {
		v.cron.Stop()
		delete(t.cronList, cronName)
	}
}

// Close 释放资源
func (t *timer) Close() {
	t.Lock()
	defer t.Unlock()
	for _, v := range t.cronList {
		v.cron.Stop()
	}
}

func NewTimerTask() Timer {
	return &timer{
		cronList: make(map[string]*taskManager),
	}
}
