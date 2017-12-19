I chose to use Go with a mysql database because I've had expereince using them
in previous personal projects.

If I had've had time to finish I would've liked to have added some middleware
for authentication and logging. I also would have like to add a json config file
and parsed into into a config object in Go and used it to set the database connection
string.

If I'd had time my method for performing the sort by name or email would have been
to add a query string to /people such as /people?sort=email or /people?sort=name.

I then would have looked for this query string in the code and sorted the array
of people appropriately.
