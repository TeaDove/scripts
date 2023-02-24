#!/bin/bash

exec uvicorn main:app --host=localhost --port 8001
