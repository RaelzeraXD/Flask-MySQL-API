from flask import Flask, jsonify, request
from pymongo import MongoClient
from bson.objectid import ObjectId

app = Flask(__name__)
connection = MongoClient("mongodb://root:pass@mongo:27017")
db = connection.flaskdb.table

@app.route('/', methods=['GET'])
def index():
    return "Welcome to my flask rest api, the operations are available in the following endpoints /users /users/:id /create /update/:id /delete/:id"

@app.route('/users', methods=['GET'])
def get_users():
    users = db.find()  # select * from table
    
    # i made this to make the output pretty instead of using dumps
    user_list = []
    for user in users:
        user_json = {
            "id": str(user['_id']),
            "name": user['name'],
            "age": user.get('age')
        }
        user_list.append(user_json)
    
    response = {
        "users": user_list
    }
    return jsonify(response)

@app.route('/users/<id>', methods=['GET'])
def getuserbyid(id):
    user = db.find_one({"_id":ObjectId(id)})
    user_json={
        "id": str(user['_id']),
        "name": user['name'],
        "age": user.get('age')}
    
    return jsonify({"user":user_json})

@app.route('/create', methods=['POST'])
def insert_data():
    data = request.get_json()
    if data:  # Checking if JSON data is provided
        inserted_id = db.insert_one(data).inserted_id
        return jsonify({"message": "user created successfully", "inserted_id": str(inserted_id)})
    else:
        return jsonify({"error": "No data provided"}), 400

@app.route('/update/<id>', methods=['PUT'])
def update(id):
    if not ObjectId.is_valid(id):
        return jsonify({"error": "Invalid id"}),400
    
    updated_data = request.get_json()

    result = db.update_one({'_id': ObjectId(id)}, {'$set': updated_data})
    return jsonify({"message": "User updated successfully"})

@app.route('/delete/<id>', methods=['DELETE'])
def delete(id):
    db.delete_one({"_id": ObjectId(id)})
    return jsonify({"success":"user deleted successfully"}) 

app.run(host="0.0.0.0",port=8080,debug=True)