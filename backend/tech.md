# Major overhauls

## The old direct database access

The old direct database access is now gone. It was replaced with a new query interface structure.
The new interface is more flexible, secure, more scalable and easier to use.

The old direct concrete implementation was:

commit: `4471caa59b4a4ab9656c44f38a40849e277af378`
Author: adaniel <adiozdaniel@gmail.com>
Date:   Thu Feb 6 14:24:43 2025 +0300

    chores(handlers): add database implementations
    - Add structs for encoding responses and requests.
