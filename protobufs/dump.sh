#!/bin/bash

set -x
BASE="/c/Program Files (x86)/Steam/steamapps/common/Artifact/"

cd "$BASE"
PROTO_FILES=$(grep -r "proto2" -l ./)
cd -

for f in $PROTO_FILES; do
	echo "Dumping $f"
	dotnet ProtobufDumper.dll "${BASE}$f"
done

#dotnet ProtobufDumper.dll ${BASE}game/dcg/bin/win64/client.dll
