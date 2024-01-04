**Generate brycpt on cli**
```
$ htpasswd -bnBC 10 "" Ch0ngeme! | tr -d ':\n'
```

-b takes the password from the second command argument

-n prints the hash to stdout instead of writing it to a file

-B instructs to use bcrypt

-C 10 sets the bcrypt cost to 10

The bare htpasswd command outputs in format <name>:<hash> followed by two newlines. 
Hence the empty string for name and tr stripping the colon and newlines.


Start a postgres database instance
$ ./dbstart.sh

Initialize the database
$ ./dbinit.sh

To test if postgres is running properly:
$ psql -h localhost -p 5432 -d keepmdb -U keepm


To Start/Stop container
$ docker stop postgres_db
$ docker start postgres_db

To destroy the instance and start over
$ docker rm -f $(docker ps -a -q)

or
$ docker rm -f container-id
