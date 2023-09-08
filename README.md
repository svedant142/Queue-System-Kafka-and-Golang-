# Queue-System-Kafka-and-Golang-

A queuing system using Go programming language and Kafka.
It includes unit, integration, benchmark tests as well.

In this simplified scenario the system includes the following parts:

API: An API where it should receive a item data and store in the database, below are the parameters that should be passed in the API - /user/item?

user_id (create users table and primary key of that table)
name
description (text)
images (array of image urls)
price (Number) Producer: After storing the item details in the database, id should be passed on to the message queue. Consumer: based on the id, images should be downloaded and compressed and stored in local. After storing, a local location path should be added as an array value in the products table in the compressed_product_images column.

Database Schema:

1. Users: (data for the table can be added manually)

id - int, primary key
name - Name of the users
mobile - Contact Number of the user
latitude - Latitude of the user’s location
longitude - Longitude of the user’s location
created_at
updated_at

2. Products: (Data should be added from the API Only)
   id - int, primary key
   name - string, Name of the product
   description - text, About your product
   images - array
   price - number
   compressed_product_images - array
   created_at
   updated_at

Testing -

The project contains unit tests , integration tests and benchmark test.
