# 100% shit-code.
# You can't understand how this works,
# trust me, because I don't understand either.
col=$(($(($(stty size | awk '{print $2}'))) / 2 - 25))
logo="\033[?25l\033[1;38;2;254;228;208m
\033[${col}G         ●●●●● ●●●●● ●●●●  ●   ● ●   ● ●   ●
\033[${col}G           ●   ●     ●   ● ●● ●● ●   ●  ● ●
\033[${col}G           ●   ●●●●  ●●●●  ● ● ● ●   ●   ●
\033[${col}G           ●   ●     ●  ●  ●   ● ●   ●  ● ●
\033[${col}G           ●   ●●●●● ●   ● ●   ●  ●●●  ●   ●



\033[${col}G ●●●   ●●●  ●   ● ●●●●●   ●    ●●●  ●   ● ●●●●● ●●●●
\033[${col}G●   ● ●   ● ●●  ●   ●    ● ●    ●   ●●  ● ●     ●   ●
\033[${col}G●     ●   ● ● ● ●   ●   ●   ●   ●   ● ● ● ●●●●  ●●●●
\033[${col}G●   ● ●   ● ●  ●●   ●   ●●●●●   ●   ●  ●● ●     ●  ●
\033[${col}G ●●●   ●●●  ●   ●   ●   ●   ●  ●●●  ●   ● ●●●●● ●   ●
"
clear
row=$(($(($(stty size | awk '{print $1}'))) / 2 - 7))
i=0
while ((i < row)); do
  i=$((i + 1))
  echo
done
printf "$logo"
sleep 1
i=0
j=$col
k1=$((row + 2))
k2=$((row + 10))
x=0
while ((x < 5)); do
  i=0
  j=$col
  while ((i < 54)); do
    i=$((i + 1))
    printf "\033[${k1};${j}H "
    printf "\033[${k2};${j}H "
    j=$((j + 1))
    sleep 0.01
  done
  k1=$((k1 + 1))
  k2=$((k2 + 1))
  x=$((x + 1))
done
clear
row=$((row + 7))
i=0
while ((i < row)); do
  i=$((i + 1))
  echo
done
x=$(stty size | awk '{print $2}')
echo $(yes "●" | sed $x'q' | tr -d '\n')
row=$((row + 1))
k=0
while ((k < 3)); do
  j=0
  i=0
  printf "\033[${row};${j}H\033[1;38;2;135;206;250m●●●●●●●●●●●●"
  while ((i < x - 12)); do
    i=$((i + 1))
    printf "\033[${row};${j}H\033[1;38;2;254;228;208m●"
    j=$((j + 1))
    printf "\033[${row};${j}H\033[1;38;2;135;206;250m●●●●●●●●●●●●●"
    sleep 0.01
  done
  while ((i <= x)); do
    i=$((i + 1))
    printf "\033[${row};${j}H\033[1;38;2;254;228;208m●"
    j=$((j + 1))
    sleep 0.01
  done
  k=$((k + 1))
done
i=0
j=0
while ((i <= x)); do
  printf "\033[${row};${j}H\033[1;38;2;135;206;250m●"
  sleep 0.01
  i=$((i + 1))
  j=$((j + 1))
done
clear
col=$(($(($(stty size | awk '{print $2}'))) / 2 - 15))
row=$(($(($(stty size | awk '{print $1}'))) / 2 - 3))
h="\033[1;38;2;254;228;208m
\033[${row};${col}H●   ●
\033[$((row + 1));${col}H●   ●
\033[$((row + 2));${col}H●●●●●
\033[$((row + 3));${col}H●   ●
\033[$((row + 4));${col}H●   ●
"
printf "$h"
sleep 0.5
e="\033[1;38;2;254;228;208m
\033[${row};$((col + 6))H●●●●●
\033[$((row + 1));$((col + 6))H●
\033[$((row + 2));$((col + 6))H●●●●
\033[$((row + 3));$((col + 6))H●
\033[$((row + 4));$((col + 6))H●●●●●
"
printf "$e"
sleep 0.5
l="\033[1;38;2;254;228;208m
\033[${row};$((col + 12))H●
\033[$((row + 1));$((col + 12))H●
\033[$((row + 2));$((col + 12))H●
\033[$((row + 3));$((col + 12))H●
\033[$((row + 4));$((col + 12))H●●●●●
"
printf "$l"
sleep 0.5
l="\033[1;38;2;254;228;208m
\033[${row};$((col + 18))H●
\033[$((row + 1));$((col + 18))H●
\033[$((row + 2));$((col + 18))H●
\033[$((row + 3));$((col + 18))H●
\033[$((row + 4));$((col + 18))H●●●●●
"
printf "$l"
sleep 0.5
o="\033[1;38;2;254;228;208m
\033[${row};$((col + 24))H ●●●
\033[$((row + 1));$((col + 24))H●   ●
\033[$((row + 2));$((col + 24))H●   ●
\033[$((row + 3));$((col + 24))H●   ●
\033[$((row + 4));$((col + 24))H ●●●
"
printf "$o"
sleep 2
clear
col=$(($(($(stty size | awk '{print $2}'))) / 2 - 21))
license="\033[1;38;2;254;228;208m
\033[${col}G●      ●●●   ●●●  ●●●●● ●   ●  ●●●● ●●●●●
\033[${col}G●       ●   ●   ● ●     ●●  ● ●     ●
\033[${col}G●       ●   ●     ●●●●  ● ● ●  ●●●  ●●●●
\033[${col}G●       ●   ●   ● ●     ●  ●●     ● ●
\033[${col}G●●●●●  ●●●   ●●●  ●●●●● ●   ● ●●●●  ●●●●●
"
printf "$license"
sleep 3
license="\033[1;38;2;254;228;208m
License of termux-container:
 
  Copyright (c) 2021-2023 Moe-hacker

  Licensed under the Apache License, Version 2.0 (the \"License\");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an \"AS IS\" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

License of ruri and rootfstool:

  MIT License
 
  Copyright (c) 2022-2023 Moe-hacker
 
  Permission is hereby granted, free of charge, to any person obtaining a copy
  of this software and associated documentation files (the \"Software\"), to deal
  in the Software without restriction, including without limitation the rights
  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
  copies of the Software, and to permit persons to whom the Software is
  furnished to do so, subject to the following conditions:
 
  The above copyright notice and this permission notice shall be included in all
  copies or substantial portions of the Software.
 
  THE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
  SOFTWARE.
"
printf "$license" | pv -qL 100
sleep 2
license="\033[33m

///WARNING///

  * Your warranty is now void.
  * I am not responsible for anything that may happen to your device by using this program.
  * You do it at your own risk and take the responsibility upon yourself.
  * This program has no super cow power.

MAKE SURE YOU HAVE READ THE LICENSE AND WARNINGS,
AND AGREE WITH THEM.
OR YOU CAN JUST PRESS CTRL-C TO EXIT.

PRESS ENTER TO CONTINUE.
"
printf "$license"
read x
clear
col=$(($(($(stty size | awk '{print $2}'))) / 2 - 12))
note="\033[1;38;2;254;228;208m
\033[${col}G●   ●  ●●●  ●●●●● ●●●●● 
\033[${col}G●●  ● ●   ●   ●   ●
\033[${col}G● ● ● ●   ●   ●   ●●●●
\033[${col}G●  ●● ●   ●   ●   ●
\033[${col}G●   ●  ●●●    ●   ●●●●●
"
printf "$note"
note="\033[1;38;2;254;228;208m

About special characters:
  termux-container uses special characters for a better look.
  But not every font supports them.
  For a test:
  If this () is not a ubuntu logo:
  please run:
  mkdir -p ~/.termux
  cp \$PREFIX/share/termux-container/font.ttf ~/.termux
  termux-reload-settings

Reporting bugs:
  Please go to:
  https://github.com/Moe-hacker/termux-container/discussions


"
printf "$note"
printf "PRESS ENTER TO CONTINUE."
read x
clear
col=$(($(($(stty size | awk '{print $2}'))) / 2 - 15))
logo="\033[1;38;2;254;228;208m
\033[${col}G   ●●  ●●          ●   ●  ●●
\033[${col}G  ●      ●●         ● ●     ●
\033[${col}G  ●        ●         ●      ●
\033[${col}G  ●      ●●         ● ●     ●
\033[${col}G   ●●  ●●    ●●●●● ●   ●  ●●

\033[${col}G●   ● ●   ●   ●●    ●   ●  ●●●
\033[${col}G●  ●  ●● ●●  ●  ●   ●  ●  ●   ●
\033[${col}G●●●   ● ● ●   ●● ●  ●●●   ●
\033[${col}G●  ●  ●   ●  ●  ●   ●  ●  ●   ●
\033[${col}G●   ● ●   ●   ●● ●  ●   ●  ●●●
"
printf "$logo"