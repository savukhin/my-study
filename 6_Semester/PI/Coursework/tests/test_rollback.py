import os
import urllib.parse
import requests
import pytest
import json

API_ADDRESS = os.getenv("API_ADDRESS", "http://localhost:8080")
EXECUTE_QUERY_ADDRESS = urllib.parse.urljoin(API_ADDRESS, "api/v1/execute-query")

def test_insert():
    query = '''
BEGIN;
    CREATE TABLE musicians(name, surname, band);
    INSERT INTO musicians(surname, band, name) values (Shinoda, LinkinPark, Mike);
COMMIT;
BEGIN;
    INSERT INTO musicians(name, surname, band) values (Chester, Bennington, LinkinPark);
COMMIT;
BEGIN;
    INSERT INTO musicians(name, surname, band) values (Maybe, Baby, MaybeBaby);
COMMIT;
SELECT * FROM musicians;
'''

    r = requests.post(EXECUTE_QUERY_ADDRESS, json={"query": query})
    print(r.content)
    assert r.status_code == 200
    parsed = json.loads(r.content.decode())
    assert parsed["message"] == "ok\n\nok\n\nok\n\nname,surname,band\nMike,Shinoda,LinkinPark\nChester,Bennington,LinkinPark\nMaybe,Baby,MaybeBaby\n"
    
    query = '''
ROLLBACK;
SELECT * FROM musicians;
'''
    r = requests.post(EXECUTE_QUERY_ADDRESS, json={"query": query})
    print(r.content)
    assert r.status_code == 200
    parsed = json.loads(r.content.decode())
    assert parsed["message"] == "ok\n\nname,surname,band\nMike,Shinoda,LinkinPark\nChester,Bennington,LinkinPark\n"
    
    query = '''
ROLLBACK;
SELECT * FROM musicians;
'''
    r = requests.post(EXECUTE_QUERY_ADDRESS, json={"query": query})
    print(r.content)
    assert r.status_code == 200
    parsed = json.loads(r.content.decode())
    assert parsed["message"] == "ok\n\nname,surname,band\nMike,Shinoda,LinkinPark\n"
    
    query = '''
ROLLBACK;
SELECT * FROM musicians;
'''
    r = requests.post(EXECUTE_QUERY_ADDRESS, json={"query": query})
    print(r.content)
    assert r.status_code == 400
    parsed = json.loads(r.content.decode())
    assert parsed["message"] == "ok\n\nno such table musicians"


if __name__ == "__main__":
    pytest.main()
