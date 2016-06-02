#/bin/bash

go build
./generator -i lead -t template.json -o ../practice/lead
./generator -i rhythm1 -t template.json -o ../practice/rhythm1
./generator -i rhythm2 -t template.json -o ../practice/rhythm2
