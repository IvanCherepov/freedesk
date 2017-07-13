#!/bin/sh

/server/bin/zendesk &
npm start --prefix /client
