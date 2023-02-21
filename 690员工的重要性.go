package GoLeetCode

/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */

 func getImportance(employees []*Employee, id int) int {
    m, queue, res := map[int]*Employee{}, []int{id}, 0  
    for _, e := range employees {
        m[e.Id] = e
    } 
    for len(queue) > 0 {
        e := m[queue[0]]
        queue = queue[1:]
        if e == nil {
            continue
        }
        res += e.Importance
        for _, id := range e.Subordinates {
            queue = append(queue, id)
        }
    } 
    return res
}
