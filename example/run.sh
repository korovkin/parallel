#!/bin/bash

# run forever:
cat commands.txt | ../parallel -j 10
