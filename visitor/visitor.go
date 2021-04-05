package visitor

import "fmt"

type Visitor interface {
	VisitEngineer(engineer *Engineer)
	VisitProductManager(productManager *ProductManager)
}

type Employee interface {
	Accept(visitor Visitor)
}

type Engineer struct {
	Name string
}

func (engineer *Engineer) Accept(visitor Visitor) {
	visitor.VisitEngineer(engineer)
}

type ProductManager struct {
	Name string
}

func (productManager *ProductManager) Accept(visitor Visitor) {
	visitor.VisitProductManager(productManager)
}

type FirstVisitor struct {
	Name string
}

func (firstVisitor *FirstVisitor) VisitEngineer(engineer *Engineer) {
	fmt.Println("我是", firstVisitor.Name, "我关注", engineer.Name, "的考勤")
}

func (firstVisitor *FirstVisitor) VisitProductManager(productManager *ProductManager) {
	fmt.Println("我是", firstVisitor.Name, "我关注", productManager.Name, "的考勤")
}

type SecondVisitor struct {
	Name string
}

func (secondVisitor *SecondVisitor) VisitEngineer(engineer *Engineer) {
	fmt.Println("我是", secondVisitor.Name, "我关注", engineer.Name, "的代码产出")
}

func (secondVisitor *SecondVisitor) VisitProductManager(productManager *ProductManager) {
	fmt.Println("我是", secondVisitor.Name, "我关注", productManager.Name, "的产品的故事产出")
}

func Run() {

	EmployeeArr := []Employee{}
	EmployeeArr = append(EmployeeArr, &ProductManager{Name: "子贺"})
	EmployeeArr = append(EmployeeArr, &Engineer{Name: "再哥"})

	leader1 := &FirstVisitor{Name: "人力资源"}
	leader2 := &SecondVisitor{Name: "CTO"}

	for _, employee := range EmployeeArr {
		employee.Accept(leader1)
		employee.Accept(leader2)
	}

}
