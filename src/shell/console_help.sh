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

# Help page for container-console.
function console_help() {
  # Usage:
  # console_help
  # I hope you can understand...
  echo -e "${COLOR}$(po_getmsg "Usage:")"
  echo "  $(po_getmsg "help                       :Show this page.")"
  echo "  $(po_getmsg "exit                       :Exit console.")"
  echo "  $(po_getmsg "new                        :Create a new container.")"
  echo "  $(po_getmsg "search [all/OS] (Arch)     :Search for images.")"
  echo "  $(po_getmsg "pull [OS] [Version] (Arch) :Download image,just save.")"
  echo "  $(po_getmsg "rmi [OS] [Version] (Arch)  :Remove an image.")"
  echo "  $(po_getmsg "cp [Name:Path] [Path]      :Copy files,like docker cp.")"
  echo "  $(po_getmsg "ls                         :List containers and images.")"
  echo "  $(po_getmsg "rm [Name]                  :Remove a container.")"
  echo "  $(po_getmsg "login [Name]               :Login to the container.")"
  echo "  $(po_getmsg "import [Image File]        :Import an image.")"
  echo "  $(po_getmsg "export [Name]              :Export a container as an image.")"
  echo "  $(po_getmsg "info                       :Show version && system info.")"
}
