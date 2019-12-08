package path

type CQueue struct {
	pHead *QueueElement
	pLastAccess *QueueElement
}

func NewCQueue() *CQueue {
	return &CQueue{}
}

/**
 * 将QueueElement根据eWeight由小到大的顺序插入队列
 * @param newElement
 */
func (c *CQueue) EnQueue(newElement QueueElement) {
	pCur := c.pHead
	var pPre *QueueElement
	for &pCur != nil && pCur.Weight < newElement.Weight {
		pPre = pCur
		pCur = pCur.Next
	}
	newElement.Next = pCur
	if pPre == nil {
		c.pHead = &newElement
	}else{
		pPre.Next = &newElement
	}
}

/**
 * 从队列中取出前面的一个元素
 * @return
 */
func (c *CQueue) DeQueue() QueueElement {
	if c.pHead == nil {
		return QueueElement{}
	}
	pRet := *c.pHead
	c.pHead = c.pHead.Next
	return pRet
}

/**
 * 读取第一个元素，但不执行DeQueue操作
 * @return
 */
func (c *CQueue) GetFirst() QueueElement {
	c.pLastAccess = c.pHead
	return *c.pLastAccess
}

/**
 * 读取上次读取后的下一个元素，不执行DeQueue操作
 * @return
 */
func (c *CQueue) GetNext() QueueElement {
	if c.pLastAccess == nil {
		c.pLastAccess = c.pLastAccess.Next
	}
	return *c.pLastAccess
}

func (c *CQueue) CanGetNext() bool {
	return c.pLastAccess.Next != nil
}

func (c *CQueue) Clear() {
	c.pHead = nil
	c.pLastAccess = nil
}
