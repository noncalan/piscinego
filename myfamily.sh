#!/bin/bash
curl -s https://01.tomorrow-school.ai/assets/superhero/all.json | jq --argjson id "$HERO_ID" '.[] | select(.id == $id) | .connections.relatives' | sed 's/^"//; s/"$//'
