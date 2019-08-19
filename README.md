# Clone the project
## Create a user called snitch with password snitch and database, incidentReport
Run the fllowing in the terminal:
* ``sudo -i -u postgres`` 
* ``createuser --interactive --pwprompt``
* ``createdb -O snitch incidentReport``

## Run the migrations with ``goose up`` command
