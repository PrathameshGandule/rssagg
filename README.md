# RSS AGGREGATOR
## Steps to start the server
- clone the repo
```bash
git clone https://github.com/PrathameshGandule/rssagg.git
cd rssagg
```
- setup .env file
```bash
PORT=3000
DB_URL=postgres://{username}:{password}@{ip_address}:{db_port}/{db_name}?sslmode=disable
```
- go inside `sql\schema` directory and run up migrations
```bash
cd sql/schema
goose postgres {db_url} up
```
- build and run
```bash
go build && ./rssagg
```
## Routes
|Method   |Route                           |Short description             |Authenticated|
|---------|--------------------------------|------------------------------|-------------|
| `GET`   | `/healthz`                     |returns empty `200` response  |:x: 
| `GET`   | `/err`                         |standard err route checking   |:x:
| `POST`  | `/users`                       |create new user               |:x:
| `GET`   | `/users`                       |get all users                 |:white_check_mark:
| `POST`  | `/feeds`                       |create a new feed             |:white_check_mark:
| `GET`   | `/feeds`                       |get all feeds                 |:x:
| `POST`  | `/feed_follows`                |follow a feed                 |:white_check_mark:
| `GET`   | `/feed_follows`                |get all your feed follows     |:white_check_mark:
| `DELETE`| `/feed_follows/{feedFollowID}` |unfollow a feed               |:white_check_mark:
| `GET`   | `/posts?limit={num}`           |get posts from followed feeds |:white_check_mark: