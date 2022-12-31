import os
from subprocess import Popen

PROJDIR = os.path.realpath(os.path.curdir)

COMMANDS = [
    ["git", "init"],
    ["git", "add", "."],
    ["git", "commit", "-a", "-m", "Initial Commit."],
]

for command in COMMANDS:
    git = Popen(command, cwd=PROJDIR)
    git.wait()
