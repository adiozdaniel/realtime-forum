# Forum

## Objectives

This project is a web forum with the following functionality:
- **Communication between users**: Users can interact by creating posts and comments.
- **Image Upload**: Users can also interact by uploading images to the posts they are trying to make.
- **Categorization of posts**: Posts can be associated with one or more categories.
- **Likes and dislikes**: Users can like or dislike posts and comments, with the counts visible to everyone.
- **Filtering posts**: Posts can be filtered by categories, user-created posts, and liked posts.

---

## SQLite

SQLite is used to store the forum's data (e.g., users, posts, likes, dislikes, comments). It is an embedded database software ideal for local storage in application software.

### Notes:

SQLite enables creating and controlling a database using queries. To learn more about SQLite, visit the [SQLite documentation](https://sqlite.org/).

---
### Traditional Authentication:

- **Registration**:
  - Users can register with a unique username and email.
  - A password is required during registration, and it is encrypted before storing.
  
- **Login**:
  - Users can log in using their email and password.
  - If the credentials are incorrect, an error response is returned.

### Sessions:

- User sessions are managed using **cookies** to keep users logged in.

## Communication

To facilitate communication among users:

- **Registered users**:
  - Can create posts and comments.
  - Posts can be associated with one or more categories (you decide the categories).
  
- **Non-registered users**:
  - Can only view posts and comments.

---

## Likes and Dislikes

- Only registered users can like or dislike posts and comments.
- The total number of likes and dislikes is visible to everyone (registered or not).

---

## Filter

The forum includes a filtering mechanism to:

- Filter posts by **categories** (like subforums for specific topics).
- Display posts created by the logged-in user (**created posts**).
- Display posts liked by the logged-in user (**liked posts**).

## Image Upload

The forum allows registered users to upload images to their respective posts. The image should not exceed 20 MB in size. Formats allowed are JPEG, SVG, PNG, and GIF.

Registered users can also upload profile images and/or edit their profile images.

### Notes:

- The "created posts" and "liked posts" filters are only available to registered users.

---

## Docker

**Docker** has been used to allow packaging the application and its dependencies into a container, ensuring consistent behavior across environments.

To build the image:

```
docker build -t forum .
```

Then to run the built image:

```
docker run -d -p 4000:4000 --name forumcontainer forum
```

---

## How to run the application

1. Clone the Repository:

   ```
   git clone https://learn.zone01kisumu.ke/git/josie-opondo/forum

   cd forum-authentication
   ```

2. Run the following command:

   ```
   make
   ```

   or

   ```
   go run ./cmd
   ```

3. On your Web Browser:

   ```
   localhost:4000
   ```

4. To run tests:

   ```
   make test
   ```

   or

   ```
   go test ./...
   ```

## Troubleshooting

If you experience any issues while using the application, try the following
  - Refresh your page.
  - If that doesn't work, clear your browser cache and try refreshing the page.
  - If the issue persists, [report it](mailto:josephineopondo05@gmail.com)

## Authors

[Adioz Eshitemi](https://github.com/adiozdaniel)

[Josephine Opondo](https://github.com/josie-opondo)

[Shayo Victor](https://github.com/worldofmakebelief)

[Edwin Nungo](https://github.com/NungoEdwin)
