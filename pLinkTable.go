
type  LinkTableNode struct
{
    pNext *LinkTableNode
}

/*
 * LinkTable Type
 */
type  LinkTable struct
{
    pHead *tLinkTableNode
    pTail *tLinkTableNode
    SumOfNode int
}
 func CreateLinkTable() *tLinkTable 
{
    pLinkTable *tLinkTable= new(tLinkTalbe)
    if pLinkTable == nil
    {
        return nil
    }
    pLinkTable.pHead = NULL
    pLinkTable.pTail = NULL
    pLinkTable.SumOfNode = 0
    return pLinkTable
}
/*
 * Delete a LinkTable
 */
func DeleteLinkTable(tLinkTable *pLinkTable) int
{
    if pLinkTable == nil
    {
        return -1;
    }
    for pLinkTable.pHead != nil
    {
        var p *tLinkTableNode = pLinkTable.pHead
        pLinkTable.pHead = pLinkTable.pHead.pNext
        pLinkTable.SumOfNode -= 1 
    }
    pLinkTable.pHead = nil
    pLinkTable.pTail = nil
    pLinkTable.SumOfNode = 0
    return 1	
}
/*
 * Add a LinkTableNode to LinkTable
 */
func AddLinkTableNode(tLinkTable *pLinkTable, tLinkTableNode * pNode) int 
{
    if pLinkTable == nil || pNode == nil
    {
        return -1
    }
    pNode.pNext = nil
    if pLinkTable.pHead == nil
    {
        pLinkTable.pHead = pNode
    }
    if pLinkTable.pTail == nil
    {
        pLinkTable.pTail = pNode
    }
    else
    {
        pLinkTable.pTail.pNext = pNode
        pLinkTable.pTail = pNode
    }
    pLinkTable.SumOfNode += 1
    return 1	
}
/*
 * Delete a LinkTableNode from LinkTable
 */
func DelLinkTableNode(tLinkTable *pLinkTable, tLinkTableNode * pNode) int 
{
    if pLinkTable == nil|| pNode == nil
    {
        return -1
    }
    if(pLinkTable.pHead == pNode)
    {
        pLinkTable.pHead = pLinkTable.pHead.pNext
        pLinkTable.SumOfNode -= 1
        if pLinkTable.SumOfNode == 0
        {
            pLinkTable.pTail = nil
        }
        return 1
    }
   	var pTempNode *tLinkTableNode = pLinkTable.pHead
    for pTempNode != nil
    {    
        if pTempNode.pNext == pNode
        {
            pTempNode.pNext = pTempNode.pNext.pNext
            pLinkTable.SumOfNode -= 1
            if pLinkTable.SumOfNode == 0
            {
                pLinkTable.pTail = nil
            }
            return 1			    
        }
        pTempNode = pTempNode.pNext
    }
    return -1		
}

/*
 * get LinkTableHead
 */
 func GetLinkTableHead(tLinkTable *pLinkTable) *tLinkTableNode
{
    if pLinkTable == nil
    {
        return nil
    }    
    return pLinkTable.pHead
}

/*
 * get next LinkTableNode
 */
func GetNextLinkTableNode(tLinkTable *pLinkTable, tLinkTableNode * pNode) * tLinkTableNode
{
    if pLinkTable == nil || pNode == nil
    {
        return nil
    }
    var pTempNode  * tLinkTableNode = pLinkTable.pHead
    for pTempNode != nil
    {    
        if pTempNode == pNode
        {
            return pTempNode.pNext			    
        }
        pTempNode = pTempNode.pNext
    }
    return nil
}