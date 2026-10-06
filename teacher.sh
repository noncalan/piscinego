#!/bin/bash
export INTERVIEW=$(sed -n '179p' streets/Buckingham_Place | tr -cd '0-9')
printf '%s\n' "$INTERVIEW"
cat "interviews/interview-$INTERVIEW"
printf '%s\n' "$MAIN_SUSPECT"
