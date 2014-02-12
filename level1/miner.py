#!/usr/bin/env python
import hashlib
import subprocess
import sys

def system(cmd, stdin=None):
  if stdin:
    stdin = open(stdin, "r")

  p = subprocess.Popen(cmd.split(), stdout=subprocess.PIPE, stdin=stdin)
  return p.stdout.read().strip()

def hash(s):
  h = hashlib.sha1()
  h.update("commit ")
  h.update(str(len(s)))
  h.update("\0")
  h.update(s)
  return h.hexdigest()

def solve(difficulty, tree, parent, tm):
  counter = 0

  while True:
    if counter % 100000 == 0:
      print "counter:", counter

    body = """
tree %(tree)s
parent %(parent)s
author CTF user <user-e49y1fsd@example.com> %(tm)s + 0000
committer CTF user <user-e49y1fsd@example.com> %(tm)s + 0000

Give me a Gitcoin

%(counter)s
""".strip() % locals()

    sha1 = hash(body)
    if sha1 < difficulty and sha1 < parent:
      body_file = open("body.txt", "w")
      body_file.write(body)
      body_file.close()

      print "Mined a Gitcoin with commit: %(sha1)s" % locals()

      print "calling hash-object"
      print system("git hash-object -t commit -w body.txt")
      print "done"
      print

      print "calling reset"
      print system("git reset --hard %(sha1)s" % locals())
      print "done"
      print
      return

    counter += 1

if __name__ == "__main__":
  difficulty = sys.argv[1]
  tree = sys.argv[2]
  parent = sys.argv[3]
  tm = sys.argv[4]

  solve(difficulty, tree, parent, tm)
