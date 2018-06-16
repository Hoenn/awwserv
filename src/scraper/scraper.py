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

def main():
  #Read from bin file to know what to scrape
  #Scrape the x top posts from y
  # 3:aww, 5:cute
  user_agent = ""
