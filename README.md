# Awwserv

An auto daily email blast taking top links from Reddit. A frontend to submit additional cuteness.

### /bin
Contains configuration files for:

Gmail SMTP: `email.cred`

Reddit API OAuth information: `scraper.cred`

### /src
#### scraper
A `python3` Reddit scraper that grabs posts to send. Requires OAuth configuration file in `scraper.cred` file in `/bin` to access Reddit API with `praw`. Writes a preformatted blob of link and description data to `src/common/scrape` to be picked up by the emailer.

#### queue
A `nodejs` webserver that hosts a static `index.html` to attach additional links to the next batch of emails. Listens to POST requests for formdata and writes to a file in `src/common/queue` to wait for the next consumption by the emailer
