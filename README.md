# RSS AGGREGATOR
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