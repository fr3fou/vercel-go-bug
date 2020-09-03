# vercel-go-bug

The newest version of the CLI ^20.0.0 (and also assuming the newest version of the @vercel/go runtime) and onwards, can't run handlers that have 3rd party dependencies. 

Here's the error output ![](https://i.imgur.com/rrVhk8p.png)
This issue doesn't occur when using CLI versions under 20.0.0
