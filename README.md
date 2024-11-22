### Readme
The project is built by gin and postgres, takes me 1 day to finish it.
- In order to run the code, first init the postgres db by using sql file in ./sql/create.sql;
- Setup config file in ./config/config.yaml, need to config the postgres db connection by using "postgres" "url";

#### Design explanation
In my understanding, this is a custody wallet using for cex.
Have 3 tables: 
- user_balance: storing user's balance;
- transaction_history: storing users' transaction history including withdraw and transfer; 
- user_info: storing user's deposit address and email;

The api related codes is in the ./api folder and database related code in ./db folder.
For reviewing codes, focus on those codes: ./api/wallet and ./db/wallet.go.
Have over 80% test coverage for db codes(./db/db_test.go).
##### API explanation
**Withdraw**
- Post: localhost:8000/wallet/withdraw
- Requst body: 
{"email":"a@gmail.com","chain":"ethereum", 
"coin_type":"eth",
"destination":"0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
"value":"0.01"}

- Response: latest balance
{
    "status": "success",
    "data": "2.6"
}

**Deposit**
- Post: localhost:8000/wallet/deposit
- Requst body: 
{"email":"a@gmail.com","chain":"ethereum"}

- Response: deposit address with request chain
`{"status": "success","data": "0x70997970C51812dc3A010C7d01b50e0d17dc79C8"}`

**Transfer**
- Post: localhost:8000/wallet/transfer
- Requst body: 
`{"email":"a@gmail.com","chain":"ethereum", "coin_type":"eth","destination_email":"b@gmail.com","value":"0.02"}`

- Response: latest balance
`{
    "status": "success",
    "data": "2.6"
}`

**Balance**
- Get: localhost:8000/wallet/balance
- Request params: email=a@gmail.com
- Response: user balance array with chain and coin type.
`{"status": "success","data": [{"coin_type": "eth","chain": "ethereum","value": "2.580000000000000000"}]} `

**History**
- Get: localhost:8000/wallet/tx_history
- Request params: email=a@gmail.com
- Response: user transaction balance array.
`{"status": "success","data": [{"from": "a@gmail.com","to": "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266","value": "0.010000000000000000","chain": "ethereum","coin_type": "eth","timestamp": 1732178122 },{"from": "a@gmail.com","to": "b@gmail.com", "value": "0.020000000000000000","chain": "ethereum","coin_type": "eth","timestamp": 1732179336} ]}`

#### Todo task
- auth middlewares with jwt for the whole 5 apis;
- a routine service to monitor deposit balance and update user balance in database;
- a service to handle coin transfer for withdraw;
- add page parameter for tx_history;