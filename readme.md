# Fam Youtube

- This curl will be used to fetch data from youtube google api. Do not need to trigger this, it will be called automatically from server every 10 seconds. I am using grpc+protobuf.
  
`curl --location --request POST 'http://localhost:9090/api/v1/trigger-job'
  `

- This curl is used to get paginated response add page size which is limit and after_token which is empty for first time for second copy after_token in request from response of first one 

`curl --location 'http://localhost:9090/api/v1/get-paginated-records' \
--header 'Content-Type: application/json' \
--data '{
    "page_context": {
        "page_size": 5,
        "after_token":""
    }
}'`

`curl --location 'http://localhost:9090/api/v1/get-paginated-records' \
--header 'Content-Type: application/json' \
--data '{
    "page_context": {
        "page_size": 5,
        "after_token": "eyJUaW1lc3RhbXAiOnsic2Vjb25kcyI6MTczNzQ4MjE3OX0sIk9mZnNldCI6NSwiSXNSZXZlcnNlIjpmYWxzZX0="
            }
}'`


-  This curl will return the videos on the basis of title and description.
if title or description is empty it will be ignored.
It will do full text search and case should match.
Copy the title and description from db itself.
Generate url from `/Users/rishusahu/go/src/github.com/rishusahu23/fam-youtube/main/main.go`
by copying title and description, generating url is necessary else it won't work.


`curl --location --request GET 'http://localhost:9090/api/v1/records?description=ranjitrophy+%23teamindia+%23rohitsharma+%23yashasvijaiswal+%23shubmangill+From+Rohit%2C+Jaiswal+to+Gill+%26+Pant%2C+Star+Indian+Players+fails+to+...&title=From+Rohit%2C+Jaiswal+to+Gill+%26amp%3B+Pant%2C+Star+Players+fails+to+impress+on+their+Ranji+Trophy+comeback%21' \
--header 'Content-Type: application/json' \
--data '{
    "page_context": {
        "page_size": 5,
        "after_token": "eyJUaW1lc3RhbXAiOnsic2Vjb25kcyI6MTczNzQ4MjE3OX0sIk9mZnNldCI6NSwiSXNSZXZlcnNlIjpmYWxzZX0="
            }
}'`

- GetPartialMatchRecords is Optimise version of GetFilteredRecords search api, so that it's able to search videos containing partial match for the search query in either video title or description.
Ex 1: A video with title `*How to make tea?*` should match for the search query `tea how`

`curl --location --request GET 'http://localhost:9090/api/v1/partial-match/records?query=MAGA%20Plantation' \
--header 'Content-Type: application/json' \
--data '{
    "page_context": {
        "page_size": 5,
        "after_token": "eyJUaW1lc3RhbXAiOnsic2Vjb25kcyI6MTczNzQ4MjE3OX0sIk9mZnNldCI6NSwiSXNSZXZlcnNlIjpmYWxzZX0="
            }
}'`