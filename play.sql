INSERT INTO employees.employees (name, last_name, patronymic,phone, position, good_job_count) VALUES ('test_name', 'test_lastname', 'test_pat', 'test_phone', 'test_pos', 1) RETURNING id; 
SELECT employees.employee_add('test_name', 'test_lastname', 'test_pat', 'test_phone', 'test_pos', 1) RETURNING id;



SELECT employees.employee_add('test_name', 'test_lastname', 'test_pat', 'test_phone', 'test_pos', 1), CURRVAL('employees.employees_id_seq');
