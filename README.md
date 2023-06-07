# Customer Relationship Management (CRM) Service

You are working in an e-commerce company and single handedly responsible with a CRM ( customer relationship management )
with the user management data. The user data should synchronize with user data from another system which from 
https://reqres.in/api/users?page=2,

## Context
For the MVP, the requirements of CRM service are defined below
User source data from https://reqres.in/api/users?page=2
There is more than one admin and only one super admin
The actor who can access the services is admin with role admin and super admin with role super-admin.
User role is the customer.

##  Docs
find the api documentation at the following postman collection [link](https://documenter.getpostman.com/view/27694518/2s93sZ8EoZ)

##  Data Flow Diagram
there are 4 table that dedicated on these services which are `customer`
`roles`, `actor` and `approval`. For
complete explaination of database structure you can follow this ddf.
if you want to know more, you can access the following [link](https://github.com/lathief/mini-project-dibimbing/tree/main/mini-project-sql).

![data flow diagram](https://github.com/lathief/mini-project-dibimbing/blob/main/mini-project-sql/erd-crm.png?raw=true)