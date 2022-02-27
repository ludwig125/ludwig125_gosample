sudo apt -y install sqlite3

# https://gitlab.com/drewolson/go_di_example/-/blob/master/create_database.sh
rm -f example.db
sqlite3 example.db 'CREATE TABLE people(id INTEGER PRIMARY KEY ASC, name TEXT, age INTEGER);'
sqlite3 example.db 'INSERT INTO people (name, age) VALUES ("drew", 35);'
sqlite3 example.db 'INSERT INTO people (name, age) VALUES ("jane", 29);'


# 以下テーブルの確認方法
# $sqlite3 example.db
# sqlite> .databases
# main: /home/ludwig125/go/src/github.com/ludwig125/ludwig125_gosample/di/di1/example.db
# sqlite> .tables
# people
# sqlite> select * from people;
# 1|drew|35
# 2|jane|29
# sqlite>

# sqlite 参考
# https://www.dbonline.jp/sqlite/database/index4.html
# https://iatlex.com/linux/first_sqlite
