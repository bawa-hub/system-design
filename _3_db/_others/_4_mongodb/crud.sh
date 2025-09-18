# show databases;
show databases;
show dbs;

# use database;
use <db_name>

# show collections;
show collections;
show tables;

# insert single 
db.products.insertOne({
    title: 'Sample product', 
    price: 99.99,
    quantity: 10,
    description: 'Sample descritption'
    })

db.products.insertOne({
    title: 'Sample product', 
    price: 99.99,
    quantity: 10,
    description: 'Sample descritption',
    _id: 'product-1a'
    })    

# insert multiple
db.categories.insertMany([
    {
        title: 'category 1',
        description: 'description',
        productsCount: 10
    },
    {
        title: 'category 2',
        description: 'description',
        productsCount: 20
    }
])

# show all
db.products.find()
db.products.find().pretty()

# find by id
db.products.findOne({_id: 'product-1a'})

# delete one
db.products.deleteOne({_id: 'product-1a'})

# delete Many
db.products.deleteMany({title: 'sample category'})

# delete all
db.products.deleteMany({})

# delete collection/table
db.products.drop()

# update one
db.products.updateOne({_id: 'product-4a'}, {$set: {ratings: 4.5}}) # to add other field

# update many
db.products.updateOne({}, {$set: {ratings: 4.5}}) # to add other field
