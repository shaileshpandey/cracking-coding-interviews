CREATE TABLE emp (employee_id,name,salary,manager_id)
AS
  VALUES
    ( 1  , 'james' ,  10000 , null ),
    ( 2  , 'alex'  ,   5000 , 1 ),
    ( 3  , 'Alice' ,   4500 , 1 ),
    ( 4  , 'Jone'  ,   3000 , 3 ),
    ( 5  , 'Omar'  ,   2200 , 2 );
    
   
WITH RECURSIVE EmployeeHierarchy AS (
    SELECT employee_id, name, manager_id, null as manager_name, 1 AS level
    FROM emp
    WHERE manager_id IS NULL -- Assuming top-level managers have NULL manager_id
    UNION ALL
    SELECT e.employee_id, e.name, e.manager_id, eh.name as manager_name, eh.level + 1
    FROM emp e
    INNER JOIN EmployeeHierarchy eh ON e.manager_id = eh.employee_id
)
SELECT *
FROM EmployeeHierarchy;
