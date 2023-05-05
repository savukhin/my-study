import os
import urllib
import requests
import pytest

API_ADDRESS = os.getenv("API_ADDRESS", "http://localhost:8080")
EXECUTE_QUERY_ADDRESS = urllib.parse.urljoin(API_ADDRESS, "api/v1/execute-query")

def test_queries():
    query1 = '''
BEGIN;
    CREATE TABLE musicians(name, surname, band);
    INSERT INTO musicians(surname, band, name) values (Shinoda, LinkinPark, Mike);
    INSERT INTO musicians(name, surname, band) values (Chester, Bennington, LinkinPark);
    INSERT INTO musicians(name, surname, band) values (Maybe, Baby, MaybeBaby);
COMMIT;
SELECT * FROM musicians;
'''

    r = requests.post(EXECUTE_QUERY_ADDRESS, json={"query": query1})
    
    assert r.status_code == 200
    assert r.content.decode() == "\n\nname,surname,band\nMike,Shinoda,LinkinPark\nChester,Bennington,LinkinPark\nMaybe,Baby,MaybeBaby\n"
    
    assert False

if __name__ == "__main__":
    pytest.main()
