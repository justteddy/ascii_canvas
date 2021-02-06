# ASCII canvas

## Start
To run service in docker containers use `make run` or `docker-compose up`

This will start the service on port 8080

## Endpoints

### Get canvas by ID
`curl --request GET 'localhost:8080/canvas/{some-canvas-ID}'`

### Draw rectangle
If the canvas doesn't exist, a new canvas will be created, and the rectangle will be drawn on it.
If it exists, it will be modified. Only single ASCII symbols are usable as `fill` and `outline` parameters. 
One of either `fill` or `outline` should not be empty.
```shell
curl --request PUT 'localhost:8080/canvas/5/drawRectangle' \
--header 'Content-Type: application/json' \
--data-raw '{
    "x": 1,
    "y": 1,
    "width": 4,
    "height": 5,
    "fill": "-",
    "outline": "*"
}'
```

### Flood fill
If the canvas doesn't exist, a new canvas will be created, and the rectangle will be drawn on it.
If it exists, it will be modified. Only single ASCII symbols are usable as `fill` parameter.
```shell
curl --request PUT 'localhost:8080/canvas/8/floodFill' \
--header 'Content-Type: application/json' \
--data-raw '{
    "x": 5,
    "y": 5,
    "fill": "-"
}'
```

## Stop
To stop service use `make down` or `docker-compose down`

## Example
Let's create new canvas with a 2 rectangles:

First rectangle:
```shell
curl --request PUT 'localhost:8080/canvas/100/drawRectangle' \
--header 'Content-Type: application/json' \
--data-raw '{
    "x": 0,
    "y": 0,
    "width": 4,
    "height": 5,
    "fill": "@",
    "outline": "*"
}'

Result:
|*|*|*|*| | | | | | | | |
|*|@|@|*| | | | | | | | |
|*|@|@|*| | | | | | | | |
|*|@|@|*| | | | | | | | |
|*|*|*|*| | | | | | | | |
| | | | | | | | | | | | |
| | | | | | | | | | | | |
| | | | | | | | | | | | |
| | | | | | | | | | | | |
| | | | | | | | | | | | |
| | | | | | | | | | | | |
| | | | | | | | | | | | |
```

Second rectangle:
```shell
curl --request PUT 'localhost:8080/canvas/100/drawRectangle' \
--header 'Content-Type: application/json' \
--data-raw '{
    "x": 6,
    "y": 6,
    "width": 4,
    "height": 4,
    "fill": "",
    "outline": "X"
}'

Result:
|*|*|*|*| | | | | | | | |
|*|@|@|*| | | | | | | | |
|*|@|@|*| | | | | | | | |
|*|@|@|*| | | | | | | | |
|*|*|*|*| | | | | | | | |
| | | | | | | | | | | | |
| | | | | | |X|X|X|X| | |
| | | | | | |X| | |X| | |
| | | | | | |X| | |X| | |
| | | | | | |X|X|X|X| | |
| | | | | | | | | | | | |
| | | | | | | | | | | | |
```

Now fill the entire area outside both rectangles:
```shell
curl --request PUT 'localhost:8080/canvas/100/floodFill' \
--header 'Content-Type: application/json' \
--data-raw '{
    "x": 11,
    "y": 11,
    "fill": "."
}'

Result:
|*|*|*|*|.|.|.|.|.|.|.|.|
|*|@|@|*|.|.|.|.|.|.|.|.|
|*|@|@|*|.|.|.|.|.|.|.|.|
|*|@|@|*|.|.|.|.|.|.|.|.|
|*|*|*|*|.|.|.|.|.|.|.|.|
|.|.|.|.|.|.|.|.|.|.|.|.|
|.|.|.|.|.|.|X|X|X|X|.|.|
|.|.|.|.|.|.|X| | |X|.|.|
|.|.|.|.|.|.|X| | |X|.|.|
|.|.|.|.|.|.|X|X|X|X|.|.|
|.|.|.|.|.|.|.|.|.|.|.|.|
|.|.|.|.|.|.|.|.|.|.|.|.|
```