package chainofresponsibility

import "fmt"

type Manager interface {
	SetNextManager(manager Manager)
	Handle(dayCount int) string
}

type TemplateManager struct {
	nextManager Manager
}

func (tm *TemplateManager) NextHandle(dayCount int) string {

	if tm.nextManager != nil {
		return tm.nextManager.Handle(dayCount)
	}

	return fmt.Sprintf("请假天数无人能审批:%d", dayCount)

}

func (tm *TemplateManager) SetNextManager(manager Manager) {
	tm.nextManager = manager
}

type LevelOneManager struct {
	TemplateManager
}

func (levelOneManager *LevelOneManager) Handle(dayCount int) string {

	if dayCount <= 3 && dayCount > 0 {
		return fmt.Sprintf("我是一级管理员，我审核通过，请假天数为：%d", dayCount)
	}

	return levelOneManager.NextHandle(dayCount)

}

type LevelTwoManager struct {
	TemplateManager
}

func (levelTwoManager *LevelTwoManager) Handle(dayCount int) string {

	if dayCount <= 7 && dayCount > 3 {
		return fmt.Sprintf("我是二级级管理员，我审核通过，请假天数为：%d", dayCount)
	}

	return levelTwoManager.NextHandle(dayCount)

}

type LevelThreeManager struct {
	TemplateManager
}

func (levelThreeManager *LevelThreeManager) Handle(dayCount int) string {

	if dayCount <= 15 && dayCount > 7 {
		return fmt.Sprintf("我是三级管理员，我审核通过，请假天数为：%d", dayCount)
	}

	return levelThreeManager.NextHandle(dayCount)

}

func Run() {
	levelOneManager := &LevelOneManager{}
	levelTwoManager := &LevelTwoManager{}
	levelThreeManager := &LevelThreeManager{}

	levelOneManager.SetNextManager(levelTwoManager)
	levelTwoManager.SetNextManager(levelThreeManager)

	for i := 0; i < 20; i++ {
		fmt.Println(levelOneManager.Handle(i))
	}

}
