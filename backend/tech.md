# Major Database Refactoring

## The first direct database access

The first direct database access is now gone. It was replaced with a new query interface structure.
The first direct concrete implementation was:

commit: `4471caa59b4a4ab9656c44f38a40849e277af378`
Author: adaniel <adiozdaniel@gmail.com>
Date:   Thu Feb 6 14:24:43 2025 +0300

    chores(handlers): add database implementations
    - Add structs for encoding responses and requests.


## The second direct database access

This is a refactored approach of the first direct database access.
At this point, the interfaces, services and seperate concrete implementations are in place but not yet consumed.

commit a14dcb7f47f056a4b2d4cbc149a49ca42e035b51 (HEAD -> bkdIntergration, origin/bkdIntergration)
Author: adioz <adiozdaniel@gmail.com>
Date:   Sat Feb 8 09:27:23 2025 +0300

    refactor(database): change database reference name
    - Changed `forumapp.database.go` database reference to Query.
    - These gives a more readable approach. For instance, to access it: `forumapp.Db.Query` implements a more straight forward approach to reference the queries provided by the database.
