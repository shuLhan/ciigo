---
Title: Example Document title
description: asciidoc document for testing style of ciigo
---

# _Example `Document` **title**_

Author A <a@a.com>; Author mid_dle B <b@b.com>
v1.2.3, 18 July 2020: remark

Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat
non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.

## Paragraph

Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat
non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.

Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat
non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.

## Heading 1

Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua.

### Heading 2

Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua.

#### Heading 3

Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua.

##### Heading 4

Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua.

###### Heading 5

Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua.

## List

### Ordered list

Example of ordered list,

1. Orange

   Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
   tempor incididunt ut labore et dolore magna aliqua.

2. Red
   1. Yellow
   2. Blue

3. Green
   Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
   tempor incididunt ut labore et dolore magna aliqua.
   1. Black

      Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
      tempor incididunt ut labore et dolore magna aliqua.
      1. White

         Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
         tempor incididunt ut labore et dolore magna aliqua.

### Unordered list

Example of unordered list,

- Orange

  Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
  tempor incididunt ut labore et dolore magna aliqua.

- Red
  - Yellow
  - Blue

- Green

  Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
  tempor incididunt ut labore et dolore magna aliqua.
  - Black

    Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
    tempor incididunt ut labore et dolore magna aliqua.
    - White

      Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
      tempor incididunt ut labore et dolore magna aliqua.

### Checklists

Example of check list,

- [x] checked
- [x] also checked
- [ ] not checked

### Description lists

CPU
: The brain of the computer.

Hard drive
: Permanent storage for operating system and/or user files.

RAM
: Temporarily stores information the CPU uses during operation.

Keyboard
: Used to enter text or control items on the screen.

Mouse
: Used to point to and select items on your computer screen.

Monitor
: Displays information in visual form using text and graphics.

## Table

| Cell in column 1, row 1 | Cell in column 2, row 1 |
| ----------------------- | ----------------------- |
| Cell in column 1, row 2 | Cell in column 2, row 2 |
| Cell in column 1, row 3 | Cell in column 2, row 3 |

| Cell in column 1, header row | Cell in column 2, header row |
| ---------------------------- | ---------------------------- |
| Cell in column 1, row 2      | Cell in column 2, row 2      |
| Cell in column 1, row 3      | Cell in column 2, row 3      |
| Cell in column 1, row 4      | Cell in column 2, row 4      |

### Table with title (not supported)

.A table with a title

| Column 1, header row    | Column 2, header row    |
| ----------------------- | ----------------------- |
| Cell in column 1, row 2 | Cell in column 2, row 2 |

### Adjust column widths (not supported)

| Column 1   | Column 2   | Column 3   |
| ---------- | ---------- | ---------- |
| width of 2 | width of 1 | width of 3 |

### Align content by column (not supported)

|      Column name       |      Column name       |
| :--------------------: | :--------------------: |
| horizontally centered. | horizontally centered. |

| Column name | Column name | Column name |
| :---------: | :---------- | ----------: |
|   Center    | Left        |       Right |

### Table Striping (not supported)

| H1  | H2  |
| --- | --- |
| A1  | B1  |
| A2  | B2  |
| A3  | B3  |
| A4  | B4  |

## Verbatim and Source blocks

Example of pre,

```
var staticfs \*memfs.MemFS

type Ciigo struct {
HTTPServer *libhttp.Server
converter *Converter
watcher \*watcher
serveOpts ServeOptions
}
```

`ciigo` is a program (as in command line interface, CLI) and a library (as
in Go package) to write static web server with embedded files using
[AsciiDoc](https://asciidoctor.org/docs/what-is-asciidoc/)
and
[Markdown](https://www.markdownguide.org/)
markup format.

## Admonitions

Sample paragraph in between admonition.

> :note: Wolpertingers are known to nest in server racks.
> Enter at your own risk.

Sample paragraph in between admonition.

> :tip: While werewolves are hardy community members, keep in mind the
> following dietary concerns:
>
> 1. They are allergic to cinnamon.
> 2. More than two glasses of orange juice in 24 hours makes them howl in harmony with alarms and sirens.
> 3. Celery makes them sad.

Sample paragraph in between admonition.

> :important: While werewolves are hardy community members, keep in mind the
> following dietary concerns:
>
> 1. They are allergic to cinnamon.
> 2. More than two glasses of orange juice in 24 hours makes them howl in
>    harmony with alarms and sirens.
> 3. Celery makes them sad.

Sample paragraph in between admonition.

> :CAUTION: .Feeding the Werewolves
>
> While werewolves are hardy community members, keep in mind the following
> dietary concerns:
>
> - They are allergic to cinnamon.
> - More than two glasses of orange juice in 24 hours makes them howl in
>   harmony with alarms and sirens.
> - Celery makes them sad.

Sample paragraph in between admonition.

> :warning: Wolpertingers are known to nest in server racks.
> Enter at your own risk.

<!-- SPDX-License-Identifier: GPL-3.0-or-later -->
<!-- SPDX-FileCopyrightText: 2026 M. Shulhan <ms@kilabit.info> -->
