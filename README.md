# Pokemon 'Go' demo

To create the database:
"""bash
cd db
docker build -t pokedb .
docker run --name pokedb -d -p 5432:5432 pokedb
"""

Then use main.go to connect.
