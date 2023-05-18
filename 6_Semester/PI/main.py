import argparse
import sys
class MyParser(argparse.ArgumentParser):
    def error(self, message):
        sys.stderr.write('error: %s\n' % message)
        self.print_help()
        sys.exit(2)
        
parser = MyParser()

parser.add_argument('--username', type=str, help='Username (default: PISql)', default="PISql")
parser.add_argument('--password', type=str, help='password (not mention if you want to input in next step)', default=None)

args = parser.parse_args()

password = input("Input password:")

help = input(">>> ")

print(
'''
add user USERNAME password PASSWORD                 Adds user with username and password
info                                                Prints information about database
create table TABLENAME (COLUMN1, ...)               Creates table with columns in brackets
insert into TABLENAME values (VALUE1, ...)          Insert into table TABLENAME values in bracket (should match table columns size)
select COLUMN1, ... from TABLENAME                  Get all rows of table TABLENAME with specific column
select COLUMN1, ... from TABLENAME limit LIMIT      Get all rows of table TABLENAME with specific column limited by LIMIT
select COLUMN1, ... from TABLENAME where CONDITION  Get all rows of table TABLENAME with specific column by condition CONDITION
drop table TABLENAME                                Drops table TABLENAME
delete from TABLENAME where CONDITION               Delete rows from table where rows satisfies condition
begin TRANSACTION                                   Starts TRANSACTION
commit TRANSACTION                                  Commits TRANSACTION end
commit                                              Commits state
rollback                                            Rollbacks to last commit
'''
)