[![Go](https://github.com/burakdenizz/getirAssignment/actions/workflows/go.yml/badge.svg)](https://github.com/burakdenizz/getirAssignment/actions/workflows/go.yml)
This is a simple two end point backend service for getir case study


# How To Run The Application
- Make sure that you have docker installed
- Make sure the port 8000 is unoccupied in docker
- Open a Terminal in the project Folder
- Simply run `docker-compose up` command

# Different Endpoints
- ## Get Key Value Pair
    - A Function for getting the specific key value pair
    - The function will take a single parameter which is Key to
    - An example curl can be found below
  - ```bash
      curl --location --request GET 'http://127.0.0.1:8000/api/in-memory' --header 'Content-Type: application/json' --data-raw '{"key": "Hello"}'
     ```
- ## Set Key Value Pair
    - A Function for setting a new key value pair if key does not exist and, if the key exist the value will be updated
    - This function will take a key value pair as a parameter
    - This function has the same endpoint with the Get Key Value pair the only difference is that this function uses post as the method
    - An example curl can be found below
    - ```bash
      curl --location --request POST 'http://127.0.0.1:8000/api/in-memory' --header 'Content-Type: application/json' --data-raw '{"key": "Hello","value": "World"}'
      ```
- ## Get Data From Mongo DB
   - A function for getting the data from mongo db using two filters
   - The first filter is filters the data according to the date it is created. 
   - The second filter filters the data according to the sum of the counts field.
  - An example curl can be found below
  - ```bash
      curl --location --request GET 'http://127.0.0.1:8000/api/get-data' --header 'Content-Type: application/json' --data-raw '{"startDate": "2016-01-26","endDate": "2021-10-12","gte":"2000","lte":"3000"}'
     ```