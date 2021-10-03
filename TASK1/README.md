# TASK 1

I've using docker to mockup the database

- Run bash setup-db.sh
- Using pgAdmin or any postgres client to connect
- Run this query : ```SELECT child.id, child.username, parent.username from public.users as child
LEFT JOIN public.users as parent on child.parent = parent.id```