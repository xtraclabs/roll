#!/bin/sh

# This script sets up the application parameters for the auth code
# callback sample. The sample assumes it's an application that
# was registered with the application roll.

export AZ_SERVER=localhost:3000
export CLIENT_ID=111-222-3333
export CLIENT_SECRET=EVNIFUt3hMFYb9aHy1N8LyEmTsLS3y+XK6xDvVbU+E0=
export REDIRECT_URI=http://localhost:2000/oauth2_callback
