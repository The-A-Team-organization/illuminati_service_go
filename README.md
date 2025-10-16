## How to run this docker

It is supposed that we are using Gmail
Make sure 2-Step Verification is enabled on your Google account.
Go to Google Account → Security → App passwords.

Edit main.go. Fill username := ""/ password := ""/ to := ""
build docker
docker build -t mailtest .
run docker
docker run --rm mailtest
check e-mail
