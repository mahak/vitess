#!/bin/bash

# Copyright 2019 The Vitess Authors.
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

# This is an example script that stops the mysqld and vttablet instances
# created by vttablet-up.sh

source "$(dirname "${BASH_SOURCE[0]:-$0}")/../env.sh"

cell=${CELL:-'test'}
uid=$TABLET_UID
printf -v alias '%s-%010d' $cell $uid
echo "Shutting down MySQL for tablet $alias..."

 #TODO: Remove underscore(_) flags in v25, replace them with dashed(-) notation
mysqlctl --tablet_uid $uid shutdown

