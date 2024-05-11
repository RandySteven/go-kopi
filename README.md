# Go Framework

<p>Hi, </p>

| Directory    | Description                                                                                                                                    |
|--------------|------------------------------------------------------------------------------------------------------------------------------------------------|
| apps         | the runer of handlers, middlewares, routers, and usecases, if there is any new API need to regis routers, usecases, and handlers               |
| cmd          | the main runner of application it self, there is main for API, drop for drop the db, migration for migrate the db, and seed to seed some datas |
| entities     | the model of application and payload for request and response data                                                                             |
| enums        |                                                                                                                                                |
| files        | for external files or configuration file such as `YAML` and `.env`                                                                                 |
| handlers     | the entrance of front end or the retreiver of request and return response                                                                      |
| interfaces   |                                                                                                                                                |
| middlewares  |                                                                                                                                                |
| pkg          |                                                                                                                                                |
| queries      | the SQL query for application                                                                                                                  |
| repositories | the DB connection to application                                                                                                               |
| usecases     | the business logic of application                                                                                                              |
|utils| the function to help the business logic or application                                                                                         |

To run the application you can run
``make run``


If you want to make a model you can run 
``make model``
