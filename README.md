# My Gram API
###### Final Project Digital Talent Scholarship FGA - Scalable Web Service with Golang (hacktiv8)

My Gram is an application that allows users to save photos and comment on other people's photos.

## Technology Stack
`Golang` `Gin Gonic` `PostgreSQL`  

## Features
- User
    - Register
    - Login
    - Update User Account
    - Delete User Account
- Photo
    - Upload Photo
    - Get Photos
    - Update Photo
    - Delete Photo
- Photo Comment
    - Add Comment
    - Get Comments
    - Edit Comment
    - Delete Comment
- Social Media Link
    - Add Social Media Link
    - Get Social Media Links
    - Edit Social Media Link
    - Delete Social Media Link

## Demo
**LIVE API** : `https://my-gram-production-fb18.up.railway.app`  
**API Documentation** : `https://documenter.getpostman.com/view/27861526/2sA35A8k9n`

## Installation
Follow these steps to install and run My Gram API on your local machine:
1. **Clone the repository:**

   ```bash
   git clone https://github.com/fazriegi/my-gram.git
   ```
   
2. **Move to cloned repository folder**

    ```bash
    cd my-gram
    ```
    
3. **Update dependecies**
    
    ```bash
    go mod tidy
    ```

4. **Copy `.env.example` to `.env`**

    ```bash
    cp .env.example .env
    ```

5. **Configure your `.env`**
6. **Run my-gram** 

    ```bash
    go run main.go
    ```


## Author

Fazri Egi - [Github](https://github.com/fazriegi)