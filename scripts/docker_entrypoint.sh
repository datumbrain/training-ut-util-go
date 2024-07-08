#!/bin/bash
go run util.go --operation merge --input-path /training-ut-util-go/input --output-path /training-ut-util-go/output

go run util.go --operation merge --input-path /training-ut-util-go/input --output-path /training-ut-util-go/output --format csv
