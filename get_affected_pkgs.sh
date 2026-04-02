#!/usr/bin/env bash

# Copyright The OpenTelemetry Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# get_affected_pkgs.sh
# Identifies unique Go modules containing changed files.

# Compare against the base branch (defaulting to main)
BASE_SHA=${1:-"origin/main"}
CURRENT_SHA=${2:-"HEAD"}

# 1. Get all changed files
# 2. Extract the directory path
# 3. For each directory, find the nearest parent containing a go.mod
# 4. Filter out duplicates and empty results
files=$(git diff --name-only "$BASE_SHA" "$CURRENT_SHA")

affected_modules=()

for file in $files; do
    dir=$(dirname "$file")
    
    # Climb up until we find a go.mod or hit the root
    while [[ "$dir" != "." && "$dir" != "/" ]]; do
        if [[ -f "$dir/go.mod" ]]; then
            affected_modules+=("$dir")
            break
        fi
        dir=$(dirname "$dir")
    done
    
    # Check root for go.mod if file was in root or no sub-mod found
    if [[ "$dir" == "." && -f "go.mod" ]]; then
        affected_modules+=(".")
    fi
done

echo "${affected_modules[@]}" | tr ' ' '\n' | sort -u | tr '\n' ' '