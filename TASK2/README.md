# TASK 2

Using Onion Architecture with Domain Driven Design 

1. **How to run** 
    - ```go run ./server/main.go```

2. **Endpoint**
    - http://localhost:8080/detail
        - query params : 
            - ```id``` : id of the movie
    - http://localhost:8080/search
        - query params : 
            - ```s``` : for searching according to the title
            - ```page``` : choose the page

3. **Unit Test**
    - Currently just completing application/services layer unit test

4. **Test gRPC**
    - ```go run ./test-rpc-client/main.go```