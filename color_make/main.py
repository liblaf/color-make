import os
import pty
import re
import shlex
import shutil
import sys
import termios
from collections.abc import Sequence
from re import Pattern
from selectors import EVENT_READ, DefaultSelector
from subprocess import Popen

from rich.console import Console

INFO_PATTERN: Pattern = re.compile(
    pattern=r"^(make|Makefile|\w+.mk)(\[\d+\])?:(\d+:)? .*$"
)
ERROR_PATTERN: Pattern = re.compile(
    pattern=r"^(make|Makefile|\w+.mk)(\[\d+\])?:(\d+:)? \*\*\* .*$"
)


def c_print(console: Console, text: str) -> None:
    for line in text.splitlines(keepends=True):
        strip: str = line.strip()
        if strip and strip.isprintable():
            if ERROR_PATTERN.match(strip):
                console.print(line, style="bold reverse red", end="", highlight=False)
            elif INFO_PATTERN.match(strip):
                console.print(line, style="bold reverse blue", end="", highlight=False)
            else:
                words: list[str] = shlex.split(line, comments=True)
                if "->" not in strip and words and shutil.which(words[0]):
                    console.print(">>>", line, style="bold green on grey15", end="")
                else:
                    console.print(line, end="")
        else:
            console.file.write(line)
        console.file.flush()


def run(args: Sequence[str]) -> int:
    out_master, out_slave = pty.openpty()
    err_master, err_slave = pty.openpty()
    termios.tcsetwinsize(out_master, termios.tcgetwinsize(sys.stdin.fileno()))
    termios.tcsetwinsize(err_master, termios.tcgetwinsize(sys.stdin.fileno()))
    process: Popen = Popen(
        args=args, stdin=sys.stdin, stdout=out_slave, stderr=err_slave, close_fds=True
    )
    os.close(out_slave)
    os.close(err_slave)
    selector: DefaultSelector = DefaultSelector()
    out_console: Console = Console()
    err_console: Console = Console(stderr=True)
    selector.register(out_master, EVENT_READ, data=out_console)
    selector.register(err_master, EVENT_READ, data=err_console)
    while selector.get_map():
        for key, mask in selector.select():
            try:
                raw: bytes = os.read(key.fd, 8192)
            except OSError:
                selector.unregister(key.fd)
            else:
                console: Console = key.data
                try:
                    text: str = raw.decode()
                    c_print(console=console, text=text)
                except:
                    os.write(console.file.fileno(), raw)
    return process.wait()


def main() -> int:
    args: list[str] = ["make", *sys.argv[1:]]
    return run(args)


if __name__ == "__main__":
    sys.exit(main())
