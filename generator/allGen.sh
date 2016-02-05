#/bin/bash

go build
./generator -i lead -t template.json -o ../practice/lead
./generator -i rhythm -t long_template.json -o ../practice/rhythm
./generator -i rhythm -t short_template.json -o ../practice/short_rhythm
