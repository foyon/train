/**
* @File:          subject.go
* @Aim:           发布事件主体，（类状态变更）
				  发布订阅模式,比如商品上架通知关注此商品的消费者
				  商品是主体事件，消费者是观察者，商品库存信息是主体
				  事件变更需要通知消费者
* @Author:        foyon
* @Created Time:  2020-09-30
*/
package main

//import ()

/*约定事件主体func*/
type subject interface {
	//注册观察者
	register(Observer observer)
	//注销观察者
	deregister(Observer observer)
	//通知
	notifyAll()
}

type item struct {
	//需要通知列表
	observerList []observer
	name         string
	//标记状态
	inStock bool
}

//具体主体事件名称
func newItem(name string) *item {
	return &item{
		name: name,
	}
}

//注册加入观察者列表
func (i *item) register(o observer) {
	i.observerList = append(i.observerList, o)
}

//注销退出观察者列表
func (i *item) deregister(o observer) {
	i.observerList = removeFromSlice(i.observerList, o)
}

func (i *item) updateSt() {
	//修改主体状态信息
	i.inStock = true
	//通知观察者们
	i.notifyAll()
}

//主体事件改变，通知所有观察者更改状态
func (i *item) notifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.name)
	}
}

//切片中移除注销的观察者
func removeFromSlice(observerList []observer, removerObserver observer) []observer { // {{{
	observerListLen := len(observerList)
	for i, observer := range observerList {
		if observer.getId() == removerObserver.getId() {
			//如果匹配到要删除的元素，将其和切片尾部交换位置
			observerList[observerListLen-1], observerList[i] = observerList[i], observerList[observerListLen-1]
			//重新从切片首部到尾部-1 即为删除过后的切片(巧妙) (原切片删除元素)
			return observerList[0 : observerListLen-1]
		}
	}

	return observerList
} // }}}
