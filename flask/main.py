import mysql.connector
from flask import Flask ,render_template,redirect,request

mydb = mysql.connector.connect(
    host='mysqlflask',
    user='root',
    password='pass',
    database='pydb',
)
app = Flask(__name__)
mycursor = mydb.cursor()

@app.route("/")
def index():
    return render_template("index.html")

@app.route('/registration',methods=['GET','POST'])
def registration():
    if request.method == 'POST':
        name = request.form.get('name')
        age = request.form.get('age')
        email = request.form.get('email')
        command='INSERT INTO pytable (name,age,email) VALUES (%s , %s, %s)'
        mycursor.execute(command,(name,age,email))
        mydb.commit()
        return render_template('index.html')
    return render_template("registration.html")

@app.route('/table', methods=['GET'])
def table():
    mycursor.execute("SELECT * FROM pytable") 
    lista=mycursor.fetchall()
    return render_template('table.html',lista=lista)

@app.route('/delete/<int:id>')
def delete(id):
    command='Delete from pytable where id = %s'
    mycursor.execute(command,(id,))
    mydb.commit()
    return render_template('index.html')

@app.route('/update/<int:id>',methods=['GET','POST'])
def update(id):
    mycursor.execute("SELECT * FROM pytable") 
    lista=mycursor.fetchall()
    if request.method == 'POST':
        name = request.form.get('name')
        age = request.form.get('age')
        email = request.form.get('email')
        command = 'UPDATE pytable SET name = %s, age = %s,email = %s WHERE id = %s'
        mycursor.execute(command,(name,age,email,id))
        mydb.commit()
        return render_template('index.html')
    return render_template('update.html',lista=lista)



app.run(host="0.0.0.0")
mycursor.close()
mydb.close()
