/*
 * SPDX-License-Identifier: Apache-2.0
 * This file is part of termux-container.
 *
 * Copyright (c) 2023 Moe-hacker
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
// This file should be statically compiled.
#include "container-console.h"
// Used for catching Ctrl-C, restart container-console.
void restart(int /*unused*/)
{
  // Restart container-console.
  system("container-console");
  // Maybe not needed.
  system("stty icanon");
  system("stty echo");
  system("stty erase '^?'");
  // Exit.
  exit(0);
}
void show_greetings()
{
  /*
   * Maybe nothing's useful here, just for fun.
   */
  // Get the size of terminal.
  struct winsize size;
  ioctl(STDOUT_FILENO, TIOCGWINSZ, &size);
  unsigned short col = size.ws_col;
  if (col % 2 == 1)
  {
    col -= 1;
  }
  // For centering output.
  char space[col / 2 + 1];
  space[0] = '\000';
  if (col > 46)
  {
    col /= 2;
    col -= 22;
    for (unsigned short i = 1; i <= col; i++)
    {
      strcat(space, " ");
    }
  }
  else
  {
    strcat(space, "");
  }
  printf("%s%s\n", space, "\033[1;38;2;66;66;66m               ▅▅▀▀▀▀▀▀▀▀▀▀▀▀▅");
  printf("%s%s\n", space, "          ▅▅▀▀▀               ▀▀▅▅");
  printf("%s%s\n", space, "     ▅▅▅▀▀            ▅           ▀▅");
  printf("%s%s\n", space, "      ▅▀      ▅▀█▅▅▀▀▅▀▅        ▅▅  ▀▅");
  printf("%s%s\n", space, "     ▅▀   █▅▀▀  ▀     ▀ ▀▀▅▅    █ ▀▀▅ █");
  printf("%s%s\n", space, "    ▅▀   ▅▀  ▅▀      ▀▅    ▀▅   █▅███▀█");
  printf("%s%s\n", space, "  ▅▅█▀▅ █ ▅▅▀          ▀▀   █   ████   █");
  printf("%s%s\n", space, "      █ █ ▅▅▅▅▅        ▅▅▅▅▅ █  ▀█▀    █");
  printf("%s%s\n", space, "      █ █▀ ▅▅▅ ▀      ▀ ▅▅▅ ▀█   █     █");
  printf("%s%s\n", space, "      █ █ █\033[40;31m█▀█\033[0m\033[1;38;2;66;66;66m█        █\033[40;31m█▀█\033[0m\033[1;38;2;66;66;66m█ █   █     █");
  printf("%s%s\n", space, "     █  █ █\033[31m███\033[1;38;2;66;66;66m█        █\033[31m███\033[1;38;2;66;66;66m█ █   █     ▀▅");
  printf("%s%s\n", space, "    ▅▀  █  ▀▀▀          ▀▀▀  █   █      █");
  printf("%s%s\n", space, "  ▅▀▅▀ █                     █   █      █");
  printf("%s%s\n", space, " █   █ ▀▅ ▅▀▅   ▅▀▅   ▅▅     █   █      ▀▅");
  printf("%s%s\n", space, "▅█▅▅██▅ ▅██  ▀███ ▅████ ▀▅█▀▅▀   █       ▀▅");
  printf("%s%s\n", space, "███████ ▀██████████████████▀▀             █");
  printf("%s%s\n", space, " █    ▀▅  ██▀ ▀██▀▀██▀▀██▀█     █▀         █");
  printf("%s%s\n", space, " ▀▅     ▀▀█              ▅▀     █          █");
  printf("%s%s\n", space, "   ▀▅    █               █     ██        ▅▀");
  printf("%s%s\n", space, "     ▀▅▅▅▀                ▀▀▀▀▀ █        █");
  printf("%s%s\n", space, "        ▀                       ▀        ▀");
  printf("%s\n", "");
  printf("%s%s\n", space, "\033[1;38;2;254;228;208m        Console of termux-container");
  printf("%s%s\n", space, "         Made with  by Moe-hacker");
  printf("%s%s\n", space, "           WARNING:  NO WARRANTY");
  printf("%s%s\n", space, "         For usage,just type `help`");
}
int main(int argc, char *argv[])
{
  /*
   * 100% shit-code here...
   * At least it works...
   */
  // Used for termux-container.
  if (argc > 1)
  {
    show_greetings();
  }
  if (geteuid() == 0)
  {
    fprintf(stderr, "%s\n", "\033[33mWarning: container-console should not be run with root privileges!");
  }
  // Check if termux-container exists.
  if (system("container -t") != 0)
  {
    fprintf(stderr, "%s\n", "\033[31mError: termux-container not installed.");
    exit(1);
  }
  // For getting input.
  char input = 0;
  // For highlighting commands.
  char arg0[1024];
  arg0[0] = '\000';
  // Will not be highlighted.
  char arg1[1024];
  arg1[0] = '\000';
  // Command to run after pressing Enter.
  char command[2048];
  char *output = NULL;
  output = arg0;
  // History file.
  FILE *history = NULL;
  char history_file[PATH_MAX];
  // Maybe it's not necessary because $HOME in termux will always be /data/data/com.termux/files/home
  char *home = getenv("HOME");
  if (!home)
  {
    fprintf(stderr, "%s\n", "\033[31mError: $HOME is not set.");
    exit(1);
  }
  strcat(history_file, home);
  strcat(history_file, "/.container_history");
  // The line to read, count from bottom to top.
  unsigned int lines = 0;
  // Will be total lines-lines, because we should read the file from the top.
  unsigned int line_to_read = 0;
  // Line number.
  unsigned int line_now = 0;
  // Catch the input, do not wait for Enter.
  system("stty -icanon");
  // Disable displaying inputs.
  system("stty -echo");
  // For catching erase key.
  system("stty erase '^h'");
  // Catch CTRL-C signel and restart.
  signal(SIGINT, restart);
  // Console
  printf("\n\033[1;38;2;254;228;208mConsole > \033[0m");
  while (true)
  {
    // Get a char in input.
    input = (char)getchar();
    switch (input)
    {
    // Get CTRL-D(EOF).
    case '\004':
      // Reset terminal configs and exit.
      system("stty erase '^?'");
      system("stty icanon");
      system("stty echo");
      printf("\n");
      exit(0);
    // Up/down key.
    case '\033':
      // Ignore `[`.
      getchar();
      switch (getchar())
      {
      // Up key.
      case 'A':
        line_now = 0;
        for (unsigned int i = 0; i < (strlen(arg0) + strlen(arg1)); i++)
        {
          // Clear commands to show.
          printf("\b");
          printf(" ");
          printf("\b");
        }
        // Open history file.
        history = fopen(history_file, "a+e");
        // Back to the head of file.
        fseek(history, 0, SEEK_SET);
        arg0[0] = '\000';
        arg1[0] = '\000';
        output = arg0;
        // Count down lines.
        line_to_read = 0;
        for (int i = 0; (i = fgetc(history)) != EOF;)
        {
          if (i == '\n')
          {
            line_to_read++;
          }
        }
        // Fix errors if the history file is empty.
        if (line_to_read == 0)
        {
          goto end;
        }
        // Set the line to read.
        if (line_to_read > lines)
        {
          lines++;
          line_to_read = line_to_read - lines;
        }
        else
        {
          line_to_read = 0;
        }
        // Back to the head of file.
        fseek(history, 0, SEEK_SET);
        // Per-character read history file.
        while (true)
        {
          input = (char)fgetc(history);
          switch (input)
          {
          // Count lines.
          case '\n':
            // Reached the line to read.
            if (line_now == line_to_read)
            {
              // Finish reading.
              fclose(history);
              goto end;
            }
            else
            {
              // Increase line number.
              line_now++;
              continue;
            }
            break;
          // Space, for spliting arg0 and arg1.
          case ' ':
            // Check if we should write to arg1
            if (line_now == line_to_read)
            {
              if (arg0[0] != '\000')
              {
                // If we should switch output to arg1
                if (arg1[0] == '\000')
                {
                  output = arg1;
                }
                strcat(output, &input);
              }
            }
            break;
          // I don't remember what's it.
          // At least it really works, so do not touch.
          case '\000':
            break;
          default:
            // Write to output.
            if (line_now == line_to_read)
            {
              strcat(output, &input);
              break;
            }
          }
        }
      end:
        break;
        // Down key.
      case 'B':
        line_now = 0;
        for (unsigned int i = 0; i < (strlen(arg0) + strlen(arg1)); i++)
        {
          printf("\b");
          printf(" ");
          printf("\b");
        }
        if (lines > 1)
        {
          lines--;
        }
        else
        {
          lines = 1;
        }
        history = fopen(history_file, "a+e");
        fseek(history, 0, SEEK_SET);
        arg0[0] = '\000';
        arg1[0] = '\000';
        output = arg0;
        line_to_read = 0;
        for (int i = 0; (i = fgetc(history)) != EOF;)
        {
          if (i == '\n')
          {
            line_to_read++;
          }
        }
        if (line_to_read == 0)
        {
          goto end1;
        }
        line_to_read = line_to_read - lines;
        fseek(history, 0, SEEK_SET);
        while (true)
        {
          input = (char)fgetc(history);
          switch (input)
          {
          case '\n':
            if (line_now == line_to_read)
            {
              fclose(history);
              goto end1;
            }
            else
            {
              line_now++;
              continue;
            }
            break;
          case ' ':
            if (line_now == line_to_read)
            {
              if (arg0[0] != '\000')
              {
                if (arg1[0] == '\000')
                {
                  output = arg1;
                }
                strcat(output, &input);
              }
            }
            break;
          case '\000':
            break;
          default:
            if (line_now == line_to_read)
            {
              strcat(output, &input);
              break;
            }
          }
        }
      }
    end1:
      break;
    // Backspace.
    case 127:
      // Check if we should write to arg1.
      if (arg1[0] == '\000')
      {
        if (strlen(output) > 0)
        {
          // Del a character.
          output[strlen(output) - 1] = '\000';
          printf("\b");
          printf(" ");
          printf("\b");
        }
      }
      else
      {
        // Check if we should switch output to arg0
        if (strlen(output) > 0)
        {
          output[strlen(output) - 1] = '\000';
          printf("\b");
          printf(" ");
          printf("\b");
        }
        else
        {
          output = arg0;
          output[strlen(output) - 1] = '\000';
          printf("\b");
          printf(" ");
          printf("\b");
        }
      }
      break;
    // Enter.
    case '\n':
      // `exit` command, same as CTRL-D
      if (strcmp(arg0, "exit") == 0)
      {
        system("stty erase '^?'");
        system("stty icanon");
        system("stty echo");
        printf("\n");
        exit(0);
      }
      // Get the command to run.
      strcat(command, "container -e container_console_main ");
      strcat(command, arg0);
      strcat(command, arg1);
      printf("\n");
      // Execute the command.
      system(command);
      // Write to history file.
      history = fopen(history_file, "a+e");
      fprintf(history, "%s%s%s", arg0, arg1, "\n");
      fclose(history);
      // Reset all variables to default values.
      arg0[0] = '\000';
      arg1[0] = '\000';
      command[0] = '\000';
      lines = 0;
      output = arg0;
      printf("\033[1;38;2;254;228;208mConsole > \033[0m");
      break;
    // Space key.
    case ' ':
      // Ignore spaces at the beginning of the line.
      if (arg0[0] != '\000')
      {
        // Check if we should switch output to arg1.
        if (arg1[0] == '\000')
        {
          output = arg1;
          strcat(output, &input);
        }
        else
        {
          strcat(output, &input);
        }
      }
      break;
    // Other characters(common characters).
    default:
      strcat(output, &input);
      break;
    }
    // Highlighting, very very shit code here.
    if (strcmp(arg0, "help") == 0 || strcmp(arg0, "exit") == 0 || strcmp(arg0, "new") == 0 || strcmp(arg0, "search") == 0 || strcmp(arg0, "pull") == 0 || strcmp(arg0, "rmi") == 0 || strcmp(arg0, "cp") == 0 || strcmp(arg0, "ls") == 0 || strcmp(arg0, "rm") == 0 || strcmp(arg0, "login") == 0 || strcmp(arg0, "import") == 0 || strcmp(arg0, "export") == 0 || strcmp(arg0, "info") == 0)
    {
      printf("\033[11G\033[1;38;2;166;227;161m%s\033[0m%s", arg0, arg1);
    }
    else
    {
      printf("\033[11G%s%s", arg0, arg1);
    }
  }
  return 0;
}
