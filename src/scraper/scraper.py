#!python3
"""
Scrape the top posts from a selection of subreddits
and write the link and post title to a file
Author: github.com/Hoenn
"""
import sys, praw

class DataEntry:
  def __init__(self):
    self.link = ""
    self.desc = ""

def auth():
  #Read from bin/scraper.cred for auth paramaters
  cred = open("../../bin/scraper.cred").read().splitlines()

  #Authenticate and set readonly	
  reddit = praw.Reddit(client_id=cred[0],
    client_secret=cred[1],
    password=cred[2],
    user_agent=cred[3],
    username=cred[4])
  reddit.read_only = True

  return reddit

def scrape(r):
  return ["abc", "def", "ghi"]

def writeOut(lines):
  with open('../scrape', 'a') as result:
    for l in lines:
      result.write(l+"\n")

  

def main():
  writeOut(scrape(auth()))
  
  
if __name__ == "__main__":
  main()
