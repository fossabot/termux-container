# Copyright 2023 moe-hacker
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
# Load configs && check the environment.
function _init() {
  # Usage:
  # _init
  if [[ ! -e /data/data/com.termux/files/usr/etc/container/global.conf ]]; then
    error "Could not find global config file!"
  fi
  source /data/data/com.termux/files/usr/etc/container/global.conf
  for i in capsh unshare chroot proot wget curl; do
    if ! command -v $i 2>&1 >/dev/null; then
      error "Missing dependencies: $i!"
    fi
  done
  if ! command -v ruri 2>&1 >/dev/null; then
    error "ruri not found, package might be broken!"
  fi
  if ! command -v container-console 2>&1 >/dev/null; then
    error "container-console not found, package might be broken!"
  fi
  kernelVersion="$(uname -r)"
  kernelMajor="${kernelVersion%%.*}"
  if ((kernelMajor < 4)); then
    echo -e "\033[33mWarning: termux-container has not been tested under linux 3.x or lower.\033[0m" >&2
  fi
}
