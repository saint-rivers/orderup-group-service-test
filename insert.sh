#!/bin/bash

curl -X POST http://localhost:3000/books \
    -d '{"name":"mb","description":"dinner party","created_at":"2022-01-01T12:12:12Z"}'
