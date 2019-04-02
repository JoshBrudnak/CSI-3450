# Endpoints

## Profile data
Gets the logged in user’s profile data
- Url: `/api/profile`
- Method: GET
- Include session id

## Dashboard data
Gets the list of categories with their entries for a user.
- Url: `/api/dashboard`
- Method: GET
- Include session id

## Login
Authenticates the user and returns a session id as a cookie
- Url: `/api/login`
- Method: POST
- Post body:
    ```
    Body: {
        Username: ,
        Password:
    }
    ```

## Logout
Deletes the user’s session
- Url: `/api/logout`
- Method: GET
- Include session id

## Create User
Creates a new user with the data given
- Url: `/api/createuser`
- Method: POST
- Post body:
    ```
    Body: {
        FirstName: ,
        LastName: ,
        Password: ,
        Repassword: ,
        Email: ,
        TotalMoney:
    }
    ```

## Delete User
Removes a user along with all of the user's data
- Url: `/api/remove/user`
- Method: GET
- Include session id

## Add Budget Entry
Adds a transaction in a budget category
- Url: `/api/add/transaction`
- Method: POST
- Include session id
- Post body:
    ```
    Body: {
        CategoryId: ,
        Value:
    }
    ```

## Edit User
Updates the user’s profile data
- Url: `/api/update/user`
- Method: POST
- Include session id
- Post body:
    ```
    Body: {
        FirstName: ,
        LastName: ,
        Email: ,
        TotalMoney:
    }
    ```
## Update User Password
Updates the user’s password
- Url: `/api/update/password`
- Method: POST
- Include session id
- Post body:
    ```
    Body: {
        Password: ,
        RePassword: ,
    }
    ```
