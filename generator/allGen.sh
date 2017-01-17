#/bin/bash

go build
./generator -i lead -t template.json -o ../practice/lead
./generator -i rhythm -t template.json -o ../practice/rhythm
./generator -i ukelele -t template.json -o ../practice/ukulele
