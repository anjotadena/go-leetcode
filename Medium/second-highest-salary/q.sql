SELECT 
    MAX(salary) as T1
FROM
    Employee
WHERE
    salary < (SELECT MAX(salary) FROM Employee)