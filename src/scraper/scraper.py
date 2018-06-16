#!python3
"""
Scrape the top posts from a selection of subreddits
and write the link and post title to a file
Author: github.com/Hoenn
"""
import sys, praw, os

class DataEntry:
  def __init__(self, link, desc):
    self.link = link
    self.desc = desc
  def string(self):
    #Delimit string with tab since they are not valid 
    #In submission url or title
    return self.link+"\t"+self.desc

def auth():
  cd = os.path.dirname(__file__)
  path = os.path.join(cd, "../../bin/scraper.cred")
  #Read from bin/scraper.cred for auth paramaters
  cred = open(path).read().splitlines()

  #Authenticate and set readonly	
  reddit = praw.Reddit(client_id=cred[0],
    client_secret=cred[1],
    password=cred[2],
    user_agent=cred[3],
    username=cred[4])
  reddit.read_only = True

  return reddit

def scrape(r):

  data = []

  submissions = r.subreddit('aww').hot(limit=5)
  for s in submissions:
      data.append(DataEntry(s.url, s.title))
  return data


def writeOut(data):
  cd = os.path.dirname(__file__)
  path = os.path.join(cd, "../common/scrape")

  with open(path, 'w+') as result:
    for d in data:
      result.write(d.string()+"\n")
    result.close()
  

  

def main():
  writeOut(scrape(auth()))
  
  
if __name__ == "__main__":
  main()
